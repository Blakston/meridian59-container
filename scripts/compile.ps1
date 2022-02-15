
# compile patcher and publisher
Write-Output "compiling patcher and publisher ..."
$env:GOOS = 'windows'
go build -ldflags "-s -w"-o bin/patch.exe patcher/cmd/patch/main.go
go build -ldflags "-s -w"-o bin/publish.exe publisher/cmd/publish/main.go
$env:GOOS = 'linux'
go build -ldflags "-s -w" -o bin/patch.bin patcher/cmd/patch/main.go
go build -ldflags "-s -w"-o bin/publish.bin publisher/cmd/publish/main.go
