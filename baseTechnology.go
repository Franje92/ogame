package ogame

import (
	"math"
	"time"
)

// BaseTechnology base struct for technologies
type BaseTechnology struct {
	BaseLevelable
}

// ConstructionTime returns the duration it takes to build given technology
func (b BaseTechnology) ConstructionTime(level, universeSpeed int, facilities Facilities) time.Duration {
	price := b.GetPrice(level)
	metalCost := float64(price.Metal)
	crystalCost := float64(price.Crystal)
	researchLabLvl := float64(facilities.ResearchLab)
	hours := (metalCost + crystalCost) / (1000 * (1 + researchLabLvl) * float64(universeSpeed))
	secs := math.Max(1, hours*3600)
	return time.Duration(int(math.Floor(secs))) * time.Second
}

// GetLevel returns current level of a technology
func (b BaseTechnology) GetLevel(resourcesBuildings ResourcesBuildings, facilities Facilities, researches Researches) int {
	return researches.ByID(b.ID)
}
