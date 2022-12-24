Mingw32Version=10.0.0_3
CXX=/opt/homebrew/Cellar/mingw-w64/${Mingw32Version}/bin/x86_64-w64-mingw32-g++
CC=/opt/homebrew/Cellar/mingw-w64/${Mingw32Version}/bin/x86_64-w64-mingw32-gcc


build-mac:
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o ./out/darwin/

build-windows:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=${CC} CXX=${CXX} go build -o ./out/windows/

build-all: build-mac build-windows
