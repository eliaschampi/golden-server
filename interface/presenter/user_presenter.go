package presenter

import (
	"golden-server/domain/entity"
	"time"
)

type userPresenter struct {
	Dni       string    `json:"dni"`
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	RolCode   string    `json:"rol_code"`
	Gender    string    `json:"gender"`
	Image     string    `json:"image"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func MapUser(user entity.User) interface{} {
	return userPresenter{
		Dni:       user.Dni,
		Name:      user.Name,
		Lastname:  user.Lastname,
		RolCode:   CodeToString(user.RolCode),
		Gender:    user.Gender,
		Image:     user.Image.String,
		Address:   user.Address.String,
		Phone:     user.Phone.String,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
}
