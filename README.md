# AtCoder workspace

This project provides a simple test automation workflow for Go-based AtCoder workspace. It includes problem setup, test execution, and output comparison, all tailored to a structured problem directory layout. A custom VS Code extension `case-runner` enhances the workflow with dynamic UI selection.

---

## ⚙️ Setup

### 🚣️ Use with VS Code DevContainer

To get started instantly with a preconfigured environment using [Dev Containers](https://code.visualstudio.com/docs/devcontainers/containers), follow these steps:

#### 1. Launch in DevContainer

Open this project in VS Code and run:

```bash
Cmd+Shift+P → Dev Containers: Reopen in Container
```

This will build a container with:

* Go:latest
* node.js
* Python3
* online-judge-tools
* atcoder-cli
* Pre-installed VS Code extensions for Go and testing

#### 2. Login to atcoder-cli and online-judge-tools with aclogin

```shell
$ aclogin
```

paste REVEL_SESSION

## 📁 Directory Structure

```plaintext
atcoder/
├── case-runner/               # VS Code extension (optional)
│   ├── src/extension.ts       # Extension source
│   └── ...
├── problems/                 # All contest problems live here
│   ├── abc405/
│   │   ├── a/
│   │   │   ├── main.go
│   │   │   ├── input/
│   │   │   │   ├── a_case01.input.txt
│   │   │   ├── expect/
│   │   │   │   ├── a_case01.expect.txt
│   │   │   ├── actual/       # Output by test.sh
│   │   │   └── diffs/        # Diff files if mismatched
│   ├── b/
│   │   └── ...
│   └── template/             # Template files for new problems
│       ├── x/
│           ├── main.go
│           ├── input/
│           └── expect/
├── test.sh                   # Run tests
├── setup.sh                  # Create new problem and download samples
├── clean.sh                  # Clean up output and diff artifacts
├── .gitignore
└── README.md
```

---

## 🚀 How to Run Tests

### 1. Set up a problem

```bash
./setup.sh abc405 a
```

Copies `template/x/` to `problems/abc405/a/`, renames test cases, and optionally downloads official sample inputs/outputs via `oj`.

### 2. Run tests (via extension)

Use the VS Code command palette:

```bash
Cmd+Shift+P → Case Runner: Run Test Case
```

Pick:

* contest (e.g. abc405)
* problem (e.g. a)
* case (e.g. a\_case02)

Or run all cases:

```bash
Cmd+Shift+P → Case Runner: Run All Cases
```

### 3. Clean up

```bash
./clean.sh
```

Removes all `.out`, `actual/`, and `diffs/` directories.

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

## 🛠 VS Code Extension: Case Runner

The `case-runner` extension allows you to dynamically select contest/problem/case via QuickPick UI and runs `test.sh` with arguments.

### Commands:

* `Case Runner: Run Test Case` – run a single case interactively
* `Case Runner: Run All Cases` – run all tests for selected problem

### Build & Install (VSIX)

```bash
cd case-runner
npm install
vsce package
```

Then in VS Code:

```bash
Extensions → ... → Install from VSIX → select `case-runner-*.vsix`
```

---

## 💡 Notes

* File names follow `a_case01.input.txt` and `a_case01.expect.txt`
* `case-runner` extension eliminates the need for `tasks.json`
* Samples can be fetched from AtCoder using `oj` during setup
* Clean workflow with everything under one repo for Go competitive programming

Feel free to customize for other languages or workflows!

## License

This project is licensed under the MIT License.

It includes code from [gosagawa/atcoder](https://github.com/gosagawa/atcoder),
which is also licensed under the MIT License.
