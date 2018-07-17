package iphash

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"sync"
	"unsafe"

	"github.com/hlts2/round-robin"
)

// ErrServersNotExists is the error that servers dose not exists
var ErrServersNotExists = errors.New("servers dose not exist")

// Servers is custom type of servers
type Servers []string

// IPHash returns ip-hash clouser
func IPHash(servers Servers) (func(string) string, error) {
	if len(servers) == 0 {
		return nil, ErrServersNotExists
	}

	rrNext, err := roundrobin.RoundRobin(roundrobin.Servers(servers))
	if err != nil {
		return nil, err
	}

	mu := new(sync.Mutex)

	m := make(map[string]string)
	prefix := strings.Join(servers, ",")

	return func(ip string) string {
		defer mu.Unlock()
		mu.Lock()

		d := prefix + ip
		hash := md5Hash(*(*[]byte)(unsafe.Pointer(&d)))

		if v, ok := m[hash]; ok {
			return v
		}

		item := rrNext()

		m[hash] = item

		return item
	}, nil
}

func md5Hash(d []byte) string {
	h := md5.New()
	h.Write(d)
	return hex.EncodeToString(h.Sum(nil))
}
