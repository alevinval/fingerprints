package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/alevinval/fingerprints/internal/matching"
)

func main() {
	a, b := os.Args[1], os.Args[2]

	var l1, l2 matching.MinutiaeList
	json.Unmarshal([]byte(a), &l1)
	json.Unmarshal([]byte(b), &l2)
	max := len(l1)
	if len(l2) > len(l1) {
		max = len(l2)
	}

	matches := matching.Match(l1, l2)
	d, _ := json.Marshal(matches)

	log.Printf("matched minutiaes: %d/%d", len(matches), max)
	log.Printf("matches: %s", d)
}
