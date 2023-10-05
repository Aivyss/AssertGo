package assert

import (
	"testing"
)

func Test(t *testing.T) *assertSetup {
	return &assertSetup{t: t}
}
