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

// GETLineItemsLineItemId200Response struct for GETLineItemsLineItemId200Response
type GETLineItemsLineItemId200Response struct {
	Data *GETLineItems200ResponseDataInner `json:"data,omitempty"`
}

// NewGETLineItemsLineItemId200Response instantiates a new GETLineItemsLineItemId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItemsLineItemId200Response() *GETLineItemsLineItemId200Response {
	this := GETLineItemsLineItemId200Response{}
	return &this
}

// NewGETLineItemsLineItemId200ResponseWithDefaults instantiates a new GETLineItemsLineItemId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItemsLineItemId200ResponseWithDefaults() *GETLineItemsLineItemId200Response {
	this := GETLineItemsLineItemId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETLineItemsLineItemId200Response) GetData() GETLineItems200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETLineItems200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemsLineItemId200Response) GetDataOk() (*GETLineItems200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETLineItemsLineItemId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETLineItems200ResponseDataInner and assigns it to the Data field.
func (o *GETLineItemsLineItemId200Response) SetData(v GETLineItems200ResponseDataInner) {
	o.Data = &v
}

func (o GETLineItemsLineItemId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItemsLineItemId200Response struct {
	value *GETLineItemsLineItemId200Response
	isSet bool
}

func (v NullableGETLineItemsLineItemId200Response) Get() *GETLineItemsLineItemId200Response {
	return v.value
}

func (v *NullableGETLineItemsLineItemId200Response) Set(val *GETLineItemsLineItemId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItemsLineItemId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItemsLineItemId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItemsLineItemId200Response(val *GETLineItemsLineItemId200Response) *NullableGETLineItemsLineItemId200Response {
	return &NullableGETLineItemsLineItemId200Response{value: val, isSet: true}
}

func (v NullableGETLineItemsLineItemId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItemsLineItemId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
