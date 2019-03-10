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
	m    *sync.Map
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
		m:    new(sync.Map),
		rr:   rr,
	}, nil
}

func (i *iphash) Next(in *url.URL) *url.URL {
	hashN := xxhash.Sum64(*(*[]byte)(unsafe.Pointer(&in.Host))) % i.cnt
	v, _ := i.m.LoadOrStore(hashN, i.rr.Next())
	return v.(*url.URL)
}
