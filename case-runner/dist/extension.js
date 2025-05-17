"use strict";
var __create = Object.create;
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __getProtoOf = Object.getPrototypeOf;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __export = (target, all) => {
  for (var name in all)
    __defProp(target, name, { get: all[name], enumerable: true });
};
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __toESM = (mod, isNodeMode, target) => (target = mod != null ? __create(__getProtoOf(mod)) : {}, __copyProps(
  // If the importer is in node compatibility mode or this is not an ESM
  // file that has been converted to a CommonJS file using a Babel-
  // compatible transform (i.e. "__esModule" has not been set), then set
  // "default" to the CommonJS "module.exports" for node compatibility.
  isNodeMode || !mod || !mod.__esModule ? __defProp(target, "default", { value: mod, enumerable: true }) : target,
  mod
));
var __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: true }), mod);

// src/extension.ts
var extension_exports = {};
__export(extension_exports, {
  activate: () => activate,
  deactivate: () => deactivate
});
module.exports = __toCommonJS(extension_exports);
var vscode = __toESM(require("vscode"));
var fs = __toESM(require("fs"));
var path = __toESM(require("path"));
function activate(context) {
  const disposable = vscode.commands.registerCommand(
    "case-runner.runTestCase",
    async () => {
      const rootPath = vscode.workspace.workspaceFolders?.[0].uri.fsPath || ".";
      const problemsDir = path.join(rootPath, "problems");
      const contests = fs.readdirSync(problemsDir).filter(
        (f) => fs.statSync(path.join(problemsDir, f)).isDirectory() && f !== "template"
      );
      const contest = await vscode.window.showQuickPick(contests, {
        placeHolder: "Select contest (e.g. abc405)"
      });
      if (!contest) {
        return;
      }
      const contestDir = path.join(problemsDir, contest);
      const problems = fs.readdirSync(contestDir).filter((f) => fs.statSync(path.join(contestDir, f)).isDirectory());
      const problem = await vscode.window.showQuickPick(problems, {
        placeHolder: "Select problem (e.g. a, b, c)"
      });
      if (!problem) {
        return;
      }
      const inputDir = path.join(contestDir, problem, "input");
      if (!fs.existsSync(inputDir)) {
        vscode.window.showWarningMessage(
          `No input directory found: ${inputDir}`
        );
        return;
      }
      const inputFiles = fs.readdirSync(inputDir).filter((f) => f.endsWith(".input.txt"));
      if (inputFiles.length === 0) {
        vscode.window.showWarningMessage(`No test cases found in ${inputDir}`);
        return;
      }
      const caseNames = inputFiles.map((f) => f.replace(".input.txt", ""));
      const picked = await vscode.window.showQuickPick(caseNames, {
        placeHolder: "Select test case"
      });
      if (!picked) {
        return;
      }
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
      const contests = fs.readdirSync(problemsDir).filter(
        (f) => fs.statSync(path.join(problemsDir, f)).isDirectory() && f !== "template"
      );
      const contest = await vscode.window.showQuickPick(contests, {
        placeHolder: "Select contest (e.g. abc405)"
      });
      if (!contest) {
        return;
      }
      const contestDir = path.join(problemsDir, contest);
      const problems = fs.readdirSync(contestDir).filter((f) => fs.statSync(path.join(contestDir, f)).isDirectory());
      const problem = await vscode.window.showQuickPick(problems, {
        placeHolder: "Select problem (e.g. a, b, c)"
      });
      if (!problem) {
        return;
      }
      const command = `./test.sh ${contest} ${problem}`;
      const terminal = vscode.window.createTerminal("Case Runner (All)");
      terminal.show();
      terminal.sendText(command);
    }
  );
  context.subscriptions.push(runAllCases);
}
function deactivate() {
}
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  activate,
  deactivate
});
//# sourceMappingURL=extension.js.map
