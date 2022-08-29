package rdb

import (
	"time"
)

// BarcodeCondition is db schema `barcode_condition`
type BarcodeCondition struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	CourierCode   string    `json:"courierCode"`
	IsCod         bool      `json:"isCod"`
	StartBarcode  string    `json:"startBarcode"`
	BatchSize     uint32    `json:"batchSize"`
	PrevCondLogID uint      `json:"prevCondLogId"`
	CondLogID     uint      `json:"condLogId"`
	CreatedAt     time.Time `gorm:"index;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"index;autoUpdateTime" `
}
