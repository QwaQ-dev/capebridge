package parsebigint

import (
	"fmt"
	"math/big"
)

func ParseBigInt(name, s string) (*big.Int, error) {
	n, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return nil, fmt.Errorf("invalid %s value: %q", name, s)
	}
	return n, nil
}
