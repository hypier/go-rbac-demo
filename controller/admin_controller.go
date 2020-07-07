package controller

import (
	"encoding/json"
	"fmt"
	"go-rbac-demo/domain"
	"go-rbac-demo/domain/entity"
	"net/http"
)

type AdminController struct {
	AdminService *domain.AdminService
}

type Result struct {
	Ret    int
	Reason string
	Data   interface{}
}

// 管理员注册
func (a *AdminController) PostRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	w.Header().Set("content-type", "application/json")

	if err := r.ParseForm(); err != nil {
		OutputJson(w, 0, err.Error(), nil)
		return
	}

	adminName := r.FormValue("adminName")
	adminPassword := r.FormValue("adminPassword")
	roleCode := r.FormValue("roleCode")

	if adminName == "" || adminPassword == "" || roleCode == "" {
		OutputJson(w, 0, "参数错误", nil)
		return
	}

	admin := &entity.Admin{
		AdminName:     adminName,
		AdminPassword: adminPassword,
		RoleCode:      roleCode,
	}

	if err := a.AdminService.CreateAdmin(admin); err != nil {
		OutputJson(w, 0, err.Error(), nil)
	} else {
		OutputJson(w, 1, "ok", admin)
	}

}

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &Result{ret, reason, i}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	_, _ = w.Write(b)
}
