package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/alevinval/fingerprints/internal/matching"
	"github.com/alevinval/fingerprints/internal/types"
)

func main() {
	a, b := os.Args[1], os.Args[2]

	var r1, r2 types.DetectionResult
	json.Unmarshal([]byte(a), &r1)
	json.Unmarshal([]byte(b), &r2)
	max := len(r1.Minutia)
	if len(r1.Minutia) > len(r1.Minutia) {
		max = len(r2.Minutia)
	}

	matches := matching.Match(r1, r2)
	d, _ := json.Marshal(matches)

	log.Printf("matched minutiaes: %d/%d", len(matches), max)
	log.Printf("matches: %s", d)
}
