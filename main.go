package main

import (
	"go-rbac-demo/controller"
	"go-rbac-demo/domain"
	"go-rbac-demo/repository"
	"log"
	"net/http"
)

func main() {
	log.Println("main")

	adminRepo := &repository.AdminRepo{}
	adminService := &domain.AdminService{AdminRepo: adminRepo}
	adminCtrl := &controller.AdminController{AdminService: adminService}

	http.HandleFunc("/admin/", adminCtrl.PostRegister)
	http.ListenAndServe(":8888", nil)
}
