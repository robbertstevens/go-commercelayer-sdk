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

// checks if the POSTParcelsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTParcelsRequest{}

// POSTParcelsRequest struct for POSTParcelsRequest
type POSTParcelsRequest struct {
	Data POSTParcelsRequestData `json:"data"`
}

// NewPOSTParcelsRequest instantiates a new POSTParcelsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTParcelsRequest(data POSTParcelsRequestData) *POSTParcelsRequest {
	this := POSTParcelsRequest{}
	this.Data = data
	return &this
}

// NewPOSTParcelsRequestWithDefaults instantiates a new POSTParcelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTParcelsRequestWithDefaults() *POSTParcelsRequest {
	this := POSTParcelsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTParcelsRequest) GetData() POSTParcelsRequestData {
	if o == nil {
		var ret POSTParcelsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTParcelsRequest) GetDataOk() (*POSTParcelsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTParcelsRequest) SetData(v POSTParcelsRequestData) {
	o.Data = v
}

func (o POSTParcelsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTParcelsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTParcelsRequest struct {
	value *POSTParcelsRequest
	isSet bool
}

func (v NullablePOSTParcelsRequest) Get() *POSTParcelsRequest {
	return v.value
}

func (v *NullablePOSTParcelsRequest) Set(val *POSTParcelsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTParcelsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTParcelsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTParcelsRequest(val *POSTParcelsRequest) *NullablePOSTParcelsRequest {
	return &NullablePOSTParcelsRequest{value: val, isSet: true}
}

func (v NullablePOSTParcelsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTParcelsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
