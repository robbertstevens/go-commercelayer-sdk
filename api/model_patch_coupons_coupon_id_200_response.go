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

// PATCHCouponsCouponId200Response struct for PATCHCouponsCouponId200Response
type PATCHCouponsCouponId200Response struct {
	Data *PATCHCouponsCouponId200ResponseData `json:"data,omitempty"`
}

// NewPATCHCouponsCouponId200Response instantiates a new PATCHCouponsCouponId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCouponsCouponId200Response() *PATCHCouponsCouponId200Response {
	this := PATCHCouponsCouponId200Response{}
	return &this
}

// NewPATCHCouponsCouponId200ResponseWithDefaults instantiates a new PATCHCouponsCouponId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCouponsCouponId200ResponseWithDefaults() *PATCHCouponsCouponId200Response {
	this := PATCHCouponsCouponId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200Response) GetData() PATCHCouponsCouponId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHCouponsCouponId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200Response) GetDataOk() (*PATCHCouponsCouponId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHCouponsCouponId200ResponseData and assigns it to the Data field.
func (o *PATCHCouponsCouponId200Response) SetData(v PATCHCouponsCouponId200ResponseData) {
	o.Data = &v
}

func (o PATCHCouponsCouponId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHCouponsCouponId200Response struct {
	value *PATCHCouponsCouponId200Response
	isSet bool
}

func (v NullablePATCHCouponsCouponId200Response) Get() *PATCHCouponsCouponId200Response {
	return v.value
}

func (v *NullablePATCHCouponsCouponId200Response) Set(val *PATCHCouponsCouponId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCouponsCouponId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCouponsCouponId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCouponsCouponId200Response(val *PATCHCouponsCouponId200Response) *NullablePATCHCouponsCouponId200Response {
	return &NullablePATCHCouponsCouponId200Response{value: val, isSet: true}
}

func (v NullablePATCHCouponsCouponId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCouponsCouponId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
