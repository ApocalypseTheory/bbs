$basePath = 'dist'

Remove-Item -Path $basePath -Recurse -Force -ErrorAction SilentlyContinue

$env:GOOS = 'windows'

go build -o "$basePath\win\bbs.exe" main.go

$env:GOOS = 'linux'

go build -o "$basePath/linux/bbs" main.go

$execPath = Join-Path $basePath 'win\bbs.exe'

& $execPath
