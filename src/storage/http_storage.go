package storage

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"testTask/src/request"
)

type HttpStorage struct {
	Token string
}

var waitGroup = &sync.WaitGroup{}
var httpClient = &http.Client{}

func (s *HttpStorage) Save(facts []request.SaveFact) {
	buffer := make(chan request.SaveFact, 1)
	go s.produce(buffer, facts)
	waitGroup.Add(1)
	go s.consume(buffer)
	waitGroup.Add(1)
	waitGroup.Wait()
}

func (s *HttpStorage) produce(buffer chan<- request.SaveFact, facts []request.SaveFact) {
	for _, fact := range facts {
		buffer <- fact
	}
	close(buffer)
	waitGroup.Done()
}

func (s *HttpStorage) consume(channel chan request.SaveFact) {
	for fact := range channel {
		if err := s.send(fact); err != nil {
			fmt.Println("error sending fact:", err)
			continue
		}
		s.checkFact(fact.GetFact)
	}
	waitGroup.Done()
}

func (s *HttpStorage) send(fact request.SaveFact) error {
	resp, err := s.sendRequest(&fact)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	r, _ := io.ReadAll(resp.Body)
	fmt.Println(string(r))
	if resp.StatusCode != 200 {
		return fmt.Errorf("received status code when sending fact %d: %s", resp.StatusCode, r)
	}

	return nil
}

func (s *HttpStorage) checkFact(fact request.GetFact) error {
	resp, err := s.sendRequest(&fact)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	r, _ := io.ReadAll(resp.Body)
	fmt.Println(string(r))
	if resp.StatusCode != 200 {
		return fmt.Errorf("received status code when check fact %d: %s", resp.StatusCode, r)
	}

	return nil
}

func (s *HttpStorage) sendRequest(fact request.Request) (*http.Response, error) {
	b := url.Values(fact.ToFormData()).Encode()
	fullUrl := fmt.Sprintf("%s?%s", fact.Url(), b)
	req, err := http.NewRequest("POST", fullUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Token))
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
