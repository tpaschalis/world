package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getUrsines() []Animal {
	animals := []Animal{
		Animal{
			Name:       "black bear",
			PluralName: "black bears",
		},
		Animal{
			Name:       "brown bear",
			PluralName: "brown bears",
		},
	}

	for _, a := range animals {
		a.AnimalType = "mammal"
		a.EatsAnimals = true
		a.EatsPlants = true
		a.IsMount = false
		a.IsPackAnimal = false
		a.IsScavenger = false
		a.MinHumidity = 0
		a.MaxHumidity = 10
		a.MinTemperature = 6
		a.MaxTemperature = 10
		a.Size = size.GetCategoryByName("medium")
		a.Resources = []resource.Resource{
			resource.Resource{
				Name:        a.Name + " hide",
				Origin:      a.Name,
				Type:        "hide",
				Commonality: 5,
			},
			resource.Resource{
				Name:        a.Name + " teeth",
				Origin:      a.Name,
				Type:        "teeth",
				Commonality: 3,
			},
			resource.Resource{
				Name:        a.Name,
				Origin:      a.Name,
				Type:        "meat",
				Commonality: 3,
			},
		}
	}

	return animals
}