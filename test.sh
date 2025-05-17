#!/bin/bash

# Usage check
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <contest_dir> <problem_letter>"
  echo "Example: $0 abc405 a"
  exit 1
fi

contest="$1"
letter="$2"
basedir="problems/$contest/$letter"
src="$basedir/main.go"
bin="$basedir/$letter.out"
indir="$basedir/input"
expectdir="$basedir/expect"
actualdir="$basedir/actual"
diffdir="$basedir/diffs"

mkdir -p "$actualdir" "$diffdir"

# Build
if ! go build -o "$bin" "$src"; then
  echo "Build failed"
  exit 1
fi

# Enable nullglob
shopt -s nullglob
inputs=("$indir"/${letter}_case*.input.txt)
shopt -u nullglob

# No inputs found
if [ ${#inputs[@]} -eq 0 ]; then
  echo "No input files found in $indir"
  exit 1
fi

pass=0
fail=0

for infile in "${inputs[@]}"; do
  case_base=$(basename "$infile" .input.txt)
  expect_file="$expectdir/${case_base}.expect.txt"
  actual_file="$actualdir/${case_base}.actual.txt"
  diff_out="$diffdir/${case_base}.diff"

  "$bin" < "$infile" > "$actual_file"

  if diff -q "$actual_file" "$expect_file" > /dev/null; then
    echo -e "$case_base: \033[0;32mPASS\033[0m"
    ((pass++))
    rm -f "$diff_out"
  else
    echo -e "$case_base: \033[0;31mFAIL\033[0m"
    ((fail++))
    diff -u --color=always "$expect_file" "$actual_file"
    diff -u "$expect_file" "$actual_file" > "$diff_out"
  fi
done

echo "========================"
echo -e "Total: $((pass + fail)), \033[0;32mPassed: $pass\033[0m, \033[0;31mFailed: $fail\033[0m"
