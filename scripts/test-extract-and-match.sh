# Runs feature extraction and passes the same input twice to
# comparison tool.
f=$(mktemp)
./fingerprint-extract &> $f
last=$(cat $f | tail -n 1)
./fingerprint-compare "$last" "$last"
