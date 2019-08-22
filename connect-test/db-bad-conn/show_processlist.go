package main

import (
	"fmt"
	"log"
)

type DBConnectionInfo struct {
	ID    []byte
	User  string
	Host  string
	DB    []byte
	Cmd   string
	Time  []byte
	State string
	Info  []byte
}

func (f *DBConnectionInfo) String() string {

	return fmt.Sprintf("ID:%s, User:%s, Host:%s, DB:%s, Cmd:%s, Time:%s, State:%s, Info:%s", string(f.ID), f.User, f.Host, string(f.DB), f.Cmd, string(f.Time), f.State, string(f.Info))
}

func showDBProcesslist() {
	queryStr := `SHOW processlist`
	rows, err := DBInstance.Query(queryStr)
	if err != nil {
		log.Println(err)
	}
	var count int
	for rows.Next() {
		// data := &DBConnectionInfo{}
		// if err := rows.Scan(&data.ID, &data.User, &data.Host, &data.DB, &data.Cmd, &data.Time, &data.State, &data.Info); err != nil {
		// 	log.Println("Scan failed:", err)
		// }

		// fmt.Println("data:", data)
		count++
	}
	log.Println("DB show processlist 數量有：", count)
}
