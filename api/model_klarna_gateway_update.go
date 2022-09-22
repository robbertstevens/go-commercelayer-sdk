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

// KlarnaGatewayUpdate struct for KlarnaGatewayUpdate
type KlarnaGatewayUpdate struct {
	Data KlarnaGatewayUpdateData `json:"data"`
}

// NewKlarnaGatewayUpdate instantiates a new KlarnaGatewayUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaGatewayUpdate(data KlarnaGatewayUpdateData) *KlarnaGatewayUpdate {
	this := KlarnaGatewayUpdate{}
	this.Data = data
	return &this
}

// NewKlarnaGatewayUpdateWithDefaults instantiates a new KlarnaGatewayUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaGatewayUpdateWithDefaults() *KlarnaGatewayUpdate {
	this := KlarnaGatewayUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *KlarnaGatewayUpdate) GetData() KlarnaGatewayUpdateData {
	if o == nil {
		var ret KlarnaGatewayUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayUpdate) GetDataOk() (*KlarnaGatewayUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *KlarnaGatewayUpdate) SetData(v KlarnaGatewayUpdateData) {
	o.Data = v
}

func (o KlarnaGatewayUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableKlarnaGatewayUpdate struct {
	value *KlarnaGatewayUpdate
	isSet bool
}

func (v NullableKlarnaGatewayUpdate) Get() *KlarnaGatewayUpdate {
	return v.value
}

func (v *NullableKlarnaGatewayUpdate) Set(val *KlarnaGatewayUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaGatewayUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaGatewayUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaGatewayUpdate(val *KlarnaGatewayUpdate) *NullableKlarnaGatewayUpdate {
	return &NullableKlarnaGatewayUpdate{value: val, isSet: true}
}

func (v NullableKlarnaGatewayUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaGatewayUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
