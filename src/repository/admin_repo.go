package repository

import (
	"go-rbac-demo/custerror"
	"go-rbac-demo/domain/entity"
)

type AdminRepository struct {
}

func (a *AdminRepository) FindOne(id int) (admin *entity.Admin, err error) {

	panic("implement me")
}

func (a *AdminRepository) FindByName(adminName string) (admin *entity.Admin, err error) {
	conn, err := connectMysql()
	if err != nil {
		return nil, err
	}
	defer func() {
		err := conn.Close()
		custerror.PrintError(err)
	}()

	var dbAdmin entity.Admin
	rows, err := conn.Query("select * from admin where admin_name = ? ", adminName)

	if err != nil {
		return nil, custerror.NewError(err)
	}

	defer func() {
		err2 := rows.Close()
		custerror.PrintError(err2)
	}()

	for rows.Next() {
		if err := rows.Scan(&dbAdmin.AdminId, &dbAdmin.AdminName, &dbAdmin.AdminPassword, &dbAdmin.RoleCode); err != nil {
			return nil, custerror.NewError(err)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, custerror.NewError(err)
	}

	return &dbAdmin, nil
}

func (a *AdminRepository) Create(admin *entity.Admin) error {
	conn, err := connectMysql()
	if err != nil {
		return err
	}
	defer func() {
		err := conn.Close()
		custerror.PrintError(err)
	}()

	res, err := conn.Exec("insert into admin(admin_name,admin_password,role_code)values(?,?,?)",
		admin.AdminName, admin.AdminPassword, admin.RoleCode)

	if err != nil {
		return custerror.NewError(err)
	}

	if id, err := res.LastInsertId(); err != nil {
		return custerror.NewError(err)
	} else {
		admin.AdminId = int(id)
		return nil
	}
}

func (a *AdminRepository) Update(admin *entity.Admin) error {
	panic("implement me")
}
