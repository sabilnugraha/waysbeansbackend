package productsdto

type CreateProductRequest struct {
	Name  string `json:"name" form:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Image int    `json:"image" form:"image" gorm:"type: varchar(255)"`
	Stock int    `json:"stock" gorm:"type: int"`
	Desc  string `json:"desc" form:"desc" gorm:"type: varchar(255)"`
}
