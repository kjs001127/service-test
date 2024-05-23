package dial_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/channel-io/ch-app-store/lib/dial"
)

func openLocalServer() {
	go func() {
		if err := gin.New().Run(":3020"); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Millisecond)
}

func TestDialLoopback(t *testing.T) {
	openLocalServer()

	localAddr := "127.0.0.1:3020"
	dialer := dial.NewIPFilteringWrapper(&net.Dialer{
		Timeout: time.Second * 5,
	})

	_, err := dialer.DialContext(context.Background(), "tcp", localAddr)
	assert.ErrorIs(t, err, dial.ErrNotPermitted)
}

func TestDialLocalhost(t *testing.T) {
	localAddr := "localhost:3020"
	dialer := dial.NewIPFilteringWrapper(&net.Dialer{
		Timeout: time.Second * 5,
	})

	_, err := dialer.DialContext(context.Background(), "tcp", localAddr)
	assert.ErrorIs(t, err, dial.ErrNotPermitted)
}

func TestDialPublicSvc(t *testing.T) {
	addr := "channel.io:80"
	dialer := dial.NewIPFilteringWrapper(&net.Dialer{
		Timeout: time.Second * 5,
	})

	conn, err := dialer.DialContext(context.Background(), "tcp", addr)
	assert.NoError(t, err)
	assert.NotNil(t, conn)
}

/*
func TestDialInternalSvc(t *testing.T) {
	addr := "admin-api.in.exp.channel.io:80"
	dialer := dial.NewIPFilteringWrapper(&net.Dialer{
		Timeout: time.Second * 5,
	})

	_, err := dialer.DialContext(context.Background(), "tcp", addr)
	assert.ErrorIs(t, err, dial.ErrNotPermitted)
}
*/
