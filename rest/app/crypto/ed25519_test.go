package crypto_test

import (
	"gin-api/rest/app/crypto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeyPair(t *testing.T) {
	pub, priv, err := crypto.GenerateKeyPair()
	assert.NoError(t, err)
	t.Logf("%s", pub)
	t.Logf("%s", priv)
}