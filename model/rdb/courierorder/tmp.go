package courierorder

import (
	"time"
)

// Tmp is db schema tmp 'couier_order_CC_ISCOD'`
type Tmp struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Barcode   *string   `json:"barcode,omitempty"`
	CreatedAt time.Time `gorm:"index;autoCreateTime"`
	UpdatedAt time.Time `gorm:"index;autoUpdateTime" `
}

// GetTableName from Condition
func GetTableName(courierName string, isCod bool) string {
	if courierName == "DHL" {
		if isCod {
			return "courier_oder_dhl_cods"
		}
		return "courier_oder_dhls"
	}
	return ""
}
