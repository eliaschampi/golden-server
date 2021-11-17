package usecase

import (
	"context"
	"golden-server/domain/entity"
)

type contactService struct {
	contactRepo entity.ContactRepository
}

func NewContactService(repo *entity.ContactRepository) entity.ContactService {
	return &contactService{*repo}
}

func (ct *contactService) GetAll(c context.Context) ([]*entity.Contact, error) {
	return ct.contactRepo.GetAll(c)
}
