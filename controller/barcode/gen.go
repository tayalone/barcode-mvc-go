package barcode

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/tayalone/barcode-mvc-go/model/rdb"
)

/*Gen is Gen New Barcode */
func Gen(id uint, courierCod string, isCod bool) (string, error) {
	myRdb, _ := rdb.GetDbInstance()
	db := myRdb.GetDb()

	bc := &rdb.BarcodeCondition{}

	cond := "courier_code = ? and is_cod = ? AND ( prev_cond_log_id <= ? AND ? <  cond_log_id)"

	r := db.Where(cond, courierCod, isCod, id, id).First(bc)

	if r.RowsAffected != 1 {
		return "", fmt.Errorf("Not Found Condition")
	}

	regFmt := "^([A-Z]{3})+([0-9]{8})+(XTH)"

	match, _ := regexp.MatchString(regFmt, bc.StartBarcode)

	if !match {
		return "", fmt.Errorf("Wrong Barcode Format")
	}

	s := regexp.MustCompile(regFmt).FindAllStringSubmatch(bc.StartBarcode, -1)

	prefix := s[0][1]
	body := s[0][2]
	suffix := s[0][3]

	c, _ := strconv.Atoi(body)
	b := c - int(bc.PrevCondLogID)
	barcode := fmt.Sprintf("%s%s%s", prefix, fmt.Sprintf("%08d", int(id)+b), suffix)

	return barcode, nil
}
