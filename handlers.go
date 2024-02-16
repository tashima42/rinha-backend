package main

type Customer struct {
	ID      int `bson:"id"`
	Limit   int `bson:"limit"`
	Balance int `bson:"balance"`
}
