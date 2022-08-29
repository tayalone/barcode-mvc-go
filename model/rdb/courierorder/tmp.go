package courierorder

import (
	"time"
)

// CourierOrder create behavior of Courier Order
type CourierOrder interface {
	GetID() uint
	GetData() Tmp
	SetBarcode(string)
}

// Tmp is db schema tmp 'couier_order_CC_ISCOD'`
type Tmp struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Barcode   *string   `json:"barcode,omitempty"`
	CreatedAt time.Time `gorm:"index;autoCreateTime"`
	UpdatedAt time.Time `gorm:"index;autoUpdateTime" `
}

// GetTableName from Condition
func GetTableName(courierCode string, isCod bool) string {
	if courierCode == "DHL" {
		if isCod {
			return "courier_oder_dhl_cods"
		}
		return "courier_oder_dhls"
	}
	return ""
}

// GetTableStruct return emptyu Struct
func GetTableStruct(courierCode string, isCod bool) CourierOrder {
	if courierCode == "DHL" {
		if isCod {
			return &CourierOderDhlCod{}
		}
		return &CourierOderDhl{}
	}
	return nil
}
