package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	errCount int
	m        = new(sync.RWMutex)
)

// DBInstance DB 連線實體
var DBInstance = setConnect()

// TimeMsg timer定時觸發
func TimeMsg() {

	d := time.Duration(time.Millisecond * 100)
	t := time.NewTicker(d)
	defer t.Stop()

	for {
		<-t.C
		showDBProcesslist()
	}
}

func main() {

	log.Println(runtime.Version())

	queryStr := `create table if not exists  badRecord66(
			id INT NOT NULL AUTO_INCREMENT,
			tutorial_title VARCHAR(100) NOT NULL,
			tutorial_author VARCHAR(40) NOT NULL,
			serial INT NOT NULL,
			PRIMARY KEY ( id )
		)ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1; `
	DBInstance.Exec(queryStr)

	go TimeMsg()

	var (
		wg sync.WaitGroup
	)
	// 總共一萬次
	for i := 0; i < 2000; i++ {

		// 每次起2000個gorutine去下語法
		for j := 0; j < 2000; j++ {
			wg.Add(1)
			go InserData(i+j, &wg)

		}

		wg.Wait()
		time.Sleep(50 * time.Millisecond)
		if i%10 == 0 {
			log.Println(i, "輪結束＝＝＝＝＝＝＝＝＝＝")
		}
	}
	log.Println("結束")
}

// InserData 寫資料進DB
func InserData(serial int, wg *sync.WaitGroup) {
	queryStr := `INSERT INTO badRecord66(tutorial_title,tutorial_author,serial) VALUES ( "小浣熊","好可愛",?)`
	_, err := DBInstance.Exec(queryStr, serial)
	if err != nil {
		log.Println("FATAL !!! serial:", serial, err.Error())
		m.Lock()
		errCount++
		m.Unlock()
	}

	wg.Done()
}

// Connection pool
func setConnect() *sql.DB {
	con := "root:qwe123@tcp(127.0.0.1:3306)/badbadtest"
	db, err := sql.Open("mysql", con)
	if err != nil {
		fmt.Println("DB Connect Error", err)
		os.Exit(3)
	}

	// Connection limit
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)

	db.SetConnMaxLifetime(30 * time.Second) // 3s
	// db.SetConnMaxLifetime(180) // 180 ns

	return db
}
