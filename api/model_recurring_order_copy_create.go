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

// checks if the RecurringOrderCopyCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecurringOrderCopyCreate{}

// RecurringOrderCopyCreate struct for RecurringOrderCopyCreate
type RecurringOrderCopyCreate struct {
	Data RecurringOrderCopyCreateData `json:"data"`
}

// NewRecurringOrderCopyCreate instantiates a new RecurringOrderCopyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurringOrderCopyCreate(data RecurringOrderCopyCreateData) *RecurringOrderCopyCreate {
	this := RecurringOrderCopyCreate{}
	this.Data = data
	return &this
}

// NewRecurringOrderCopyCreateWithDefaults instantiates a new RecurringOrderCopyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringOrderCopyCreateWithDefaults() *RecurringOrderCopyCreate {
	this := RecurringOrderCopyCreate{}
	return &this
}

// GetData returns the Data field value
func (o *RecurringOrderCopyCreate) GetData() RecurringOrderCopyCreateData {
	if o == nil {
		var ret RecurringOrderCopyCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RecurringOrderCopyCreate) GetDataOk() (*RecurringOrderCopyCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RecurringOrderCopyCreate) SetData(v RecurringOrderCopyCreateData) {
	o.Data = v
}

func (o RecurringOrderCopyCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecurringOrderCopyCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableRecurringOrderCopyCreate struct {
	value *RecurringOrderCopyCreate
	isSet bool
}

func (v NullableRecurringOrderCopyCreate) Get() *RecurringOrderCopyCreate {
	return v.value
}

func (v *NullableRecurringOrderCopyCreate) Set(val *RecurringOrderCopyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringOrderCopyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringOrderCopyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringOrderCopyCreate(val *RecurringOrderCopyCreate) *NullableRecurringOrderCopyCreate {
	return &NullableRecurringOrderCopyCreate{value: val, isSet: true}
}

func (v NullableRecurringOrderCopyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringOrderCopyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
