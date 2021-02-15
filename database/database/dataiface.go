package main

type DataStore interface {
	InsertEmployee(employee Employee)
	CloseDBConn()
}

var datastore DataStore

func SetDataStore(data DataStore) {
	datastore = data
}

func GetDataStore() DataStore {
	return datastore
}
