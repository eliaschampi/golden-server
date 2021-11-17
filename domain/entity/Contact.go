package entity

import (
	"context"
	"golden-server/domain/rules"
	"time"
)

type Contact struct {
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Email       *string   `json:"email"`
	Phone       *string   `json:"phone"`
	Address     *string   `json:"address"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

//ContactRepository
type ContactRepository interface {
	GetAll(c context.Context) ([]*Contact, error)
	Create(c context.Context, contact *rules.ContactStruct) (string, error)
}

//ContactRepository
type ContactService interface {
	GetAll(c context.Context) ([]*Contact, error)
	Create(c context.Context, contact *rules.ContactStruct) (string, []*rules.ErrorResponse, error)
}
