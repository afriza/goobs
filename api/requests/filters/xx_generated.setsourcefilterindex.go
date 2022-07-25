// This file has been automatically generated. Don't edit it.

package filters

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSourceFilterIndexParams represents the params body for the "SetSourceFilterIndex" request.
Sets the index position of a filter on a source.
*/
type SetSourceFilterIndexParams struct {
	requests.ParamsBasic

	// New index position of the filter
	FilterIndex float64 `json:"filterIndex,omitempty"`

	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "SetSourceFilterIndex".
func (o *SetSourceFilterIndexParams) GetSelfName() string {
	return "SetSourceFilterIndex"
}

/*
SetSourceFilterIndexResponse represents the response body for the "SetSourceFilterIndex" request.
Sets the index position of a filter on a source.
*/
type SetSourceFilterIndexResponse struct {
	requests.ResponseBasic
}

// SetSourceFilterIndex sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceFilterIndex(params *SetSourceFilterIndexParams) (*SetSourceFilterIndexResponse, error) {
	data := &SetSourceFilterIndexResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
