package orderapi

type OrderRepositoryDeps struct {
	Db *Db
}

type OrderRepository struct {
	Db *Db
}

func NewOrderRepository(db *Db) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (repo *OrderRepository) Create(order *Order) (*Order, error) {
	res := repo.Db.Create(order)
	if res.Error != nil {
		return nil, res.Error
	}
	return order, nil
}

func (repo *OrderRepository) GetById(id uint, userId uint) (*Order, error) {
	var ord Order
	res := repo.Db.Where("user_id=?", userId).First(&ord, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &ord, nil
}

func (repo *OrderRepository) GetByUser(userId uint) ([]Order, error) {
	var orders []Order
	res := repo.Db.Where("user_id=?", userId).Find(&orders)
	if res.Error != nil {
		return nil, res.Error
	}
	return orders, nil
}
