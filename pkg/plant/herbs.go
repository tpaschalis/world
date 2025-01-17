package plant

import "github.com/ironarachne/world/pkg/resource"

func getHerbs() []Plant {
	herbs := []Plant{
		{
			Name:           "parsley",
			PluralName:     "parsleys",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "parsley",
					Origin:       "parsley",
					MainMaterial: "parsley",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "brahmi",
			PluralName:     "brahmis",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "brahmi",
					Origin:       "brahmi",
					MainMaterial: "brahmi",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "basil",
			PluralName:     "basils",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "basil",
					Origin:       "basil",
					MainMaterial: "basil",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "cilantro",
			PluralName:     "cilantros",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "cilantro",
					Origin:       "cilantro",
					MainMaterial: "cilantro",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "saffron",
			PluralName:     "saffrons",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "saffron",
					Origin:       "saffron",
					MainMaterial: "saffron",
					Tags: []string{
						"herb",
					},
					Commonality: 3,
					Value:       1,
				},
			},
		},
		{
			Name:           "rosemary",
			PluralName:     "rosemarys",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "rosemary",
					Origin:       "rosemary",
					MainMaterial: "rosemary",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "thyme",
			PluralName:     "thymes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "thyme",
					Origin:       "thyme",
					MainMaterial: "thyme",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "black pepper",
			PluralName:     "black peppers",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "black pepper",
					Origin:       "black pepper",
					MainMaterial: "black pepper",
					Tags: []string{
						"spice",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "oregano",
			PluralName:     "oreganos",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "oregano",
					Origin:       "oregano",
					MainMaterial: "oregano",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "mint",
			PluralName:     "mints",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "mint",
					Origin:       "mint",
					MainMaterial: "mint",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "lemongrass",
			PluralName:     "lemongrasses",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "lemongrass",
					Origin:       "lemongrass",
					MainMaterial: "lemongrass",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "coriander",
			PluralName:     "corianders",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "coriander",
					Origin:       "coriander",
					MainMaterial: "coriander",
					Tags: []string{
						"spice",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "clove",
			PluralName:     "cloves",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "clove",
					Origin:       "clove",
					MainMaterial: "clove",
					Tags: []string{
						"spice",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "sage",
			PluralName:     "sages",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "sage",
					Origin:       "sage",
					MainMaterial: "sage",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "turmeric",
			PluralName:     "turmerics",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "turmeric",
					Origin:       "turmeric",
					MainMaterial: "turmeric",
					Tags: []string{
						"spice",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "vanilla",
			PluralName:     "vanillas",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "vanilla",
					Origin:       "vanilla",
					MainMaterial: "vanilla",
					Tags: []string{
						"herb",
					},
					Commonality: 3,
					Value:       1,
				},
			},
		},
		{
			Name:           "galangal",
			PluralName:     "galangals",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "galangal",
					Origin:       "galangal",
					MainMaterial: "galangal",
					Tags: []string{
						"spice",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "ginger",
			PluralName:     "gingers",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "ginger",
					Origin:       "ginger",
					MainMaterial: "ginger",
					Tags: []string{
						"spice",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
	}

	return herbs
}
