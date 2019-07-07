#!/bin/bash
BASH_PATH=$(dirname $0)

echo "Waiting for mysql to get up"
sleep 60
echo "Create MySQL servers (master / slave repl)"
echo "-----------"

echo "* Create replication user"

mysql --host mysqlslave -uroot -p$MySQL_SLAVE_PASSWORD -AN -e 'STOP SLAVE;';
mysql --host mysqlslave -uroot -p$MySQL_MASTER_PASSWORD -AN -e 'RESET SLAVE ALL;';

mysql --host mysqlmaster -uroot -p$MySQL_MASTER_PASSWORD -AN -e "CREATE USER '$MYSQL_REPLICATION_USER'@'%';";
mysql --host mysqlmaster -uroot -p$MySQL_MASTER_PASSWORD -AN -e "GRANT REPLICATION SLAVE ON *.* TO 
    '$MYSQL_REPLICATION_USER'@'%' IDENTIFIED BY '$MYSQL_REPLICATION_PASSWORD';";
mysql --host mysqlmaster -uroot -p$MySQL_MASTER_PASSWORD -AN -e 'flush privileges;';

echo "* Set MySQL01 as master on MySQL02"

MYSQL01_Position=$(eval "mysql --host mysqlmaster -uroot -p$MYSQL_MASTER_PASSWORD -e 'show master status \G' | grep Position | sed -n -e 's/^.*: //p'")
MYSQL01_File=$(eval "mysql --host mysqlmaster -uroot -p$MYSQL_MASTER_PASSWORD -e 'show master status \G' | grep File | sed -n -e 's/^.*: //p'")
MASTER_IP=$(eval "getent hosts mysqlmaster | awk '{pring \$1}'")
echo $MASTER_IP
mysql --host mysqlslave -uroot -p$MySQL_SLAVE_PASSWORD -AN -e "CHANGE MASTER TO master_host='mysqlmaster', master_port=3306, \
    master_user='$MYSQL_REPLICATION_USER', master_password='$MYSQL_REPLICATION_PASSWORD', master_log_file='$MYSQL01_File', \
    master_log_pos=$MYSQL_02_Position;"

echo "* Start Slave on both Servers"
mysql --host mysqlslave -uroot -p$MySQL_SLAVE_PASSWORD -AN -e 'star slave';

echo "Increase the max_connectons to 2000"
mysql --host mysqlmaster -uroot -p$MySQL_MASTER_PASSWORD -AN -e "set GLOBAL max_connections=2000";
mysql --host mysqlslave -uroot -p$MySQL_SLAVE_PASSWORD -AN -e "set GLOBAL max_connections=2000";

mysql --host mysqlslave -uroot -p$MySQL_SLAVE_PASSWORD -AN -e "show slave status \G";

echo "MySQL servers created!"
echo "--------------"
echo
echo Variables avaliable fo you :-
echo
echo MYSQL01_IP      : mysqlmaster
echo MYSQL02_IP      : mysqlslave