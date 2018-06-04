package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAliases(t *testing.T) {

	tt := []struct {
		Path  string
		Alias string
	}{
		{Path: "secretbox/*/production", Alias: "alias/secretbox/_all_/production"},
		{Path: "secretbox/test/ci", Alias: "alias/secretbox/test/ci"},
	}

	for _, tc := range tt {
		assert.Equal(t, aliasPath(tc.Path), tc.Alias)
		assert.Equal(t, unaliasPath(tc.Alias), tc.Path)
	}
}
