package elasticsearch

import "testing"

func TestCreateIngestPipeline(t *testing.T) {
	host := "http://localhost:9200"
	pipeline := "test"
	username := "elastic"
	password := "elastic"
	err := CreateIngestPipeline(host, pipeline, username, password)
	if err != nil {
		t.Fatal(err)
	}
}
