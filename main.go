package main

import (
	"go-rbac-demo/controller"
	"go-rbac-demo/domain"
	"go-rbac-demo/repository"
	"net/http"
)

func main() {

	adminRepo := &repository.AdminRepo{}
	adminService := &domain.AdminService{AdminRepo: adminRepo}
	adminCtrl := &controller.AdminController{AdminService: adminService}

	http.HandleFunc("/admin", adminCtrl.PostRegister)
	http.HandleFunc("/login", adminCtrl.PostLogin)
	http.ListenAndServe(":8888", nil)
}
