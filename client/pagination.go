package client

import (
	"bytes"
	"encoding/json"
	"net/url"
)

// listPage is the envelope paginated list endpoints return.
type listPage struct {
	Next     string          `json:"next"`
	Previous string          `json:"previous"`
	Count    int             `json:"count"`
	Results  json.RawMessage `json:"results"`
}

// decodeList unmarshals a bare JSON array or a {results, next} envelope into v
// and returns the next-page link, if any.
func decodeList(body []byte, v interface{}) (next string, err error) {
	trimmed := bytes.TrimLeft(body, " \t\r\n")
	if len(trimmed) > 0 && trimmed[0] == '{' {
		var p listPage
		if err := json.Unmarshal(body, &p); err != nil {
			return "", err
		}
		if len(p.Results) == 0 {
			return "", nil
		}
		return p.Next, json.Unmarshal(p.Results, v)
	}
	return "", json.Unmarshal(body, v)
}

// nextPagePath keeps only the path and query of a "next" link so it resolves
// against the client's base URL.
func nextPagePath(next string) (string, error) {
	u, err := url.Parse(next)
	if err != nil {
		return "", err
	}
	path := u.Path
	if u.RawQuery != "" {
		path += "?" + u.RawQuery
	}
	return path, nil
}

// getAll fetches every page of a list endpoint, passing each raw page array to appendPage.
func (c *Client) getAll(path string, appendPage func(results []byte) error) error {
	seen := map[string]bool{}
	for path != "" && !seen[path] {
		seen[path] = true
		body, err := c.newRequestDo("GET", path, nil)
		if err != nil {
			return err
		}
		var raw json.RawMessage
		next, err := decodeList(body.BodyBytes, &raw)
		if err != nil {
			return err
		}
		if len(raw) > 0 {
			if err := appendPage(raw); err != nil {
				return err
			}
		}
		if next == "" {
			return nil
		}
		if path, err = nextPagePath(next); err != nil {
			return err
		}
	}
	return nil
}
