package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoConnect(ctx context.Context, uri string) (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(uri))
	// coll := client.Database("sample_mflix").Collection("movies")
	// title := "Back to the Future"
	// var result bson.M
	// err = coll.FindOne(ctx, bson.D{{"title", title}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the title %s\n", title)
	// 	return
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)
}

func insertCustomer(ctx context.Context, client *mongo.Client, customer Customer) error {
	coll := client.Database("crebito").Collection("customers")
	_, err := coll.InsertOne(ctx, customer)
	return err
}

func populateCustomers(ctx context.Context, client *mongo.Client) error {
	if err := insertCustomer(ctx, client, Customer{
		ID:      1,
		Limit:   100000,
		Balance: 0,
	}); err != nil {
		return err
	}
	if err := insertCustomer(ctx, client, Customer{
		ID:      2,
		Limit:   80000,
		Balance: 0,
	}); err != nil {
		return err
	}
	if err := insertCustomer(ctx, client, Customer{
		ID:      3,
		Limit:   1000000,
		Balance: 0,
	}); err != nil {
		return err
	}
	if err := insertCustomer(ctx, client, Customer{
		ID:      4,
		Limit:   10000000,
		Balance: 0,
	}); err != nil {
		return err
	}
	if err := insertCustomer(ctx, client, Customer{
		ID:      5,
		Limit:   500000,
		Balance: 0,
	}); err != nil {
		return err
	}
	return nil
}
