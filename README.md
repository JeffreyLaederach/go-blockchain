# Overview

**This is an extremely simple blockchain I used to experiment with the Go programming language for the first time.**

## Go Programming Language Setup:

1. Download and install Go (1.23.0) from this link: [Download Go](https://go.dev/doc/install)

2. Verify that Go is successfully installed by opening Command Prompt and typing the command `go version` and it should return "`go version go1.23.0 windows/amd64`"

3. Enable dependency tracking by using the `go mod init` command plus the file path to this GitHub repository on your computer so the full command should look something like "`go mod init OneDrive/Documents/GitHub/go-blockchain/blockchain`" and you should see a "go.mod" file appear in the root folder (if there isn't one already)

4. Install the Visual Studio Code "Go" Extension

5. You are all ready to run Go files in the integrated terminal in Visual Studio Code by simply using the appropriate commands below:

### Compile/Create Executable:

(Creates a .exe file named after the folder in which the .go files are located in)

`go build`

### Run Code:

`go run .`
