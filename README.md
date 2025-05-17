# Competitive Programming Test Runner

This project provides a simple test automation script for Go-based competitive programming problems. It allows you to run test cases with sample inputs and compare the outputs against expected results using a unified diff format.

---

## 📁 Directory Structure

```
.
├── problems/            # Root directory for all problem sets
│   ├── abc405/
│   │   ├── a.go         # Go source file to test
│   │   ├── in/          # Input test cases
│   │   │   ├── a1
│   │   │   └── ...
│   │   └── expect/      # Expected outputs
│   │       ├── a1
│   │       └── ...
├── diffs/               # Generated unified diffs for failed cases
│   ├── a1.diff
│   └── ...
├── test.sh              # Shell script to automate testing
└── README.md            # This file
```

---

## 🚀 How to Run Tests

### 1. Prerequisites

* Go must be installed.
* For macOS users who want colored unified diff:

```bash
brew install diffutils
alias diff=gdiff
```

### 2. Usage

```bash
./test.sh <directory> <problem>
```

#### Example:

```bash
./test.sh abc405 a
```

This will run `go build` on `problems/abc405/a.go`, then execute the binary with each input from `in/` and compare with corresponding output from `expect/`.

---

## ✅ Test Output

* `PASS`: Output matches expected.
* `FAIL`: Output differs. A unified diff is printed and saved in the `diffs/` directory.

### Diff Example:

```diff
--- expect/a1
+++ tmp_a1
@@ -1 +1 @@
-expected output
+actual output
```

---

## 💡 Notes

* Input and expected files must have matching filenames (e.g., `a1`).
* The script loops over all test cases in the `in/` directory.
* Failed diffs are saved to `diffs/<testcase>.diff`.
