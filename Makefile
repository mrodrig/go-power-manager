.setup-builddir:
	mkdir -p build

windows: .setup-builddir main.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc fyne package -os windows
	mv *.exe build/