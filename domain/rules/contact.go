package rules

type ContactStruct struct {
	Name        string `validate:"required,lte=255"`
	Type        string `validate:"required"`
	Email       string `validate:"omitempty,email"`
	Phone       string `validate:"lte=30"`
	Address     string `validate:"lte=120"`
	Description string `validate:"lte=300"`
}
