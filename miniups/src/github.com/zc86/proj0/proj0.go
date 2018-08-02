package main 

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
	"log"
	"github.com/golang/protobuf/proto"
	// protobuf
	"UPS"
	"UAComm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// type locations []*UPS.UDeliveryLocation

func checkError(err error) {
	if err != nil {
		panic(err.Error())
		os.Exit(1)
	}
}

func sTruckInit(conn net.Conn, db *sqlx.DB) {
	cnt := &UPS.UConnect{}
	var numtruck int32 = 10
	fmt.Println("Enter Number of Trucks: ")
	// for {
	// 	fmt.Scanf("%d\n", &numtruck)
	// 	if numtruck <= 0 {
	// 		fmt.Println("Truck numbers should not be negative!\n", "Please Retry: ")
	// 		continue
	// 	} else {
	// 		cnt.Numtrucksinit = proto.Int32(numtruck)
	// 		break
	// 	}
	// }
	cnt.Numtrucksinit = proto.Int32(numtruck)
	// initialize truck table
	var i int32
	for i = 0; i < numtruck; i++ {
		Insert_truck(db, i, 0, 0, 0)
	}
	out, err := proto.Marshal(cnt)
	if err != nil {
		fmt.Println("Failed to encode UConnect")
	}
	n := proto.Size(cnt)
	out = append(proto.EncodeVarint(uint64(n)), out...)
	conn.Write(out)
}

func rConnect(conn net.Conn) *UPS.UConnected {
	cnted := &UPS.UConnected{}
	mesg := make([]byte, 1024)
	size, err := conn.Read(mesg)
	checkError(err)
	mesg = mesg[:size]
	_, n := proto.DecodeVarint(mesg) 
	mesg = mesg[n:]
	// func Unmarshal(buf []byte, pb Message) error
	if err := proto.Unmarshal(mesg, cnted); err != nil {
		fmt.Println("Failed to parse UConnected")
	}
	fmt.Println("UConnected Worldid:\n", *cnted.Worldid)
	return cnted
}

func addGoPickup(truckid, whid int32, ucmd *UPS.UCommands) {
	pkup := &UPS.UGoPickup{}
	pkup.Truckid = proto.Int32(truckid)
	pkup.Whid = proto.Int32(whid)
	ucmd.Pickups = append(ucmd.Pickups, pkup)
}

// GoDeliver Field => 0 or 1  --- No need provide add function

// initiate UGoDeliver with truck id, then call this function to add deliveries
func addDeliverLocation(pid int64, x, y int32, udl *[]*UPS.UDeliveryLocation) {
	dlctn := &UPS.UDeliveryLocation{}
	dlctn.Packageid = proto.Int64(pid)
	dlctn.X = proto.Int32(x)
	dlctn.Y = proto.Int32(y)
	*udl = append(*udl, dlctn)
}

func addPkg(pid int64, status string, ua *UAComm.UACommands) {
	p := &UAComm.PackageStatus{}
	p.Pkgid = proto.Int64(pid)
	p.Status = proto.String(status)
	ua.Pks = append(ua.Pks, p)
}

// func addStartLoad(pkgids *[]*int64, pkgid int64, ua *UAComm.UACommands) {
// 	sl := &UAComm.StartLoad{}
// 	sl.Whid = proto.Int32(whid)
// 	sl.Packageid = proto.Int64(pkgid)
// 	sl.Truckid = proto.Int32(truckid)
// 	ua.Sl = append(ua.Sl, sl)
// }

func main() {
	// Connect to MySQL
	db, err := sqlx.Connect("mysql", "root:123@tcp(db:3306)/mysql")
	//db, err := sqlx.Connect("postgres", "user=jichun dbname=ups sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	Setup_db(db)
	
	// Establish Connection with World
	wconn, err := net.Dial("tcp", "152.3.77.189:12345")
	if err != nil {
		fmt.Println("Connect to World Failed")
	}
	fmt.Println("-----Connected to World-----")

	// pickupCmd(conn, 0, 1)

	// Accept the connection request from Amazon
	fmt.Println("-----Start UPS Server-----")
	ln, err := net.Listen("tcp", ":34567")
	checkError(err)
	// for {
	aconn, err := ln.Accept()
    if err != nil {
    	fmt.Println("Failed to Accept")
	}
	fmt.Println("-----Connected to Amazon-----")

	for {
		sTruckInit(wconn, db)
		if ucnted := rConnect(wconn); ucnted.Error == nil {
			uacmd := &UAComm.UACommands{}
			uacmd.Worldid = proto.Int64(*ucnted.Worldid)
			sendAmazon(aconn, uacmd)
			break
		}
	}
	// sTruckInit(wconn, db)
	cw := make(chan struct{})
	ca := make(chan struct{})

	go rcvWorld(aconn, wconn, db, cw)
	go rcvAmazon(aconn, wconn, db, ca)
	// }

	// sendADisconn(aconn)
	<-ca 
	aconn.Close()
	
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan, os.Interrupt)
	// Block until a signal is received.
	s := <-signal_chan
	fmt.Println("Got signal:", s)
	go sendWDisconn(wconn)
	<-cw
	wconn.Close()
	
	// for {
		// infinite loop
	// }
	
  	ln.Close()
}

func sendWDisconn(wconn net.Conn) {
	ucmd := &UPS.UCommands{}
	ucmd.Disconnect = proto.Bool(true)
	sendWorld(wconn, ucmd)
}

func sendAFinished(aconn net.Conn) {
	uacmd := &UAComm.UACommands{}
	uacmd.Finished = proto.Bool(true)
	sendAmazon(aconn, uacmd)
}

func rcvWorld(aconn, wconn net.Conn, db *sqlx.DB, c chan struct{}) {
	fmt.Println("------Waiting World Response------")
	// defer wconn.Close()
	for {
		data := make([]byte, 8192)
		size, err := wconn.Read(data)
		if err != nil {
			fmt.Println("Failed Read from World:\n", err.Error())
			// continue
			break
		}
		mesg := data[:size]
		// x -> size, n -> number of bytes
		_, n := proto.DecodeVarint(mesg) 
		mesg = mesg[n:]
		fmt.Println("Original World UReponses :\n", mesg)
		go handleWorld(aconn, db, mesg, c)
	}
}

func handleWorld(aconn net.Conn, db *sqlx.DB, mesg []byte, c chan struct{}) {
	rsp := &UPS.UResponses{}
	if err := proto.Unmarshal(mesg, rsp); err != nil {
		fmt.Println("Failed to parse UResponses")
	}
	fmt.Println("Unmarshal World UReponses :\n", *rsp)
	// Finished from world
	if rsp.Finished != nil && *rsp.Finished == true {
		fmt.Println("Receive Finished from World")
		close(c)
	}
	// Completions
	for _, completion := range rsp.Completions {
		Process_finish_from_W(aconn, db, *completion.Truckid, *completion.X, *completion.Y)
	}
	// DeliveryMade
	for _, delivery := range rsp.Delivered {
		Process_deliverymade_from_W(aconn, db, *delivery.Packageid)
	}
	if rsp.Error != nil {
		fmt.Println("Error from World:\n", *rsp.Error)
	}
}

func rcvAmazon(aconn, wconn net.Conn, db *sqlx.DB, c chan struct{}) {
	fmt.Println("------Waiting Amazon Response------")
	// defer aconn.Close()
	for {
		data := make([]byte, 8192)
		size, err := aconn.Read(data)
		if err != nil {
			fmt.Println("Failed Read from Amazon:\n", err.Error())
			// continue
			break
		}
		mesg := data[:size]
		// x -> size, n -> number of bytes
		_, n := proto.DecodeVarint(mesg) 
		mesg = mesg[n:]
		fmt.Println("Original Amazon UReponses :\n", mesg)
		go handleAmazon(aconn, wconn, db, mesg, c)
	}
}

func handleAmazon(aconn, wconn net.Conn, db *sqlx.DB, mesg []byte, c chan struct{}) {
	au := &UAComm.AUCommands{}
	if err := proto.Unmarshal(mesg, au); err != nil {
		fmt.Println("Failed to parse AUCommands")
	}
	fmt.Println("Unmarshal Amazon Message:\n", *au)
	// The First Message received from Amazon
	if len(au.Whinfo) != 0 {
		for _, whinfo := range au.Whinfo {
			Insert_warehouse(db, *whinfo.Whid, *whinfo.Whx, *whinfo.Why)
		}
	}
	// Finished from Amazon
	if au.Disconnect != nil && *au.Disconnect == true {
		for {
			res := checkOrderDelivered(db) // 
			if res == true {
				sendAFinished(aconn)
				break
			} 
			time.Sleep(100 * time.Millisecond)
		}
		close(c)
	}
	if au.Pc != nil {
		purchase := au.Pc
		var priority int32 = 0
		if *purchase.Isprime == false {
			priority = 1
		}
		// Insert_warehouse(db, *purchase.Whid, *purchase.Wh_x, *purchase.Wh_y) 
		Process_order_from_A(aconn, wconn, db, *purchase.Orderid, *purchase.Upsuserid, *purchase.Whid, *purchase.X, *purchase.Y, priority, purchase.Things)
	}
	if au.Ldd != nil {
		loaded := au.Ldd
		Process_loaded_from_A(aconn, wconn, db, *loaded.Packageid)
	}
	if au.Chdes != nil {
		chdest := au.Chdes
		res := Process_change_destination_from_A(db, *chdest.Pkgid, *chdest.X, *chdest.Y)
		sendChDesResult(aconn, *chdest.Pkgid, res)
	}
}

func sendChDesResult(aconn net.Conn, pkgid int64, res bool) {
	uacmd := &UAComm.UACommands{}
	dcr := &UAComm.DesChangeResult{}
	dcr.Pkgid = proto.Int64(pkgid)
	dcr.Res = proto.Bool(res)
	uacmd.Dcr = dcr
	sendAmazon(aconn, uacmd)
}

func loadInform(aconn net.Conn, whid, truckid int32, pkgs []ORDER) {
	ua := &UAComm.UACommands{}
	sl := &UAComm.StartLoad{}
	sl.Whid = proto.Int32(whid)
	sl.Truckid = proto.Int32(truckid)
	for _, pkg := range pkgs {
		// addStartLoad(&sl.Packageid, pkg.Order_id)
		sl.Packageid = append(sl.Packageid, pkg.Order_id)
	}
	ua.Sl = append(ua.Sl, sl)
	// ua.whid = whid
	// ua.packageid = pkgid
	// ua.truckid = truckid
	sendAmazon(aconn, ua)
}

func pkgIdInform(conn net.Conn, pkgid, amazonOrderid int64) {
	ua := &UAComm.UACommands{}
	pidstruct := &UAComm.PackageID{}
	pidstruct.Pkgid = proto.Int64(pkgid)
	pidstruct.Orderid = proto.Int64(amazonOrderid)
	ua.Pid = pidstruct
	sendAmazon(conn, ua)
}

func pickupCmd(conn net.Conn, truckid, whid int32) {
	ucmd := &UPS.UCommands{}
	ucmd.Simspeed = proto.Uint32(100000)
	addGoPickup(truckid, whid, ucmd)
	// fmt.Println("Simspeed:\n", *ucmd.Simspeed)
	fmt.Println("Send UGoPickup:\n", *ucmd)
	sendWorld(conn, ucmd)
}

func sendWorld(conn net.Conn, ucmd *UPS.UCommands) {
	out, err := proto.Marshal(ucmd)
	if err != nil {
		fmt.Println("Failed to encode UCommands")
	}
	n := proto.Size(ucmd)
	out = append(proto.EncodeVarint(uint64(n)), out...)
	conn.Write(out)
}

func sendAmazon(conn net.Conn, uacmd *UAComm.UACommands) {
	fmt.Println("Send to Amazon Message:\n", *uacmd)
	out, err := proto.Marshal(uacmd)
	if err != nil {
		fmt.Println("Failed to encode UACommands")
	}
	n := proto.Size(uacmd)
	out = append(proto.EncodeVarint(uint64(n)), out...)
	conn.Write(out)
}

func addDeliver(conn net.Conn, truck_id int32, orders []ORDER) {
	ucmd := &UPS.UCommands{}
	ugd := &UPS.UGoDeliver{}
	ugd.Truckid = proto.Int32(truck_id)
	// ucmd.Deliveries = append(ucmd.Deliveries, ugd)
	for _, order := range orders {
		addDeliverLocation(order.Order_id, order.Des_x, order.Des_y, &ugd.Packages) // ???
	}
	ucmd.Deliveries = append(ucmd.Deliveries, ugd)
	fmt.Println("GoDeliver Message to World:\n", *ucmd)
	sendWorld(conn, ucmd)
}

func sendStatus(conn net.Conn, orders []ORDER) {
	uacmd := &UAComm.UACommands{}
	for _, order := range orders {
		addPkg(order.Order_id, "out for delivery", uacmd)
	}
	sendAmazon(conn, uacmd)
}