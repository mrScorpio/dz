package orderapi

import (
	"fmt"
)

func CrTbl() {
	conf := LoadConfig()
	db := NewDb(conf)

	db.AutoMigrate(&Product{}, &User{}, &Order{})
}

func TestDb() {
	conf := LoadConfig()
	db := NewDb(conf)
	repo := NewProdRepository(db)
	repo.Create(&Product{
		Name:        "testname",
		Description: "testdescr",
		PriceRub:    666,
	})
	prods, err := repo.GetSlice(6)
	if err != nil {
		panic(err)
	}
	fmt.Println(prods)
}
