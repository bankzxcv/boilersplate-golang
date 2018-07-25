package db

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var (
	s *mgo.Session
)

func InitDb() {
	// @TODO Add Config for Mongodb
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	if s != nil {
		fmt.Println("Connect MONGDB Success")
	}
}

func GetSession() *mgo.Session {
	return s
}
