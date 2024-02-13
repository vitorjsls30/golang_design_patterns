// Singleton Pattern
// The Singleton Pattern allow us to create only one instance o a given struct throughout the whole Application.
// It's imoportant to note that the instance must be globally accessible through the whole application.

// Why?
// It is very usefull when we need to handle operations that must preferably be executed once, like: Database creation,
// global configurations, caches and so on.

// How?
// Basically we create a struct that represents ou Singleton only once via a constructor method that garanties that
// only one item will be created and returned if it already exists.

package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

// Our singleton struct
type DataBaseHandler struct {
}

var dbHandler *DataBaseHandler

func (d *DataBaseHandler) query(sql string) {
	fmt.Printf("[DatabaseHandler/query] %s\n", sql)
}

func GetDBInstance() *DataBaseHandler {
	// 1 - check to return in case it already exists...
	if dbHandler == nil {
		// 2 - lock the creation process to avoid that a concurrent goroutine is also entering the creatin process...
		lock.Lock()
		defer lock.Unlock()
		// 3 - check again in case any goroutine bypassed the first check...
		if dbHandler == nil {
			fmt.Println("[DatabaseHandler/GetInstance] creating instance...")
			dbHandler = &DataBaseHandler{}
		} else {
			fmt.Println("[DatabaseHandler/GetInstance] [2] returning already created instance...")
		}
	} else {
		fmt.Println("[DatabaseHandler/GetInstance] [1] returning already created instance...")
	}
	return dbHandler
}

func main() {
	fmt.Println("[main] trying to get a DB instance...")
	for i := 0; i < 30; i++ {
		go GetDBInstance()
	}
	fmt.Scanln()
}
