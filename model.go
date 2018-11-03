package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"

	"github.com/kr/pretty"

	_ "github.com/lib/pq"
	kallax "gopkg.in/src-d/go-kallax.v1"
)

type TestModel struct {
	kallax.Model `table:"test_models"`
	ID           int64 `pk:"autoincr"`
	Data         []byte
	Counter      int64
}

var dbconn *sql.DB

func init() {
	var err error

	dbURL := "postgres://localhost/kallax_test?sslmode=disable"

	dbconn, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalln("db", err)
	}

}

// localhost/kallax_test?sslmode=disable
func test() (err error) {
	store := NewTestModelStore(dbconn)

	var data [32]byte
	_, err = rand.Read(data[:])
	if err != nil {
		return
	}

	obj := NewTestModel()
	obj.Data = data[:]
	err = store.Insert(obj)
	if err != nil {
		return
	}

	pretty.Println(obj)
	fmt.Printf("obj.Data: %x\n", obj.Data)

	obj.Counter += 1
	_, err = store.Update(obj, Schema.TestModel.Counter)
	if err != nil {
		return
	}

	pretty.Println("after update", obj)
	fmt.Printf("after update obj.Data: %x\n", obj.Data)

	return nil
}

func main() {
	err := test()
	if err != nil {
		log.Fatalln("main err:", err)
	}
}
