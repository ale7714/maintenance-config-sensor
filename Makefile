
bin/maintenanceconfigsensor: *.go cmd/module/*.go go.*
	go build -o bin/maintenanceconfigsensor cmd/module/cmd.go

bin/remoteserver: *.go cmd/remote/*.go go.*
	go build -o bin/remoteserver cmd/remote/cmd.go

lint:
	gofmt -w -s .

updaterdk:
	go get go.viam.com/rdk@latest
	go mod tidy

module: bin/maintenanceconfigsensor
	tar czf module.tar.gz bin/maintenanceconfigsensor