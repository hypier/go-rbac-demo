package entity

type Admin struct {
	//BaseEntity
	AdminId       int    `db:admin_id`
	AdminName     string `admin_name`
	AdminPassword string `admin_password`
	RoleCode      string `role_code`
}
