package entity 

import (
	"error"
	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
)  

func Product struct {
	gorm.Model

	Name 		string `json:"name"`
	Price 	float64 `json:"price"`
	Stock 	int			`json:"stokc"`
	Description		string `json:"description"`
}


