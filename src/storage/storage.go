package storage

import "testTask/src/request"

type Storage interface {
	Save([]request.SaveFact)
}
