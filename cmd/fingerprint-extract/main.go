package main

import (
	"encoding/json"
	"os"

	"github.com/alevinval/fingerprints/internal/cmdhelper"
	"github.com/alevinval/fingerprints/internal/extraction"
)

func main() {
	path := os.Args[1]
	_, m := cmdhelper.LoadImage(path)
	result := extraction.DetectionResult(m)
	d, _ := json.Marshal(result)
	println(string(d))
}
