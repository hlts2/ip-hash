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
iph, err := iphash.New([]string{
    "server-1",
    "server-2",
    "server-3",
})

iph.Next("192.168.33.10") // server-1
iph.Next("192.168.33.10") // server-1
iph.Next("192.168.33.11") // server-2
iph.Next("192.168.33.11") // server-2
```

## Author
[hlts2](https://github.com/hlts2)

## LICENSE
ip-hash released under MIT license, refer [LICENSE](https://github.com/hlts2/ip-hash/blob/master/LICENSE) file.
