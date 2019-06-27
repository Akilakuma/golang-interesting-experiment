package main

import (
	"log"
	"math/rand"
	"sync"
)

// UserMap 測試用Map
type UserMap map[int]int

var muLock sync.RWMutex

// Copy 複製一份新的UserMap
func (macMap *UserMap) Copy() UserMap {

	macMapDeReff := *macMap
	muLock.Lock()
	defer muLock.Unlock()

	newMap := make(UserMap, 0)
	for k, v := range macMapDeReff {
		newMap[k] = v
	}
	// return macMapDeReff
	return newMap
}

func main() {
	// 用assign的方式，把m指向n
	test1()
	// 用new一個新map的方式，把m的內容逐一複製到n
	test2()
	// 用assign的方式，把m指向，但是寫入map和range map前後加上lock
	test3()
}


// test1 用assign的方式，把m指向n
// 結果: fatal error: concurrent map iteration and map write
func test1() {
	m := make(UserMap)

	for i := 0; i < 6400000; i++ {
		key := rand.Intn(128)
		m[i] = key
	}
	n := m
	// n := m.Copy()
	go func() {
		for {
			key := rand.Intn(128)

			m[key] = key
		}
	}()

	for key, value := range n {
		log.Println(key, value)
	}
}

// test2 用new一個新map的方式，把m的內容逐一複製到n
// 結果: OK沒問題
func test2() {
	m := make(UserMap)

	for i := 0; i < 6400000; i++ {
		key := rand.Intn(128)
		m[i] = key
	}
	// n := m
	n := m.Copy()
	go func() {
		for {
			key := rand.Intn(128)
			m[key] = key
		}
	}()

	for key, value := range n {
		log.Println(key, value)
	}
}

// test3 用assign的方式，把m指向，但是寫入map和range map前後加上lock
// 結果: 不會爆炸，但是寫入會卡住range，range會卡住寫入，完全無法一起做事
func test3() {
	m := make(UserMap)

	for i := 0; i < 6400000; i++ {
		key := rand.Intn(128)
		m[i] = key
	}
	n := m
	// n := m.Copy()
	go func() {
		for {
			key := rand.Intn(128)
			muLock.Lock()
			m[key] = key
			muLock.Unlock()
			log.Println("set once")
		}
	}()
	muLock.Lock()
	for key, value := range n {
		log.Println(key, value)
	}
	muLock.Unlock()
}
