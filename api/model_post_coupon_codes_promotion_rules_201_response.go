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

// checks if the POSTCouponCodesPromotionRules201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCouponCodesPromotionRules201Response{}

// POSTCouponCodesPromotionRules201Response struct for POSTCouponCodesPromotionRules201Response
type POSTCouponCodesPromotionRules201Response struct {
	Data *POSTCouponCodesPromotionRules201ResponseData `json:"data,omitempty"`
}

// NewPOSTCouponCodesPromotionRules201Response instantiates a new POSTCouponCodesPromotionRules201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCouponCodesPromotionRules201Response() *POSTCouponCodesPromotionRules201Response {
	this := POSTCouponCodesPromotionRules201Response{}
	return &this
}

// NewPOSTCouponCodesPromotionRules201ResponseWithDefaults instantiates a new POSTCouponCodesPromotionRules201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCouponCodesPromotionRules201ResponseWithDefaults() *POSTCouponCodesPromotionRules201Response {
	this := POSTCouponCodesPromotionRules201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCouponCodesPromotionRules201Response) GetData() POSTCouponCodesPromotionRules201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCouponCodesPromotionRules201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCouponCodesPromotionRules201Response) GetDataOk() (*POSTCouponCodesPromotionRules201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCouponCodesPromotionRules201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCouponCodesPromotionRules201ResponseData and assigns it to the Data field.
func (o *POSTCouponCodesPromotionRules201Response) SetData(v POSTCouponCodesPromotionRules201ResponseData) {
	o.Data = &v
}

func (o POSTCouponCodesPromotionRules201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCouponCodesPromotionRules201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCouponCodesPromotionRules201Response struct {
	value *POSTCouponCodesPromotionRules201Response
	isSet bool
}

func (v NullablePOSTCouponCodesPromotionRules201Response) Get() *POSTCouponCodesPromotionRules201Response {
	return v.value
}

func (v *NullablePOSTCouponCodesPromotionRules201Response) Set(val *POSTCouponCodesPromotionRules201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCouponCodesPromotionRules201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCouponCodesPromotionRules201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCouponCodesPromotionRules201Response(val *POSTCouponCodesPromotionRules201Response) *NullablePOSTCouponCodesPromotionRules201Response {
	return &NullablePOSTCouponCodesPromotionRules201Response{value: val, isSet: true}
}

func (v NullablePOSTCouponCodesPromotionRules201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCouponCodesPromotionRules201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
