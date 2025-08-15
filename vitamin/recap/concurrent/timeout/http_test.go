package timeout

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestHTTPClientSuccess(t *testing.T) {
	// service setup
	go func() {
		_ = NewRouter().Run(":8080")
	}()
	time.Sleep(50 * time.Millisecond)

	// request
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"http://127.0.0.1:8080/work?t=100", nil)
	require.NoError(t, err)

	// response
	resp, err := (&http.Client{Timeout: 500 * time.Millisecond}).Do(req)
	require.NoError(t, err)

	var res Response
	err = json.NewDecoder(resp.Body).Decode(&res)
	require.NoError(t, err)
	require.Equal(t, Response{Msg: "success"}, res)
}

func TestHTTPClientDeadlineExceeded(t *testing.T) {
	// service setup
	go func() {
		_ = NewRouter().Run(":8080")
	}()

	// request
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"http://127.0.0.1:8080/work?t=800", nil)
	require.NoError(t, err)

	// 超时
	_, err = (&http.Client{Timeout: 1 * time.Millisecond}).Do(req)
	require.Error(t, err)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestHTTPClientContextCanceled(t *testing.T) {
	// service setup
	go func() {
		_ = NewRouter().Run(":8080")
	}()
	time.Sleep(50 * time.Millisecond)

	// request
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(350 * time.Millisecond)
		cancel()
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"http://127.0.0.1:8080/work?t=800", nil)
	require.NoError(t, err)

	// response
	_, err = (&http.Client{Timeout: 500 * time.Millisecond}).Do(req)
	require.Error(t, err)
	assert.ErrorIs(t, err, context.Canceled)
}
