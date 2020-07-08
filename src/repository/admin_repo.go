package repository

import (
	"database/sql"
	"go-rbac-demo/custerror"
	"go-rbac-demo/domain/entity"
)

var conn *sql.DB

func init() {
	conn, _ = connectMysql()
}

type AdminRepository struct {
}

func (a *AdminRepository) FindOne(id int) (admin *entity.Admin, err error) {

	panic("implement me")
}

func (a *AdminRepository) FindByName(adminName string) (admin *entity.Admin, err error) {

	var dbAdmin entity.Admin
	rows, err := conn.Query("select * from admin where admin_name = ? ", adminName)

	if err != nil {
		return nil, custerror.NewError(err)
	}

	defer func() {
		err2 := rows.Close()
		custerror.PrintError(err2)
	}()

	var roleCode sql.NullString

	for rows.Next() {
		if err := rows.Scan(&dbAdmin.AdminId, &dbAdmin.AdminName, &dbAdmin.AdminPassword, &roleCode, &dbAdmin.Role); err != nil {
			return nil, custerror.NewError(err)
		}

		dbAdmin.RoleCode = roleCode.String
	}

	if err = rows.Err(); err != nil {
		return nil, custerror.NewError(err)
	}

	return &dbAdmin, nil
}

func (a *AdminRepository) Create(admin *entity.Admin) error {

	res, err := conn.Exec("insert into admin(admin_name,admin_password,role_code,role)values(?,?,?,?)",
		admin.AdminName, admin.AdminPassword, admin.RoleCode, admin.Role)

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
