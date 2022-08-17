package resourceredis

import (
	"context"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResource_Start(t *testing.T) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "failed connecting to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("redis", "7.0.4-alpine3.16", []string{})
	require.NoError(t, err, "failed starting redis")
	t.Cleanup(func() {
		resource.Close()
	})

	rsc := New(&PlatformConfig{
		Addrs:    []string{resource.GetHostPort("6379/tcp")},
		Username: "",
		Password: "",
	})

	require.Eventuallyf(t, func() bool {
		ctx := context.Background()
		err := rsc.Start(ctx)
		return assert.NoError(t, err)
	}, 10*time.Second, 1*time.Second, "redis is not ready")

	require.NoError(t, rsc.Close(), "failed closing resource")
}
