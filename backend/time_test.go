package main

import (
	"log"
	"testing"
	"time"
)

func TestThing(t *testing.T) {
	_, err := time.Parse("2006-01-02", "2022-12-31")
	if err != nil {
		log.Fatalf("Unable to parse due date: %v", err)
	}
}
