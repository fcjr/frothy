build: clean build_mac

build_mac:
	go1.16rc1 build -o bin/frothy_mac cmd/main.go
	cp -r assets/mac/Frothy.app bin/
	mkdir -p "bin/Frothy.app/Contents/MacOS"
	cp bin/frothy_mac "bin/Frothy.app/Contents/MacOS/Frothy"
	chmod +x "bin/Frothy.app/Contents/MacOS/Frothy"

clean:
	rm -rf bin