package main

import (
	"encoding/json"
	"os"

	"github.com/alevinval/fingerprints/internal/cmdhelper"
	"github.com/alevinval/fingerprints/internal/matching"
)

func main() {
	path := os.Args[1]
	_, m := cmdhelper.LoadImage(path)
	minutia := matching.ExtractFeatures(m)
	d, _ := json.Marshal(minutia)
	println(string(d))
}
