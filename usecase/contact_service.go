package usecase

import (
	"context"
	"golden-server/domain/entity"
	"golden-server/domain/rules"
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

func (ct *contactService) Create(c context.Context, contact *rules.ContactStruct) (string, []*rules.ErrorResponse, error) {

	validatedErr := rules.ValidateStruct(contact)

	if validatedErr != nil {
		return "", validatedErr, nil
	}
	code, err := ct.contactRepo.Create(c, contact)
	if err != nil {
		return "", nil, err
	}
	return code, nil, nil

}
