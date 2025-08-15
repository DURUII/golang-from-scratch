package ch00

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

const BaseURL = "https://api.deepseek.com"

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type APIClient struct {
	HTTPClient Client
}

type RequestBody struct {
	Messages []map[string]string `json:"messages"`
	Model    string              `json:"model"`
	Stream   bool                `json:"stream"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found - using system environment variables")
	}
}

func NewAPIClient() *APIClient {
	return &APIClient{
		HTTPClient: http.DefaultClient,
	}
}

func (c *APIClient) SendRequest(reqBody RequestBody, apiKey string) (*http.Response, error) {
	url := BaseURL + "/chat/completions"
	payload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	return c.HTTPClient.Do(req)
}

func TestDeepSeekSSE(t *testing.T) {
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	require.NotEmpty(t, apiKey, "DEEPSEEK_API_KEY environment variable not set")

	client := NewAPIClient()

	reqBody := RequestBody{
		Messages: []map[string]string{
			{"role": "system", "content": "You are a warm supporter"},
			{"role": "user", "content": "我好难受"},
		},
		Model:  "deepseek-chat",
		Stream: true,
	}

	resp, err := client.SendRequest(reqBody, apiKey)
	require.NoError(t, err)
	defer func() {
		if err := resp.Body.Close(); err != nil {
			t.Fatalf("failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("API request failed: %d\n%s", resp.StatusCode, body)
	}

	var fullResponse strings.Builder
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "data:") {
			data := strings.TrimSpace(strings.TrimPrefix(line, "data:"))

			if data == "[DONE]" {
				break
			}

			var chunk struct {
				Choices []struct {
					Delta struct {
						Content string `json:"content"`
					} `json:"delta"`
				} `json:"choices"`
			}

			if err := json.Unmarshal([]byte(data), &chunk); err != nil {
				t.Logf("JSON parse error: %v", err)
				continue
			}

			for _, choice := range chunk.Choices {
				if content := choice.Delta.Content; content != "" {
					fmt.Print(content)
					fullResponse.WriteString(content)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("Error reading response: %v", err)
	}

	output := map[string]string{"response": fullResponse.String()}
	file, err := os.Create("output.json")
	require.NoError(t, err)
	defer func() {
		if err := file.Close(); err != nil {
			t.Fatalf("failed to close file: %v", err)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	require.NoError(t, encoder.Encode(output))

	fmt.Println("\nOutput saved to output.json")
}