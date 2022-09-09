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

// GETShippingWeightTiersShippingWeightTierId200Response struct for GETShippingWeightTiersShippingWeightTierId200Response
type GETShippingWeightTiersShippingWeightTierId200Response struct {
	Data *GETShippingWeightTiers200ResponseDataInner `json:"data,omitempty"`
}

// NewGETShippingWeightTiersShippingWeightTierId200Response instantiates a new GETShippingWeightTiersShippingWeightTierId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingWeightTiersShippingWeightTierId200Response() *GETShippingWeightTiersShippingWeightTierId200Response {
	this := GETShippingWeightTiersShippingWeightTierId200Response{}
	return &this
}

// NewGETShippingWeightTiersShippingWeightTierId200ResponseWithDefaults instantiates a new GETShippingWeightTiersShippingWeightTierId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingWeightTiersShippingWeightTierId200ResponseWithDefaults() *GETShippingWeightTiersShippingWeightTierId200Response {
	this := GETShippingWeightTiersShippingWeightTierId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETShippingWeightTiersShippingWeightTierId200Response) GetData() GETShippingWeightTiers200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETShippingWeightTiers200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingWeightTiersShippingWeightTierId200Response) GetDataOk() (*GETShippingWeightTiers200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETShippingWeightTiersShippingWeightTierId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETShippingWeightTiers200ResponseDataInner and assigns it to the Data field.
func (o *GETShippingWeightTiersShippingWeightTierId200Response) SetData(v GETShippingWeightTiers200ResponseDataInner) {
	o.Data = &v
}

func (o GETShippingWeightTiersShippingWeightTierId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETShippingWeightTiersShippingWeightTierId200Response struct {
	value *GETShippingWeightTiersShippingWeightTierId200Response
	isSet bool
}

func (v NullableGETShippingWeightTiersShippingWeightTierId200Response) Get() *GETShippingWeightTiersShippingWeightTierId200Response {
	return v.value
}

func (v *NullableGETShippingWeightTiersShippingWeightTierId200Response) Set(val *GETShippingWeightTiersShippingWeightTierId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingWeightTiersShippingWeightTierId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingWeightTiersShippingWeightTierId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingWeightTiersShippingWeightTierId200Response(val *GETShippingWeightTiersShippingWeightTierId200Response) *NullableGETShippingWeightTiersShippingWeightTierId200Response {
	return &NullableGETShippingWeightTiersShippingWeightTierId200Response{value: val, isSet: true}
}

func (v NullableGETShippingWeightTiersShippingWeightTierId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingWeightTiersShippingWeightTierId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


