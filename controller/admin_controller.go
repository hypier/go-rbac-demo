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

	adminName := r.FormValue("name")
	adminPassword := r.FormValue("password")
	roleCode := r.FormValue("role")

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

// 管理员登陆
func (a *AdminController) PostLogin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	if err := r.ParseForm(); err != nil {
		OutputJson(w, 0, err.Error(), nil)
		return
	}

	adminName := r.FormValue("name")
	adminPassword := r.FormValue("password")

	if adminName == "" || adminPassword == "" {
		OutputJson(w, 0, "参数错误", nil)
		return
	}

	admin, err := a.AdminService.CheckAdmin(adminName, adminPassword)
	if err != nil {
		OutputJson(w, 0, err.Error(), nil)
		return
	}

	addCookie("admin_name", admin.AdminName, w)

	OutputJson(w, 1, "登陆成功", admin.AdminName)
}

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &Result{ret, reason, i}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	_, _ = w.Write(b)
}

// 存入cookie,使用cookie存储
func addCookie(name string, value string, w http.ResponseWriter) {

	cookie := http.Cookie{Name: name, Value: value, Path: "/"}
	http.SetCookie(w, &cookie)
}
