package main

import (
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"errors"
)

type Auth struct {
	DB     *mgo.Database
	Secret string
}

func (auth Auth) User(token string) (*User, error) {
	if len(token) == 0 {
		return nil, errors.New("Empty token.")
	}

	signed, err := jwt.Parse(token, func(passed_token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte(auth.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims := signed.Claims.(jwt.MapClaims)
	sid := claims["user_id"].(string)
	oid := bson.ObjectIdHex(sid)

	var user User
	err = auth.DB.C("users").FindId(oid).One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type User struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserName     string        `bson:"username" json:"username"`
	Image        string        `bson:"image" json:"image,omitempty"`
}