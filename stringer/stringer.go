package stringer

import (
	"strconv"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipa IPAddr) String() string {
	s := make([]string, len(ipa))
	for i := range ipa {
		s[i] = strconv.Itoa(int(ipa[i]))
	}
	return strings.Join(s, ".")
}

// Stringer exercise main
// hosts := map[string]stringer.IPAddr{
// 	"loopback":  {127, 0, 0, 1},
// 	"googleDNS": {8, 8, 8, 8},
// }
// for name, ip := range hosts {
// 	fmt.Printf("%v: %v\n", name, ip)
// }
