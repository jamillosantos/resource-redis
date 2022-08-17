package resourceredis

import (
	"context"

	"github.com/go-redis/redis/v9"
)

type Resource struct {
	redis.UniversalClient

	name   string
	config *PlatformConfig
}

type opts struct {
	name string
}

type Option func(*opts)

func defaultOpts() opts {
	return opts{
		name: "Redis",
	}
}

func WithName(name string) Option {
	return func(o *opts) {
		o.name = name
	}
}

func New(config *PlatformConfig, option ...Option) *Resource {
	opts := defaultOpts()
	for _, o := range option {
		o(&opts)
	}
	return &Resource{
		name:   opts.name,
		config: config,
	}
}

func (r Resource) Name() string {
	return r.name
}

func (r *Resource) Start(ctx context.Context) error {
	clusterOpts := defaultClusterOpts()
	err := applyConfig(clusterOpts, r.config)
	if err != nil {
		return err
	}
	client := redis.NewUniversalClient(&clusterOpts)
	result := client.Ping(ctx)
	if err := result.Err(); err != nil {
		return err
	}
	r.UniversalClient = client
	return nil
}

func applyConfig(clusterOpts redis.UniversalOptions, config *PlatformConfig) error {
	clusterOpts.Addrs = config.Addrs
	clusterOpts.Username = config.Username
	clusterOpts.Password = config.Password

	if config.DialTimeout != 0 {
		clusterOpts.DialTimeout = config.DialTimeout
	}

	if config.ReadTimeout != 0 {
		clusterOpts.ReadTimeout = config.ReadTimeout
	}
	if config.WriteTimeout != 0 {
		clusterOpts.WriteTimeout = config.WriteTimeout
	}

	if config.Pool.Size != 0 {
		clusterOpts.PoolSize = config.Pool.Size
	}
	if clusterOpts.PoolTimeout != 0 {
		clusterOpts.PoolTimeout = config.Pool.Timeout
	}
	if clusterOpts.MinIdleConns != 0 {
		clusterOpts.MinIdleConns = config.Pool.MinIdleConns
	}
	if clusterOpts.MaxIdleConns != 0 {
		clusterOpts.MaxIdleConns = config.Pool.MaxIdleConns
	}
	if clusterOpts.ConnMaxLifetime != 0 {
		clusterOpts.ConnMaxLifetime = config.Pool.ConnMaxLifetime
	}
	if clusterOpts.ConnMaxIdleTime != 0 {
		clusterOpts.ConnMaxIdleTime = config.Pool.ConnMaxIdleTime
	}
	return nil
}

func defaultClusterOpts() redis.UniversalOptions {
	return redis.UniversalOptions{}
}

func (r *Resource) Stop(ctx context.Context) error {
	return r.Close()
}
