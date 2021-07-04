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

type DB interface {
	AddListItem(ctx context.Context, i ListItem) error
	DeleteListItem(ctx context.Context, i ListItem) (int64, error)
	GetList(ctx context.Context) ([]ListItem, error)
}

type Mongo struct {
	C           *mongo.Client
	dbName      string
	collections []string
}

func NewMongo(connString string, dbName string, collections []string) (*Mongo, error) {
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
		collections: collections,
	}, nil
}

func (m *Mongo) AddListItem(ctx context.Context, i ListItem) error {
	coll := m.C.Database(m.dbName).Collection(listCollection)
	_, err := coll.InsertOne(ctx, i)

	return err
}

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
