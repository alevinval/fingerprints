.PHONY: corpus

corpus:
	go build ./cmd/fingerprint-corpus
	rm -rf out
	mkdir out
	./fingerprint-corpus out
