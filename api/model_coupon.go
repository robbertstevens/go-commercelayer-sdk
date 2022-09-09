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

// Coupon struct for Coupon
type Coupon struct {
	Data CouponData `json:"data"`
}

// NewCoupon instantiates a new Coupon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoupon(data CouponData) *Coupon {
	this := Coupon{}
	this.Data = data
	return &this
}

// NewCouponWithDefaults instantiates a new Coupon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponWithDefaults() *Coupon {
	this := Coupon{}
	return &this
}

// GetData returns the Data field value
func (o *Coupon) GetData() CouponData {
	if o == nil {
		var ret CouponData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Coupon) GetDataOk() (*CouponData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Coupon) SetData(v CouponData) {
	o.Data = v
}

func (o Coupon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCoupon struct {
	value *Coupon
	isSet bool
}

func (v NullableCoupon) Get() *Coupon {
	return v.value
}

func (v *NullableCoupon) Set(val *Coupon) {
	v.value = val
	v.isSet = true
}

func (v NullableCoupon) IsSet() bool {
	return v.isSet
}

func (v *NullableCoupon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoupon(val *Coupon) *NullableCoupon {
	return &NullableCoupon{value: val, isSet: true}
}

func (v NullableCoupon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoupon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


