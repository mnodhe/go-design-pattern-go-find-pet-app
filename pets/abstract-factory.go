package pets

import (
	"fmt"
	"go-design-pattern-go-find-pet-app/models"
)

type AnimalInterface interface {
	Show() string
}
type DogFromFactory struct {
	Pet *models.Dog
}

func (dff *DogFromFactory) Show() string {
	return fmt.Sprint("This animal is a %s" + dff.Pet.Breed.Breed)
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (dff *CatFromFactory) Show() string {
	return fmt.Sprint("This animal is a %s" + dff.Pet.Breed.Breed)
}
