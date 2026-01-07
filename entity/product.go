package entity

import (
	"strconv"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)


type Product struct {
	gorm.Model
	Name        string  `json:"name" valid:"required~Name is required"`
	Price       float64 `json:"price" valid:"required~Price must be greater than 0,gt0~Price must be greater than 0"`
	Stock       int     `json:"stock" valid:"range(0|2147483647)~Stock cannot be negative"`
	Description string  `json:"description" valid:"stringlength(10|10000)~Description must be at least 10 characters"`
}

func init() {
	govalidator.TagMap["gt0"] = govalidator.Validator(func(str string) bool {
		v, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return false
		}
		return v > 0
	})
}

func (p *Product) Validate() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
