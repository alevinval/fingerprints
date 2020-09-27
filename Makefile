.PHONY: corpus clean

corpus: clean
	mkdir out
	go build ./cmd/fingerprint-corpus
	./fingerprint-corpus out

clean:
	rm -rf out
	rm -f fingerprint-corpus
