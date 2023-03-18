package associations_models

type UserData struct {
	ID_USER              uint64    `json:"id_user" gorm:"column:id_user; primaryKey"`
	ID_USER_ROL          uint64    `json:"-" gorm:"column:id_user_rol"`
	USER_NAME            string    `json:"user_name" gorm:"column:user_name"`
	USER_SURNAME         string    `json:"user_surname" gorm:"column:user_surname"`
	USER_DOCUMENT_TYPE   string    `json:"user_document_type" gorm:"column:user_document_type"`
	USER_DOCUMENT_NUMBER string    `json:"user_document_number" gorm:"column:user_document_number"`
	USER_CELLPHONE       string    `json:"user_cellphone" gorm:"column:user_cellphone"`
	USER_EMAIL           string    `json:"user_email" gorm:"column:user_email"`
	USER_ROLE            UserRoles `json:"user_role" gorm:"foreignKey:ID_USER_ROL;references:ID_USER_ROL"`
}

type UserRoles struct {
	ID_USER_ROL   uint64 `json:"id" gorm:"column:id_user_rol; primaryKey"`
	USER_ROL_NAME string `json:"name" gorm:"column:user_rol_name"`
}
