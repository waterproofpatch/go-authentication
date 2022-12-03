package go_authentication

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gDb *gorm.DB

// initialize the database, conditionally dropping tables
func resetDb(dropTables bool) {
	log.Printf("Resetting db...")
	db := GetDb()

	// Migrate the schema
	if dropTables {
		log.Printf("Dropping tables...")
		db.Migrator().DropTable(&User{})
	}
	db.AutoMigrate(&User{})
	db.Create(&User{Email: GetConfig().DefaultAdminUser, Password: GetConfig().DefaultAdminPass, IsVerified: true, IsAdmin: true, VerificationCode: "", RegistrationDate: ""})
}

// getDb returns the database object
func GetDb() *gorm.DB {
	if gDb == nil {
		panic("gDb is not initialized!")
	}
	return gDb
}
func InitDb(dbUrl string, dropTables bool) {
	database, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Read
	gDb = database
	resetDb(dropTables == true)
}
