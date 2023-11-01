// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Platform/internal/sdk/pkg/models/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

type DeleteNpaRulesIDRequest struct {
	// npa policy id
	ID int `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteNpaRulesIDRequest) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

type DeleteNpaRulesID200ApplicationJSONStatus string

const (
	DeleteNpaRulesID200ApplicationJSONStatusSuccess DeleteNpaRulesID200ApplicationJSONStatus = "success"
	DeleteNpaRulesID200ApplicationJSONStatusError   DeleteNpaRulesID200ApplicationJSONStatus = "error"
)

func (e DeleteNpaRulesID200ApplicationJSONStatus) ToPointer() *DeleteNpaRulesID200ApplicationJSONStatus {
	return &e
}

func (e *DeleteNpaRulesID200ApplicationJSONStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteNpaRulesID200ApplicationJSONStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteNpaRulesID200ApplicationJSONStatus: %v", v)
	}
}

// DeleteNpaRulesID200ApplicationJSON - successful operation
type DeleteNpaRulesID200ApplicationJSON struct {
	Data   *shared.NpaPolicyResponseItem             `json:"data,omitempty"`
	Status *DeleteNpaRulesID200ApplicationJSONStatus `json:"status,omitempty"`
}

func (o *DeleteNpaRulesID200ApplicationJSON) GetData() *shared.NpaPolicyResponseItem {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *DeleteNpaRulesID200ApplicationJSON) GetStatus() *DeleteNpaRulesID200ApplicationJSONStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteNpaRulesIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful operation
	DeleteNpaRulesID200ApplicationJSONObject *DeleteNpaRulesID200ApplicationJSON
	// Invalid request
	NpaPolicyResponse400 *shared.NpaPolicyResponse400
}

func (o *DeleteNpaRulesIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteNpaRulesIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteNpaRulesIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteNpaRulesIDResponse) GetDeleteNpaRulesID200ApplicationJSONObject() *DeleteNpaRulesID200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteNpaRulesID200ApplicationJSONObject
}

func (o *DeleteNpaRulesIDResponse) GetNpaPolicyResponse400() *shared.NpaPolicyResponse400 {
	if o == nil {
		return nil
	}
	return o.NpaPolicyResponse400
}
