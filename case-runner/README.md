# Case Runner

This is a VS Code extension for competitive programming.

You can select a contest, problem, and case file from the UI, and it will run `test.sh` with the appropriate arguments in the terminal.

## Usage

- Press `Cmd+Shift+P` → `Case Runner: Run Test Case`
- Select contest (e.g., abc405)
- Select problem (e.g., a)
- Select case (e.g., a_case01)

The extension will run `./test.sh abc405 a --case a_case01`.

## License

MIT
