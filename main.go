package main

import (
	"fmt"

	"github.com/Lastjigsaw/SE68-LabExam-final.git/entity"
)

func main() {
	p := entity.Product{
		Name:        "fksdfjajdf",
		Price:       100,
		Stock:       10,
		Description: "short",
	}

	err := p.Validate()
	if err != nil {
		fmt.Println("Validation error:", err.Error())
		return
	}

	fmt.Println("Product is valid")
}
