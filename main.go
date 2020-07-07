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

	http.HandleFunc("/reg", adminCtrl.PostRegister)

	http.HandleFunc("/login", adminCtrl.PostLogin)
	http.HandleFunc("/logout", adminCtrl.PostLogout)

	http.HandleFunc("/admin", adminCtrl.GetAdmin)
	http.HandleFunc("/user", adminCtrl.GetUser)

	http.ListenAndServe(":8888", nil)
}
