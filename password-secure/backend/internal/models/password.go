package models

type Password struct {
	HashedPassword string `bson:"hashed_password"`
}