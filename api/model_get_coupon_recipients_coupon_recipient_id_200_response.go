/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETCouponRecipientsCouponRecipientId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCouponRecipientsCouponRecipientId200Response{}

// GETCouponRecipientsCouponRecipientId200Response struct for GETCouponRecipientsCouponRecipientId200Response
type GETCouponRecipientsCouponRecipientId200Response struct {
	Data *GETCouponRecipientsCouponRecipientId200ResponseData `json:"data,omitempty"`
}

// NewGETCouponRecipientsCouponRecipientId200Response instantiates a new GETCouponRecipientsCouponRecipientId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCouponRecipientsCouponRecipientId200Response() *GETCouponRecipientsCouponRecipientId200Response {
	this := GETCouponRecipientsCouponRecipientId200Response{}
	return &this
}

// NewGETCouponRecipientsCouponRecipientId200ResponseWithDefaults instantiates a new GETCouponRecipientsCouponRecipientId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCouponRecipientsCouponRecipientId200ResponseWithDefaults() *GETCouponRecipientsCouponRecipientId200Response {
	this := GETCouponRecipientsCouponRecipientId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCouponRecipientsCouponRecipientId200Response) GetData() GETCouponRecipientsCouponRecipientId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETCouponRecipientsCouponRecipientId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCouponRecipientsCouponRecipientId200Response) GetDataOk() (*GETCouponRecipientsCouponRecipientId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCouponRecipientsCouponRecipientId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCouponRecipientsCouponRecipientId200ResponseData and assigns it to the Data field.
func (o *GETCouponRecipientsCouponRecipientId200Response) SetData(v GETCouponRecipientsCouponRecipientId200ResponseData) {
	o.Data = &v
}

func (o GETCouponRecipientsCouponRecipientId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCouponRecipientsCouponRecipientId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETCouponRecipientsCouponRecipientId200Response struct {
	value *GETCouponRecipientsCouponRecipientId200Response
	isSet bool
}

func (v NullableGETCouponRecipientsCouponRecipientId200Response) Get() *GETCouponRecipientsCouponRecipientId200Response {
	return v.value
}

func (v *NullableGETCouponRecipientsCouponRecipientId200Response) Set(val *GETCouponRecipientsCouponRecipientId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCouponRecipientsCouponRecipientId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCouponRecipientsCouponRecipientId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCouponRecipientsCouponRecipientId200Response(val *GETCouponRecipientsCouponRecipientId200Response) *NullableGETCouponRecipientsCouponRecipientId200Response {
	return &NullableGETCouponRecipientsCouponRecipientId200Response{value: val, isSet: true}
}

func (v NullableGETCouponRecipientsCouponRecipientId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCouponRecipientsCouponRecipientId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
