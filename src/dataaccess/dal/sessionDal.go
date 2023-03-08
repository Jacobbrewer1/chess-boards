package dal

import (
	"context"
	"encoding/json"
	"github.com/Jacobbrewer1/chess-boards/src/session"
	"github.com/go-redis/redis"
	"time"
)

type ISessionDal interface {
	SaveSession(toSave session.Session) error
	GetSession(key string) (session.Session, error)
}

type sessionDal struct {
	ctx    context.Context
	client *redis.Client
}

func (s *sessionDal) SaveSession(toSave session.Session) error {
	data, err := json.Marshal(toSave)
	if err != nil {
		return err
	}

	timeout := session.Timeout.Seconds()

	if err := s.client.WithContext(s.ctx).Set(toSave.Key, data, time.Duration(timeout)*time.Second).Err(); err != nil {
		return err
	}

	return nil
}

func (s *sessionDal) GetSession(key string) (session.Session, error) {
	data, err := s.client.Get(key).Result()
	if err != nil {
		return session.Session{}, err
	}

	var sess session.Session
	if err := json.Unmarshal([]byte(data), &sess); err != nil {
		return session.Session{}, err
	}

	return sess, nil
}

func NewSessionDal(database int) ISessionDal {
	return NewSessionDalWithContext(context.Background(), database)
}

func NewSessionDalWithContext(ctx context.Context, database int) ISessionDal {
	return &sessionDal{
		ctx:    ctx,
		client: Connections.RedisDb().Client(database),
	}
}
