package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getSnakes() []Animal {
	animals := []Animal{
		Animal{
			Name:       "asp",
			PluralName: "asps",
		},
		Animal{
			Name:       "cobra",
			PluralName: "cobras",
		},
		Animal{
			Name:       "rattlesnake",
			PluralName: "rattlesnakes",
		},
		Animal{
			Name:       "black mamba",
			PluralName: "black mambas",
		},
		Animal{
			Name:       "coral snake",
			PluralName: "coral snakes",
		},
		Animal{
			Name:       "boa constrictor",
			PluralName: "boa constrictors",
		},
	}

	for _, a := range animals {
		a.AnimalType = "mammal"
		a.EatsAnimals = true
		a.EatsPlants = false
		a.IsMount = false
		a.IsPackAnimal = false
		a.IsScavenger = false
		a.MinHumidity = 0
		a.MaxHumidity = 10
		a.MinTemperature = 6
		a.MaxTemperature = 10
		a.Size = size.GetCategoryByName("tiny")
		a.Resources = []resource.Resource{
			resource.Resource{
				Name:        a.Name + " hide",
				Origin:      a.Name,
				Type:        "hide",
				Commonality: 3,
			},
			resource.Resource{
				Name:        a.Name + " fangs",
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
			resource.Resource{
				Name:        a.Name + " heart",
				Origin:      a.Name,
				Type:        "meat",
				Commonality: 1,
			},
			resource.Resource{
				Name:        a.Name + " venom",
				Origin:      a.Name,
				Type:        "venom",
				Commonality: 3,
			},
		}
	}

	return animals
}