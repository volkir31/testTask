package storage

import "testTask/src/request"

// Storage хранилище
type Storage interface {
	Save([]request.SaveFact)
}
