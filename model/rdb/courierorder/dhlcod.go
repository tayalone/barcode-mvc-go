package courierorder

/*
CourierOderDhlCod is db schema tmp 'couier_order_dhl_cod'`
use `Composition Technic` https://golangbot.com/inheritance/
*/
type CourierOderDhlCod struct {
	Tmp
}

// GetID get pk
func (c *CourierOderDhlCod) GetID() uint {
	return c.ID
}

// GetData get Tmp
func (c *CourierOderDhlCod) GetData() Tmp {
	return Tmp{
		ID:        c.ID,
		Barcode:   c.Barcode,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func (c *CourierOderDhlCod) SetBarcode(barcode string) {
	c.Barcode = &barcode
}
