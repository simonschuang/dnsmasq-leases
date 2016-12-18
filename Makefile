
all:
	CGO_ENABLED=0 go build dnsmasq-leases.go

clean:
	$(RM) dnsmasq-leases
