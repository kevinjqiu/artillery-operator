.PHONY: clean

clean:
	rm -fr build

build/artillery-operator:
	go build -o build/artillery-operator
