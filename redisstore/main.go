package redisstore

import (
	"github.com/boj/redistore"
	gSessions "github.com/gorilla/sessions"
	nSessions "github.com/yinshuwei/negroni-sessions"
)

//New returns a new Redis store
func New(size int, network, address, password string, keyPairs ...[]byte) (nSessions.Store, error) {
	store, err := redistore.NewRediStore(size, network, address, password, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &rediStore{store}, nil
}

type rediStore struct {
	*redistore.RediStore
}

func (c *rediStore) Options(options nSessions.Options) {
	c.RediStore.Options = &gSessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HTTPOnly,
	}
}
