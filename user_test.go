package snakecasefields

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	usersvr "github.com/akm/goa_v2_post_payload_including_snake_case_fields/gen/http/user/server"
	user "github.com/akm/goa_v2_post_payload_including_snake_case_fields/gen/user"
	goahttp "goa.design/goa/http"
	"goa.design/goa/http/middleware"
)

func TestUserCreate(t *testing.T) {
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
		dec := goahttp.RequestDecoder
		enc := goahttp.ResponseEncoder
		logger := log.New(os.Stderr, "LOG: ", log.Lshortfile)
		userSvc := NewUser(logger)
		userEndpoints := user.NewEndpoints(userSvc)
		eh := errorHandler(logger)
		userServer := usersvr.New(userEndpoints, mux, dec, enc, eh)
		usersvr.Mount(mux, userServer)
	}

	testUser := &TestUser{
		Firstname: "Foo",
		Lastname:  "Bar",
	}
	userJson, err := json.Marshal(testUser)
	if err != nil {
		t.Errorf("Failed to json.Marshal %v because of %v\n", testUser, err)
		return
	}

	w := httptest.NewRecorder()
	{
		b := bytes.NewBuffer(userJson)
		path := "/users"
		r, err := http.NewRequest("POST", path, b)
		if err != nil {
			t.Errorf("Failed to http.NewRequest because of %v\n", err)
			return
		}
		r.RequestURI = path
		r.URL = &url.URL{Path: path}
		mux.ServeHTTP(w, r)
	}

	w.Flush()
	if w.Code != http.StatusCreated {
		t.Errorf("Status code must be %v but was %v. Response body: %s\n", http.StatusCreated, w.Code, w.Body.String())
		return
	}

	var userRes TestUser
	if err := json.Unmarshal(w.Body.Bytes(), &userRes); err != nil {
		t.Errorf("Failed to Unmarshal because of %v\n", err)
		return
	}

	if userRes.Firstname != testUser.Firstname {
		t.Errorf("Firstname must be %v but was %v\n", testUser.Firstname, userRes.Firstname)
	}
	if userRes.Lastname != testUser.Lastname {
		t.Errorf("Lastname must be %v but was %v\n", testUser.Lastname, userRes.Lastname)
	}
}

type TestUser struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
