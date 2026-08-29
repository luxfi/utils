// Copyright (C) 2019-2025, Lux Industries, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package json

import (
	"encoding/json"
	"testing"

	"github.com/holiman/uint256"
)

func TestUint256RoundTrip(t *testing.T) {
	for _, dec := range []string{
		"0",
		"18446744073709551615",   // uint64 max
		"18446744073709551616",   // one past it
		"1000000000000000000000", // 1000 tokens at 18 decimals
		"115792089237316195423570985008687907853269984665640564039457584007913129639935", // 2^256-1
	} {
		u := Uint256(*uint256.MustFromDecimal(dec))
		b, err := json.Marshal(u)
		if err != nil {
			t.Fatalf("%s: marshal: %v", dec, err)
		}
		if string(b) != `"`+dec+`"` {
			t.Fatalf("%s: got %s", dec, b)
		}
		var back Uint256
		if err := json.Unmarshal(b, &back); err != nil {
			t.Fatalf("%s: unmarshal: %v", dec, err)
		}
		if v := uint256.Int(back); v.Dec() != dec {
			t.Errorf("%s: round-tripped to %s", dec, v.Dec())
		}
	}
}

// The reason this type exists: the narrower one silently cannot hold the value.
func TestUint64CannotHoldWhatUint256Can(t *testing.T) {
	const wei = "1000000000000000000000"
	var small Uint64
	if err := json.Unmarshal([]byte(`"`+wei+`"`), &small); err == nil {
		t.Fatal("Uint64 accepted a 256-bit value")
	}
	var big Uint256
	if err := json.Unmarshal([]byte(`"`+wei+`"`), &big); err != nil {
		t.Fatalf("Uint256 refused it: %v", err)
	}
}

func TestUint256NullLeavesValue(t *testing.T) {
	u := Uint256(*uint256.NewInt(7))
	if err := u.UnmarshalJSON([]byte(Null)); err != nil {
		t.Fatal(err)
	}
	if v := uint256.Int(u); v.Dec() != "7" {
		t.Errorf("null overwrote the value: %s", v.Dec())
	}
}

func TestUint256RejectsGarbage(t *testing.T) {
	var u Uint256
	if err := json.Unmarshal([]byte(`"not-a-number"`), &u); err == nil {
		t.Error("accepted garbage")
	}
}
