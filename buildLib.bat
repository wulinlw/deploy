copy /y C:\go_test\deploy\deploy\upload\upload.go C:\Go_path\src\deploy\upload
copy /y C:\go_test\deploy\deploy\command\command.go C:\Go_path\src\deploy\command
go install deploy/upload
go install deploy/command
pause