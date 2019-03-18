package storage

import (
	"requestbingo/models"
	"time"
)

type MemoryStorage struct {
	BinTtl       int
	Bins         map[string]*models.Bin
	RequestCount int
}

const INTERVAL int = 3600

func CreateBin(storage *MemoryStorage, p bool) *models.Bin {
	bin := models.CreateBin(p)
	storage.Bins[bin.Name] = bin
	return bin
}

func CreateRequest(storage *MemoryStorage, p bool) *models.Request {
	request := models.CreateRequest()
	storage.RequestCount += 1
	return request
}

func CountBins(storage *MemoryStorage) int {
	return len(storage.Bins)
}

func LookupBin(storage *MemoryStorage, name string) *models.Bin {
	return storage.Bins[name]
}

func expireBins(storage *MemoryStorage) {
	expiry := time.Now().Second() - storage.BinTtl
	for name, bin := range storage.Bins {
		if bin.Created.Second() < expiry {
			delete(storage.Bins, name)
		}
	}
}
