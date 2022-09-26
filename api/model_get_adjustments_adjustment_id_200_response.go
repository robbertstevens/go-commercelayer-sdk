/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETAdjustmentsAdjustmentId200Response struct for GETAdjustmentsAdjustmentId200Response
type GETAdjustmentsAdjustmentId200Response struct {
	Data *GETAdjustments200ResponseDataInner `json:"data,omitempty"`
}

// NewGETAdjustmentsAdjustmentId200Response instantiates a new GETAdjustmentsAdjustmentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAdjustmentsAdjustmentId200Response() *GETAdjustmentsAdjustmentId200Response {
	this := GETAdjustmentsAdjustmentId200Response{}
	return &this
}

// NewGETAdjustmentsAdjustmentId200ResponseWithDefaults instantiates a new GETAdjustmentsAdjustmentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAdjustmentsAdjustmentId200ResponseWithDefaults() *GETAdjustmentsAdjustmentId200Response {
	this := GETAdjustmentsAdjustmentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAdjustmentsAdjustmentId200Response) GetData() GETAdjustments200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETAdjustments200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdjustmentsAdjustmentId200Response) GetDataOk() (*GETAdjustments200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAdjustmentsAdjustmentId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAdjustments200ResponseDataInner and assigns it to the Data field.
func (o *GETAdjustmentsAdjustmentId200Response) SetData(v GETAdjustments200ResponseDataInner) {
	o.Data = &v
}

func (o GETAdjustmentsAdjustmentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETAdjustmentsAdjustmentId200Response struct {
	value *GETAdjustmentsAdjustmentId200Response
	isSet bool
}

func (v NullableGETAdjustmentsAdjustmentId200Response) Get() *GETAdjustmentsAdjustmentId200Response {
	return v.value
}

func (v *NullableGETAdjustmentsAdjustmentId200Response) Set(val *GETAdjustmentsAdjustmentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAdjustmentsAdjustmentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAdjustmentsAdjustmentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAdjustmentsAdjustmentId200Response(val *GETAdjustmentsAdjustmentId200Response) *NullableGETAdjustmentsAdjustmentId200Response {
	return &NullableGETAdjustmentsAdjustmentId200Response{value: val, isSet: true}
}

func (v NullableGETAdjustmentsAdjustmentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAdjustmentsAdjustmentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
