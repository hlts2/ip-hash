package iphash

import (
	"reflect"
	"testing"
)

func TestIPHash(t *testing.T) {
	tests := []struct {
		servers  Servers
		ips      []string
		errExist bool
		expected Servers
	}{
		{
			servers: Servers{
				"server-1",
				"server-2",
				"server-3",
			},
			ips: []string{
				"192.168.33.10",
				"192.168.33.10",
				"192.168.33.11",
				"192.168.33.11",
			},
			errExist: false,
			expected: Servers{
				"server-1",
				"server-1",
				"server-2",
				"server-2",
			},
		},
		{
			servers: Servers{},
			ips: []string{
				"192.168.33.10",
				"192.168.33.10",
				"192.168.33.11",
				"192.168.33.11",
			},
			errExist: true,
		},
	}

	for _, test := range tests {
		got := make(Servers, 0, len(test.expected))
		assign, err := IPHash(test.servers)

		errExist := !(err == nil)
		if errExist != test.errExist {
			t.Fatalf("IPHash err is wrong. expected: %v, got: %v", test.errExist, errExist)
		}

		if err != nil {
			continue
		}

		for _, ip := range test.ips {
			got = append(got, assign(ip))
		}

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("IPHash is wrong. expected: %v, got: %v", test.expected, got)
		}
	}
}
