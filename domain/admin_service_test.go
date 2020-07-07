package domain

import (
	"fmt"
	"go-rbac-demo/domain/entity"
	"go-rbac-demo/repository"
	"testing"
)

func TestAdminService_CreateAdmin(t *testing.T) {

	admin := &entity.Admin{AdminName: "heyong", AdminPassword: "123456", RoleCode: "Admin"}
	adminRepo := &repository.AdminRepo{}

	adminService := &AdminService{AdminRepo: adminRepo}

	if err := adminService.CreateAdmin(admin); err != nil {
		fmt.Println(admin)
	}
}
