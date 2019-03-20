package storage

import (
	"requestbingo/models"
)

type MemoryStorage struct {
	BinTtl       int
	Bins         map[string]*models.Bin
	RequestCount int
}

var Mem *MemoryStorage

func init() {
	Mem = new(MemoryStorage)
	Mem.Bins = make(map[string]*models.Bin)
}

func CreateBin(bin *models.Bin) *models.Bin {
	Mem.Bins[bin.Name] = bin
	return bin
}

func CreateRequest(bin *models.Bin, request *models.Request) *models.Request {
	bin.Requests = append(bin.Requests, *request)
	Mem.RequestCount += 1
	return request
}

func CountBins() int {
	return len(Mem.Bins)
}

func LookupBin(name string) *models.Bin {
	return Mem.Bins[name]
}
