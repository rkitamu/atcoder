#!/bin/bash

# Usage: ./test.sh <contest> <problem> [--case case_name]
if [ "$#" -lt 2 ]; then
  echo "Usage: $0 <contest> <problem> [--case case_name]"
  echo "Example: $0 abc405 a"
  echo "         $0 abc405 a --case a_case02"
  exit 1
fi

contest="$1"
letter="$2"
shift 2

case_filter=""
if [ "$1" == "--case" ]; then
  case_filter="$2"
  shift 2
fi

basedir="problems/$contest/$letter"
src="$basedir/main.go"
bin="$basedir/$letter.out"
indir="$basedir/input"
expectdir="$basedir/expect"
actualdir="$basedir/actual"
diffdir="$basedir/diffs"

mkdir -p "$actualdir"
diffdir_created=0  # defer creation of diffs/

# Build
if ! go build -o "$bin" "$src"; then
  echo "Build failed"
  exit 1
fi

# Enable nullglob
shopt -s nullglob
inputs=()
if [ -n "$case_filter" ]; then
  inputs=("$indir/${case_filter}.input.txt")
else
  inputs=("$indir/${letter}_case"*.input.txt)
fi
shopt -u nullglob

# No inputs
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

    if [ "$diffdir_created" -eq 0 ]; then
      mkdir -p "$diffdir"
      diffdir_created=1
    fi

    diff -u --color=always "$expect_file" "$actual_file"
    diff -u "$expect_file" "$actual_file" > "$diff_out"
  fi
done

echo "========================"
echo -e "Total: $((pass + fail)), \033[0;32mPassed: $pass\033[0m, \033[0;31mFailed: $fail\033[0m"
