package rdb

import (
	"errors"
	"fmt"
	"log"
	"os"

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
