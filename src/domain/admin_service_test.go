package domain

import (
	"go-rbac-demo/domain/entity"
	"go-rbac-demo/repository"
	"strconv"
	"testing"
)

func TestAdminService_CreateAdmin(t *testing.T) {

	admin := &entity.Admin{AdminName: "heyong7p1", AdminPassword: "123456", RoleCode: "Admin"}
	adminRepo := &repository.AdminRepository{}

	adminService := &AdminService{adminRepo}

	if err := adminService.CreateAdmin(admin); err != nil {
		t.Error(err.Error())
	} else {
		t.Log(admin)
	}
}

func TestAdminService_CheckAdmin(t *testing.T) {

	adminRepo := &repository.AdminRepository{}
	adminService := &AdminService{adminRepo}

	adminName, adminPassword := "admin", "admin"
	admin, err := adminService.CheckAdmin(adminName, adminPassword)

	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(admin)
	}
}

func TestAdminService_GetAdmin(t *testing.T) {
	adminRepo := &repository.AdminRepository{}
	adminService := &AdminService{adminRepo}

	adminName := "admin"
	admin, err := adminService.GetAdmin(adminName)

	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(admin)
	}
}

func BenchmarkAdminService_CheckAdmin(b *testing.B) {

	adminRepo := &repository.AdminRepository{}
	adminService := &AdminService{adminRepo}

	adminName, adminPassword := "admin", "admin"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = adminService.CheckAdmin(adminName, adminPassword)
	}
	b.StopTimer()
}

func BenchmarkAdminService_CreateAdmin(b *testing.B) {

	adminRepo := &repository.AdminRepository{}

	adminService := &AdminService{adminRepo}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		admin := &entity.Admin{AdminName: "heyonga" + strconv.Itoa(i), AdminPassword: "123456", RoleCode: "Admin"}
		_ = adminService.CreateAdmin(admin)
	}
	b.StopTimer()

}
