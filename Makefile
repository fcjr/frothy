.DEFAULT_GOAL := all

all: clean install-tools darwin windows

download:
	@echo Download go.mod dependencies
	@go1.16 mod download

install-tools: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go1.16 install %

darwin: darwin_binary darwin_app darwin_dmg

darwin_binary:
	@GOOS=darwin GOARCH=amd64 go1.16 build -o bin/frothy_darwin_amd64 cmd/desktop/frothy.go
	# TODO look into implementing variadic in macdriver to support apple silicon
	# GOOS=darwin GOARCH=arm64 go1.16 build -o bin/frothy_darwin_arm64 cmd/desktop/frothy.go
	# lipo -create bin/frothy_darwin_amd64 bin/frothy_darwin_amd64 -output bin/frothy_darwin_universal

darwin_app: darwin_binary
	@cp -r assets/mac/Frothy.app bin/
	@mkdir -p "bin/Frothy.app/Contents/MacOS"
	@cp bin/frothy_darwin_amd64 "bin/Frothy.app/Contents/MacOS/Frothy"
	@chmod +x "bin/Frothy.app/Contents/MacOS/Frothy"

darwin_dmg: darwin_app
	@mkdir -p bin/dmg_source
	@cp -r bin/Frothy.app bin/dmg_source/
	@create-dmg \
		--volname "Frothy Install" \
		--volicon "assets/images/installer.icns" \
		--background "assets/images/dmg_background.png" \
		--window-pos 200 120 \
		--window-size 658 526 \
		--icon-size 100 \
		--icon "Frothy.app" 219 350 \
		--hide-extension "Frothy.app" \
		--app-drop-link 438 350 \
		"bin/Frothy.dmg" \
		"bin/dmg_source/"
	@rm -rf bin/dmg_source

windows: windows_binary windows_nsis_installer 

windows_binary:
	@rsrc -manifest frothy.exe.manifest -ico assets/images/logo.ico -o rsrc.syso
	@GOOS=windows GOARCH=amd64 go1.16 build -ldflags="-H windowsgui" -o bin/Frothy.exe cmd/desktop/frothy.go

windows_nsis_installer: windows_binary
	makensis \
		/DOUTPUTFILE=../../bin/FrothyInstaller.exe \
		/DMAJORVERSION=0 \
		/DMINORVERSION=0 \
		/DBUILDVERSION=1 \
		assets/windows/frothy.nsi

clean:
	@rm -rf bin