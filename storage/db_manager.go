package storage

import (
	"github.com/MAXXKUMAR/Exp/models"
	"errors"
	"math/rand"
)

// Simulated data for mode counts
var simulatedData = map[string]map[string]int{
	"123": {"mode1": rand.Intn(100), "mode2": rand.Intn(100)},
	"456": {"mode1": rand.Intn(100), "mode3": rand.Intn(100)},
}

func GetModeCounts(areaCode string) (map[string]int, error) {
	// Simulate fetching mode counts from the database
	if data, ok := simulatedData[areaCode]; ok {
		return data, nil
	}
	return nil, errors.New("mode counts not found for the specified area code")
}
