package ports

import "github.com/tayalone/barcode-hex-go/barcode"

// MainRepository is ....
type MainRepository interface {
	ReadAll(query interface{}) []barcode.Config
	ReadById(id uint) barcode.Config
	Create(input barcode.Input) barcode.Config
	Update(input barcode.Update) error
	DeleteById(id uint) error
	ReadByCond(logID uint64, couierCode string, isCod bool) barcode.Config
}
