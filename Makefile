.PHONY: build compile clean

full:	clean compile synth

build:	clean compile

synth:
	npx cdk synth

compile:
	echo "compiling lambdas"
	GOOS=linux go build -o bin/hello/handler ./app/handlers/hello

clean:
	echo "cleaning bin/ directory"
	rm -rf ./bin/