package database

import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

type MgoDatabase struct{}

func (self MgoDatabase) DialDatabase() Session {
	var err error
	var mgoSession *mgo.Session

	info := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "gobro",
		Timeout:  60 * time.Second}

	mgoSession, err = mgo.DialWithInfo(info)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	return MgoSession{mgoSession}
}
