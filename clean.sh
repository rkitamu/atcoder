#!/bin/bash

echo "Cleaning all test artifacts..."

# Remove compiled binaries (*.out)
find problems -type f -name '*.out' -exec rm -v {} +

# Remove entire actual/ and diffs/ directories (recursively)
find problems -type d -name 'actual' -exec rm -rv {} +
find problems -type d -name 'diffs' -exec rm -rv {} +

echo "Done."
