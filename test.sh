#!/bin/bash

# 色定義
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # 色リセット

# 引数確認
if [ $# -ne 2 ]; then
  echo "Usage: $0 <directory> <problem>"
  echo "Example: $0 abc405 a"
  exit 1
fi

dir="problems/$1"
name="$2"
src="$dir/$name.go"
bin="$dir/$name.out"
indir="$dir/in"
expectdir="$dir/expect"

# diffs 保存用ディレクトリ
mkdir -p diffs

# ビルド（失敗時は終了）
if ! go build -o "$bin" "$src"; then
  echo -e "${RED}Build failed${NC}"
  exit 1
fi

pass_count=0
fail_count=0

for infile in "$indir"/*; do
  base=$(basename "$infile")
  expect_file="$expectdir/$base"
  tmp_out="tmp_$base"
  diff_out="diffs/$base.diff"

  "$bin" < "$infile" > "$tmp_out"

  if diff -q "$tmp_out" "$expect_file" > /dev/null; then
    echo -e "$base: ${GREEN}PASS${NC}"
    ((pass_count++))
    rm -f "$diff_out"
  else
    echo -e "$base: ${RED}FAIL${NC}"
    ((fail_count++))

    # diff to stdout
    echo "Unified diff:"
    diff -u --color=always "$expect_file" "$tmp_out"

    # save diff to file
    diff -u "$expect_file" "$tmp_out" > "$diff_out"
  fi

  rm -f "$tmp_out"
done

total=$((pass_count + fail_count))
echo "========================"
echo -e "Total: $total, ${GREEN}Passed: $pass_count${NC}, ${RED}Failed: $fail_count${NC}"
