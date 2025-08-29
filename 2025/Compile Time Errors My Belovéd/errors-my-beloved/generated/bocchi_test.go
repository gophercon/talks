package bocchi_test

import (
	"unsafe"

	"git.sunturtle.xyz/zephyr/errors-my-beloved/generated/models"
)

// If a compiler error happens here, then the database schema has changed,
// so you need to update unit tests.
var _ [0]struct{} = [unsafe.Sizeof(models.Bocchi{}) - 56]struct{}{}
