package services

import (
	"github.com/MAXXKUMAR/Exp/storage"
	"github.com/MAXXKUMAR/Exp/cache"
	"github.com/MAXXKUMAR/Exp/storage"
)

func GetModeCounts(areaCode string) (map[string]int, error) {
	// Check the cache first
	if cachedData, found := cache.GetFromCache(areaCode); found {
		return cachedData, nil
	}

	// Fetch mode counts from the database
	modeCounts, err := storage.GetModeCounts(areaCode)
	if err != nil {
		return nil, err
	}

	// Cache the retrieved data for future requests
	cache.StoreInCache(areaCode, modeCounts)

	return modeCounts, nil
}
