package domain

import (
	"errors"
	"go-rbac-demo/domain/entity"
	"go-rbac-demo/repository"
	"go-rbac-demo/util"
)

type AdminService struct {
	AdminRepo repository.AdminRepository
}

func (a *AdminService) CreateAdmin(admin *entity.Admin) error {

	if admin == nil || admin.AdminName == "" {
		return errors.New("用户名不能为空")
	}

	if admin.AdminPassword == "" {
		return errors.New("用户密码不能为空")
	}

	if admin.RoleCode == "" {
		return errors.New("用户角色不能为空")
	}

	if b, err := a.AdminRepo.FindByName(admin.AdminName); b != nil && err == nil {
		return errors.New("用户已经存在")
	}

	admin.AdminPassword = util.MD5(admin.AdminPassword)

	return a.AdminRepo.Create(admin)
}

// 登陆检查
func (a *AdminService) CheckAdmin(adminName string, adminPassword string) (*entity.Admin, error) {
	if adminName == "" {
		return nil, errors.New("用户名不能为空")
	}

	if adminPassword == "" {
		return nil, errors.New("用户密码不能为空")
	}

	encryptedPassword := util.MD5(adminPassword)

	dbAdmin, err := a.AdminRepo.FindByName(adminName)
	if err != nil {
		return nil, errors.New("系统报错")
	}

	if dbAdmin == nil {
		return nil, errors.New("没有找到此用户")
	}

	if dbAdmin.AdminPassword != encryptedPassword {
		return nil, errors.New("用户密码不正确")
	}

	return dbAdmin, nil
}
