/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PATCHExternalGatewaysExternalGatewayId200Response struct for PATCHExternalGatewaysExternalGatewayId200Response
type PATCHExternalGatewaysExternalGatewayId200Response struct {
	Data *PATCHExternalGatewaysExternalGatewayId200ResponseData `json:"data,omitempty"`
}

// NewPATCHExternalGatewaysExternalGatewayId200Response instantiates a new PATCHExternalGatewaysExternalGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHExternalGatewaysExternalGatewayId200Response() *PATCHExternalGatewaysExternalGatewayId200Response {
	this := PATCHExternalGatewaysExternalGatewayId200Response{}
	return &this
}

// NewPATCHExternalGatewaysExternalGatewayId200ResponseWithDefaults instantiates a new PATCHExternalGatewaysExternalGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHExternalGatewaysExternalGatewayId200ResponseWithDefaults() *PATCHExternalGatewaysExternalGatewayId200Response {
	this := PATCHExternalGatewaysExternalGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHExternalGatewaysExternalGatewayId200Response) GetData() PATCHExternalGatewaysExternalGatewayId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHExternalGatewaysExternalGatewayId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalGatewaysExternalGatewayId200Response) GetDataOk() (*PATCHExternalGatewaysExternalGatewayId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHExternalGatewaysExternalGatewayId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHExternalGatewaysExternalGatewayId200ResponseData and assigns it to the Data field.
func (o *PATCHExternalGatewaysExternalGatewayId200Response) SetData(v PATCHExternalGatewaysExternalGatewayId200ResponseData) {
	o.Data = &v
}

func (o PATCHExternalGatewaysExternalGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHExternalGatewaysExternalGatewayId200Response struct {
	value *PATCHExternalGatewaysExternalGatewayId200Response
	isSet bool
}

func (v NullablePATCHExternalGatewaysExternalGatewayId200Response) Get() *PATCHExternalGatewaysExternalGatewayId200Response {
	return v.value
}

func (v *NullablePATCHExternalGatewaysExternalGatewayId200Response) Set(val *PATCHExternalGatewaysExternalGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHExternalGatewaysExternalGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHExternalGatewaysExternalGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHExternalGatewaysExternalGatewayId200Response(val *PATCHExternalGatewaysExternalGatewayId200Response) *NullablePATCHExternalGatewaysExternalGatewayId200Response {
	return &NullablePATCHExternalGatewaysExternalGatewayId200Response{value: val, isSet: true}
}

func (v NullablePATCHExternalGatewaysExternalGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHExternalGatewaysExternalGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
