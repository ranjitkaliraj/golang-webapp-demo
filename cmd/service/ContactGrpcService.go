package service

import (
	"context"
	"log"

	"github.com/ranjitkaliraj/go-test-crm-service/cmd/contacts"
	pb "github.com/ranjitkaliraj/go-test-crm-service/cmd/grpc"
)

type ContactGrpcService struct {
	commandHandler *contacts.ContactCommandHandler
	query          *contacts.ContactQuery
	pb.UnimplementedContactServiceServer
}

func NewContactGrpcService() *ContactGrpcService {
	repository := contacts.NewContactRepository()
	return &ContactGrpcService{commandHandler: contacts.NewContactCommandHandler(repository), query: contacts.NewContactQuery(repository)}
}

func (service *ContactGrpcService) CreateContact(ctx context.Context, req *pb.Contact) (*pb.ContactResponse, error) {
	command := contacts.AddContactCommand{
		Contact: contacts.Contact{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
			Address: contacts.Address{
				Street: req.Address.Street,
				City:   req.Address.City,
				State:  req.Address.State,
				Zip:    req.Address.Zip,
			},
		},
	}
	contact, err := service.commandHandler.HandleCreateCommand(ctx, &command)
	if err != nil {
		log.Printf("Error creating contact: %v", err)
		return nil, err
	}
	return &pb.ContactResponse{Contact: toContactProto(*contact)}, nil
}

func (service *ContactGrpcService) UpdateContact(ctx context.Context, req *pb.Contact) (*pb.ContactResponse, error) {
	command := contacts.UpdateContactCommand{
		Id: req.Id,
		Contact: contacts.Contact{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
			Address: contacts.Address{
				Street: req.Address.Street,
				City:   req.Address.City,
				State:  req.Address.State,
				Zip:    req.Address.Zip,
			},
		},
	}
	contact, err := service.commandHandler.HandleUpdateCommand(ctx, &command)
	if err != nil {
		log.Printf("Error updating contact: %v", err)
		return nil, err
	}
	return &pb.ContactResponse{Contact: toContactProto(*contact)}, nil
}

func (service *ContactGrpcService) DeleteContact(ctx context.Context, req *pb.ContactId) (*pb.Empty, error) {
	command := contacts.DeleteContactCommand{ContactId: req.Id}
	err := service.commandHandler.HandleDeleteCommand(ctx, &command)
	if err != nil {
		log.Printf("Error deleting contact: %v", err)
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (service *ContactGrpcService) GetContact(ctx context.Context, req *pb.ContactId) (*pb.ContactResponse, error) {
	contact, err := service.query.GetContact(ctx, req.Id)
	if err != nil {
		log.Printf("Error getting contact: %v", err)
		return nil, err
	}
	return &pb.ContactResponse{Contact: toContactProto(*contact)}, nil
}

func (service *ContactGrpcService) ListContacts(ctx context.Context, req *pb.Empty) (*pb.ContactList, error) {
	contacts, err := service.query.GetContacts(ctx)
	if err != nil {
		log.Printf("Error listing contacts: %v", err)
		return nil, err
	}
	var protoContacts []*pb.Contact
	for _, contact := range contacts {
		protoContacts = append(protoContacts, toContactProto(*contact))
	}
	return &pb.ContactList{Contacts: protoContacts}, nil
}

func toContactProto(c contacts.Contact) *pb.Contact {
	return &pb.Contact{
		Id:        c.ID.Hex(),
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
		Address:   toAddressProto(c.Address),
	}
}

func toAddressProto(c contacts.Address) *pb.Address {
	return &pb.Address{
		Street: c.Street,
		City:   c.City,
		State:  c.State,
		Zip:    c.Zip,
	}
}
