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

// checks if the GETFreeGiftPromotionsFreeGiftPromotionId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETFreeGiftPromotionsFreeGiftPromotionId200Response{}

// GETFreeGiftPromotionsFreeGiftPromotionId200Response struct for GETFreeGiftPromotionsFreeGiftPromotionId200Response
type GETFreeGiftPromotionsFreeGiftPromotionId200Response struct {
	Data *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseData `json:"data,omitempty"`
}

// NewGETFreeGiftPromotionsFreeGiftPromotionId200Response instantiates a new GETFreeGiftPromotionsFreeGiftPromotionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETFreeGiftPromotionsFreeGiftPromotionId200Response() *GETFreeGiftPromotionsFreeGiftPromotionId200Response {
	this := GETFreeGiftPromotionsFreeGiftPromotionId200Response{}
	return &this
}

// NewGETFreeGiftPromotionsFreeGiftPromotionId200ResponseWithDefaults instantiates a new GETFreeGiftPromotionsFreeGiftPromotionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETFreeGiftPromotionsFreeGiftPromotionId200ResponseWithDefaults() *GETFreeGiftPromotionsFreeGiftPromotionId200Response {
	this := GETFreeGiftPromotionsFreeGiftPromotionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETFreeGiftPromotionsFreeGiftPromotionId200Response) GetData() GETFreeGiftPromotionsFreeGiftPromotionId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETFreeGiftPromotionsFreeGiftPromotionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFreeGiftPromotionsFreeGiftPromotionId200Response) GetDataOk() (*GETFreeGiftPromotionsFreeGiftPromotionId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETFreeGiftPromotionsFreeGiftPromotionId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETFreeGiftPromotionsFreeGiftPromotionId200ResponseData and assigns it to the Data field.
func (o *GETFreeGiftPromotionsFreeGiftPromotionId200Response) SetData(v GETFreeGiftPromotionsFreeGiftPromotionId200ResponseData) {
	o.Data = &v
}

func (o GETFreeGiftPromotionsFreeGiftPromotionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETFreeGiftPromotionsFreeGiftPromotionId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETFreeGiftPromotionsFreeGiftPromotionId200Response struct {
	value *GETFreeGiftPromotionsFreeGiftPromotionId200Response
	isSet bool
}

func (v NullableGETFreeGiftPromotionsFreeGiftPromotionId200Response) Get() *GETFreeGiftPromotionsFreeGiftPromotionId200Response {
	return v.value
}

func (v *NullableGETFreeGiftPromotionsFreeGiftPromotionId200Response) Set(val *GETFreeGiftPromotionsFreeGiftPromotionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETFreeGiftPromotionsFreeGiftPromotionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETFreeGiftPromotionsFreeGiftPromotionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETFreeGiftPromotionsFreeGiftPromotionId200Response(val *GETFreeGiftPromotionsFreeGiftPromotionId200Response) *NullableGETFreeGiftPromotionsFreeGiftPromotionId200Response {
	return &NullableGETFreeGiftPromotionsFreeGiftPromotionId200Response{value: val, isSet: true}
}

func (v NullableGETFreeGiftPromotionsFreeGiftPromotionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETFreeGiftPromotionsFreeGiftPromotionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
