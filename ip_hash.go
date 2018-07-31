package iphash

import (
	"errors"
	"strconv"

	lockfree "github.com/hlts2/lock-free"
	"github.com/hlts2/round-robin"
)

// ErrServersNotExists is the error that servers dose not exists
var ErrServersNotExists = errors.New("servers dose not exist")

// IPHash is base ip-hash interface
type IPHash interface {
	Next(addr string) string
}

type iphash struct {
	addrs []string
	lf    lockfree.LockFree
	m     map[string]string
	rr    roundrobin.RoundRobin
}

// New returns IPHash(*iphash) object
func New(addrs []string) (IPHash, error) {
	if len(addrs) == 0 {
		return nil, ErrServersNotExists
	}

	rr, _ := roundrobin.New(addrs)

	return &iphash{
		addrs: addrs,
		m:     make(map[string]string),
		lf:    lockfree.New(),
		rr:    rr,
	}, nil
}

func (i *iphash) Next(addr string) string {
	hash := strconv.Itoa(int(fnv32(addr)) % len(i.addrs))

	i.lf.Wait()
	addr, ok := i.m[hash]
	if ok {
		i.lf.Signal()
		return addr
	}

	addr = i.rr.Next()
	i.m[hash] = addr

	i.lf.Signal()

	return addr
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
