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

// GETWebhooks200Response struct for GETWebhooks200Response
type GETWebhooks200Response struct {
	Data []GETWebhooks200ResponseDataInner `json:"data,omitempty"`
}

// NewGETWebhooks200Response instantiates a new GETWebhooks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETWebhooks200Response() *GETWebhooks200Response {
	this := GETWebhooks200Response{}
	return &this
}

// NewGETWebhooks200ResponseWithDefaults instantiates a new GETWebhooks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETWebhooks200ResponseWithDefaults() *GETWebhooks200Response {
	this := GETWebhooks200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETWebhooks200Response) GetData() []GETWebhooks200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETWebhooks200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooks200Response) GetDataOk() ([]GETWebhooks200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETWebhooks200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETWebhooks200ResponseDataInner and assigns it to the Data field.
func (o *GETWebhooks200Response) SetData(v []GETWebhooks200ResponseDataInner) {
	o.Data = v
}

func (o GETWebhooks200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETWebhooks200Response struct {
	value *GETWebhooks200Response
	isSet bool
}

func (v NullableGETWebhooks200Response) Get() *GETWebhooks200Response {
	return v.value
}

func (v *NullableGETWebhooks200Response) Set(val *GETWebhooks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETWebhooks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETWebhooks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETWebhooks200Response(val *GETWebhooks200Response) *NullableGETWebhooks200Response {
	return &NullableGETWebhooks200Response{value: val, isSet: true}
}

func (v NullableGETWebhooks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETWebhooks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
