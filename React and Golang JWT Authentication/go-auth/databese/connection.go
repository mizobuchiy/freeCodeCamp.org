package databese

import (
	"os"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func Connect()  {
	user := os.Getenv("MYSQL_USER")
  pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
  dbname := os.Getenv("MYSQL_DATABASE")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, dbname)

	_, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
