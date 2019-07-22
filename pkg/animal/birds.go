package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getBirds() []Animal {
	birds := []Animal{
		{
			Name:           "flamingo",
			PluralName:     "flamingos",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      true,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "flamingo feathers",
					Origin: "flamingo",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
		{
			Name:           "peacock",
			PluralName:     "peacocks",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      true,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "peacock feathers",
					Origin: "peacock",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
	}

	return birds
}
