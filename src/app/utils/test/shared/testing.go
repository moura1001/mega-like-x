package utilstesting

import (
	"fmt"
	"io"
	"moura1001/mega_like_x/src/app/model"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubGameStore struct {
	likes     map[string]int
	likeCalls []string
	polling   model.Polling
}

func (s *StubGameStore) GetGameLikes(name string) int {
	return s.likes[name]
}

func (s *StubGameStore) RecordLike(name string) {
	s.likeCalls = append(s.likeCalls, name)
}

func (s *StubGameStore) GetPolling() model.Polling {
	return s.polling
}

func GetNewStubGameStore(likes map[string]int, likeCalls []string, polling model.Polling) StubGameStore {
	return StubGameStore{likes, likeCalls, polling}
}

func AssertPolling(t *testing.T, got, want model.Polling) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertLikesValue(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct likes, got %d, want %d", got, want)
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("didnt expect an error but got one: %v", err)
	}
}

func AssertGameLike(t *testing.T, store *StubGameStore, game string) {
	t.Helper()

	if len(store.likeCalls) != 1 {
		t.Errorf("got %d calls to RecordLike, want %d", len(store.likeCalls), 1)
	}

	if store.likeCalls[0] != game {
		t.Errorf("did not store correct liked game, got '%s' want '%s'", store.likeCalls[0], game)
	}
}

func NewGetLikesRequest(game string) *http.Request {
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/likes/%s", game), nil)
	return req
}

func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body error, got '%s', want '%s'", got, want)
	}
}

func AssertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got '%d', want '%d'", got, want)
	}
}

func NewPostLikeRequest(game string) *http.Request {
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/likes/%s", game), nil)
	return req
}

func NewGetPollingRequest() *http.Request {
	req := httptest.NewRequest(http.MethodGet, "/games", nil)
	return req
}

func GetPollingFromResponse(t *testing.T, body io.Reader) (polling model.Polling) {
	t.Helper()

	polling, err := model.NewGamePolling(body)
	if err != nil {
		t.Fatalf("unable to parse response from server '%v' to get polling: '%v'", body, err)
	}

	return
}

const JsonContentType = "application/json"

func AssertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Body)
	}
}
