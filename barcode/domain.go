package barcode

import "time"

/*Input is a structue of ...*/
type Input struct {
	CourierCode  string `json:"courierCode"`
	IsCod        bool   `json:"isCod"`
	StartBarcode string `json:"startBarcode"`
	BatchSize    uint32 `json:"batchSize"`
}

/*Update is a structue of ...*/
type Update struct {
	CourierCode string `json:"courierCode"`
	IsCod       bool   `json:"isCod"`
	BatchSize   uint32 `json:"batchSize"`
}

/*Config is a structue of barcode Config*/
type Config struct {
	ID uint64 `json:"id"`
	Input
	FirstLogID uint64    `json:"firstLogId"`
	LastLogID  uint64    `json:"lastLogId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdateAt   time.Time `json:"updateAt"`
}
