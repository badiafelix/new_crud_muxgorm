package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB //penamaan var jika di file atau folder lain dan dipanggil di main harus pake awal huruf besar

func ConnectDb() { //penamaan function jika di file atau folder lain dan dipanggil di main harus pake awal huruf besar
	dsn := "host=localhost user=postgres password=root1234 dbname=Go_database port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil { 

		fmt.Println(err)
	}
	Db = conn

}
