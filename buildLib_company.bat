copy /y D:\gotest\grpc_my\deploy\upload\upload.go C:\go_path\src\deploy\upload
copy /y D:\gotest\grpc_my\deploy\command\command.go C:\go_path\src\deploy\command
go install deploy/upload
go install deploy/command
pause