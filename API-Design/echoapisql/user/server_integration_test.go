package user

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUser(t *testing.T) {
	seedUser(t)
	var us []User
	res := request(http.MethodGet, uri("users"), nil)
	err := res.Decode(&us)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Greater(t, len(us), 0)

}

func TestCreateUser(t *testing.T) {
	body := bytes.NewBufferString(`{"name":"AnuchitO"}`)

	var c User
	res := request(http.MethodPost, uri("users"), body)
	err := res.Decode(&c)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.Equal(t, "AnuchitO", c.Name)
	assert.Greater(t, c.ID, 0)
}

func TestGetUserByID(t *testing.T) {
	c := seedUser(t)
	var u User
	res := request(http.MethodGet, uri("users", strconv.Itoa(c.ID)), nil)
	err := res.Decode(&u)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, c.ID, u.ID)
	assert.Equal(t, c.Name, u.Name)
}

func uri(paths ...string) string {
	host := "http://localhost:2566"
	if paths == nil {
		return host
	}
	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func seedUser(t *testing.T) User {
	var c User
	body := bytes.NewBufferString(`{"name":"AnuchitO"}`)
	err := request(http.MethodPost, uri("users"), body).Decode(&c)
	if err != nil {
		t.Error(err)
	}
	return c
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	return json.NewDecoder(r.Body).Decode(v)

}
func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Authorization",)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	return &Response{res, err}
}
