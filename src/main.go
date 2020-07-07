package main

import (
	"fmt"
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
	http.HandleFunc("/logout", adminCtrl.GetLogout)

	http.HandleFunc("/admin", adminCtrl.GetAdmin)
	http.HandleFunc("/user", adminCtrl.GetUser)

	fmt.Println("Start Server(8888)...")
	http.ListenAndServe(":8888", nil)
}
