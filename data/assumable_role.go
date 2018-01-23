package data

import (
	"fmt"
	"regexp"
)

// arnRegex will find the account number and role in an IAM Role ARN
var arnRegex = regexp.MustCompile(`arn:aws:iam::(?P<accNum>\d+):role/(?P<role>.+)`)

// AssumableRole contains the ARN of a role that can be assumed and a name to more easily identify it
type AssumableRole struct {
	Name    string
	RoleARN string
}

// SetARNFromParts will build and set an ARN from the role name and account number
func (r *AssumableRole) SetARNFromParts(roleName, accountNum string) {
	r.RoleARN = fmt.Sprintf("arn:aws:iam::%s:role/%s", accountNum, roleName)
}

// GetPartsFromARN will get the role name and account number parts out of the ARN
func (r *AssumableRole) GetPartsFromARN() (roleName, accountNum string, err error) {
	//TODO: add empty string and no match checks

	// make sure we have an ARN to work with

	// run the match
	match := arnRegex.FindStringSubmatch(r.RoleARN)

	// convert into a map so that we can access the capture groups by name
	result := make(map[string]string)
	for i, name := range arnRegex.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}

	// get the values from the map
	roleName = result["role"]
	accountNum = result["accNum"]

	return
}
