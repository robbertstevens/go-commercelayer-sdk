/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SatispayPaymentUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SatispayPaymentUpdate{}

// SatispayPaymentUpdate struct for SatispayPaymentUpdate
type SatispayPaymentUpdate struct {
	Data SatispayPaymentUpdateData `json:"data"`
}

// NewSatispayPaymentUpdate instantiates a new SatispayPaymentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSatispayPaymentUpdate(data SatispayPaymentUpdateData) *SatispayPaymentUpdate {
	this := SatispayPaymentUpdate{}
	this.Data = data
	return &this
}

// NewSatispayPaymentUpdateWithDefaults instantiates a new SatispayPaymentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSatispayPaymentUpdateWithDefaults() *SatispayPaymentUpdate {
	this := SatispayPaymentUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *SatispayPaymentUpdate) GetData() SatispayPaymentUpdateData {
	if o == nil {
		var ret SatispayPaymentUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SatispayPaymentUpdate) GetDataOk() (*SatispayPaymentUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SatispayPaymentUpdate) SetData(v SatispayPaymentUpdateData) {
	o.Data = v
}

func (o SatispayPaymentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SatispayPaymentUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSatispayPaymentUpdate struct {
	value *SatispayPaymentUpdate
	isSet bool
}

func (v NullableSatispayPaymentUpdate) Get() *SatispayPaymentUpdate {
	return v.value
}

func (v *NullableSatispayPaymentUpdate) Set(val *SatispayPaymentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSatispayPaymentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSatispayPaymentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSatispayPaymentUpdate(val *SatispayPaymentUpdate) *NullableSatispayPaymentUpdate {
	return &NullableSatispayPaymentUpdate{value: val, isSet: true}
}

func (v NullableSatispayPaymentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSatispayPaymentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
