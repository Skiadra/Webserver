@echo off
echo Building and running server...
CompileDaemon --build="go build -o ./main.exe ./cmd/main.go" --command="./main.exe"