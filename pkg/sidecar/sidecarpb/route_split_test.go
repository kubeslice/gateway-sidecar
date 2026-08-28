package sidecar

import (
	"net"
	"strings"
	"testing"

	"github.com/vishvananda/netlink"
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
	cases := []struct {
		in   string
		want []string
	}{
		{"10.11.0.0/16", []string{"10.11.0.0/17", "10.11.128.0/17"}},
		{"10.11.0.0/20", []string{"10.11.0.0/21", "10.11.8.0/21"}},
		{"10.11.32.3/32", []string{"10.11.32.3/32"}},               // host route unchanged
		{"10.11.0.0/31", []string{"10.11.0.0/32", "10.11.0.1/32"}}, // smallest splittable v4
		{"fd00::/48", []string{"fd00::/49", "fd00:0:0:8000::/49"}}, // IPv6 splits too
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

// TestStaleTunnelRouteKeys verifies the teardown diff: keys present in the
// previously-programmed set but absent from the desired set are returned for
// removal, so a topology/subnet change withdraws the old routes.
func TestStaleTunnelRouteKeys(t *testing.T) {
	route := func(cidr string) netlink.Route {
		_, n, _ := net.ParseCIDR(cidr)
		return netlink.Route{Dst: n}
	}
	prev := map[string]netlink.Route{
		"10.11.0.0/17":   route("10.11.0.0/17"),
		"10.11.128.0/17": route("10.11.128.0/17"),
	}
	// Flip to a peer /24 split (entire-slice route no longer desired).
	desired := []netlink.Route{route("10.20.5.0/25"), route("10.20.5.128/25")}
	stale := staleTunnelRouteKeys(prev, desired)
	if len(stale) != 2 {
		t.Fatalf("expected both old /17s stale, got %v", stale)
	}

	// No change: nothing stale.
	same := []netlink.Route{route("10.11.0.0/17"), route("10.11.128.0/17")}
	if got := staleTunnelRouteKeys(prev, same); len(got) != 0 {
		t.Fatalf("expected nothing stale when desired == previous, got %v", got)
	}

	// Partial overlap: only the non-desired key is stale.
	partial := []netlink.Route{route("10.11.0.0/17"), route("10.11.64.0/18")}
	got := staleTunnelRouteKeys(prev, partial)
	if len(got) != 1 || got[0] != "10.11.128.0/17" {
		t.Fatalf("expected only 10.11.128.0/17 stale, got %v", got)
	}
}
