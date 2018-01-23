package data

import (
	"testing"
)

//TODO: add tests that actually test error states and handling

// setup some test data that can be reused for both the set and get tests
var cases = []struct {
	role, accNum, arn string
}{
	{"RoleName1", "123456", "arn:aws:iam::123456:role/RoleName1"},
	{"RoleName2", "654321", "arn:aws:iam::654321:role/RoleName2"},
}

func TestSetARNFromParts(t *testing.T) {
	for _, c := range cases {
		ar := AssumableRole{}
		ar.SetARNFromParts(c.role, c.accNum)
		if ar.RoleARN != c.arn {
			t.Errorf("Generated %s from r: %s, a: %s, wanted %s", ar.RoleARN, c.role, c.accNum, c.arn)
		}
	}
}

func TestGetPartsFromARN(t *testing.T) {
	for _, c := range cases {
		ar := AssumableRole{RoleARN: c.arn}
		role, accNum, _ := ar.GetPartsFromARN()
		if role != c.role || accNum != c.accNum {
			t.Errorf("Found r: %s, a: %s from %s; expected r: %s, a: %s", role, accNum, ar.RoleARN, c.role, c.accNum)
		}
	}
}
