# AtCoder workspace

Go-based AtCoder workspace.  

---

## ⚙️ Setup

To get started instantly with a preconfigured environment using [Dev Containers](https://code.visualstudio.com/docs/devcontainers/containers), follow these steps:

### 1. Launch in DevContainer

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

### 2. Login to the tools

1. Login to [AtCoder](https://atcoder.jp/login) on your browser.  
2. Copy the REVEL_SESSION from developer tools -> Application -> Storage -> Cookies -> https://atcoder.jp -> REVEL_SESSION.
3. Run `aclogin`.  

```shell
$ aclogin
```

Paste copied REVEL_SESSION

4. Run VSCode task: `Update atcoder-cli template`
5. Complete

## 🚀 Usage

### Basic Workflow

1. **Create a new contest workspace:**

```bash
$ acc new abc123 --template default
```

2. **Write your solution:**
Edit the generated `main.go` file in each problem directory.

3. **Test your solution:**

Open Edited `main.go` file and run the test task: Run Test Cases

4. **Submit your solution:**

TODO: online-judge-tools is not supported yet.

```bash
$ acc submit
```

## 📁 Directory Structure

```plaintext
atcoder/
├── problems/                 # All contest problems live here
│   ├── {ContestID}/
│   │   ├── a/
├── scripts                   # utility scripts (call by vscode tasks)
├── settings                  # each tool settings
├── tools                     # utility tools
├── .gitignore
└── README.md
```

## License

This project is licensed under the MIT License.

It includes code from [gosagawa/atcoder](https://github.com/gosagawa/atcoder),
which is also licensed under the MIT License.
