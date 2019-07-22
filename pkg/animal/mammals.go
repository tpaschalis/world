package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getMammals() []Animal {
	animals := []Animal{
		{
			Name:           "beaver",
			PluralName:     "beavers",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      true,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 6,
			Resources: []resource.Resource{
				{
					Name:        "beaver hide",
					Origin:      "beaver",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "beaver teeth",
					Origin:      "beaver",
					Tags: []string{
						"teeth",
					},
					Commonality: 5,
				},
				{
					Name:        "beaver meat",
					Origin:      "beaver",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("small"),
		},
		{
			Name:           "deer",
			PluralName:     "deer",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "deer hide",
					Origin:      "deer",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "deer teeth",
					Origin:      "deer",
					Tags: []string{
						"teeth",
					},
					Commonality: 5,
				},
				{
					Name:        "venison",
					Origin:      "deer",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "deer antler",
					Origin:      "deer",
					Tags: []string{
						"antler",
						"bone",
					},
					Commonality: 5,
				},
				{
					Name:   "deer sinew",
					Origin: "deer",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "squirrel",
			PluralName:     "squirrels",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "squirrel hide",
					Origin:      "squirrel",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "squirrel",
					Origin:      "squirrel",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
		{
			Name:           "camel",
			PluralName:     "camels",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        true,
			IsPackAnimal:   true,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "camel hide",
					Origin:      "camel",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "camel teeth",
					Origin:      "camel",
					Tags: []string{
						"teeth",
					},
					Commonality: 5,
				},
				{
					Name:        "camel",
					Origin:      "camel",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "camel milk",
					Origin:      "camel",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
				},
				{
					Name:   "camel sinew",
					Origin: "camel",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "bison",
			PluralName:     "bison",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 6,
			Resources: []resource.Resource{
				{
					Name:        "bison hide",
					Origin:      "bison",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "bison teeth",
					Origin:      "bison",
					Tags: []string{
						"teeth",
					},
					Commonality: 5,
				},
				{
					Name:        "bison",
					Origin:      "bison",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "bison bone",
					Origin:      "bison",
					Tags: []string{
						"bone",
					},
					Commonality: 5,
				},
				{
					Name:   "bison sinew",
					Origin: "bison",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "cow",
			PluralName:     "cows",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name: "calf brains",
					Origin: "cow",
					Tags: []string{
						"brains",
						"meat",
					},
				},
				{
					Name:        "cow hide",
					Origin:      "cow",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "cow teeth",
					Origin:      "cow",
					Tags: []string{
						"teeth",
					},
					Commonality: 5,
				},
				{
					Name:        "beef",
					Origin:      "cow",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "beef loin",
					Origin:      "cow",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "beef ribs",
					Origin:      "cow",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "cow bone",
					Origin:      "cow",
					Tags: []string{
						"bone",
					},
					Commonality: 5,
				},
				{
					Name:   "cow sinew",
					Origin: "cow",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "elephant",
			PluralName:     "elephants",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        true,
			IsPackAnimal:   true,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "elephant hide",
					Origin:      "elephant",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "elephant milk",
					Origin:      "elephant",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
				},
				{
					Name:        "elephant",
					Origin:      "elephant",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "ivory",
					Origin:      "elephant",
					Tags: []string{
						"bone",
					},
					Commonality: 5,
				},
				{
					Name:   "elephant sinew",
					Origin: "elephant",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("large"),
		},
		{
			Name:           "goat",
			PluralName:     "goats",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "goat hide",
					Origin:      "goat",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "goat's milk",
					Origin:      "goat",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
				},
				{
					Name:        "goat",
					Origin:      "goat",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:   "goat sinew",
					Origin: "goat",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "sheep",
			PluralName:     "sheep",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "wool",
					Origin:      "sheep",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "sheep's milk",
					Origin:      "sheep",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
				},
				{
					Name:        "lamb",
					Origin:      "sheep",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:   "sheep sinew",
					Origin: "sheep",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "alpaca",
			PluralName:     "alpacas",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   true,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 4,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:        "alpaca wool",
					Origin:      "alpaca",
					Tags: []string{
						"hide",
						"wool",
					},
					Commonality: 5,
				},
				{
					Name:        "alpaca milk",
					Origin:      "alpaca",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
				},
				{
					Name:        "alpaca",
					Origin:      "alpaca",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:   "alpaca sinew",
					Origin: "alpaca",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "llama",
			PluralName:     "llamas",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   true,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:        "llama wool",
					Origin:      "llama",
					Tags: []string{
						"hide",
						"wool",
					},
					Commonality: 5,
				},
				{
					Name:        "llama",
					Origin:      "llama",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "llama milk",
					Origin:      "llama",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
				},
				{
					Name:   "llama sinew",
					Origin: "llama",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "hippopotamus",
			PluralName:     "hippopotamuses",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      true,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "hippopotamus hide",
					Origin:      "hippopotamus",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "hippopotamus",
					Origin:      "hippopotamus",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "hippopotamus milk",
					Origin:      "hippopotamus",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("large"),
		},
		{
			Name:           "antelope",
			PluralName:     "antelopes",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "antelope hide",
					Origin:      "antelope",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "antelope",
					Origin:      "antelope",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "antelope milk",
					Origin:      "antelope",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
				},
				{
					Name:   "antelope sinew",
					Origin: "antelope",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "gazelle",
			PluralName:     "gazelles",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "gazelle hide",
					Origin:      "gazelle",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "gazelle",
					Origin:      "gazelle",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "gazelle milk",
					Origin:      "gazelle",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
				},
				{
					Name:   "gazelle sinew",
					Origin: "gazelle",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "rabbit",
			PluralName:     "rabbits",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "rabbit hide",
					Origin:      "rabbit",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "rabbit",
					Origin:      "rabbit",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
		{
			Name:           "ermine",
			PluralName:     "ermines",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 6,
			Resources: []resource.Resource{
				{
					Name:        "ermine fur",
					Origin:      "ermine",
					Tags: []string{
						"fur",
					},
					Commonality: 5,
				},
				{
					Name:        "ermine",
					Origin:      "ermine",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("small"),
		},
		{
			Name:           "mink",
			PluralName:     "minks",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 6,
			Resources: []resource.Resource{
				{
					Name:        "mink fur",
					Origin:      "mink",
					Tags: []string{
						"fur",
					},
					Commonality: 5,
				},
				{
					Name:        "mink",
					Origin:      "mink",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("small"),
		},
		{
			Name:           "pig",
			PluralName:     "pigs",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "pig hide",
					Origin:      "pig",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "pork",
					Origin:      "pig",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "pork loin",
					Origin:      "pig",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "pork intestine",
					Origin:      "pig",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "pork ribs",
					Origin:      "pig",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "bacon",
					Origin:      "pig",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:   "pig sinew",
					Origin: "pig",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "raccoon",
			PluralName:     "raccoons",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    true,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:        "raccoon hide",
					Origin:      "raccoon",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "raccoon",
					Origin:      "raccoon",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("small"),
		},
		{
			Name:           "reindeer",
			PluralName:     "reindeer",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   true,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 4,
			Resources: []resource.Resource{
				{
					Name:        "reindeer hide",
					Origin:      "reindeer",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
				},
				{
					Name:        "reindeer",
					Origin:      "reindeer",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
				},
				{
					Name:        "reindeer milk",
					Origin:      "reindeer",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
				},
				{
					Name:        "reindeer bone",
					Origin:      "reindeer",
					Tags: []string{
						"bone",
					},
					Commonality: 5,
				},
				{
					Name:   "reindeer sinew",
					Origin: "reindeer",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
	}

	return animals
}
