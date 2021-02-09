build: clean build_mac build_windows

build_mac:
	GOOS=darwin GOARCH=amd64 go1.16rc1 build -o bin/frothy_darwin_amd64 cmd/frothy.go
	cp -r assets/mac/Frothy.app bin/
	mkdir -p "bin/Frothy.app/Contents/MacOS"
	cp bin/frothy_darwin_amd64 "bin/Frothy.app/Contents/MacOS/Frothy"
	chmod +x "bin/Frothy.app/Contents/MacOS/Frothy"

build_windows:
	GOOS=windows GOARCH=amd64 go1.16rc1 build -ldflags="-H windowsgui" -o bin/frothy.exe cmd/frothy.go

clean:
	rm -rf bin