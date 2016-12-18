

default:
	CGO_ENABLED=0 go build dnsmasq-leases.go

clean:
	$(RM) dnsmasq-leases

build:
	docker run -v $(PWD):/dnsmasq-leases -w /dnsmasq-leases golang:1.8-alpine go build
