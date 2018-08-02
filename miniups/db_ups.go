package main

import (
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"fmt"
	"net"
	"time"
	// "UPS"
	"UAComm"
)

type TRUCK struct {
	Truck_id int32
	T_status  string
	X  int32
	Y   int32
	Towhid int32
}

type ORDER struct{
	Order_id int64
	O_status string
	User_name string
	Truck_id int32
	Time string
	Whid int32
	Wh_x int32
	Wh_y int32
	Des_x int32
	Des_y int32
	Priority int32
}

type WH struct{
	Whid int32
	X int32
	Y int32
}

type ITEM struct{
	item_id int64
	description string
	count int32
	order_id int64
}

/*
func main() {
	db, err := sqlx.Connect("mysql", "root:SJCzq@tcp(localhost:3306)/ups")
	//db, err := sqlx.Connect("postgres", "user=jichun dbname=ups sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	Setup_db(db)
	multiple_order_same_warehouse_test(db)
}
*/
/*
func basic_test(db *sqlx.DB){
	Insert_truck(db, 1, 1, 1, 0)
	Insert_truck(db, 2, 2, 2, 0)
	Insert_truck(db, 3, 3, 3, 0)
	Insert_warehouse(db, 1, 1,1)
	Insert_warehouse(db, 2, 2,2)
	Insert_warehouse(db, 3, 3,3)
	Process_order_from_A(db, 1, 1, 10, 10)
	Process_order_from_A(db, 1, 2, 11, 11)
	Process_order_from_A(db, 1, 3, 12, 12)
	Process_finish_from_W(db, 1, 1,1)
	Process_change_destination_from_A(db, 1, 99,99)
	Process_loaded_from_A(db, 1)
	Process_deliverymade_from_W(db, 1)
	Process_finish_from_W(db, 1, 100, 100)
}

func multiple_order_same_warehouse_test(db *sqlx.DB){
	//multiple orders at the same warehouse, either come during the truck en route to warehouse or during the loading at that warehouse
	Insert_truck(db, 1, 1, 1, 0)
	Insert_warehouse(db, 1, 1,1)
	Insert_warehouse(db, 2, 2,2)
	Insert_warehouse(db, 3, 3,3)
	Process_order_from_A(db, 1, 1, 10, 10)
	Process_order_from_A(db, 1, 1, 11, 11)//during the truck en route to warehouse
	Process_finish_from_W(db, 1, 1,1)
	Process_order_from_A(db, 1, 1, 12, 12)//during the loading at that warehouse
	Process_loaded_from_A(db, 1)
	Process_loaded_from_A(db, 2)
	Process_change_destination_from_A(db, 1, 99, 99)//can change
	Process_loaded_from_A(db, 3)
	Process_change_destination_from_A(db, 2, 98, 98)//cannot change
	Process_deliverymade_from_W(db, 1)
	Process_deliverymade_from_W(db, 2)
	Process_deliverymade_from_W(db, 3)
	Process_finish_from_W(db, 1, 100, 100)
}

func multiple_different_pickup_test(db *sqlx.DB){
	//multiple different warehouse pickup assigned to one truck
	Insert_truck(db, 1, 1, 1, 0)
	Insert_warehouse(db, 1, 1,1)
	Insert_warehouse(db, 2, 2,2)
	Insert_warehouse(db, 3, 3,3)
	Process_order_from_A(db, 1, 1, 10, 10)
	Process_finish_from_W(db, 1, 1, 1)
	Process_order_from_A(db, 1, 2, 12, 12)//come a order when truck 1 waiting
	Process_loaded_from_A(db, 1)//pick up all order at warehouse 1 then pickup other assigned orders
	Process_finish_from_W(db, 1, 2, 2)
	Process_loaded_from_A(db, 2)
	Process_deliverymade_from_W(db, 1)
	Process_deliverymade_from_W(db, 2)
	Process_finish_from_W(db, 1, 100, 100)
}
*/
func Setup_db(db *sqlx.DB){
	Drop_table(db)
	schema0 := `CREATE TABLE truck (
    truck_id INTEGER,
    t_status VARCHAR(100),
    x INTEGER,
    y INTEGER,
    towhid INTEGER,
    PRIMARY KEY ( truck_id )
	);`
	db.Exec(schema0)

	schema1 := `CREATE TABLE orders (
    order_id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    o_status VARCHAR(100),
    user_name VARCHAR(100),
    truck_id INTEGER REFERENCES truck(truck_id),
    time DATETIME,
    whid INTEGER,
    wh_x INTEGER,
    wh_y INTEGER,
  	des_x INTEGER,
  	des_y INTEGER,
  	priority INTEGER,
    PRIMARY KEY ( order_id ),
          	FOREIGN KEY  (truck_id)
      		REFERENCES truck(truck_id)
      		ON UPDATE CASCADE ON DELETE RESTRICT
	);`
	db.Exec(schema1)

	schema2 := `CREATE TABLE warehouse (
    whid INTEGER,
    x INTEGER,
    y INTEGER,
    PRIMARY KEY ( whid )
	);`
	db.Exec(schema2)

	schema3 := `CREATE TABLE item (
	iid INTEGER UNSIGNED AUTO_INCREMENT,
    item_id INTEGER NOT NULL,
    description VARCHAR(255) NOT NULL,
    count INTEGER NOT NULL,
    order_id INTEGER UNSIGNED NOT NULL,
    PRIMARY KEY (iid),
  	FOREIGN KEY  (order_id)
		REFERENCES orders(order_id)
		ON UPDATE CASCADE ON DELETE RESTRICT
	);`
	_, err := db.Exec(schema3)
	if err != nil {
		panic(err.Error())
		os.Exit(1)
	}

}

func Drop_table(db *sqlx.DB){
	_, err := db.Exec("DROP TABLE IF EXISTS item")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("DROP TABLE IF EXISTS orders")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("DROP TABLE IF EXISTS truck")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS warehouse")
	if err != nil {
		panic(err)
	}

}

func Insert_truck (db *sqlx.DB, truck_id int32, x int32, y int32, towhid int32) bool{
	//initialize t_status as "free"
	tx, err := db.Begin()
	defer tx.Commit()
	_, err= tx.Exec("INSERT INTO truck (truck_id, t_status, x, y, towhid) VALUES (?, ?, ?, ?, ?)", truck_id, "idle", x, y, towhid)
	if err!=nil{
		return false
	}
	return true
}

func Insert_warehouse (db *sqlx.DB, whid int32, x int32, y int32) bool{
	//initialize t_status as "free"
	tx, err := db.Begin()
	defer tx.Commit()
	_, err= tx.Exec("INSERT INTO warehouse (whid, x, y) VALUES (?, ?, ?)", whid, x, y)
	if err!=nil{
		return false
	}
	return true
}

func Get_wh_xy(db *sqlx.DB, tx *sqlx.Tx, whid int32) WH{
	var warehouse WH
	err := tx.Get(&warehouse,"SELECT * FROM warehouse WHERE whid = ?;", whid)
	if err!=nil{
		log.Fatalln(err)
	}
	return warehouse
}

func Insert_order (db *sqlx.DB, tx *sqlx.Tx, user_name string, whid int32, des_x int32, des_y int32, priority int32) int64{
	//no commit, need commit!!!
	//if not specified user_id, passing 0 to user_id
	//initialize o_status as "wait arrange pick up"
	//initialize truck_id as 0, need to assign a free truck
	//return auto_increment order_id
	warehouse := Get_wh_xy(db, tx, whid)
	wh_x := warehouse.X
	wh_y := warehouse.Y
	t := time.Now().Format("2006-01-02 15:04:05")
	res, err := tx.Exec("INSERT INTO orders (o_status, user_name, time, whid, wh_x, wh_y, des_x, des_y, priority) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", "created", user_name, t, whid, wh_x, wh_y, des_x, des_y, priority)
	if err!=nil{
		log.Fatalln(err)
	}
	order_id, err := res.LastInsertId()
	return order_id
}

func Insert_item (db *sqlx.DB, tx *sqlx.Tx, item_id int64, description string, count int32, order_id int64){//@@@
	//no commit, need commit!!!
	_, err := tx.Exec("INSERT INTO item (item_id, description, count, order_id) VALUES (?, ?, ?, ?)", item_id, description, count, order_id)
	if err!=nil{
		log.Fatalln(err)
	}
}

func Update_tstatus(db *sqlx.DB, tx *sqlx.Tx, u_status string, truck_id int32, x int32, y int32, towhid int32)bool{
	//update truck status (idle, to warehouse, waiting for pick up, delivering).
	//if delivering, xy is meaningless, set random or 0,0
	//if waiting for pick up, xy no need to update, set random or 0,0
	var status string
	err := tx.Get(&status,"SELECT truck.t_status FROM truck WHERE truck_id = ? FOR UPDATE;", truck_id)
	if err!=nil{
		return false
		log.Fatalln(err)
	}
	if u_status!="waiting for pick up" {
		_, err = tx.Exec("UPDATE truck SET truck.t_status = ?, truck.x = ?, truck.y = ?, truck.towhid = ?  WHERE truck.truck_id = ?;", u_status, x, y, towhid, truck_id)
		//err = tx.Commit()
		fmt.Println("update truck: ", truck_id, " status from ", status, " to ", u_status, " location as x=", x, " y=", y, " on the way to warehouse:", towhid)
		if err != nil {
			log.Fatalln(err)
		}
	}else{
		_, err = tx.Exec("UPDATE truck SET truck.t_status = ? WHERE truck.truck_id = ?;", u_status, truck_id)
		//err = tx.Commit()
		fmt.Println("update truck ", truck_id, " status to ", u_status)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return true
}

func Update_ostatus(db *sqlx.DB, tx *sqlx.Tx, o_status string, order_id int64)bool{
	//created, truck en route to warehouse, waiting pick up, picked up, out for delivery, delivered
	//no commit, need commit!!!
	var status string
	err := tx.Get(&status,"SELECT orders.o_status FROM orders WHERE order_id = ? FOR UPDATE;", order_id)
	if err!=nil{
		return false
		log.Fatalln(err)
	}
	_, err = tx.Exec("UPDATE orders SET orders.o_status = ? WHERE orders.order_id = ?;", o_status, order_id)
	fmt.Println("update order ", order_id, " status from ", status, " to ", o_status)
	if err != nil {
		log.Fatalln(err)
	}
	return true
}

func Update_otruckid(db *sqlx.DB, tx *sqlx.Tx, truck_id int32, order_id int64)bool{
	//no commit, need commit!!!
	var id int64
	err := tx.Get(&id,"SELECT orders.order_id FROM orders WHERE order_id = ? FOR UPDATE;", order_id)
	if err!=nil{
		log.Fatalln(err)
		return false
	}
	_, err = tx.Exec("UPDATE orders SET orders.truck_id = ? WHERE orders.order_id = ?;" ,truck_id ,order_id)
	fmt.Println("assign order ", order_id, " to truck:", truck_id)
	if err != nil {
		log.Fatalln(err)
	}
	return true
}

func Get_truckid_byorder(db *sqlx.DB, tx *sqlx.Tx, order_id int64) int32 {
	var truck_id int32
	tx.Get(&truck_id,"SELECT orders.truck_id FROM orders WHERE orders.order_id = ?;",order_id)
	return truck_id
}

func Get_unsentloading_orders_ontruck_and_send(aconn net.Conn, db *sqlx.DB, tx *sqlx.Tx, truck_id int32){
	orders := []ORDER{}
	tx.Select(&orders, "select * from orders WHERE orders.truck_id = ? AND orders.o_status = ? FOR UPDATE; ", truck_id, "truck en route to pick up")
	if len(orders)!=0{
		//send load for orders
		_, err := tx.Exec("UPDATE orders SET orders.o_status = ? WHERE orders.truck_id = ? AND orders.o_status = ?;" ,"waiting pick up", truck_id ,"truck en route to pick up")
		fmt.Println("send load to Amazon for orders:", orders)
		loadInform(aconn, orders[0].Whid, truck_id, orders)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func Get_order_destination(db *sqlx.DB, tx *sqlx.Tx, order_id int64)(int32, int32){
	var order ORDER
	tx.Get(&order, "select * from orders WHERE orders.order_id = ? FOR UPDATE ", order_id)
	return order.Des_x, order.Des_y
}

func checkOrderDelivered(db *sqlx.DB) bool {
	// orders := []ORDER{}
	// tx.Select(&orders, "select * from orders WHERE NOT o_status = ?;", "delivered")
	rows, _ := db.Queryx("select * from orders WHERE NOT o_status = ?", "delivered")
	if rows.Next() {
		return false
	} else {
		return true
	}
}

func Get_free_truck(wconn net.Conn, db *sqlx.DB, order_id int64, whid int32, wh_x int32, wh_y int32)bool{
	//no commit, need commit!!!
	tx,_ := db.Beginx()
	defer tx.Commit()
	trucks := []TRUCK{}
	tx.Select(&trucks, "select * from truck WHERE (truck.t_status = ? OR truck.t_status = ?) AND truck.x = ? AND truck.y = ? FOR UPDATE ", "to warehouse", "waiting for pick up", wh_x, wh_y)
	if len(trucks)!=0 {
		//update orders.truck_id
		Update_otruckid(db, tx, trucks[0].Truck_id, order_id)
		//update ostatus
		Update_ostatus(db, tx, "truck en route to pick up", order_id)
		return true
	}
	tx.Select(&trucks, "select * from truck WHERE truck.t_status = ? FOR UPDATE ", "idle")
	if len(trucks) == 0 {
		tx.Select(&trucks, "select * from truck WHERE truck.t_status = ? FOR UPDATE ", "waiting for pick up")
		if len(trucks) == 0{
			return false
		}
		//update orders.truck_id
		Update_otruckid(db, tx, trucks[0].Truck_id, order_id)
	}else{
		//find idle truck, update tstatus to busy
		Update_tstatus(db, tx, "to warehouse", trucks[0].Truck_id, wh_x, wh_y, whid)
		//update orders.truck_id
		Update_otruckid(db, tx, trucks[0].Truck_id, order_id)
		//update ostatus
		Update_ostatus(db, tx, "truck en route to pick up", order_id)
		//send pickup command to trucks[0].Truck_id
		fmt.Println("send Gopickup for order:", order_id, " to truck:", trucks[0].Truck_id)
		//@@@
		// ucmd := &UCommands{}
		pickupCmd(wconn, trucks[0].Truck_id, whid)
	}
	return true
}

func Process_order_from_A(aconn, wconn net.Conn, db *sqlx.DB, amazonOrderid int64, user_name string, whid int32, des_x int32, des_y int32, priority int32, items []*UAComm.Product){
	//insert new order as created, not assigned a truck
	tx, _ :=db.Beginx()
	order_id := Insert_order(db, tx, user_name, whid, des_x, des_y, priority)
	for i:=0; i<len(items); i++{
		Insert_item(db, tx, *items[i].Id, *items[i].Description, *items[i].Count, order_id)
	}
	warehouse := Get_wh_xy(db, tx, whid)
	if (priority == 0){
		fmt.Println("First priority, we will process your order first!")
	}
	tx.Commit()
	// Send package ID to Amazon
	pkgIdInform(aconn, order_id, amazonOrderid)
	for{
		fmt.Println("order: ", order_id, "try to get a free truck")
		if(Get_free_truck(wconn, db, order_id, whid, warehouse.X, warehouse.Y)){
			fmt.Println("order: ", order_id, "get a free truck!")
			break;
		}
		time.Sleep(1000*time.Millisecond)
	}
}


func Process_loaded_from_A(aconn, wconn net.Conn, db *sqlx.DB, order_id int64){//assume A send loaded+package_id
	tx, _ :=db.Beginx()
	defer tx.Commit()
	truck_id := Get_truckid_byorder(db, tx, order_id)
	Update_ostatus(db, tx, "picked up", order_id)
	orders := []ORDER{}
	orders_insamewh := []ORDER{}
	tx.Select(&orders, "select * from orders WHERE (orders.o_status = ? OR orders.o_status = ?) AND orders.truck_id = ? FOR UPDATE ", "waiting pick up", "truck en route to warehouse", truck_id)
	if len(orders) == 0{//all orders have been picked up at this warehouse, this is the last loaded order
		fmt.Println("all orders have been picked up by truck:", truck_id)
		tx.Select(&orders, "select * from orders WHERE orders.o_status = ? AND orders.truck_id = ? ORDER BY orders.time*orders.priority ASC FOR UPDATE ", "created", truck_id)
		if len(orders)!=0 {//has other assigned orders to pick up
			tx.Select(&orders_insamewh, "select * from orders WHERE orders.o_status = ? AND orders.truck_id = ? AND orders.whid = ? ORDER BY orders.time*orders.priority ASC FOR UPDATE ", "created", truck_id, orders[0].Whid)
			Update_tstatus(db, tx, "to warehouse", truck_id, orders[0].Wh_x, orders[0].Wh_y, orders[0].Whid)
			// Update_ostatus(db, tx, "truck en route to pick up", orders[0].Order_id)
			//send Gopickup orders[0].Order_id to truck_id
			fmt.Println("continue pick up, send Gopickup for order:", orders_insamewh, " to truck:", truck_id)
			//@@
			pickupCmd(wconn, truck_id, orders[0].Whid)
			for i :=0; i<len(orders_insamewh); i++{
				Update_ostatus(db, tx, "truck en route to pick up", orders_insamewh[i].Order_id)
			}

		} else{//all assigned orders have been picked up, start delivery
			Update_tstatus(db, tx, "delivering", truck_id, 0, 0, 0)
			tx.Select(&orders, "select * from orders WHERE orders.truck_id = ? AND orders.o_status = ? FOR UPDATE ", truck_id, "picked up")
			_, err := tx.Exec("UPDATE orders SET orders.o_status = ? WHERE orders.truck_id = ? AND orders.o_status = ?;", "out for delivery", truck_id, "picked up")
			fmt.Println("update order:", orders, " status from picked up to out for delivery")
			if err != nil {
				log.Fatalln(err)
			}
			// tx.Commit()
			//send Godelivery for all Order_id in orders to world
			fmt.Println("send Godelivery for orders:", orders, " to truck:", truck_id)
			//@@
			addDeliver(wconn, truck_id, orders)
			//send out for delivery status for all Order_id in orders to Amazon
			//@@
			// fmt.Println("Send to Amazon in loaded from A")
			sendStatus(aconn, orders)
		}
	}else {//this is not the last loaded order
		fmt.Println("this is not the last loaded order by truck:", truck_id)
		//check unsent load order
		//send load
		//update ostatus as waiting pick up
		Get_unsentloading_orders_ontruck_and_send(aconn, db, tx, truck_id)
		return
	}
}

func Process_change_destination_from_A(db *sqlx.DB, order_id int64, des_x int32, des_y int32)bool{
	tx, _ :=db.Beginx()
	defer tx.Commit()
	var o_status string
	tx.Get(&o_status,"SELECT orders.o_status FROM orders WHERE orders.order_id = ? FOR UPDATE ;",order_id)
	if o_status != "out for delivery"{
		_, err := tx.Exec("UPDATE orders SET orders.des_x = ?, orders.des_y = ? WHERE orders.order_id = ?;", des_x, des_y, order_id)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("change destination for order:", order_id, " to (", des_x, ", ", des_y, ")")
		return true
	}else{
		fmt.Println("cannot change destination for order:", order_id, " it is delivering")
		return false
	}
}

func Process_deliverymade_from_W(aconn net.Conn, db *sqlx.DB, order_id int64){
	tx, _ :=db.Beginx()
	defer tx.Commit()
	Update_ostatus(db, tx, "delivered", order_id)
	x, y := Get_order_destination(db, tx, order_id)
	fmt.Println("order:", order_id, "delivered at (", x, ",", y, ")")
	//send order_id delivered to A
	//@@
	uacmd := &UAComm.UACommands{}
	addPkg(order_id, "delivered", uacmd)
	fmt.Println("Send to Amazon in deliverymade()")
	sendAmazon(aconn, uacmd)
}

func Process_finish_from_W(conn net.Conn, db *sqlx.DB, truck_id int32, x int32, y int32){
	tx, _ :=db.Beginx()
	defer tx.Commit()
	var truck TRUCK
	err := tx.Get(&truck, "select * from truck WHERE truck.truck_id = ? AND truck.t_status =? AND truck.x = ? AND truck.y = ?", truck_id, "to warehouse", x, y)
	if err==nil {
		//arrive warehouse
		fmt.Println("truck: ", truck_id, "arrive warehouse:", truck.Towhid)
		Update_tstatus(db, tx, "waiting for pick up", truck_id, 0, 0, 0)
		Get_unsentloading_orders_ontruck_and_send(conn, db, tx, truck_id)
	} else{
		//finish
		Update_tstatus(db, tx, "idle", truck_id, x, y, 0)
	}
}