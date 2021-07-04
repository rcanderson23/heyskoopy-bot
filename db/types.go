package db

// ListItem contains information necessary for adding and deleting items from 'The List'
type ListItem struct {
	Name    string `bson:"name"`
	Creator string `bson:"creator"`
}
