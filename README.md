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

---

## 🔧 Customization

You can extend or modify the script to support:

* Filtering specific test cases
* Running multiple problem sets (e.g., `abc405`, `abc406`, ...)
* Enhanced diff output with context

Feel free to adjust the script to your workflow!

---

## 🐞 Debugging in VS Code

To debug a Go file from `problems/<dir>/<file>.go` interactively in VS Code:

1. Create a `.vscode/launch.json` with the following:

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Debug Go (select problem)",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/problems/${input:dir}/${input:file}",
      "args": [],
      "cwd": "${workspaceFolder}/problems/${input:dir}",
      "console": "integratedTerminal"
    }
  ],
  "inputs": [
    {
      "id": "dir",
      "type": "promptString",
      "description": "Enter problem directory (e.g. abc405)",
      "default": "abc405"
    },
    {
      "id": "file",
      "type": "promptString",
      "description": "Enter filename (e.g. a.go)",
      "default": "a.go"
    }
  ]
}
```

2. Open the Run & Debug tab, select "Debug Go (select problem)"
3. Input the directory and file name when prompted
4. Set breakpoints and start debugging 🚀
