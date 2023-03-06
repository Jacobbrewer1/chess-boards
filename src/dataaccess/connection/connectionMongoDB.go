package connection

import (
	"context"
	"fmt"
	"github.com/Jacobbrewer1/chess-boards/src/custom"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type MongoDB struct {
	client      *mongo.Client
	Host        string                     `json:"host,omitempty" yaml:"host,omitempty"`
	Port        string                     `json:"port,omitempty" yaml:"port,omitempty"`
	Database    string                     `json:"database,omitempty" yaml:"database,omitempty"`
	Collections custom.Map[string, string] `json:"collections,omitempty" yaml:"collections,omitempty"`
	sync.RWMutex
}

func (m *MongoDB) Client() *mongo.Client {
	if m.client == nil {
		m.generateClient()
	}

	return m.client
}

func (m *MongoDB) GetCollection(collectionName string) (*mongo.Collection, error) {
	if !m.Collections.Has(collectionName) {
		return nil, fmt.Errorf("collection %s does not exist in config", collectionName)
	}

	coll := m.Collections.Get(collectionName)
	if coll == nil {
		// Should never be hit but just in case
		return nil, fmt.Errorf("collection %s does not exist in config", collectionName)
	}

	return m.Client().Database(m.Database).Collection(*coll), nil
}

func (m *MongoDB) Ping() {
	if m.client == nil {
		m.generateClient()
	}

	m.RLock()
	defer m.RUnlock()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := m.client.Ping(ctx, nil); err != nil {
		m.client = nil
		panic(err)
	}
}

func (m *MongoDB) generateClient() {
	m.RLock()
	defer m.RUnlock()

	if m.Host == "" && m.Port == "" || m.Database == "" {
		panic("invalid connection")
	}

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", m.Host, m.Port))

	var err error
	m.client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		m.client = nil
		panic(err)
	}
}
