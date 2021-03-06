package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"

	"github.com/kr/pretty"
	_ "github.com/lib/pq"

	"database/sql/driver"

	_ "github.com/jackc/pgx/stdlib"
	kallax "gopkg.in/src-d/go-kallax.v1"
)

// ByteArray represents a byte array `bytea`.
type ByteArray []byte

// Scan implements the sql.Scanner interface.
func (a *ByteArray) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		// lib/pq reuses buffer. Need to copy instead.
		// *(*[]byte)(a) = src
		srcCopy := make([]byte, len(src), len(src))
		copy(srcCopy, src)
		*a = srcCopy
		return nil
	case string:
		*a = []byte(src)
		return nil
	case nil:
		*a = nil
		return nil
	}

	return fmt.Errorf("kallax: cannot convert %T to ByteArray", src)
}

// Value implements the driver.Valuer interface.
func (a ByteArray) Value() (driver.Value, error) {
	return ([]byte)(a), nil
}

type TestModel struct {
	kallax.Model `table:"test_models"`
	ID           int64 `pk:"autoincr"`
	// Data         []byte
	// Data2        []byte
	Data    ByteArray `sqltype:"bytea"`
	Data2   ByteArray `sqltype:"bytea"`
	Counter int64
}

var dbconn *sql.DB

func init() {
	var err error

	dbURL := "postgres://localhost/kallax_test?sslmode=disable"
	// pgconfig, err := pgx.ParseConnectionString(dbURL)
	// if err != nil {
	// 	log.Fatalln("db", err)
	// }

	// dbconn, err = pgx.Connect(pgconfig)

	// dbconn, err = sql.Open("pgx", dbURL)
	dbconn, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalln("db", err)
	}
}

// localhost/kallax_test?sslmode=disable
func test() (err error) {
	store := NewTestModelStore(dbconn)

	var data [32]byte
	var data2 [32]byte
	_, err = rand.Read(data[:])
	_, err = rand.Read(data2[:])
	if err != nil {
		return
	}

	obj := NewTestModel()
	obj.Data = data[:]
	obj.Data2 = data2[:]
	err = store.Insert(obj)
	if err != nil {
		return
	}

	obj = store.MustFindOne(NewTestModelQuery().FindByID(obj.ID))

	pretty.Println(obj)
	fmt.Printf("obj.Data: %x\n", obj.Data)
	fmt.Printf("obj.Data2: %x\n", obj.Data2)

	obj.Counter += 1
	_, err = store.Update(obj, Schema.TestModel.Counter)
	if err != nil {
		return
	}

	pretty.Println("after update", obj)
	fmt.Printf("after update obj.Data: %x\n", obj.Data)
	fmt.Printf("after update obj.Data2: %x\n", obj.Data2)

	return nil
}

func main() {
	err := test()
	if err != nil {
		log.Fatalln("main err:", err)
	}
}
