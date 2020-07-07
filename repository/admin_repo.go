package repository

import (
	"fmt"
	"go-rbac-demo/domain/entity"
)

func init() {

}

type AdminRepo struct {
}

func (a *AdminRepo) FindOne(id int) (admin *entity.Admin, err error) {

	panic("implement me")
}

func (a *AdminRepo) FindByName(adminName string) (admin *entity.Admin, err error) {
	var conn = connectMysql()
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var dbAdmin entity.Admin
	res, err := conn.Query("select * from admin where admin_name = ? ", adminName)

	if err != nil {
		return nil, err
	}

	for res.Next() {
		err := res.Scan(&dbAdmin.AdminId, &dbAdmin.AdminName, &dbAdmin.AdminPassword, &dbAdmin.RoleCode)
		if err != nil {
			return nil, err
		}
	}

	return &dbAdmin, nil
}

func (a *AdminRepo) Create(admin *entity.Admin) bool {
	var conn = connectMysql()
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	res, err := conn.Exec("insert into admin(admin_name,admin_password)values(?,?)", admin.AdminName, admin.AdminPassword)

	if err != nil {
		return false
	}

	id, err := res.LastInsertId()

	if err == nil {
		admin.AdminId = int(id)
		return true
	} else {
		return false
	}
}

func (a *AdminRepo) Update(admin *entity.Admin) error {
	panic("implement me")
}
