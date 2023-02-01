/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BraintreeGateway struct for BraintreeGateway
type BraintreeGateway struct {
	Data *BraintreeGatewayData `json:"data,omitempty"`
}

// NewBraintreeGateway instantiates a new BraintreeGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBraintreeGateway() *BraintreeGateway {
	this := BraintreeGateway{}
	return &this
}

// NewBraintreeGatewayWithDefaults instantiates a new BraintreeGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBraintreeGatewayWithDefaults() *BraintreeGateway {
	this := BraintreeGateway{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BraintreeGateway) GetData() BraintreeGatewayData {
	if o == nil || o.Data == nil {
		var ret BraintreeGatewayData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreeGateway) GetDataOk() (*BraintreeGatewayData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BraintreeGateway) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given BraintreeGatewayData and assigns it to the Data field.
func (o *BraintreeGateway) SetData(v BraintreeGatewayData) {
	o.Data = &v
}

func (o BraintreeGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBraintreeGateway struct {
	value *BraintreeGateway
	isSet bool
}

func (v NullableBraintreeGateway) Get() *BraintreeGateway {
	return v.value
}

func (v *NullableBraintreeGateway) Set(val *BraintreeGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableBraintreeGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableBraintreeGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBraintreeGateway(val *BraintreeGateway) *NullableBraintreeGateway {
	return &NullableBraintreeGateway{value: val, isSet: true}
}

func (v NullableBraintreeGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBraintreeGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
