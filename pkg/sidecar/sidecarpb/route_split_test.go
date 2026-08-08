package sidecar

import (
	"net"
	"testing"
)

func TestMoreSpecificHalves(t *testing.T) {
	cases := []struct{ in string; want []string }{
		{"10.11.0.0/16", []string{"10.11.0.0/17", "10.11.128.0/17"}},
		{"10.11.0.0/20", []string{"10.11.0.0/21", "10.11.8.0/21"}},
		{"10.11.32.3/32", []string{"10.11.32.3/32"}}, // host route unchanged
	}
	for _, c := range cases {
		_, n, _ := net.ParseCIDR(c.in)
		got := moreSpecificHalves(n)
		if len(got) != len(c.want) {
			t.Fatalf("%s: got %d halves, want %d", c.in, len(got), len(c.want))
		}
		for i, g := range got {
			if g.String() != c.want[i] {
				t.Errorf("%s[%d]: got %s want %s", c.in, i, g.String(), c.want[i])
			}
		}
	}
}
