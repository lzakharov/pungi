package pungi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		value := struct {
			Bool                 bool
			NullableBool         bool `pungi:"nullable"`
			Int                  int
			NullableInt          uint `pungi:"nullable"`
			Uint                 uint
			NullableUint         uint `pungi:"nullable"`
			Float32              float32
			NullableFloat32      float32 `pungi:"nullable"`
			String               string
			NullableString       string `pungi:"nullable"`
			StructPtr            *struct{ String string }
			NullableStructPtr    *struct{ String string } `pungi:"nullable"`
			NullableNilStructPtr *struct{ String string } `pungi:"nullable"`
		}{
			Bool:                 true,
			NullableBool:         false,
			Int:                  1,
			NullableInt:          0,
			Uint:                 1,
			NullableUint:         0,
			Float32:              1,
			NullableFloat32:      0,
			String:               "s",
			NullableString:       "",
			StructPtr:            &struct{ String string }{String: "s"},
			NullableStructPtr:    &struct{ String string }{String: "s"},
			NullableNilStructPtr: nil,
		}
		err := IsValid(value)
		require.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		type ProxyConfig struct {
			URL      string
			Username string `pungi:"nullable"`
			Password string `pungi:"nullable"`
		}

		type Config struct {
			Token   string
			Offset  int `pungi:"nullable"`
			Timeout int
			Proxy   *ProxyConfig `pungi:"nullable"`
		}

		invalidConfig := &Config{
			Token:   "",
			Offset:  0,
			Timeout: 60,
			Proxy: &ProxyConfig{
				URL: "127.0.0.1:9050",
			},
		}
		expected := "'Config.Token' has zero value"

		err := IsValid(invalidConfig)
		require.Error(t, err)
		require.Equal(t, expected, err.Error())
	})
}
