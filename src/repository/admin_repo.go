package repository

import (
	"go-rbac-demo/domain/entity"
	"log"
)

type AdminRepository struct {
}

func (a *AdminRepository) FindOne(id int) (admin *entity.Admin, err error) {

	panic("implement me")
}

func (a *AdminRepository) FindByName(adminName string) (admin *entity.Admin, err error) {
	var conn = connectMysql()
	defer func() {
		err := conn.Close()
		checkErr(err)
	}()

	var dbAdmin entity.Admin
	rows, err := conn.Query("select * from admin where admin_name = ? ", adminName)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func() {
		err2 := rows.Close()
		checkErr(err2)
	}()

	for rows.Next() {
		err := rows.Scan(&dbAdmin.AdminId, &dbAdmin.AdminName, &dbAdmin.AdminPassword, &dbAdmin.RoleCode)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &dbAdmin, nil
}

func (a *AdminRepository) Create(admin *entity.Admin) bool {
	var conn = connectMysql()
	defer func() {
		err := conn.Close()
		checkErr(err)
	}()

	res, err := conn.Exec("insert into admin(admin_name,admin_password,role_code)values(?,?,?)",
		admin.AdminName, admin.AdminPassword, admin.RoleCode)

	if err != nil {
		log.Fatal(err)
		return false
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
		return false
	} else {
		admin.AdminId = int(id)
		return true
	}
}

func (a *AdminRepository) Update(admin *entity.Admin) bool {
	panic("implement me")
}
