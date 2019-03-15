package storage

import "requestbingo/models"

type MemoryStorage struct {
	BinTtl       int
	Bins         map[string]models.Bin
	RequestCount int
}

func create() {

}

func count() {

}

func lookup() {

}

func expire() {

}
