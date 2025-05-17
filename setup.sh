#!/bin/bash

# Usage check
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <contest_dir> <problem_letter>"
  echo "Example: $0 abc405 a"
  exit 1
fi

contest="$1"
letter="$2"

template_dir="problems/template/x"
problem_dir="problems/$contest/$letter"
in_dir="$problem_dir/in"
expect_dir="$problem_dir/expect"
main_target="$problem_dir/main.go"

# Create target directories
mkdir -p "$in_dir" "$expect_dir"

# Copy and rename input files
for f in "$template_dir/input/"*.input.txt; do
  base=$(basename "$f")
  case_num=$(echo "$base" | sed -E 's/^x_(case[0-9]+)\.input\.txt$/\1/')
  cp "$f" "$in_dir/${letter}_${case_num}.input.txt"
done

# Copy and rename expect files
for f in "$template_dir/expect/"*.expect.txt; do
  base=$(basename "$f")
  case_num=$(echo "$base" | sed -E 's/^x_(case[0-9]+)\.expect\.txt$/\1/')
  cp "$f" "$expect_dir/${letter}_${case_num}.expect.txt"
done

# Copy main.go
cp "$template_dir/main.go" "$main_target"

echo "Setup complete:"
echo "  Source: $main_target"
echo "  Input:  $in_dir"
echo "  Expect: $expect_dir"
