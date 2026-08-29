// Copyright (C) 2019-2025, Lux Industries, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package json

import "github.com/holiman/uint256"

// Uint256 is a 256-bit unsigned integer that marshals/unmarshals as a JSON
// string, like the narrower widths beside it.
//
// This is the width the EVM carries. A wei value, a token balance or a supply
// does not fit in 64 bits — an 18-decimal balance passes 2^64 at about 18.4
// tokens — so a field of that shape reaching this boundary previously had to
// be a bare string or a hand-rolled big.Int at every site.
//
// uint256.Int rather than big.Int: a fixed-size value cannot be nil, so a
// zero Uint256 is zero rather than a panic, and it does not allocate.
type Uint256 uint256.Int

func (u Uint256) MarshalJSON() ([]byte, error) {
	v := uint256.Int(u)
	return []byte(`"` + v.Dec() + `"`), nil
}

func (u *Uint256) UnmarshalJSON(b []byte) error {
	str := string(b)
	if str == Null {
		return nil
	}
	if len(str) >= 2 {
		if lastIndex := len(str) - 1; str[0] == '"' && str[lastIndex] == '"' {
			str = str[1:lastIndex]
		}
	}
	v, err := uint256.FromDecimal(str)
	if err != nil {
		return err
	}
	*u = Uint256(*v)
	return nil
}
