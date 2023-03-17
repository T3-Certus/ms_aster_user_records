package models

import (
	"time"

	"gorm.io/gorm"
)

type UserData struct {
	ID_USER              uint64 `json:"id_user" gorm:"column:id_user; primaryKey"`
	ID_USER_ROL          uint64 `json:"id_user_rol" gorm:"column:id_user_rol"`
	USER_NAME            string `json:"user_name" gorm:"column:user_name"`
	USER_SURNAME         string `json:"user_surname" gorm:"column:user_surname"`
	USER_DOCUMENT_TYPE   string `json:"user_document_type" gorm:"column:user_document_type"`
	USER_DOCUMENT_NUMBER string `json:"user_document_number" gorm:"column:user_document_number"`
	USER_CELLPHONE       string `json:"user_cellphone" gorm:"column:user_cellphone"`
	USER_EMAIL           string `json:"user_email" gorm:"column:user_email"`
	// USER_PASSWORD        string    `gorm:"column:user_password"`

}

type UserOrderModel struct {
	ID_USER_ORDER        uint64                 `gorm:"column:id_user_order"`
	TOTAL_PRICE          uint32                 `gorm:"column:total_price"`
	TOTAL_PRICE_WDISC    uint32                 `gorm:"column:total_price_wdisc"`
	ORDER_DATE           time.Time              `gorm:"column:order_date"`
	DELIVERY_DATE        time.Time              `gorm:"column:delivery_date"`
	TOTAL_ITEMS_QUANTITY uint16                 `gorm:"column:total_items_quantity"`
	PRODUCTS             map[string]interface{} `gorm:"column:products"`
	SHIPPING_ADDRESS     map[string]interface{} `gorm:"column:shipping_address"`
	SHIPPING_PRICE       uint32                 `gorm:"column:shipping_price"`
	ID_TRANSACTION_PAY   string                 `gorm:"column:id_transaction_pay"`
}

type UserRoleModel struct {
	gorm.Model

	ID_USER_ROL   uint64 `gorm:"column:id_user_rol "`
	USER_ROL_NAME string `gorm:"column:user_rol_name"`
}

type OrderStateModel struct {
	gorm.Model

	ID_ORDER_STATE uint64 `gorm:"column:id_order_state"`
	ORDER_STATE    string `gorm:"column:order_state"`
}
