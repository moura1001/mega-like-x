package webserver

import "moura1001/mega_like_x/src/app/model"

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
