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

// checks if the CustomerAddressUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerAddressUpdate{}

// CustomerAddressUpdate struct for CustomerAddressUpdate
type CustomerAddressUpdate struct {
	Data PATCHCustomerAddressesCustomerAddressIdRequestData `json:"data"`
}

// NewCustomerAddressUpdate instantiates a new CustomerAddressUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddressUpdate(data PATCHCustomerAddressesCustomerAddressIdRequestData) *CustomerAddressUpdate {
	this := CustomerAddressUpdate{}
	this.Data = data
	return &this
}

// NewCustomerAddressUpdateWithDefaults instantiates a new CustomerAddressUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressUpdateWithDefaults() *CustomerAddressUpdate {
	this := CustomerAddressUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerAddressUpdate) GetData() PATCHCustomerAddressesCustomerAddressIdRequestData {
	if o == nil {
		var ret PATCHCustomerAddressesCustomerAddressIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerAddressUpdate) GetDataOk() (*PATCHCustomerAddressesCustomerAddressIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerAddressUpdate) SetData(v PATCHCustomerAddressesCustomerAddressIdRequestData) {
	o.Data = v
}

func (o CustomerAddressUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerAddressUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCustomerAddressUpdate struct {
	value *CustomerAddressUpdate
	isSet bool
}

func (v NullableCustomerAddressUpdate) Get() *CustomerAddressUpdate {
	return v.value
}

func (v *NullableCustomerAddressUpdate) Set(val *CustomerAddressUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressUpdate(val *CustomerAddressUpdate) *NullableCustomerAddressUpdate {
	return &NullableCustomerAddressUpdate{value: val, isSet: true}
}

func (v NullableCustomerAddressUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
