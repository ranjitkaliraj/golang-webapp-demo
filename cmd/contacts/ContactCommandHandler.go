package contacts

import (
	"context"
)

type AddContactCommand struct {
	Contact Contact
}

type UpdateContactCommand struct {
	Id      string
	Contact Contact
}

type DeleteContactCommand struct {
	ContactId string
}

type ContactCommandHandler struct {
	repository *Repository
}

func NewContactCommandHandler(repository *Repository) *ContactCommandHandler {
	return &ContactCommandHandler{repository: repository}
}

func (handler *ContactCommandHandler) HandleCreateCommand(ctx context.Context, command *AddContactCommand) (*Contact, error) {
	return handler.repository.Create(ctx, &command.Contact)
}

func (handler *ContactCommandHandler) HandleUpdateCommand(ctx context.Context, command *UpdateContactCommand) (*Contact, error) {
	return handler.repository.Update(ctx, command.Id, &command.Contact)
}

func (handler *ContactCommandHandler) HandleDeleteCommand(ctx context.Context, command *DeleteContactCommand) error {
	return handler.repository.Delete(ctx, command.ContactId)
}
