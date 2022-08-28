package ports

import "github.com/tayalone/barcode-hex-go/barcode"

// BarcodeService is ....
type BarcodeService interface {
	GetAll(query interface{}) []barcode.Config
	GetById(id uint) barcode.Config
	Create(input barcode.Input) barcode.Config
	UpdateById(id uint, input barcode.Update) error
	RemoveById(id uint) error

	GetBarCode(logID uint64, couierCode string, isCod bool) string
}
