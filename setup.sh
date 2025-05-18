#!/bin/bash

# Usage check
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <contest> <problem>"
  echo "Example: $0 abc405 a"
  exit 1
fi

contest="$1"
letter="$2"
url="https://atcoder.jp/contests/${contest}/tasks/${contest}_${letter}"

template_dir="problems/template/x"
problem_dir="problems/$contest/$letter"
in_dir="$problem_dir/input"
expect_dir="$problem_dir/expect"
main_target="$problem_dir/main.go"

# Create directories
mkdir -p "$in_dir" "$expect_dir"

# Copy main.go from template
cp "$template_dir/main.go" "$main_target"

# Fetch and rename samples using inline Python
python3 - "$contest" "$letter" "$url" <<'EOF'
import os, sys, subprocess, glob, shutil

# for devcontainer
os.environ["PATH"] = os.path.expanduser("~/.local/bin") + ":" + os.environ["PATH"]

contest = sys.argv[1]
problem = sys.argv[2]
url = sys.argv[3]

basedir = f"problems/{contest}/{problem}"
input_dir = os.path.join(basedir, "input")
expect_dir = os.path.join(basedir, "expect")

# Download in the problem directory directly into .
subprocess.run(["oj", "download", "-d", ".", url], check=True, cwd=basedir)

# Rename and move sample files
inputs = sorted(glob.glob(os.path.join(basedir, "sample-*.in")))
outputs = sorted(glob.glob(os.path.join(basedir, "sample-*.out")))

for i, (inf, outf) in enumerate(zip(inputs, outputs), 1):
    base = f"{problem}_case{i:02}"
    shutil.move(inf, os.path.join(input_dir, f"{base}.input.txt"))
    shutil.move(outf, os.path.join(expect_dir, f"{base}.expect.txt"))
EOF

echo "Setup complete:"
echo "  Source: $main_target"
echo "  Input:  $in_dir"
echo "  Expect: $expect_dir"
