package religion

import (
	"fmt"
	"github.com/ironarachne/world/pkg/random"
)

// Class is a type of religion
type Class struct {
	Name            string
	Commonality     int
	FounderTitle    string
	LeaderTitle     string
	PantheonMinSize int
	PantheonMaxSize int
	GatheringPlaces []string
}

func getAllClasses() []Class {
	return []Class{
		{
			Name:            "monotheistic",
			Commonality:     5,
			FounderTitle:    "prophet",
			LeaderTitle:     "priest",
			PantheonMinSize: 1,
			PantheonMaxSize: 1,
			GatheringPlaces: []string{
				"temple",
				"shrine",
				"church",
			},
		},
		{
			Name:            "duotheistic",
			Commonality:     1,
			FounderTitle:    "prophet",
			LeaderTitle:     "priest",
			PantheonMinSize: 2,
			PantheonMaxSize: 2,
			GatheringPlaces: []string{
				"temple",
				"shrine",
				"church",
			},
		},
		{
			Name:            "polytheistic",
			Commonality:     12,
			FounderTitle:    "prophet",
			LeaderTitle:     "priest",
			PantheonMinSize: 10,
			PantheonMaxSize: 30,
			GatheringPlaces: []string{
				"temple",
				"shrine",
				"church",
			},
		},
		{
			Name:            "shamanistic",
			Commonality:     2,
			FounderTitle:    "shaman",
			LeaderTitle:     "shaman",
			PantheonMinSize: 0,
			PantheonMaxSize: 0,
			GatheringPlaces: []string{
				"medicine lodge",
				"sweat lodge",
				"spirit lodge",
			},
		},
		{
			Name:            "philosophical",
			Commonality:     1,
			FounderTitle:    "philosopher",
			LeaderTitle:     "sage",
			PantheonMinSize: 0,
			PantheonMaxSize: 0,
			GatheringPlaces: []string{
				"great hall",
				"forum",
				"public house",
				"town square",
			},
		},
	}
}

func getWeightedClass() (Class, error) {
	classes := getAllClasses()

	weights := map[string]int{}

	for _, c := range classes {
		weights[c.Name] = c.Commonality
	}

	name, err := random.StringFromThresholdMap(weights)
	if err != nil {
		err = fmt.Errorf("Failed to get random weighted religion class: %w", err)
		return Class{}, err
	}

	for _, c := range classes {
		if c.Name == name {
			return c, nil
		}
	}

	err = fmt.Errorf("Failed to get random weighted religion class!")
	return Class{}, err
}
