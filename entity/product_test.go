package entity_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/Lastjigsaw/SE68-LabExam-final.git/entity"
)

func TestEntityProduct(t *testing.T) {

	t.Run("valid product", func(t *testing.T) {
		g := NewWithT(t)

		p := entity.Product{
			Name:        "Book",
			Price:       100,
			Stock:       10,
			Description: "This is a good product",
		}
		err := p.Validate()
		g.Expect(err).To(BeNil())
	})

	t.Run("name is empty", func(t *testing.T) {
		g := NewWithT(t)

		p := entity.Product{
			Name:        "",
			Price:       100,
			Stock:       10,
			Description: "This is a good product",
		}
		err := p.Validate()
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("Name is required"))
	})

	t.Run("price <= 0", func(t *testing.T) {
		g := NewWithT(t)

		p := entity.Product{
			Name:        "Book",
			Price:       0,
			Stock:       10,
			Description: "This is a good product",
		}
		err := p.Validate()
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("Price must be greater than 0"))
	})

	t.Run("stock < 0", func(t *testing.T) {
		g := NewWithT(t)

		p := entity.Product{
			Name:        "Book",
			Price:       100,
			Stock:       -1,
			Description: "This is a good product",
		}
		err := p.Validate()
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("Stock cannot be negative"))
	})

	t.Run("description too short", func(t *testing.T) {
		g := NewWithT(t)

		p := entity.Product{
			Name:        "Book",
			Price:       100,
			Stock:       10,
			Description: "short",
		}
		err := p.Validate()
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("Description must be at least 10 characters"))
	})
}
