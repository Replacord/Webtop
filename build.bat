go get

set GOOS=windows
set GOARCH=386
go build -ldflags -H=windowsgui -o bin/Win/WebtopWin386.exe Webtop.go
set GOARCH=amd64
go build -ldflags -H=windowsgui -o bin/Win/WebtopWinAmd64.exe Webtop.go
set GOARCH=arm
go build -ldflags -H=windowsgui -o bin/Win/WebtopWinArm.exe Webtop.go

set GOOS=linux
set GOARCH=386
go build -ldflags -H=windowsgui -o bin/Linux/WebtopLinux386 Webtop.go
set GOARCH=amd64
go build -ldflags -H=windowsgui -o bin/Linux/WebtopLinuxAmd64 Webtop.go
set GOARCH=arm
go build -ldflags -H=windowsgui -o bin/Linux/WebtopLinuxArm Webtop.go

set GOOS=darwin
set GOARCH=amd64
go build -ldflags -H=windowsgui -o bin/Mac/WebtopMacAmd64 Webtop.go

