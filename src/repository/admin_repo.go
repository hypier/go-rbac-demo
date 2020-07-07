package repository

import (
	"fmt"
	"go-rbac-demo/domain/entity"
)

type AdminRepo struct {
}

func (a *AdminRepo) FindOne(id int) (admin *entity.Admin, err error) {

	panic("implement me")
}

func (a *AdminRepo) FindByName(adminName string) (admin *entity.Admin, err error) {
	var conn = connectMysql()
	defer func() {
		err := conn.Close()
		checkErr(err)
	}()

	var dbAdmin entity.Admin
	res, err := conn.Query("select * from admin where admin_name = ? ", adminName)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for res.Next() {
		err := res.Scan(&dbAdmin.AdminId, &dbAdmin.AdminName, &dbAdmin.AdminPassword, &dbAdmin.RoleCode)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}

	return &dbAdmin, nil
}

func (a *AdminRepo) Create(admin *entity.Admin) bool {
	var conn = connectMysql()
	defer func() {
		err := conn.Close()
		checkErr(err)
	}()

	res, err := conn.Exec("insert into admin(admin_name,admin_password,role_code)values(?,?,?)",
		admin.AdminName, admin.AdminPassword, admin.RoleCode)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		admin.AdminId = int(id)
		return true
	}
}

func (a *AdminRepo) Update(admin *entity.Admin) bool {
	panic("implement me")
}
