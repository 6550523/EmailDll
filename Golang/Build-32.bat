#copy the commands to Windows PowerShell and execute
SET GOOS=windows
SET GOARCH=386
go build -ldflags "-s -w" -buildmode=c-shared -o ../Release/email-win-32.dll

SET GOOS=linux
SET GOARCH=386
go build -ldflags "-s -w" -buildmode=c-shared -o ../Release/email-linux-32.so

SET GOOS=darwin
SET GOARCH=386
go build -ldflags "-s -w" -buildmode=c-shared -o ../Release/email-darwin-32.so