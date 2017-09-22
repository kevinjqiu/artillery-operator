.PHONY: clean

build/artillery-operator:
	go build -o build/artillery-operator

clean:
	rm -fr build
