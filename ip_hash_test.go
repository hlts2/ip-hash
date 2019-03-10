package iphash

// import (
// 	"reflect"
// 	"testing"
// )
//
// func TestIPHash(t *testing.T) {
// 	tests := []struct {
// 		servers  []string
// 		ips      []string
// 		errExist bool
// 		expected []string
// 	}{
// 		{
// 			servers: []string{
// 				"server-1",
// 				"server-2",
// 				"server-3",
// 			},
// 			ips: []string{
// 				"192.168.33.10",
// 				"192.168.33.10",
// 				"192.168.33.11",
// 				"192.168.33.11",
// 			},
// 			errExist: false,
// 			expected: []string{
// 				"server-1",
// 				"server-1",
// 				"server-2",
// 				"server-2",
// 			},
// 		},
// 		{
// 			servers: []string{},
// 			ips: []string{
// 				"192.168.33.10",
// 				"192.168.33.10",
// 				"192.168.33.11",
// 				"192.168.33.11",
// 			},
// 			errExist: true,
// 		},
// 	}
//
// 	for _, test := range tests {
// 		got := make([]string, 0, len(test.expected))
// 		iphash, err := New(test.servers)
//
// 		errExist := !(err == nil)
// 		if errExist != test.errExist {
// 			t.Fatalf("IPHash err is wrong. expected: %v, got: %v", test.errExist, errExist)
// 		}
//
// 		if err != nil {
// 			continue
// 		}
//
// 		for _, ip := range test.ips {
// 			got = append(got, iphash.Next(ip))
// 		}
//
// 		if !reflect.DeepEqual(test.expected, got) {
// 			t.Errorf("IPHash is wrong. expected: %v, got: %v", test.expected, got)
// 		}
// 	}
// }
