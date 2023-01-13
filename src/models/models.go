package models

type UserDataModel struct {
	ID_USER              uint64 `gorm:"column:id_user; primaryKey"`
	USER_NAME            string `gorm:"column:user_name"`
	USER_SURNAME         string `gorm:"column:user_surname"`
	USER_DOCUMENT_TYPE   string `gorm:"column:user_document_type"`
	USER_DOCUMENT_NUMBER string `gorm:"column:user_document_number"`
	USER_CELLPHONE       string `gorm:"column:user_cellphone"`
	USER_EMAIL           string `gorm:"column:user_email"`
	USER_PASSWORD        string `gorm:"column:user_password"`
}
