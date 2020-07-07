package domain

import (
	"go-rbac-demo/domain/entity"
	"go-rbac-demo/repository"
	"testing"
)

func TestAdminService_CreateAdmin(t *testing.T) {

	admin := &entity.Admin{AdminName: "heyong1", AdminPassword: "123456", RoleCode: "Admin"}
	adminRepo := &repository.AdminRepo{}

	adminService := &AdminService{adminRepo}

	if err := adminService.CreateAdmin(admin); err != nil {
		t.Error(err.Error())
	} else {
		t.Log(admin)
	}
}
