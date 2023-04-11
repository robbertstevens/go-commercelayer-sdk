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

// checks if the WireTransferCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WireTransferCreate{}

// WireTransferCreate struct for WireTransferCreate
type WireTransferCreate struct {
	Data POSTWireTransfersRequestData `json:"data"`
}

// NewWireTransferCreate instantiates a new WireTransferCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireTransferCreate(data POSTWireTransfersRequestData) *WireTransferCreate {
	this := WireTransferCreate{}
	this.Data = data
	return &this
}

// NewWireTransferCreateWithDefaults instantiates a new WireTransferCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireTransferCreateWithDefaults() *WireTransferCreate {
	this := WireTransferCreate{}
	return &this
}

// GetData returns the Data field value
func (o *WireTransferCreate) GetData() POSTWireTransfersRequestData {
	if o == nil {
		var ret POSTWireTransfersRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WireTransferCreate) GetDataOk() (*POSTWireTransfersRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *WireTransferCreate) SetData(v POSTWireTransfersRequestData) {
	o.Data = v
}

func (o WireTransferCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WireTransferCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableWireTransferCreate struct {
	value *WireTransferCreate
	isSet bool
}

func (v NullableWireTransferCreate) Get() *WireTransferCreate {
	return v.value
}

func (v *NullableWireTransferCreate) Set(val *WireTransferCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableWireTransferCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableWireTransferCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireTransferCreate(val *WireTransferCreate) *NullableWireTransferCreate {
	return &NullableWireTransferCreate{value: val, isSet: true}
}

func (v NullableWireTransferCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireTransferCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
