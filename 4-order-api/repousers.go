package orderapi

import "gorm.io/gorm/clause"

type UserRepositoryDeps struct {
	Db *Db
}

type UserRepository struct {
	Db *Db
}

func NewUserRepository(db *Db) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	res := repo.Db.DB.Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (repo *UserRepository) GetUserByPhone(phone string) (*User, error) {
	var user User
	res := repo.Db.DB.First(&user, "phone=?", phone)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (repo *UserRepository) Update(user *User) (*User, error) {
	res := repo.Db.DB.Clauses(clause.Returning{}).Updates(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (repo *UserRepository) GetBySession(sessionId string) (*User, error) {
	var user User
	res := repo.Db.DB.First(&user, "session_id=?", sessionId)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
