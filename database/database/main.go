package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	bytes, err := ioutil.ReadFile("database_config.yaml")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	dbStore, err := NewStoreImpl(string(bytes))
	if err != nil {
		panic(err)
	}
	emp := Employee{"mary"}
	dbStore.InsertEmployee(emp)

	dbStore.CloseDBConn()
}
