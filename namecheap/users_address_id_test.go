package namecheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddressID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		id    AddressID
		isSet bool
		value int
	}{
		{"zero value is not set", AddressID{}, false, 0},
		{"zero is a set id", NewAddressID(0), true, 0},
		{"positive id", NewAddressID(777), true, 777},
		{"negative id is kept for the method to reject", NewAddressID(-1), true, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.isSet, tt.id.IsSet())
			assert.Equal(t, tt.value, tt.id.Int())
		})
	}

	t.Run("an unset id differs from a set zero id", func(t *testing.T) {
		t.Parallel()
		assert.NotEqual(t, AddressID{}, NewAddressID(0))
	})
}
