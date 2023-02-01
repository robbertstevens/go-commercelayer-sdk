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

// GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response struct for GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response
type GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response struct {
	Data *GETCouponCodesPromotionRules200ResponseDataInner `json:"data,omitempty"`
}

// NewGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response instantiates a new GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response() *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response {
	this := GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response{}
	return &this
}

// NewGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseWithDefaults instantiates a new GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseWithDefaults() *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response {
	this := GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) GetData() GETCouponCodesPromotionRules200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETCouponCodesPromotionRules200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) GetDataOk() (*GETCouponCodesPromotionRules200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCouponCodesPromotionRules200ResponseDataInner and assigns it to the Data field.
func (o *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) SetData(v GETCouponCodesPromotionRules200ResponseDataInner) {
	o.Data = &v
}

func (o GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response struct {
	value *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response
	isSet bool
}

func (v NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) Get() *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response {
	return v.value
}

func (v *NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) Set(val *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response(val *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) *NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response {
	return &NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response{value: val, isSet: true}
}

func (v NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
