package main

import (
	"gopkg.in/mgo.v2"
)

func mongo(address, db string) (*mgo.Database, error) {
	session, err := mgo.Dial(address)
	if err != nil {
		return nil, err
	}

	return session.DB(db), nil
}
