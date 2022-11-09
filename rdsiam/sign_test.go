package rdsiam

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractRegion(t *testing.T) {

	type test struct {
		input    string
		expected string
	}

	tests := []test{
		{input: "some-rds-proxy.proxy-withid.eu-west-1.rds.amazonaws.com", expected: "eu-west-1"},
		{input: "some-other-url.com", expected: ""},
	}

	for _, tc := range tests {
		region := extractRegion(tc.input)
		assert.Equal(t, tc.expected, region)
	}
}
