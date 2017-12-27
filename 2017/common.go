package main

import (
	"strconv"
	"log"
)

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to convert %v to int: %v", s, err)
	}
	return i
}