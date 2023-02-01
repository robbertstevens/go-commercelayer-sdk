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

// AdyenGateway struct for AdyenGateway
type AdyenGateway struct {
	Data *AdyenGatewayData `json:"data,omitempty"`
}

// NewAdyenGateway instantiates a new AdyenGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenGateway() *AdyenGateway {
	this := AdyenGateway{}
	return &this
}

// NewAdyenGatewayWithDefaults instantiates a new AdyenGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenGatewayWithDefaults() *AdyenGateway {
	this := AdyenGateway{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AdyenGateway) GetData() AdyenGatewayData {
	if o == nil || o.Data == nil {
		var ret AdyenGatewayData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGateway) GetDataOk() (*AdyenGatewayData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AdyenGateway) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AdyenGatewayData and assigns it to the Data field.
func (o *AdyenGateway) SetData(v AdyenGatewayData) {
	o.Data = &v
}

func (o AdyenGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAdyenGateway struct {
	value *AdyenGateway
	isSet bool
}

func (v NullableAdyenGateway) Get() *AdyenGateway {
	return v.value
}

func (v *NullableAdyenGateway) Set(val *AdyenGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenGateway(val *AdyenGateway) *NullableAdyenGateway {
	return &NullableAdyenGateway{value: val, isSet: true}
}

func (v NullableAdyenGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
