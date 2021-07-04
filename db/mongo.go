package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	listCollection = "list"
	//tagCollection = "tag"
)

// DB is an interface to store and retrieve stateful data
type DB interface {
	// Adds an item to the List
	AddListItem(ctx context.Context, i ListItem) error

	// Deletes and item from the List
	DeleteListItem(ctx context.Context, i ListItem) (int64, error)

	// Retrieve all items on the List
	GetList(ctx context.Context) ([]ListItem, error)
}

// Mongo implements the DB interface
type Mongo struct {
	C           *mongo.Client
	dbName      string
}

// NewMongo creates a new Mongo object to be used by the bot
func NewMongo(connString string, dbName string) (*Mongo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	return &Mongo{
		C:           client,
		dbName:      dbName,
	}, nil
}

// AddListItem takes in a context and ListItem and returns an error if the item is not added to the list
func (m *Mongo) AddListItem(ctx context.Context, i ListItem) error {
	coll := m.C.Database(m.dbName).Collection(listCollection)
	_, err := coll.InsertOne(ctx, i)

	return err
}

// DeleteListItem takes in a context and ListItem, returns the number of items deleted and an error
func (m *Mongo) DeleteListItem(ctx context.Context, i ListItem) (int64, error) {
	var dCount int64

	coll := m.C.Database(m.dbName).Collection(listCollection)
	f := bson.M{"name": i.Name}

	del, err := coll.DeleteOne(ctx, f)
	if del != nil {
		dCount = del.DeletedCount
	}

	return dCount, err
}

// GetList takes a context and returns slice of all ListItems and an error
func (m *Mongo) GetList(ctx context.Context) ([]ListItem, error) {
	coll := m.C.Database(m.dbName).Collection(listCollection)

	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var list []ListItem

	for cursor.Next(ctx) {
		var item ListItem

		err = cursor.Decode(&item)
		if err != nil {
			return nil, err
		}

		list = append(list, item)
	}

	return list, nil
}
