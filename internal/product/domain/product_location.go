package domain

type ProductLocation struct {
	BaseModel

	ID        uint64 `gorm:"primary_key"`
	UserID    uint64
	ProductID uint64 `gorm:"index"`

	Text     string
	Country  string `gorm:"index"`
	Province string
	City     string `gorm:"index"`
	Area     string
	Street   string
	Building string
	Store    string

	Longitude float64
	Latitude  float64
}

func (ProductLocation) TableName() string {
	return "product_locations"
}
