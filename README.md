# Competitive Programming Test Runner

This project provides a simple test automation workflow for Go-based competitive programming problems. It includes problem setup, test execution, and output comparison, all tailored to a structured problem directory layout.

---

## 📁 Directory Structure

```plaintext
.
├── problems/                     # All contest problems live here
│   ├── abc405/                   # Contest name
│   │   ├── a/                    # Problem name (per letter)
│   │   │   ├── main.go           # Problem source code
│   │   │   ├── input/            # Input test cases
│   │   │   │   ├── a_case01.input.txt
│   │   │   │   └── ...
│   │   │   ├── expect/           # Expected outputs
│   │   │   │   ├── a_case01.expect.txt
│   │   │   │   └── ...
│   │   │   ├── actual/           # Actual outputs from latest test run
│   │   │   │   ├── a_case01.actual.txt
│   │   │   ├── diffs/            # Diffs if any test fails
│   │   │   │   ├── a_case01.diff
│   │   │   └── ...
│   │   └── ...
│   └── template/                 # Template files for new problems
│       └── x/
│           ├── main.go
│           ├── input/
│           │   ├── x_case01.input.txt
│           ├── expect/
│           │   ├── x_case01.expect.txt
├── test.sh                       # Run tests for a problem
├── setup.sh                      # Create new problem from template
├── clean.sh                      # Clean up test artifacts
├── .gitignore                    # Ignore actual/ and diffs/
└── README.md
```

---

## 🚀 How to Run Tests

### 1. Set up a problem

```bash
./setup.sh abc405 a
```

This will copy `template/x/` to `problems/abc405/a/`, rename files to match `a_caseNN`, and place them in the proper folders.

### 2. Run tests

```bash
./test.sh abc405 a
```

This builds `main.go`, runs each test input, compares against expected output, prints diffs, and saves them if mismatched.

### 3. Clean up

```bash
./clean.sh
```

This removes all `*.out`, `actual/`, and `diffs/` directories under `problems/`.

---

## ✅ Test Output

* **PASS**: Output matches expected
* **FAIL**: Output differs → diff is shown and saved

```diff
--- expect/a_case01.expect.txt
+++ actual/a_case01.actual.txt
@@ -1 +1 @@
-expected output
+actual output
```

---

## 🛠 VS Code Integration

### tasks.json

You can run tasks from VS Code with argument prompts:

* **Run Tests (prompt)** → calls `test.sh` with user input
* **Setup Problem (prompt)** → calls `setup.sh` with user input
* **Clean All** → calls `clean.sh`

### launch.json

Debug any problem file by specifying:

```json
{
  "name": "Debug Go (prompt)",
  "type": "go",
  "request": "launch",
  "program": "${workspaceFolder}/problems/${input:contest}/${input:problem}/main.go"
}
```

---

## 🧼 .gitignore

```gitignore
**/actual/
**/diffs/
*.out
```

---

## 💡 Notes

* File names must follow: `a_case01.input.txt`, `a_case01.expect.txt`
* Template files are located in `problems/template/x/`
* Everything is problem-isolated: no cross-contest conflicts

Feel free to extend with case descriptions, metadata, or CI integration!
