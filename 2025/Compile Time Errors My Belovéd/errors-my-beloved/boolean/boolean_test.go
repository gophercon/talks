// boolean_test.go
package bocchi_test

import "git.sunturtle.xyz/zephyr/errors-my-beloved/boolean/challenge"

var _ = map[bool]struct{}{
	false:                      {},
	challenge.Flag == "gc25{}": {}, // ERROR: ./boolean_test.go:7:2: duplicate key false in map literal
}

var _ = map[bool]struct{}{
	false:                          {},
	challenge.Flag[13:] != "CRDy}": {}, // silently succeeds, even though it evaluates to false
}
