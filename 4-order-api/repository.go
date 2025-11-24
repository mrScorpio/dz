package orderapi

import "gorm.io/gorm/clause"

type ProdRepositoryDeps struct {
	Db *Db
}

type ProdRepository struct {
	Db *Db
}

func NewProdRepository(db *Db) *ProdRepository {
	return &ProdRepository{
		Db: db,
	}
}

func (repo *ProdRepository) Create(prod *Product) (*Product, error) {
	res := repo.Db.DB.Create(prod)
	if res.Error != nil {
		return nil, res.Error
	}
	return prod, nil
}

func (repo *ProdRepository) Update(prod *Product) (*Product, error) {
	res := repo.Db.DB.Clauses(clause.Returning{}).Updates(prod)
	if res.Error != nil {
		return nil, res.Error
	}
	return prod, nil
}

func (repo *ProdRepository) Delete(id uint) error {
	res := repo.Db.DB.Delete(&Product{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *ProdRepository) GetById(id uint) (*Product, error) {
	var prod Product
	res := repo.Db.DB.First(&prod, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &prod, nil
}

func (repo *ProdRepository) GetSlice(num int) ([]Product, error) {
	var prodSl []Product
	res := repo.Db.DB.Limit(num).Find(&prodSl)
	if res.Error != nil {
		return nil, res.Error
	}
	return prodSl, nil
}
