import * as vscode from "vscode";
import * as fs from "fs";
import * as path from "path";

export function activate(context: vscode.ExtensionContext) {
  const disposable = vscode.commands.registerCommand(
    "case-runner.runTestCase",
    async () => {
      const rootPath = vscode.workspace.workspaceFolders?.[0].uri.fsPath || ".";
      const problemsDir = path.join(rootPath, "problems");

      // 🔹 Contest picker (exclude 'template')
      const contests = fs
        .readdirSync(problemsDir)
        .filter(
          (f) =>
            fs.statSync(path.join(problemsDir, f)).isDirectory() &&
            f !== "template"
        );
      const contest = await vscode.window.showQuickPick(contests, {
        placeHolder: "Select contest (e.g. abc405)",
      });
      if (!contest) {
        return;
      }

      // 🔹 Problem picker (e.g. a, b)
      const contestDir = path.join(problemsDir, contest);
      const problems = fs
        .readdirSync(contestDir)
        .filter((f) => fs.statSync(path.join(contestDir, f)).isDirectory());
      const problem = await vscode.window.showQuickPick(problems, {
        placeHolder: "Select problem (e.g. a, b, c)",
      });
      if (!problem) {
        return;
      }

      // 🔹 Case picker (e.g. a_case01)
      const inputDir = path.join(contestDir, problem, "input");
      if (!fs.existsSync(inputDir)) {
        vscode.window.showWarningMessage(
          `No input directory found: ${inputDir}`
        );
        return;
      }

      const inputFiles = fs
        .readdirSync(inputDir)
        .filter((f) => f.endsWith(".input.txt"));
      if (inputFiles.length === 0) {
        vscode.window.showWarningMessage(`No test cases found in ${inputDir}`);
        return;
      }

      const caseNames = inputFiles.map((f) => f.replace(".input.txt", ""));
      const picked = await vscode.window.showQuickPick(caseNames, {
        placeHolder: "Select test case",
      });
      if (!picked) {
        return;
      }

      // 🔹 Execute test.sh in terminal
      const command = `./test.sh ${contest} ${problem} --case ${picked}`;
      const terminal = vscode.window.createTerminal("Case Runner");
      terminal.show();
      terminal.sendText(command);
    }
  );

  context.subscriptions.push(disposable);

  const runAllCases = vscode.commands.registerCommand(
    "case-runner.runAllCases",
    async () => {
      const rootPath = vscode.workspace.workspaceFolders?.[0].uri.fsPath || ".";
      const problemsDir = path.join(rootPath, "problems");

      // Contest picker
      const contests = fs
        .readdirSync(problemsDir)
        .filter(
          (f) =>
            fs.statSync(path.join(problemsDir, f)).isDirectory() &&
            f !== "template"
        );
      const contest = await vscode.window.showQuickPick(contests, {
        placeHolder: "Select contest (e.g. abc405)",
      });
      if (!contest) { return; }

      // Problem picker
      const contestDir = path.join(problemsDir, contest);
      const problems = fs
        .readdirSync(contestDir)
        .filter((f) => fs.statSync(path.join(contestDir, f)).isDirectory());
      const problem = await vscode.window.showQuickPick(problems, {
        placeHolder: "Select problem (e.g. a, b, c)",
      });
      if (!problem) { return; }

      // Run test.sh without --case
      const command = `./test.sh ${contest} ${problem}`;
      const terminal = vscode.window.createTerminal("Case Runner (All)");
      terminal.show();
      terminal.sendText(command);
    }
  );
  context.subscriptions.push(runAllCases);
}

export function deactivate() {}
