package mock_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/facktoreal/ip/app/lib/providers/mock"
	"github.com/facktoreal/ip/app/lib/repositories"
)

func TestMock_Health_NewHealthRepository(t *testing.T) {
	r := mock.NewHealthRepository()

	assert.Implements(t, (*repositories.HealthRepository)(nil), r)
}

func TestMock_Health_Check(t *testing.T) {
	r := mock.NewHealthRepository()

	t.Run("Existing key", func(t *testing.T) {
		assert.NoError(t, r.Check(nil))
	})
}
