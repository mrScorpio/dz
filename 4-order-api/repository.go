package orderapi

import "gorm.io/gorm/clause"

type LinkRepositoryDeps struct {
	Db *Db
}

type LinkRepository struct {
	Db *Db
}

func NewLinkRepository(db *Db) *LinkRepository {
	return &LinkRepository{
		Db: db,
	}
}

func (repo *LinkRepository) Create(prod *Product) (*Product, error) {
	res := repo.Db.DB.Create(prod)
	if res.Error != nil {
		return nil, res.Error
	}
	return prod, nil
}

func (repo *LinkRepository) Update(prod *Product) (*Product, error) {
	res := repo.Db.DB.Clauses(clause.Returning{}).Updates(prod)
	if res.Error != nil {
		return nil, res.Error
	}
	return prod, nil
}

func (repo *LinkRepository) Delete(id uint) error {
	res := repo.Db.DB.Delete(&Product{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *LinkRepository) GetById(id uint) (*Product, error) {
	var prod Product
	res := repo.Db.DB.First(&prod, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &prod, nil
}

func (repo *LinkRepository) GetSlice(num int) ([]Product, error) {
	var prodSl []Product
	res := repo.Db.DB.Limit(num).Find(&prodSl)
	if res.Error != nil {
		return nil, res.Error
	}
	return prodSl, nil
}
