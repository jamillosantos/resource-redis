package resourceredis

import (
	"time"
)

type PlatformConfigPool struct {
	Size            int           `config:"size"`
	MinIdleConns    int           `config:"min_idle_conns"`
	MaxIdleConns    int           `config:"max_idle_conns"`
	Timeout         time.Duration `config:"pool_timeout"`
	ConnMaxLifetime time.Duration `config:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `config:"conn_max_idle_time"`
}

type PlatformConfig struct {
	Addrs    []string `config:"addrs,secret"`
	Username string   `config:"username,secret"`
	Password string   `config:"password,secret"`

	DialTimeout  time.Duration      `config:"dial_timeout"`
	ReadTimeout  time.Duration      `config:"read_timeout"`
	WriteTimeout time.Duration      `config:"write_timeout"`
	Pool         PlatformConfigPool `config:"pool"`
}
