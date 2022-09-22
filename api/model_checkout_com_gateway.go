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

// CheckoutComGateway struct for CheckoutComGateway
type CheckoutComGateway struct {
	Data CheckoutComGatewayData `json:"data"`
}

// NewCheckoutComGateway instantiates a new CheckoutComGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComGateway(data CheckoutComGatewayData) *CheckoutComGateway {
	this := CheckoutComGateway{}
	this.Data = data
	return &this
}

// NewCheckoutComGatewayWithDefaults instantiates a new CheckoutComGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComGatewayWithDefaults() *CheckoutComGateway {
	this := CheckoutComGateway{}
	return &this
}

// GetData returns the Data field value
func (o *CheckoutComGateway) GetData() CheckoutComGatewayData {
	if o == nil {
		var ret CheckoutComGatewayData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CheckoutComGateway) GetDataOk() (*CheckoutComGatewayData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CheckoutComGateway) SetData(v CheckoutComGatewayData) {
	o.Data = v
}

func (o CheckoutComGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutComGateway struct {
	value *CheckoutComGateway
	isSet bool
}

func (v NullableCheckoutComGateway) Get() *CheckoutComGateway {
	return v.value
}

func (v *NullableCheckoutComGateway) Set(val *CheckoutComGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComGateway(val *CheckoutComGateway) *NullableCheckoutComGateway {
	return &NullableCheckoutComGateway{value: val, isSet: true}
}

func (v NullableCheckoutComGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
