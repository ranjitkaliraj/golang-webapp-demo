package contacts

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	Street string `bson:"street"`
	City   string `bson:"city"`
	State  string `bson:"state"`
	Zip    string `bson:"zip"`
}

type Contact struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Email     string             `bson:"email"`
	Phone     string             `bson:"phone"`
	Address   Address            `bson:"address"`
}
