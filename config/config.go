package config
import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func SetupDB() *gorm.DB {
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
 dsn := "host=localhost user=vianto password=Vianto1125 dbname=db_golang port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	panic(err.Error())
	}
	return db
}