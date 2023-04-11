/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTMerchantsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMerchantsRequest{}

// POSTMerchantsRequest struct for POSTMerchantsRequest
type POSTMerchantsRequest struct {
	Data POSTMerchantsRequestData `json:"data"`
}

// NewPOSTMerchantsRequest instantiates a new POSTMerchantsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMerchantsRequest(data POSTMerchantsRequestData) *POSTMerchantsRequest {
	this := POSTMerchantsRequest{}
	this.Data = data
	return &this
}

// NewPOSTMerchantsRequestWithDefaults instantiates a new POSTMerchantsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMerchantsRequestWithDefaults() *POSTMerchantsRequest {
	this := POSTMerchantsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTMerchantsRequest) GetData() POSTMerchantsRequestData {
	if o == nil {
		var ret POSTMerchantsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTMerchantsRequest) GetDataOk() (*POSTMerchantsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTMerchantsRequest) SetData(v POSTMerchantsRequestData) {
	o.Data = v
}

func (o POSTMerchantsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMerchantsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTMerchantsRequest struct {
	value *POSTMerchantsRequest
	isSet bool
}

func (v NullablePOSTMerchantsRequest) Get() *POSTMerchantsRequest {
	return v.value
}

func (v *NullablePOSTMerchantsRequest) Set(val *POSTMerchantsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMerchantsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMerchantsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMerchantsRequest(val *POSTMerchantsRequest) *NullablePOSTMerchantsRequest {
	return &NullablePOSTMerchantsRequest{value: val, isSet: true}
}

func (v NullablePOSTMerchantsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMerchantsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
