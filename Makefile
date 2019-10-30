compile:
	go build -o "tmg" ./src/main

test: compile
	./tmg

clean:
	rm tmg
