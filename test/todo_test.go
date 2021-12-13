//go:build e2e

package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetTodos(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/todo")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostTodo(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"name": "Big Data lecture 2 Assignment", "Description": "Page 21 of lecture 2", "Is_complete": false}`).
		Post(BASE_URL + "/api/todo")
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}
