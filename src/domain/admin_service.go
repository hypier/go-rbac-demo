package domain

import (
	"go-rbac-demo/custerror"
	"go-rbac-demo/domain/entity"
	"go-rbac-demo/repository"
	"go-rbac-demo/util"
)

type AdminService struct {
	AdminRepo repository.AdminRepo
}

// 创建管理员
func (a *AdminService) CreateAdmin(admin *entity.Admin) error {

	if admin == nil || admin.AdminName == "" {
		return custerror.New("用户名不能为空")
	}

	if admin.AdminPassword == "" {
		return custerror.New("用户密码不能为空")
	}

	if admin.RoleCode == "" {
		return custerror.New("用户角色不能为空")
	}

	if b, err := a.AdminRepo.FindByName(admin.AdminName); err == nil && b.AdminId != 0 {
		return custerror.New("用户已经存在")
	}

	admin.AdminPassword = util.MD5(admin.AdminPassword)

	return a.AdminRepo.Create(admin)
}

// 登陆检查
func (a *AdminService) CheckAdmin(adminName string, adminPassword string) (*entity.Admin, error) {
	if adminName == "" {
		return nil, custerror.New("用户名不能为空")
	}

	if adminPassword == "" {
		return nil, custerror.New("用户密码不能为空")
	}

	encryptedPassword := util.MD5(adminPassword)

	dbAdmin, err := a.AdminRepo.FindByName(adminName)
	if err != nil {
		return nil, custerror.New("系统报错")
	}

	if dbAdmin == nil || dbAdmin.AdminId == 0 {
		return nil, custerror.New("没有找到此用户")
	}

	if dbAdmin.AdminPassword != encryptedPassword {
		return nil, custerror.New("用户密码不正确")
	}

	return dbAdmin, nil
}

// 管理员查询
func (a *AdminService) GetAdmin(adminName string) (*entity.Admin, error) {
	if adminName == "" {
		return nil, custerror.New("用户名不能为空")
	}

	return a.AdminRepo.FindByName(adminName)
}
