package courierorder

/*
CourierOderDhl is db schema tmp 'couier_order_dhl'`
use `Composition Technic` https://golangbot.com/inheritance/
*/
type CourierOderDhl struct {
	Tmp
}

// GetID get pk
func (c *CourierOderDhl) GetID() uint {
	return c.ID
}

// GetData get Tmp
func (c *CourierOderDhl) GetData() Tmp {
	return Tmp{
		ID:        c.ID,
		Barcode:   c.Barcode,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

// SetBarcode set initil Barcode
func (c *CourierOderDhl) SetBarcode(barcode string) {
	c.Barcode = &barcode
}
