package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	// FindCarts() ([]models.Cart, error)
	// GetCart(ID int) (models.Cart, error)
	CreateCart(Cart models.Cart) (models.Cart, error)
	// UpdateCart(Cart models.Cart) (models.Cart, error)
	// DeleteCart(Cart models.Cart) (models.Cart, error)
	// CreateTransactionID(transaction models.Transaction) (models.Transaction, error)
	// FindCartsTransaction(TrxID int) ([]models.Cart, error)
	FindCartId(CartId int) ([]models.Cart, error)
	// UpdateCartTrans(models.Cart) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}

func (r *repository) FindCartId(UserID int) ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Product").Find(&cart, "user_id = ?", UserID).Error

	return cart, err
}
