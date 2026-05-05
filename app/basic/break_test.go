package main

import (
	"testing"
)

func TestBreak(t *testing.T) {
l1:
	for {
		for {
			break l1
		}
	}
}