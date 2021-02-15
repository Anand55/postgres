package main

import (
	"database/sql"
	"fmt"

	"github.com/ghodss/yaml"
	_ "github.com/ghodss/yaml"
	_ "github.com/lib/pq"
)

func NewStoreImpl(jsonCongfig string) (DataStore, error) {
	config := DatabaseConfig{}
	err := yaml.Unmarshal([]byte(jsonCongfig), &config)
	if err != nil {
		return nil, err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.BaseUrl, config.Port, config.User, config.Password, config.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	svc := &DbStoreImpl{config: config, db: db}
	SetDataStore(svc)
	fmt.Println("Successfully connected")
	return svc, nil
}

type DbStoreImpl struct {
	config   DatabaseConfig
	employee Employee
	db       *sql.DB
}

func (s *DbStoreImpl) InsertEmployee(employee Employee) {
	insertStmt := `insert into "employ"("employ_name") values($1)`
	_, e := s.db.Exec(insertStmt, employee.Name)
	if e != nil {
		panic(e)
	}
}

func (s *DbStoreImpl) CloseDBConn() {
	s.db.Close()
}
