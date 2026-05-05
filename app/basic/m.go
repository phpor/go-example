package main

import "strings"

func matchCondition(osType, os, model, brand, amodel string) bool {
	// os_type='ios'
	if osType == "ios" {
		return true
	}

	// lower(os) rlike 'ios|iphone|ipad'
	if strings.Contains(strings.ToLower(os), "ios") ||
		strings.Contains(strings.ToLower(os), "iphone") ||
		strings.Contains(strings.ToLower(os), "ipad") {
		return true
	}

	// model rlike '^ios|^iphone|^ipad'
	if strings.HasPrefix(model, "ios") ||
		strings.HasPrefix(model, "iphone") ||
		strings.HasPrefix(model, "ipad") {
		return true
	}

	// lower(brand) rlike '^apple|^iphone|^ipad|^ipod'
	if strings.HasPrefix(strings.ToLower(brand), "apple") ||
		strings.HasPrefix(strings.ToLower(brand), "iphone") ||
		strings.HasPrefix(strings.ToLower(brand), "ipad") ||
		strings.HasPrefix(strings.ToLower(brand), "ipod") {
		return true
	}

	// lower(amodel) rlike '^apple|^iphone|^ipad|^ipod'
	if strings.HasPrefix(strings.ToLower(amodel), "apple") ||
		strings.HasPrefix(strings.ToLower(amodel), "iphone") ||
		strings.HasPrefix(strings.ToLower(amodel), "ipad") ||
		strings.HasPrefix(strings.ToLower(amodel), "ipod") {
		return true
	}

	return false
}