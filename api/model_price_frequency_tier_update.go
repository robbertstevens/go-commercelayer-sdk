/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PriceFrequencyTierUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceFrequencyTierUpdate{}

// PriceFrequencyTierUpdate struct for PriceFrequencyTierUpdate
type PriceFrequencyTierUpdate struct {
	Data PriceFrequencyTierUpdateData `json:"data"`
}

// NewPriceFrequencyTierUpdate instantiates a new PriceFrequencyTierUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceFrequencyTierUpdate(data PriceFrequencyTierUpdateData) *PriceFrequencyTierUpdate {
	this := PriceFrequencyTierUpdate{}
	this.Data = data
	return &this
}

// NewPriceFrequencyTierUpdateWithDefaults instantiates a new PriceFrequencyTierUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceFrequencyTierUpdateWithDefaults() *PriceFrequencyTierUpdate {
	this := PriceFrequencyTierUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *PriceFrequencyTierUpdate) GetData() PriceFrequencyTierUpdateData {
	if o == nil {
		var ret PriceFrequencyTierUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PriceFrequencyTierUpdate) GetDataOk() (*PriceFrequencyTierUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PriceFrequencyTierUpdate) SetData(v PriceFrequencyTierUpdateData) {
	o.Data = v
}

func (o PriceFrequencyTierUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceFrequencyTierUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePriceFrequencyTierUpdate struct {
	value *PriceFrequencyTierUpdate
	isSet bool
}

func (v NullablePriceFrequencyTierUpdate) Get() *PriceFrequencyTierUpdate {
	return v.value
}

func (v *NullablePriceFrequencyTierUpdate) Set(val *PriceFrequencyTierUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceFrequencyTierUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceFrequencyTierUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceFrequencyTierUpdate(val *PriceFrequencyTierUpdate) *NullablePriceFrequencyTierUpdate {
	return &NullablePriceFrequencyTierUpdate{value: val, isSet: true}
}

func (v NullablePriceFrequencyTierUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceFrequencyTierUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
