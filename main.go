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

	http.HandleFunc("/admin/", adminCtrl.PostRegister)
	http.HandleFunc("/admin1/", Post)
	http.ListenAndServe(":8888", nil)
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	w.Write([]byte(r.Method))
}
