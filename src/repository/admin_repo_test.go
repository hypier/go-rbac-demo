package repository

import (
	"fmt"
	"go-rbac-demo/domain/entity"
	"testing"
)

func TestAdminRepository_FindByName(t *testing.T) {
	adminRepo := &AdminRepository{}
	name, _ := adminRepo.FindByName("heyong7p21")

	fmt.Println(name)
}

func TestAdminRepository_Create(t *testing.T) {
	adminRepo := &AdminRepository{}
	role := entity.Role{RoleId: 1, RoleCode: "Admin", RoleName: "管理员"}
	admin := &entity.Admin{AdminName: "heyong7p21", AdminPassword: "123456", RoleCode: "Admin", Role: &role}

	_ = adminRepo.Create(admin)

	fmt.Println(admin)
}

func BenchmarkAdminRepository_FindByName(b *testing.B) {

	adminRepo := &AdminRepository{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = adminRepo.FindByName("admin")
	}
	b.StopTimer()
}
