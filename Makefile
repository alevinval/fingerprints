.PHONY: corpus clean

build: clean
	go build ./cmd/fingerprint-corpus
	go build ./cmd/fingerprint-extract
	go build ./cmd/fingerprint-compare

clean:
	rm -rf out
	rm -f fingerprint-corpus
	rm -f fingerprint-extract
	rm -f fingerprint-compare

corpus: clean build
	mkdir out
	./fingerprint-corpus out

test-extract-match:
	./scripts/test-extract-and-match.sh
