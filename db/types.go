package db

type List struct {
	ListItems []ListItem
}

type ListItem struct {
	Name    string `bson:"name"`
	Creator string `bson:"creator"`
}
