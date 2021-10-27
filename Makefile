start:
	go build -buildmode=plugin -o $@.so modules/$@/main.go
	strip $@.so

modules: start 
