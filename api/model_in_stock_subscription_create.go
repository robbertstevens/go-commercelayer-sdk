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

// checks if the InStockSubscriptionCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InStockSubscriptionCreate{}

// InStockSubscriptionCreate struct for InStockSubscriptionCreate
type InStockSubscriptionCreate struct {
	Data InStockSubscriptionCreateData `json:"data"`
}

// NewInStockSubscriptionCreate instantiates a new InStockSubscriptionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionCreate(data InStockSubscriptionCreateData) *InStockSubscriptionCreate {
	this := InStockSubscriptionCreate{}
	this.Data = data
	return &this
}

// NewInStockSubscriptionCreateWithDefaults instantiates a new InStockSubscriptionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionCreateWithDefaults() *InStockSubscriptionCreate {
	this := InStockSubscriptionCreate{}
	return &this
}

// GetData returns the Data field value
func (o *InStockSubscriptionCreate) GetData() InStockSubscriptionCreateData {
	if o == nil {
		var ret InStockSubscriptionCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreate) GetDataOk() (*InStockSubscriptionCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InStockSubscriptionCreate) SetData(v InStockSubscriptionCreateData) {
	o.Data = v
}

func (o InStockSubscriptionCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InStockSubscriptionCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableInStockSubscriptionCreate struct {
	value *InStockSubscriptionCreate
	isSet bool
}

func (v NullableInStockSubscriptionCreate) Get() *InStockSubscriptionCreate {
	return v.value
}

func (v *NullableInStockSubscriptionCreate) Set(val *InStockSubscriptionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionCreate(val *InStockSubscriptionCreate) *NullableInStockSubscriptionCreate {
	return &NullableInStockSubscriptionCreate{value: val, isSet: true}
}

func (v NullableInStockSubscriptionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
