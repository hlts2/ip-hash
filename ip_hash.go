package iphash

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"sync/atomic"
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

	var flg int32

	m := make(map[string]string)
	prefix := strings.Join(servers, ",")

	return func(ip string) string {
		for {
			if flg == 0 && atomic.CompareAndSwapInt32(&flg, 0, 1) {
				break
			}
		}

		d := prefix + ip
		hash := md5Hash(*(*[]byte)(unsafe.Pointer(&d)))

		if v, ok := m[hash]; ok {
			// I do not use defer, decause defer is slow
			flg = 0
			return v
		}

		item := rrNext()

		m[hash] = item

		flg = 0
		return item
	}, nil
}

func md5Hash(d []byte) string {
	h := md5.New()
	h.Write(d)
	return hex.EncodeToString(h.Sum(nil))
}
