package main

import (
	"encoding/json"
	"os"

	"github.com/alevinval/fingerprints/internal/cmdhelper"
	"github.com/alevinval/fingerprints/internal/matching"
)

func main() {
	path := os.Args[1]
	img := cmdhelper.LoadImage(path)
	minutia := matching.ExtractFeatures(img)
	d, _ := json.Marshal(minutia)
	println(string(d))
}
