package config
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func SetupDB() *gorm.DB {
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	panic(err.Error())
	}
	return db
}