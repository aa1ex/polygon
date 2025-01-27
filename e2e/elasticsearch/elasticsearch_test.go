package elasticsearch

import (
	"testing"
	"time"
)

func TestCreateIngestPipeline(t *testing.T) {
	host := "http://elasticsearch:9200"
	pipeline := "test"
	username := "elastic"
	password := "elastic"
	err := WaitElasticReady(host, username, password, 3, time.Second*10)
	if err != nil {
		t.Fatal(err)
	}
	err = CreateIngestPipeline(host, pipeline, username, password)
	if err != nil {
		t.Fatal(err)
	}
}
