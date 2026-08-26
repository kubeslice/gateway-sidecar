package sidecar

import (
	"net"
	"strings"
	"testing"
)

// TestTunnelMSSClampCommands verifies the MSS-clamp iptables rule targets the
// mangle table's FORWARD chain on the given interface with clamp-mss-to-pmtu,
// and that -C (check) and -A (add) differ only in that verb. The interface is
// parameterized (not hard-wired to tun0) so the rule is correct for any tunnel.
func TestTunnelMSSClampCommands(t *testing.T) {
	for _, iface := range []string{"tun0", "wg0"} {
		check, add := tunnelMSSClampCommands(iface)
		for _, cmd := range []string{check, add} {
			for _, want := range []string{"iptables", "-t mangle", "FORWARD", "-o " + iface, "-p tcp", "--tcp-flags SYN,RST SYN", "TCPMSS", "--clamp-mss-to-pmtu"} {
				if !strings.Contains(cmd, want) {
					t.Errorf("iface %s: command %q missing %q", iface, cmd, want)
				}
			}
		}
		if !strings.Contains(check, "-C FORWARD") {
			t.Errorf("iface %s: check command should use -C: %q", iface, check)
		}
		if !strings.Contains(add, "-A FORWARD") {
			t.Errorf("iface %s: add command should use -A: %q", iface, add)
		}
		if strings.Replace(check, "-C", "-A", 1) != add {
			t.Errorf("iface %s: check and add should differ only in the -C/-A verb:\n check=%q\n add=%q", iface, check, add)
		}
	}
}

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
