package repository

import (
	"go-rbac-demo/domain/entity"
)

type AdminRepository interface {
	FindOne(id int) (admin *entity.Admin, err error)
	FindByName(adminName string) (admin *entity.Admin, err error)
	Create(admin *entity.Admin) bool
	Update(admin *entity.Admin) bool
}

type RoleRepository interface {
	FindOne(id int) (role *entity.Role, err error)
	FindByCode(roleCode string) (role *entity.Role, err error)
	Create(role *entity.Role) bool
	Update(role *entity.Role) bool
}
