#!/bin/bash

./fingerprint-corpus corpus/nist3.jpg out
cp out/Normalized.png examples/example-input-1.png
cp out/Debug.png examples/example-output-1.png

./fingerprint-corpus corpus/nist4.png out
cp out/Normalized.png examples/example-input-2.png
cp out/Debug.png examples/example-output-2.png
