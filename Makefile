.PHONY: corpus clean

build: clean
	go build ./cmd/fingerprint-corpus
	go build ./cmd/fingerprint-extract
	go build ./cmd/fingerprint-compare

clean:
	rm -f fingerprint-corpus
	rm -f fingerprint-extract
	rm -f fingerprint-compare

corpus: clean build
	mkdir -p out
	./fingerprint-corpus corpus/nist1.jpg out

generate-readme-examples: clean build
	bash -c "./scripts/generate-readme-examples.sh"

test-extract-match:
	bash -c "./scripts/test-extract-and-match.sh"
