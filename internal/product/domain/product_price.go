package domain

type ProductPrice struct {
	BaseModel

	ID        uint64 `gorm:"primary_key"`
	UserID    uint64
	ProductID uint64 `gorm:"index"`

	Currency         string
	Price            uint64
	IsPriceEstimated bool
}

func (ProductPrice) TableName() string {
	return "product_prices"
}
