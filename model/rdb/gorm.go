package rdb

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/tayalone/barcode-mvc-go/model/rdb/courierorder"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*Connection is Definition of DbConnection Behavior */
type Connection interface {
	GetDb() *gorm.DB
	GetStatus() bool
}

/*RDB is Definition of Value */
type RDB struct {
	db     *gorm.DB
	errMsg string
}

var myRdb = RDB{}

// Connect is Trigger gorm connect postgres DB
func Connect() *RDB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("RDM_HOST"),
		os.Getenv("RDM_USER"),
		os.Getenv("RDM_PASSWORD"),
		os.Getenv("RDM_DB"),
		os.Getenv("RDM_PORT"),
		os.Getenv("TIME_ZONE"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		log.Println("FAIL: Connect RDB Error", err.Error())

		myRdb.db = nil
		myRdb.errMsg = err.Error()
		return &myRdb

	}
	log.Println("Connect RDB Success!!!")
	fmt.Println("db", db)
	myRdb.db = db
	myRdb.errMsg = ""
	fmt.Println("myRdb", myRdb)
	fmt.Println("myRdb.db", myRdb.db)
	fmt.Println("myRdb.errMsg", myRdb.errMsg)

	return &myRdb
}

// GetDbInstance return pointer of Db Connection Instance
func GetDbInstance() (*RDB, error) {
	fmt.Println(myRdb)
	if myRdb.GetStatus() {
		return &myRdb, nil
	}
	return nil, errors.New("Db is Connecting Fail")
}

// GetDb return gorm db
func (rdb *RDB) GetDb() *gorm.DB {
	return rdb.db
}

// GetStatus return status connect db status
func (rdb *RDB) GetStatus() bool {
	return myRdb.db != nil && myRdb.errMsg == ""
}

// AutoMigrate Watch And Validate Data
func (rdb *RDB) AutoMigrate() {
	db := rdb.db
	// db.Set("gorm:table_options", "ENGINE=InnoDB")

	// /  about 'barcode_condition'
	if (db.Migrator().HasTable(&BarcodeCondition{})) {
		log.Println("Table Existing, Drop IT")

		db.Migrator().DropTable(&BarcodeCondition{})
	}
	db.AutoMigrate(&BarcodeCondition{})
	log.Println("Create 'barcode_conditions'")

	// / Add Initail Data
	initBCC := []BarcodeCondition{
		{
			CourierCode:   "DHL",
			IsCod:         true,
			StartBarcode:  "DCA00000001XTH",
			BatchSize:     100,
			PrevCondLogID: 1,
			CondLogID:     101,
		}, {
			CourierCode:   "DHL",
			IsCod:         false,
			StartBarcode:  "DNA00000001XTH",
			BatchSize:     300,
			PrevCondLogID: 1,
			CondLogID:     301,
		},
	}
	db.Create(initBCC)

	log.Println("Create Initil Data")

	// /  about 'courier_order_dhl'
	if (db.Migrator().HasTable(&courierorder.CourierOderDhl{})) {
		log.Println("Table Existing, Drop IT")

		db.Migrator().DropTable(&courierorder.CourierOderDhl{})
	}
	db.AutoMigrate(&courierorder.CourierOderDhl{})
	log.Println("Create 'courier_order_dhl'")

	// /  about 'courier_order_dhl_cod'
	if (db.Migrator().HasTable(&courierorder.CourierOderDhlCod{})) {
		log.Println("Table Existing, Drop IT")

		db.Migrator().DropTable(&courierorder.CourierOderDhlCod{})
	}
	db.AutoMigrate(&courierorder.CourierOderDhlCod{})
	log.Println("Create 'courier_order_dhl_cod'")
}
