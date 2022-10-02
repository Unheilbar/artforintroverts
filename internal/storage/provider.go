package storage

import (
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
	return &Provider{
		collection: collection,
		cache: &cache{
			mx: &sync.Mutex{},
		},
	}
}

func (p *Provider) RefreshCache() error {
	if p.cache.IsValid() {
		return nil
	}
	p.cache.lock()
	defer p.cache.unlock()
	users, err := p.Storage().GetAll()
	if err != nil {
		return err
	}

	p.cache.refresh(users)
	return nil
}
