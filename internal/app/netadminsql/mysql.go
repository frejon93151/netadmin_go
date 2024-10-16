/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package netadminsql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func GetVLANRange(opt82 string) (vlan []int, err error) {
	cfg := mysql.Config{
		User:                 os.Getenv("NETADMIN__MYSQL_USER"),
		Passwd:               os.Getenv("NETADMIN__MYSQL_PASSWD"),
		Addr:                 os.Getenv("NETADMIN__MYSQL_ADDR"),
		Net:                  "tcp",
		DBName:               "netadmin",
		AllowNativePasswords: true,
	}
	dsn := cfg.FormatDSN()
	fmt.Println("cfg.FormatDSN() = ", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("sql.Open() = err", err.Error())
		return
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT SUBSTRING(fqpn, 7, 10) FROM netadmin.inv_physicalinterface WHERE fqpn = '%s' ORDER BY SUBSTRING(fqpn, 7, 10)", opt82)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("db.Query() = err", err.Error())
		return
	}

	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&vlan); err != nil {
			fmt.Println("rows.Next() = err", err.Error())
		}
	}

	return
}
