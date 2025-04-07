package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Data struct {
	Id   uint
	Name string `gorm:"unique"`
}

var db *gorm.DB

func main() {
	dsn := `host=127.0.0.1 user=coltea password=coltea dbname=coltea port=5432 sslmode=disable TimeZone=Asia/Shanghai`

	var newLogger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             2000 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})
	db1, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	db = db1
	if err := db.AutoMigrate(
		&Data{},
	); err != nil {
		panic(err)
	}
	data := make([]Data, 0)
	for i := 0; i < 1000; i++ {
		data = append(data, Data{
			Name: rand.Int(),
		})
	}

	db.CreateInBatches([]Data{})

}
