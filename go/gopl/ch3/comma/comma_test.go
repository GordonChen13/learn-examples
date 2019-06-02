package comma

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComma(t *testing.T) {
	cs := []struct {
		name string
		input string
		want string
	}{
		{"return 100 when input 100", "100", "100"},
		{"return 999,999 when input 99999", "999999", "999,999"},
		{"return 100,000,000 when input 100000000", "100000000", "100,000,000"},
		}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			get := Comma(c.input)
            assert.Equal(t, c.want, get)
		})
	}
}

