package orderapi

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initDb() *Db {
	conf := LoadConfig("../fortest.env")
	return NewDb(conf)
}

func initData(db *Db) (string, []uint) {
	user := NewUser("666666")
	newJwt := NewJWT(user.Secret)
	token, _ := newJwt.Create(user.Phone)

	db.DB.Create(user)
	db.DB.Create(NewProduct("каша", "гречка", 6))
	db.DB.Create(NewProduct("каша", "перловка", 6))
	db.DB.Create(NewProduct("конфета", "медунок", 6))
	prods := []Product{}
	db.DB.Find(&prods)
	prodIds := make([]uint, len(prods))
	for i, v := range prods {
		prodIds[i] = v.ID
	}
	return token, prodIds
}

func clearData(db *Db) {
	var order Order
	db.DB.Find(&order)
	db.DB.Unscoped().Table("order_products").Where("order_id=?", order.ID).Delete("*")
	db.DB.Unscoped().Where("address=?", "домой").Delete(&Order{})
	db.DB.Unscoped().Where("price_rub=?", 6).Delete(&Product{})
	db.DB.Unscoped().Where("phone=?", "666666").Delete(&User{})
}

func TestCrOrder(t *testing.T) {
	db := initDb()
	token, prodIds := initData(db)
	defer clearData(db)

	data, _ := json.Marshal(&NewOrderReq{
		Address:    "домой",
		ProductIds: prodIds,
	})

	ts := httptest.NewServer(ProdServ(StatusTest))
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL+"/order", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Fatal(err)
	}
	body, _ := io.ReadAll(res.Body)
	log.Println(string(body))
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("want %d,got %d", http.StatusOK, res.StatusCode)
	}
}
