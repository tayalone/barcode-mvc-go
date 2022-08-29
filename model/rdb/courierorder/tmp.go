package courierorder

import (
	"time"
)

// CourierOderTmp is db schema tmp 'couier_order_CC_ISCOD'`
type CourierOderTmp struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Barcode   *string   `json:"barcode,omitempty"`
	CreatedAt time.Time `gorm:"index;autoCreateTime"`
	UpdatedAt time.Time `gorm:"index;autoUpdateTime" `
}
