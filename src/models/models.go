package models

import (
	"time"
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
}

type UserOrders struct {
	ID_USER_ORDER        uint64    `json:"id_user_order" gorm:"column:id_user_order; primaryKey"`
	ID_USER              uint64    `json:"id_user" gorm:"column:id_user"`
	ID_ORDER_STATE       uint64    `json:"id_order_state" gorm:"column:id_order_state"`
	TOTAL_PRICE          uint32    `json:"total_price" gorm:"column:total_price"`
	TOTAL_PRICE_WDISC    uint32    `json:"total_price_wdisc" gorm:"column:total_price_wdisc"`
	ORDER_DATE           time.Time `json:"order_date" gorm:"column:order_date"`
	DELIVERY_DATE        time.Time `json:"delivery_date" gorm:"column:delivery_date"`
	TOTAL_ITEMS_QUANTITY uint16    `json:"total_items_quantity" gorm:"column:total_items_quantity"`
	// PRODUCTS             map[string]interface{} `json:"products" gorm:"column:products"`
	// SHIPPING_ADDRESS     map[string]interface{} `json:"shipping_address" gorm:"column:shipping_address"`
	SHIPPING_PRICE     uint32 `json:"shipping_price" gorm:"column:shipping_price"`
	ID_TRANSACTION_PAY string `json:"id_transaction_pay" gorm:"column:id_transaction_pay"`
}

// type ProductsMap map[string]interface{}

// func (p *ProductsMap) Scan(value interface{}) error {
// 	if value == nil {
// 		*p = nil
// 		return nil
// 	}

// 	data, ok := value.([]byte)
// 	if !ok {
// 		return fmt.Errorf("failed to unmarshal JSON value: %v", value)
// 	}
// }

type UserRoles struct {
	ID_USER_ROL   uint64 `json:"id_user_rol" gorm:"column:id_user_rol; primaryKey"`
	USER_ROL_NAME string `json:"user_rol_name" gorm:"column:user_rol_name"`
}

type OrderStates struct {
	ID_ORDER_STATE uint64 `gorm:"column:id_order_state; primaryKey"`
	ORDER_STATE    string `gorm:"column:order_state"`
}
