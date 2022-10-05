package storage

import (
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type Provider struct {
	collection *mongo.Collection
	cache      *cache
}

func (p *Provider) Cache() *cache {
	return p.cache
}

func (p *Provider) Storage() *persistentStorage {
	return &persistentStorage{p.collection}
}

func NewStorageProvider(client *mongo.Client, dbName string, collectionName string) *Provider {
	collection := client.Database(dbName).Collection(collectionName)

	provider := &Provider{
		collection: collection,
	}

	provider.initCache()

	return provider
}

func (p *Provider) initCache() {
	users, err := p.Storage().GetAll()
	if err != nil {
		fmt.Printf("err accured during cache init operation %s", err)
	}

	cap := uint(len(users))
	cache := &cache{
		mx:       &sync.Mutex{},
		capacity: cap,
		users:    make(map[string]*node, cap),
		queue:    &linkedList{size: 0},
	}

	for _, user := range users {
		cache.Set(user)
	}

	p.cache = cache
}
