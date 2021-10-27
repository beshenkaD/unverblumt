start:
	go build -buildmode=plugin -o start.so modules/start/main.go

plugins: start 
