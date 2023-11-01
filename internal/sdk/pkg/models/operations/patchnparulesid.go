// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Platform/internal/sdk/pkg/models/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

// PatchNpaRulesIDSilent - flag to skip output except status code
type PatchNpaRulesIDSilent string

const (
	PatchNpaRulesIDSilentOne  PatchNpaRulesIDSilent = "1"
	PatchNpaRulesIDSilentZero PatchNpaRulesIDSilent = "0"
)

func (e PatchNpaRulesIDSilent) ToPointer() *PatchNpaRulesIDSilent {
	return &e
}

func (e *PatchNpaRulesIDSilent) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "1":
		fallthrough
	case "0":
		*e = PatchNpaRulesIDSilent(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PatchNpaRulesIDSilent: %v", v)
	}
}

type PatchNpaRulesIDRequest struct {
	// policy rule id
	ID               int                     `pathParam:"style=simple,explode=false,name=id"`
	NpaPolicyRequest shared.NpaPolicyRequest `request:"mediaType=application/json"`
	// flag to skip output except status code
	Silent *PatchNpaRulesIDSilent `queryParam:"style=form,explode=true,name=silent"`
}

func (o *PatchNpaRulesIDRequest) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *PatchNpaRulesIDRequest) GetNpaPolicyRequest() shared.NpaPolicyRequest {
	if o == nil {
		return shared.NpaPolicyRequest{}
	}
	return o.NpaPolicyRequest
}

func (o *PatchNpaRulesIDRequest) GetSilent() *PatchNpaRulesIDSilent {
	if o == nil {
		return nil
	}
	return o.Silent
}

type PatchNpaRulesID200ApplicationJSONStatus string

const (
	PatchNpaRulesID200ApplicationJSONStatusSuccess PatchNpaRulesID200ApplicationJSONStatus = "success"
	PatchNpaRulesID200ApplicationJSONStatusError   PatchNpaRulesID200ApplicationJSONStatus = "error"
)

func (e PatchNpaRulesID200ApplicationJSONStatus) ToPointer() *PatchNpaRulesID200ApplicationJSONStatus {
	return &e
}

func (e *PatchNpaRulesID200ApplicationJSONStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = PatchNpaRulesID200ApplicationJSONStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PatchNpaRulesID200ApplicationJSONStatus: %v", v)
	}
}

// PatchNpaRulesID200ApplicationJSON - successful operation
type PatchNpaRulesID200ApplicationJSON struct {
	Data   *shared.NpaPolicyResponseItem            `json:"data,omitempty"`
	Status *PatchNpaRulesID200ApplicationJSONStatus `json:"status,omitempty"`
}

func (o *PatchNpaRulesID200ApplicationJSON) GetData() *shared.NpaPolicyResponseItem {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *PatchNpaRulesID200ApplicationJSON) GetStatus() *PatchNpaRulesID200ApplicationJSONStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type PatchNpaRulesIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid request
	NpaPolicyResponse400 *shared.NpaPolicyResponse400
	// successful operation
	PatchNpaRulesID200ApplicationJSONObject *PatchNpaRulesID200ApplicationJSON
}

func (o *PatchNpaRulesIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchNpaRulesIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchNpaRulesIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchNpaRulesIDResponse) GetNpaPolicyResponse400() *shared.NpaPolicyResponse400 {
	if o == nil {
		return nil
	}
	return o.NpaPolicyResponse400
}

func (o *PatchNpaRulesIDResponse) GetPatchNpaRulesID200ApplicationJSONObject() *PatchNpaRulesID200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchNpaRulesID200ApplicationJSONObject
}
