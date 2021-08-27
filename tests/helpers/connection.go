package helpers

import (
	"database/sql"
	mocket "github.com/selvatico/go-mocket"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	sqlDB, err := sql.Open(mocket.DriverName, "connection_string")
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
