package main

import (
	"encoding/json"

	"github.com/alevinval/fingerprints/internal/cmdhelper"
	"github.com/alevinval/fingerprints/internal/matching"
)

func main() {
	img := cmdhelper.LoadImage("corpus/nist3.jpg")
	minutia := matching.ExtractFeatures(img)
	d, _ := json.Marshal(minutia)
	println(string(d))
}
