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

// checks if the POSTAdjustmentsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAdjustmentsRequest{}

// POSTAdjustmentsRequest struct for POSTAdjustmentsRequest
type POSTAdjustmentsRequest struct {
	Data POSTAdjustmentsRequestData `json:"data"`
}

// NewPOSTAdjustmentsRequest instantiates a new POSTAdjustmentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAdjustmentsRequest(data POSTAdjustmentsRequestData) *POSTAdjustmentsRequest {
	this := POSTAdjustmentsRequest{}
	this.Data = data
	return &this
}

// NewPOSTAdjustmentsRequestWithDefaults instantiates a new POSTAdjustmentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAdjustmentsRequestWithDefaults() *POSTAdjustmentsRequest {
	this := POSTAdjustmentsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTAdjustmentsRequest) GetData() POSTAdjustmentsRequestData {
	if o == nil {
		var ret POSTAdjustmentsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTAdjustmentsRequest) GetDataOk() (*POSTAdjustmentsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTAdjustmentsRequest) SetData(v POSTAdjustmentsRequestData) {
	o.Data = v
}

func (o POSTAdjustmentsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAdjustmentsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTAdjustmentsRequest struct {
	value *POSTAdjustmentsRequest
	isSet bool
}

func (v NullablePOSTAdjustmentsRequest) Get() *POSTAdjustmentsRequest {
	return v.value
}

func (v *NullablePOSTAdjustmentsRequest) Set(val *POSTAdjustmentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAdjustmentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAdjustmentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAdjustmentsRequest(val *POSTAdjustmentsRequest) *NullablePOSTAdjustmentsRequest {
	return &NullablePOSTAdjustmentsRequest{value: val, isSet: true}
}

func (v NullablePOSTAdjustmentsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAdjustmentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
