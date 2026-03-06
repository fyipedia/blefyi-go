// Package blefyi provides a Go client for the BLEFYI API.
//
// BLEFYI is a comprehensive Bluetooth Low Energy reference covering chips,
// profiles, versions, beacons, use cases, and manufacturers. This client
// requires no authentication and has zero external dependencies.
//
// Usage:
//
//	client := blefyi.NewClient()
//	result, err := client.Search("heart rate")
package blefyi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DefaultBaseURL is the default base URL for the BLEFYI API.
const DefaultBaseURL = "https://blefyi.com/api"

// Client is a BLEFYI API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new BLEFYI API client with default settings.
func NewClient() *Client {
	return &Client{
		BaseURL:    DefaultBaseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) get(path string, result interface{}) error {
	resp, err := c.HTTPClient.Get(c.BaseURL + path)
	if err != nil {
		return fmt.Errorf("blefyi: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("blefyi: HTTP %d: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("blefyi: decode failed: %w", err)
	}
	return nil
}

// Search searches across BLE chips, profiles, and glossary terms.
func (c *Client) Search(query string) (*SearchResult, error) {
	var result SearchResult
	path := fmt.Sprintf("/search/?q=%s", url.QueryEscape(query))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Chip returns details for a BLE chip by slug.
func (c *Client) Chip(slug string) (*ChipDetail, error) {
	var result ChipDetail
	if err := c.get("/chip/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Profile returns details for a BLE profile by slug.
func (c *Client) Profile(slug string) (*ProfileDetail, error) {
	var result ProfileDetail
	if err := c.get("/profile/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Version returns details for a BLE version by slug.
func (c *Client) Version(slug string) (*VersionDetail, error) {
	var result VersionDetail
	if err := c.get("/version/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Beacon returns details for a BLE beacon type by slug.
func (c *Client) Beacon(slug string) (*BeaconDetail, error) {
	var result BeaconDetail
	if err := c.get("/beacon/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Usecase returns details for a BLE use case by slug.
func (c *Client) Usecase(slug string) (*UsecaseDetail, error) {
	var result UsecaseDetail
	if err := c.get("/use-case/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Manufacturer returns details for a BLE manufacturer by slug.
func (c *Client) Manufacturer(slug string) (*ManufacturerDetail, error) {
	var result ManufacturerDetail
	if err := c.get("/manufacturer/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GlossaryTerm returns a glossary term by slug.
func (c *Client) GlossaryTerm(slug string) (*GlossaryTerm, error) {
	var result GlossaryTerm
	if err := c.get("/glossary/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Compare compares two BLE chips.
func (c *Client) Compare(slugA, slugB string) (*CompareResult, error) {
	var result CompareResult
	path := fmt.Sprintf("/compare/?a=%s&b=%s", url.QueryEscape(slugA), url.QueryEscape(slugB))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Random returns a random BLE chip.
func (c *Client) Random() (*ChipDetail, error) {
	var result ChipDetail
	if err := c.get("/random/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}
