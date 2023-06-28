all: run

run:
	go run cmd/main/main.go

push:
	git push git@github.com:RB-PRO/SmartLogisterNotificationBot.git

pull:
	git pull git@github.com:RB-PRO/SmartLogisterNotificationBot.git

pushW:
	git pushW https://github.com/RB-PRO/SmartLogisterNotificationBot.git

pullW:
	git pull https://github.com/RB-PRO/SmartLogisterNotificationBot.git

pushCar:
	scp -r main authorization.json root@194.87.107.129:go/SmartLogisterNotificationBot/

build-config:
	go env GOOS GOARCH

build-linux-osx:
	export GOARCH=arm
	export GOOS=linux
	go env GOOS GOARCH
	go build ./cmd/main/main.go  

build-windows-linux:
	set GOARCH=amd64
	set GOOS=linux
	set CGO_ENABLED=0
	go env GOOS GOARCH
	go build cmd/main/main.go

build-linux-windows:
	export GOARCH=amd64
	export GOOS=windows
	go env GOOS GOARCH
	go build ./cmd/main/main.go