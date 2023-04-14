package contacts

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	collection *mongo.Collection
}

func NewContactRepository() *Repository {
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	//defer client.Disconnect(context.Background())
	db := client.Database("test")
	return &Repository{collection: db.Collection("contacts")}
}

func (cs *Repository) Create(ctx context.Context, contact *Contact) (*Contact, error) {
	log.Printf("Creating contact...")
	res, err := cs.collection.InsertOne(ctx, contact)
	if err != nil {
		log.Printf("Failed to create contact: %v\n", err)
		return nil, err
	}
	contact.ID = res.InsertedID.(primitive.ObjectID)
	return contact, nil
}

func (cs *Repository) Update(ctx context.Context, id string, update *Contact) (*Contact, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	updateDoc := bson.M{"$set": bson.M{
		"first_name": update.FirstName,
		"last_name":  update.LastName,
		"email":      update.Email,
		"phone":      update.Phone,
		"address": bson.M{
			"street": update.Address.Street,
			"city":   update.Address.City,
			"state":  update.Address.State,
			"zip":    update.Address.Zip,
		},
	}}

	var contact Contact
	err = cs.collection.FindOneAndUpdate(ctx, filter, updateDoc, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&contact)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Contact not found with ID %s\n", id)
			return nil, fmt.Errorf("contact not found")
		}
		log.Printf("Failed to update contact: %v\n", err)
		return nil, err
	}
	return &contact, nil
}

func (s *Repository) Delete(ctx context.Context, id string) error {
	objectID, err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		log.Printf("Invalid ID: %v\n", err1)
		return nil
	}
	filter := bson.M{"_id": objectID}
	_, err := s.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Error deleting contact: %v", err)
		return err
	}
	return nil
}

func (cs *Repository) FindById(ctx context.Context, id string) (*Contact, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	var contact Contact
	err = cs.collection.FindOne(ctx, filter).Decode(&contact)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Contact not found with ID %s\n", id)
			return nil, fmt.Errorf("contact not found")
		}
		log.Printf("Failed to get contact: %v\n", err)
		return nil, err
	}
	return &contact, nil
}

func (s *Repository) Find(ctx context.Context) ([]*Contact, error) {
	var results []*Contact
	cur, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error listing contacts: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var contact Contact
		err := cur.Decode(&contact)
		if err != nil {
			log.Printf("Error decoding contact: %v", err)
			return nil, err
		}
		results = append(results, &contact)
	}
	if err := cur.Err(); err != nil {
		log.Printf("Error iterating over contacts: %v", err)
		return nil, err
	}
	return results, nil
}
