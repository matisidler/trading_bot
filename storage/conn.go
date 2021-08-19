package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

type driver string

const MySQL driver = "mysql"

/* const Postgres driver = "psql" */

func NewConnection(d string) *gorm.DB {
	switch d {
	case "mysql":
		db = NewMySql()
		return db
	/* case "psql":
	db = NewPsql()
	return db */
	default:
		log.Fatal("not valid driver")
		return nil
	}
}

/* func NewPsql() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/school2?sslmode=disable"), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open DB %v", err)
			return
		}
		fmt.Println("connected to postgres")
	})
	return db
} */

func NewMySql() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/trades?parseTime=true"), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open DB %v", err)
			return
		}
		fmt.Println("connected to mysql")
	})
	return db
}

func StringToNull(s string) sql.NullString {
	var nullString sql.NullString
	if s == "" {
		nullString.Valid = false
	} else {
		nullString.Valid = true
		nullString.String = s
	}
	return nullString
}
