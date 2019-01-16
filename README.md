# Pungi ![](https://img.shields.io/github/license/lzakharov/pungi.svg)

Tiny Go struct validator.

## Examples

```go
package main

import (
	"fmt"
	"github.com/lzakharov/pungi"
)

type Config struct {
	Token   string
	Offset  int `pungi:"nullable"`
	Timeout int
	Proxy   *ProxyConfig `pungi:"nullable"`
}

type ProxyConfig struct {
	URL      string
	Username string `pungi:"nullable"`
	Password string `pungi:"nullable"`
}

func main() {
	validConfig := &Config{
		Token:   "TOKEN",
		Offset:  0,
		Timeout: 60,
		Proxy: &ProxyConfig{
			URL:      "127.0.0.1:9050",
		},
	}

	invalidConfig := &Config{
		Token:   "", // empty string!
		Offset:  0,
		Timeout: 60,
		Proxy: &ProxyConfig{
			URL:      "127.0.0.1:9050",
		},
	}


	fmt.Println(pungi.IsValid(validConfig)) // nil
	fmt.Println(pungi.IsValid(invalidConfig).Error()) // 'Config.Token' has a zero value
}
```

## Issue tracker
Please report any bugs and enhancement ideas using the **Pungi** issue tracker:

https://github.com/lzakharov/pungi/issues

Feel free to also ask questions on the tracker.

## License

Copyright (c) 2019 Lev Zakharov. Licensed under [the MIT License](https://raw.githubusercontent.com/lzakharov/pungi/master/LICENSE).