package api_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vbogretsov/sgmock"
	"github.com/vbogretsov/sgmock/api"
)

func TestAPI(t *testing.T) {
	m := sgmock.New()
	e := api.New(apiKey, m)
	c := client{e: e}

	t.Run("SendReturns401IfNoToken", func(t *testing.T) {
		err := c.send("", msg1)

		require.NotNil(t, err, "expected error bit got nil")
		httpErr, ok := err.(httpError)
		require.True(t, ok, "expected http error bot git %v", err)
		require.Equal(t, http.StatusUnauthorized, httpErr.code)
		require.Equal(t, 0, len(m.List()))
	})
	t.Run("SendReturns401IfTokenInvalid", func(t *testing.T) {
		err := c.send("invalid", msg1)

		require.NotNil(t, err, "expected error bit got nil")
		httpErr, ok := err.(httpError)
		require.True(t, ok, "expected http error bot git %v", err)
		require.Equal(t, http.StatusUnauthorized, httpErr.code)
		require.Equal(t, 0, len(m.List()))
	})
	t.Run("SendAddsMessageToMockIf200", func(t *testing.T) {
		m.Clear()
		defer m.Clear()

		err := c.send(apiKey, msg1)
		require.Nil(t, err)

		require.Equal(t, 1, len(m.List()))
		require.Equal(t, msg1, m.List()[0])
	})
	t.Run("ListReturnsAllMessages", func(t *testing.T) {
		m.Clear()
		defer m.Clear()

		require.Nil(t, c.send(apiKey, msg1))
		require.Nil(t, c.send(apiKey, msg2))

		require.Equal(t, 2, len(m.List()))
		require.Equal(t, msg1, m.List()[0])
		require.Equal(t, msg2, m.List()[1])
	})
	t.Run("ClearClearsAllMessages", func(t *testing.T) {
		m.Clear()
		defer m.Clear()

		require.Nil(t, c.send(apiKey, msg1))
		require.Nil(t, c.send(apiKey, msg2))
		require.Nil(t, c.clear())
		require.Equal(t, 0, len(m.List()))
	})
}
