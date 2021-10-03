package domain

type Product struct {
	BaseModel

	ID          uint64 `gorm:"primary_key"`
	Name        string
	Description string
	ImagesURLs  string // comma separated
	Weight      string
	Dimensions  string
	Categories  string // comma separated

	// Foreign keys
	Prices    []ProductPrice    `gorm:"foreignKey:ProductID"`
	Locations []ProductLocation `gorm:"foreignKey:ProductID"`
}

func (Product) TableName() string {
	return "products"
}
