build: clean build_mac build_windows

build_mac:
	GOOS=darwin GOARCH=amd64 go1.16rc1 build -o bin/frothy_darwin_amd64 cmd/frothy.go
	# TODO look into implementing variadic in macdriver to support apple silicon
	# GOOS=darwin GOARCH=arm64 go1.16rc1 build -o bin/frothy_darwin_arm64 cmd/frothy.go
	# lipo -create bin/frothy_darwin_amd64 bin/frothy_darwin_amd64 -output bin/frothy_darwin_universal

	cp -r assets/mac/Frothy.app bin/
	mkdir -p "bin/Frothy.app/Contents/MacOS"
	cp bin/frothy_darwin_amd64 "bin/Frothy.app/Contents/MacOS/Frothy"
	chmod +x "bin/Frothy.app/Contents/MacOS/Frothy"

build_windows:
	GOOS=windows GOARCH=amd64 go1.16rc1 build -ldflags="-H windowsgui" -o bin/frothy.exe cmd/frothy.go

clean:
	rm -rf bin