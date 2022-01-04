TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
NAMESPACE=edu
NAME=nftower
BINARY=terraform-provider-${NAME}
VERSION=0.2
OS_ARCH=darwin_amd64
DIR=terraform-provider-${NAME}_${VERSION}_${OS_ARCH}

default: install

build:
	go build -o ${BINARY}
clean:
	rm -rf release
 
release:
	rm -rf release
	mkdir -p release/${DIR}
	GOOS=darwin GOARCH=amd64 go build -o ${DIR}/${BINARY}_${VERSION}_darwin_amd64
	zip -r release/${DIR}.zip ${DIR}
	rm -rf release/${DIR}
	cp terraform-registry-manifest.json release/terraform-provider-${NAME}_${VERSION}_manifest.json
	shasum -a 256 *.zip > release/terraform-provider-${NAME}_${VERSION}_SHA256SUMS
	gpg --detach-sign release/terraform-provider-${NAME}_${VERSION}_SHA256SUMS
	
	# GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	# GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	# GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	# GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	# GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	# GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	# GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	# GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	# GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	# GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	# GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64
	

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m   