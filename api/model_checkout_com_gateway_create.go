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

// checks if the CheckoutComGatewayCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutComGatewayCreate{}

// CheckoutComGatewayCreate struct for CheckoutComGatewayCreate
type CheckoutComGatewayCreate struct {
	Data POSTCheckoutComGatewaysRequestData `json:"data"`
}

// NewCheckoutComGatewayCreate instantiates a new CheckoutComGatewayCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComGatewayCreate(data POSTCheckoutComGatewaysRequestData) *CheckoutComGatewayCreate {
	this := CheckoutComGatewayCreate{}
	this.Data = data
	return &this
}

// NewCheckoutComGatewayCreateWithDefaults instantiates a new CheckoutComGatewayCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComGatewayCreateWithDefaults() *CheckoutComGatewayCreate {
	this := CheckoutComGatewayCreate{}
	return &this
}

// GetData returns the Data field value
func (o *CheckoutComGatewayCreate) GetData() POSTCheckoutComGatewaysRequestData {
	if o == nil {
		var ret POSTCheckoutComGatewaysRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayCreate) GetDataOk() (*POSTCheckoutComGatewaysRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CheckoutComGatewayCreate) SetData(v POSTCheckoutComGatewaysRequestData) {
	o.Data = v
}

func (o CheckoutComGatewayCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutComGatewayCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCheckoutComGatewayCreate struct {
	value *CheckoutComGatewayCreate
	isSet bool
}

func (v NullableCheckoutComGatewayCreate) Get() *CheckoutComGatewayCreate {
	return v.value
}

func (v *NullableCheckoutComGatewayCreate) Set(val *CheckoutComGatewayCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComGatewayCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComGatewayCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComGatewayCreate(val *CheckoutComGatewayCreate) *NullableCheckoutComGatewayCreate {
	return &NullableCheckoutComGatewayCreate{value: val, isSet: true}
}

func (v NullableCheckoutComGatewayCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComGatewayCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
