package webapi

import "io/ioutil"

const textGeneratorUrl = "https://baconipsum.com/api/?type=all-meat&sentences=1&format=text"

type DummyTextGenerator struct {
	url string
}

func NewDummyTextGenerator() IDummyTextGenerator {
	return &DummyTextGenerator{
		url: textGeneratorUrl,
	}
}

func (s *DummyTextGenerator) Get() (string, error) {
	resp, err := makeGetRequest(s.url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}
