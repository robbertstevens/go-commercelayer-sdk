/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PATCHCapturesCaptureId200Response struct for PATCHCapturesCaptureId200Response
type PATCHCapturesCaptureId200Response struct {
	Data *PATCHCapturesCaptureId200ResponseData `json:"data,omitempty"`
}

// NewPATCHCapturesCaptureId200Response instantiates a new PATCHCapturesCaptureId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCapturesCaptureId200Response() *PATCHCapturesCaptureId200Response {
	this := PATCHCapturesCaptureId200Response{}
	return &this
}

// NewPATCHCapturesCaptureId200ResponseWithDefaults instantiates a new PATCHCapturesCaptureId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCapturesCaptureId200ResponseWithDefaults() *PATCHCapturesCaptureId200Response {
	this := PATCHCapturesCaptureId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHCapturesCaptureId200Response) GetData() PATCHCapturesCaptureId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHCapturesCaptureId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCapturesCaptureId200Response) GetDataOk() (*PATCHCapturesCaptureId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHCapturesCaptureId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHCapturesCaptureId200ResponseData and assigns it to the Data field.
func (o *PATCHCapturesCaptureId200Response) SetData(v PATCHCapturesCaptureId200ResponseData) {
	o.Data = &v
}

func (o PATCHCapturesCaptureId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHCapturesCaptureId200Response struct {
	value *PATCHCapturesCaptureId200Response
	isSet bool
}

func (v NullablePATCHCapturesCaptureId200Response) Get() *PATCHCapturesCaptureId200Response {
	return v.value
}

func (v *NullablePATCHCapturesCaptureId200Response) Set(val *PATCHCapturesCaptureId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCapturesCaptureId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCapturesCaptureId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCapturesCaptureId200Response(val *PATCHCapturesCaptureId200Response) *NullablePATCHCapturesCaptureId200Response {
	return &NullablePATCHCapturesCaptureId200Response{value: val, isSet: true}
}

func (v NullablePATCHCapturesCaptureId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCapturesCaptureId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
