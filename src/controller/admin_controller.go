package controller

import (
	"encoding/json"
	"go-rbac-demo/custerror"
	"go-rbac-demo/domain"
	"go-rbac-demo/domain/entity"
	"go-rbac-demo/util"
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
		custerror.PrintError(err)
		OutputJson(w, 0, err.Error(), nil)
	} else {
		OutputJson(w, 1, "注册成功", admin)
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

	util.AddCookie("admin_name", admin.AdminName, w)

	OutputJson(w, 1, "登陆成功", admin.AdminName)
}

// 管理员退出
func (a *AdminController) GetLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if util.ClearCookie("admin_name", r, w) {
		OutputJson(w, 1, "退出成功", nil)
	} else {
		// 返回状态码
		w.WriteHeader(401)
		return
	}

}

// 管理员页面
func (a *AdminController) GetAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	admin, ok := a.auth(w, r)
	if !ok {
		return
	}

	if admin.RoleCode == "Admin" {
		OutputJson(w, 1, "ok", admin)
	} else {
		w.WriteHeader(403)
	}

}

// 普通页面
func (a *AdminController) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if admin, ok := a.auth(w, r); !ok {
		return
	} else {
		OutputJson(w, 1, "ok", admin)
	}
}

// 登陆验证
func (a *AdminController) auth(w http.ResponseWriter, r *http.Request) (*entity.Admin, bool) {
	adminName, err := util.GetCookie("admin_name", r)
	if err != nil {
		// 返回状态码
		w.WriteHeader(401)
		return nil, false
	}

	admin, err := a.AdminService.GetAdmin(adminName)
	if err != nil {
		// 返回错误
		w.WriteHeader(401)
		return nil, false
	}

	return admin, true
}

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &Result{ret, reason, i}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	_, _ = w.Write(b)
}
