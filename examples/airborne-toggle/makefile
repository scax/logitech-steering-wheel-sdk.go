.PHONY: build prebuild build-dev-runner

build: prebuild
	go build -o dist/airborne-toggle.exe main.go

build-dev-runner: prebuild
	echo @echo off > ./dist/airborne.cmd
	echo start ./airborne-toggle.exe >> ./dist/airborne.cmd
	echo pause >> ./dist/airborne.cmd

prebuild:
	IF NOT EXIST ./dist mkdir dist
	
clean:
	rmdir /s /q "./dist"