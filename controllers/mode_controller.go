package controllers

import (
	"github.com/MAXXKUMAR/Exp/services"
)

func GetModeCounts(areaCode string) (map[string]int, error) {
	// Call the service to get mode counts for the specified area code
	return services.GetModeCounts(areaCode)
}
