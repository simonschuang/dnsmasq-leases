
## Build
build source code from local
$ make 

build source code in docker
$ make build

cleanup binary
$ make clean

## Usages
./dnsmasq-leases -h

./dnsmasq-leases dhcp.lease
```
{
    "Lease": [
        {
            "Timestamp": "2000-01-01T01:19:35Z",
            "Mac": "00:00:00:00:00:05",
            "Ip": "192.168.1.155",
            "Host": "wdt",
            "Id": "01:00:00:00:00:00:05"
        },
        {
            "Timestamp": "2000-01-01T01:18:42Z",
            "Mac": "00:00:00:00:00:04",
            "Ip": "192.168.1.237",
            "Host": "*",
            "Id": "01:00:00:00:00:00:04"
        },
        {
            "Timestamp": "2000-01-01T01:15:51Z",
            "Mac": "00:0f:b0:3a:b5:0b",
            "Ip": "192.168.1.208",
            "Host": "colinux",
            "Id": "01:00:0f:b0:3a:b5:0b"
        },
        {
            "Timestamp": "2000-01-01T01:18:13Z",
            "Mac": "02:0f:b0:3a:b5:0b",
            "Ip": "192.168.1.199",
            "Host": "*",
            "Id": "01:02:0f:b0:3a:b5:0b"
        }
    ]
}
```
