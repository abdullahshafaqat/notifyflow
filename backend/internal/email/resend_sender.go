package email

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const resendEndpoint = "https://api.resend.com/emails"

type resendRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Text    string   `json:"text"`
}

type resendErrorResponse struct {
	Message string `json:"message"`
}

type ResendSender struct {
	apiKey string
	from   string
	http   *http.Client
}

func NewResendSender(apiKey, from string) (*ResendSender, error) {
	if apiKey == "" {
		return nil, errors.New("resend api key is required")
	}
	if from == "" {
		return nil, errors.New("email from is required")
	}

	return &ResendSender{
		apiKey: apiKey,
		from:   from,
		http: &http.Client{
			Timeout: 15 * time.Second,
		},
	}, nil
}

func (s *ResendSender) Send(ctx context.Context, to, subject, body string) error {
	if to == "" {
		return errors.New("recipient address is required")
	}

	payload := resendRequest{
		From:    s.from,
		To:      []string{to},
		Subject: subject,
		Text:    body,
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to encode resend request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, resendEndpoint, bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("failed to build resend request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.http.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call resend: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	var errResp resendErrorResponse
	if decodeErr := json.NewDecoder(resp.Body).Decode(&errResp); decodeErr == nil && errResp.Message != "" {
		return fmt.Errorf("resend api error: %s", errResp.Message)
	}

	return fmt.Errorf("resend api returned status %d", resp.StatusCode)
}
