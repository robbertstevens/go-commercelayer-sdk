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

// checks if the DeliveryLeadTimeCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryLeadTimeCreate{}

// DeliveryLeadTimeCreate struct for DeliveryLeadTimeCreate
type DeliveryLeadTimeCreate struct {
	Data POSTDeliveryLeadTimesRequestData `json:"data"`
}

// NewDeliveryLeadTimeCreate instantiates a new DeliveryLeadTimeCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeCreate(data POSTDeliveryLeadTimesRequestData) *DeliveryLeadTimeCreate {
	this := DeliveryLeadTimeCreate{}
	this.Data = data
	return &this
}

// NewDeliveryLeadTimeCreateWithDefaults instantiates a new DeliveryLeadTimeCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeCreateWithDefaults() *DeliveryLeadTimeCreate {
	this := DeliveryLeadTimeCreate{}
	return &this
}

// GetData returns the Data field value
func (o *DeliveryLeadTimeCreate) GetData() POSTDeliveryLeadTimesRequestData {
	if o == nil {
		var ret POSTDeliveryLeadTimesRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeCreate) GetDataOk() (*POSTDeliveryLeadTimesRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeliveryLeadTimeCreate) SetData(v POSTDeliveryLeadTimesRequestData) {
	o.Data = v
}

func (o DeliveryLeadTimeCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryLeadTimeCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableDeliveryLeadTimeCreate struct {
	value *DeliveryLeadTimeCreate
	isSet bool
}

func (v NullableDeliveryLeadTimeCreate) Get() *DeliveryLeadTimeCreate {
	return v.value
}

func (v *NullableDeliveryLeadTimeCreate) Set(val *DeliveryLeadTimeCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeCreate(val *DeliveryLeadTimeCreate) *NullableDeliveryLeadTimeCreate {
	return &NullableDeliveryLeadTimeCreate{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
