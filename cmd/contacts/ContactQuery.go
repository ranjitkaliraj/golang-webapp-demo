package contacts

import (
	"context"
)

type ContactQuery struct {
	repository *Repository
}

func NewContactQuery(repository *Repository) *ContactQuery {
	return &ContactQuery{repository: repository}
}

func (query *ContactQuery) GetContact(ctx context.Context, id string) (*Contact, error) {
	return query.repository.FindById(ctx, id)
}

func (query *ContactQuery) GetContacts(ctx context.Context) ([]*Contact, error) {
	return query.repository.Find(ctx)
}
