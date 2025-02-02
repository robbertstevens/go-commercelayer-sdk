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

// checks if the OrderDataRelationshipsOrderCopies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDataRelationshipsOrderCopies{}

// OrderDataRelationshipsOrderCopies struct for OrderDataRelationshipsOrderCopies
type OrderDataRelationshipsOrderCopies struct {
	Data *OrderDataRelationshipsOrderCopiesData `json:"data,omitempty"`
}

// NewOrderDataRelationshipsOrderCopies instantiates a new OrderDataRelationshipsOrderCopies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDataRelationshipsOrderCopies() *OrderDataRelationshipsOrderCopies {
	this := OrderDataRelationshipsOrderCopies{}
	return &this
}

// NewOrderDataRelationshipsOrderCopiesWithDefaults instantiates a new OrderDataRelationshipsOrderCopies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataRelationshipsOrderCopiesWithDefaults() *OrderDataRelationshipsOrderCopies {
	this := OrderDataRelationshipsOrderCopies{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderDataRelationshipsOrderCopies) GetData() OrderDataRelationshipsOrderCopiesData {
	if o == nil || IsNil(o.Data) {
		var ret OrderDataRelationshipsOrderCopiesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationshipsOrderCopies) GetDataOk() (*OrderDataRelationshipsOrderCopiesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderDataRelationshipsOrderCopies) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given OrderDataRelationshipsOrderCopiesData and assigns it to the Data field.
func (o *OrderDataRelationshipsOrderCopies) SetData(v OrderDataRelationshipsOrderCopiesData) {
	o.Data = &v
}

func (o OrderDataRelationshipsOrderCopies) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDataRelationshipsOrderCopies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableOrderDataRelationshipsOrderCopies struct {
	value *OrderDataRelationshipsOrderCopies
	isSet bool
}

func (v NullableOrderDataRelationshipsOrderCopies) Get() *OrderDataRelationshipsOrderCopies {
	return v.value
}

func (v *NullableOrderDataRelationshipsOrderCopies) Set(val *OrderDataRelationshipsOrderCopies) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataRelationshipsOrderCopies) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataRelationshipsOrderCopies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataRelationshipsOrderCopies(val *OrderDataRelationshipsOrderCopies) *NullableOrderDataRelationshipsOrderCopies {
	return &NullableOrderDataRelationshipsOrderCopies{value: val, isSet: true}
}

func (v NullableOrderDataRelationshipsOrderCopies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataRelationshipsOrderCopies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
