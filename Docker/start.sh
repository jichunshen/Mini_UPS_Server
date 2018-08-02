#!/bin/bash
PSQLPORT=5432
DB="packagesim"
/docker-entrypoint.sh postgres &
echo "---Checking for postgres to be ready..."

dbexists=`psql -lt`
ans="$?"
while [ "$ans" != "0" ]
do
    echo "---Waiting for postgres..."
    sleep 3
    dbexists=`psql -lt`
    ans="$?"
done
echo "---Postgres is ready..."
dbexists=`echo $dbexists | cut -f1 -d'|' | grep ^$DB\$$`

if [ "$dbexists" == "" ]
then

    echo "Create DB..."
    # create db
    createdb $DB
    psql $DB < createtables.sql
fi
echo "---Starting World Sim..."
./sim
echo "---World Sim Exited ($?)..."
