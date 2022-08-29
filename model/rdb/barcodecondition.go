package rdb

import (
	"time"
)

// BarcodeCondition is db schema `barcode_condition`
type BarcodeCondition struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	CourierCode   string    `gorm:"not null;index;index:coureier_code_is_code" json:"courierCode"`
	IsCod         bool      `gorm:"not null;index;index:coureier_code_is_code" json:"isCod"`
	StartBarcode  string    `gorm:"not null" json:"startBarcode"`
	BatchSize     uint32    `gorm:"not null" json:"batchSize"`
	PrevCondLogID uint      `gorm:"not null;index" json:"prevCondLogId"`
	CondLogID     uint      `gorm:"not null;index" json:"condLogId"`
	CreatedAt     time.Time `gorm:"index;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"index;autoUpdateTime" `
}
