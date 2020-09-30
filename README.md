# Fingerprints

Fingerprint image analysis in Go. Identify, compare and match biometric
fingerprint images.

# Running the library

This is a good summary of what happens when you run this software
on a fingerprint image. On the right is the input image, on the left
is a debug output image.

![input](examples/example-input-1.png)
![output](examples/example-output-1.png)
![input](examples/example-input-2.png)
![output](examples/example-output-2.png)

The library detects features of the fingerprint (bifurcations, terminations ...)
which represent the "identity" of the fingerprint. It provides matching functions
to assert whether two fingerprint images are from the same person, or not.
