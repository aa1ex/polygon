package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func WaitElasticReady(elasticURL, username, password string, retries int, delay time.Duration) error {
	for i := 0; i < retries; i++ {
		req, err := http.NewRequest("GET", elasticURL, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		if username != "" && password != "" {
			req.SetBasicAuth(username, password)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			time.Sleep(delay)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			return nil
		}

		time.Sleep(delay)
	}

	return fmt.Errorf("elasticsearch is not ready after %d retries", retries)
}

func CreateIngestPipeline(elasticURL, pipelineID, username, password string) error {
	url := fmt.Sprintf("%s/_ingest/pipeline/%s", elasticURL, pipelineID)

	pipelineBody := map[string]interface{}{
		"description": "test ingest pipeline",
		"processors": []interface{}{
			map[string]interface{}{
				"set": map[string]interface{}{
					"field": "processed_at",
					"value": "{{_ingest.timestamp}}",
				},
			},
		},
	}

	body, err := json.Marshal(pipelineBody)
	if err != nil {
		return fmt.Errorf("failed to marshal body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	client := &http.Client{Timeout: time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make HTTP request: %w", err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body response: %w", err)
	}
	_ = resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status: %d, body: %s", resp.StatusCode, string(respBody))
	}

	return nil
}
