package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	hiroTestnet = "https://api.testnet.hiro.so"
	hiroMainnet = "https://api.hiro.so"
)

type HiroClient struct {
	baseURL    string
	signerURL  string
	httpClient *http.Client
}

func NewHiroClient(env string, signerURL string) *HiroClient {
	baseURL := hiroTestnet
	if env == "mainnet" {
		baseURL = hiroMainnet
	}
	return &HiroClient{
		baseURL:    baseURL,
		signerURL:  signerURL,
		httpClient: http.DefaultClient,
	}
}

func (c *HiroClient) get(ctx context.Context, url string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("hiro status %d: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("decode: %w", err)
	}
	return nil
}

// Fetching contract events with offset= and limit=
func (c *HiroClient) FetchContractEvents(ctx context.Context, contractID string, offset, limit int) (*HiroEventsResponse, error) {
	url := fmt.Sprintf("%s/extended/v1/contract/%s/events?offset=%d&limit=%d",
		c.baseURL, contractID, offset, limit)

	var result HiroEventsResponse
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetching transaction info
func (c *HiroClient) FetchTxInfo(ctx context.Context, txID string) (*HiroTxResponse, error) {
	url := fmt.Sprintf("%s/extended/v1/tx/%s", c.baseURL, txID)

	var result HiroTxResponse
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// For checking is already relayed
func (c *HiroClient) ReadOnlyCall(ctx context.Context, contractAddr, contractName, functionName string, args []string) (*ReadOnlyCallResponse, error) {
	url := fmt.Sprintf("%s/v2/contracts/call-read/%s/%s/%s", c.baseURL, contractAddr, contractName, functionName)

	requestBody := ReadOnlyCallRequest{
		Sender:    contractAddr,
		Arguments: args,
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http post failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	respStr := string(respBody)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Hiro call-read failed: HTTP %d - %s (url: %s)", resp.StatusCode, respStr, url)
	}

	var result ReadOnlyCallResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %w - raw response: %s", err, respStr[:300])
	}

	return &result, nil
}

// For error handling
func (c *HiroClient) BroadcastTx(ctx context.Context, txBytes []byte) (string, error) {
	url := fmt.Sprintf("%s/v2/transactions", c.baseURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(txBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res struct {
		TxID  string `json:"txid"`
		Error string `json:"error"`
		Code  string `json:"code"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	if res.Error != "" {
		return "", fmt.Errorf("broadcast error: %s (code: %s)", res.Error, res.Code)
	}

	return res.TxID, nil
}

// For sign transactions at stacks
func (c *HiroClient) CallSigner(ctx context.Context, contract string, function string, args []map[string]any) (string, error) {
	if c.signerURL == "" {
		return "", fmt.Errorf("signerURL is empty - set STACKS_SIGNER_URL env var")
	}

	parts := strings.Split(contract, ".")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid contract format, need addr.name")
	}

	payload := map[string]any{
		"contractAddress": parts[0],
		"contractName":    parts[1],
		"functionName":    function,
		"functionArgs":    args,
	}

	body, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/sign-contract-call", c.signerURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("signer request failed: %w (url: %s)", err, url)
	}
	defer resp.Body.Close()

	var res struct {
		TxID  string `json:"txid"`
		Error string `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		bodyBytes, _ := io.ReadAll(resp.Body) // на случай ошибки
		return "", fmt.Errorf("signer decode failed: %w - raw: %s", err, string(bodyBytes))
	}

	if res.Error != "" {
		return "", fmt.Errorf("signer error: %s", res.Error)
	}

	return res.TxID, nil
}
