me:
	go build -buildmode=plugin -o $@.so modules/$@/*.go
	strip $@.so

weather:
	go build -buildmode=plugin -o $@.so modules/$@/*.go
	strip $@.so

modules: me weather
