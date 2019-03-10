package iphash

import (
	"errors"
	"net/url"
	"sync"
	"unsafe"

	"github.com/cespare/xxhash"
	"github.com/hlts2/round-robin"
)

// ErrServersNotExists is the error that servers dose not exists
var ErrServersNotExists = errors.New("servers dose not exist")

// IPHash is base ip-hash interface
type IPHash interface {
	Next(in *url.URL) *url.URL
}

type iphash struct {
	urls []*url.URL
	cnt  uint64
	m    map[uint64]*url.URL
	mu   *sync.Mutex
	rr   roundrobin.RoundRobin
}

// New returns IPHash(*iphash) object
func New(urls []*url.URL) (IPHash, error) {
	if len(urls) == 0 {
		return nil, ErrServersNotExists
	}

	rr, _ := roundrobin.New(urls)

	return &iphash{
		urls: urls,
		cnt:  uint64(len(urls)),
		m:    make(map[uint64]*url.URL),
		mu:   new(sync.Mutex),
		rr:   rr,
	}, nil
}

func (i *iphash) Next(in *url.URL) *url.URL {
	hashN := xxhash.Sum64(*(*[]byte)(unsafe.Pointer(&in.Host))) % i.cnt

	i.mu.Lock()
	defer i.mu.Unlock()

	if url, ok := i.m[hashN]; ok {
		return url
	}

	url := i.rr.Next()
	i.m[hashN] = url

	return url
}
