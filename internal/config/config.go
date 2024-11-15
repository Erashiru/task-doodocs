package config

import "flag"

type Config struct {
	Address string
}

func Loader() *Config {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	conf := Config{
		Address: *addr,
	}

	return &conf
}
