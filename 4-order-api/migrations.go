package orderapi

func CrTbl() {
	conf := LoadConfig()
	db := NewDb(conf)

	db.AutoMigrate(&Product{})
}
