# ip-hash
ip-hash is balancing algorithm, based on [round-robin](https://github.com/hlts2/round-robin).

## Requrement

Go (>= 1.8)

## Installation

```shell
go get github.com/hlts2/ip-hash
```

## Example

```go

// Servers is []string type
assign, err := iphash.IPHash(iphash.Servers{
    "server-1",
    "server-2",
    "server-3",
})

assign("192.168.33.10") // server-1
assign("192.168.33.10") // server-1
assign("192.168.33.11") // server-2
assign("192.168.33.11") // server-2
```
