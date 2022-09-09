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

// CheckoutComGatewayUpdate struct for CheckoutComGatewayUpdate
type CheckoutComGatewayUpdate struct {
	Data CheckoutComGatewayUpdateData `json:"data"`
}

// NewCheckoutComGatewayUpdate instantiates a new CheckoutComGatewayUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComGatewayUpdate(data CheckoutComGatewayUpdateData) *CheckoutComGatewayUpdate {
	this := CheckoutComGatewayUpdate{}
	this.Data = data
	return &this
}

// NewCheckoutComGatewayUpdateWithDefaults instantiates a new CheckoutComGatewayUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComGatewayUpdateWithDefaults() *CheckoutComGatewayUpdate {
	this := CheckoutComGatewayUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *CheckoutComGatewayUpdate) GetData() CheckoutComGatewayUpdateData {
	if o == nil {
		var ret CheckoutComGatewayUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayUpdate) GetDataOk() (*CheckoutComGatewayUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CheckoutComGatewayUpdate) SetData(v CheckoutComGatewayUpdateData) {
	o.Data = v
}

func (o CheckoutComGatewayUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutComGatewayUpdate struct {
	value *CheckoutComGatewayUpdate
	isSet bool
}

func (v NullableCheckoutComGatewayUpdate) Get() *CheckoutComGatewayUpdate {
	return v.value
}

func (v *NullableCheckoutComGatewayUpdate) Set(val *CheckoutComGatewayUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComGatewayUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComGatewayUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComGatewayUpdate(val *CheckoutComGatewayUpdate) *NullableCheckoutComGatewayUpdate {
	return &NullableCheckoutComGatewayUpdate{value: val, isSet: true}
}

func (v NullableCheckoutComGatewayUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComGatewayUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


