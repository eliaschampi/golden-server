package presenter

import "golden-server/domain/entity"

type rolPresenter struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Description string `json:"description"`
}

func MapRol(rol domainentity.Rol) interface{} {
	return rolPresenter{
		Code:        CodeToString(rol.Code),
		Name:        rol.Name,
		Description: rol.Description,
	}
}


