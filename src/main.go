package main

import (
	"math/rand"
	"os"
	"testTask/src/misc"
	"testTask/src/request"
	"testTask/src/storage"
	"time"
)

func main() {
	httpStorage := storage.HttpStorage{Token: os.Getenv("TOKEN")}
	httpStorage.Save(getSaveFacts())
}

func getSaveFacts() []request.SaveFact {
	var saveFacts []request.SaveFact
	for i := 0; i < 1; i++ {
		saveFacts = append(saveFacts, request.SaveFact{
			GetFact: request.GetFact{
				PeriodStart:     time.Now(),
				PeriodEnd:       time.Now(),
				PeriodKey:       misc.MONTH,
				IndicatorToMoId: rand.Int(),
			},
			IndicatorToMoFactId: rand.Int(),
			Value:               rand.Int(),
			FactTime:            time.Now(),
			IsPlan:              rand.Int()%2 == 0,
			AuthUserId:          rand.Int(),
			Comment:             "buffer Mikhailichenko",
		})
	}

	return saveFacts
}
