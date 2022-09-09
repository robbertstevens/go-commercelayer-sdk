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

// GETSkuListItemsSkuListItemId200Response struct for GETSkuListItemsSkuListItemId200Response
type GETSkuListItemsSkuListItemId200Response struct {
	Data *GETSkuListItems200ResponseDataInner `json:"data,omitempty"`
}

// NewGETSkuListItemsSkuListItemId200Response instantiates a new GETSkuListItemsSkuListItemId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkuListItemsSkuListItemId200Response() *GETSkuListItemsSkuListItemId200Response {
	this := GETSkuListItemsSkuListItemId200Response{}
	return &this
}

// NewGETSkuListItemsSkuListItemId200ResponseWithDefaults instantiates a new GETSkuListItemsSkuListItemId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkuListItemsSkuListItemId200ResponseWithDefaults() *GETSkuListItemsSkuListItemId200Response {
	this := GETSkuListItemsSkuListItemId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETSkuListItemsSkuListItemId200Response) GetData() GETSkuListItems200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETSkuListItems200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuListItemsSkuListItemId200Response) GetDataOk() (*GETSkuListItems200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETSkuListItemsSkuListItemId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETSkuListItems200ResponseDataInner and assigns it to the Data field.
func (o *GETSkuListItemsSkuListItemId200Response) SetData(v GETSkuListItems200ResponseDataInner) {
	o.Data = &v
}

func (o GETSkuListItemsSkuListItemId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETSkuListItemsSkuListItemId200Response struct {
	value *GETSkuListItemsSkuListItemId200Response
	isSet bool
}

func (v NullableGETSkuListItemsSkuListItemId200Response) Get() *GETSkuListItemsSkuListItemId200Response {
	return v.value
}

func (v *NullableGETSkuListItemsSkuListItemId200Response) Set(val *GETSkuListItemsSkuListItemId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkuListItemsSkuListItemId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkuListItemsSkuListItemId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkuListItemsSkuListItemId200Response(val *GETSkuListItemsSkuListItemId200Response) *NullableGETSkuListItemsSkuListItemId200Response {
	return &NullableGETSkuListItemsSkuListItemId200Response{value: val, isSet: true}
}

func (v NullableGETSkuListItemsSkuListItemId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkuListItemsSkuListItemId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


