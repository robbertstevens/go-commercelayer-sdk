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

// GETStockItemsStockItemId200Response struct for GETStockItemsStockItemId200Response
type GETStockItemsStockItemId200Response struct {
	Data *GETStockItems200ResponseDataInner `json:"data,omitempty"`
}

// NewGETStockItemsStockItemId200Response instantiates a new GETStockItemsStockItemId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockItemsStockItemId200Response() *GETStockItemsStockItemId200Response {
	this := GETStockItemsStockItemId200Response{}
	return &this
}

// NewGETStockItemsStockItemId200ResponseWithDefaults instantiates a new GETStockItemsStockItemId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockItemsStockItemId200ResponseWithDefaults() *GETStockItemsStockItemId200Response {
	this := GETStockItemsStockItemId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETStockItemsStockItemId200Response) GetData() GETStockItems200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETStockItems200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockItemsStockItemId200Response) GetDataOk() (*GETStockItems200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETStockItemsStockItemId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETStockItems200ResponseDataInner and assigns it to the Data field.
func (o *GETStockItemsStockItemId200Response) SetData(v GETStockItems200ResponseDataInner) {
	o.Data = &v
}

func (o GETStockItemsStockItemId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETStockItemsStockItemId200Response struct {
	value *GETStockItemsStockItemId200Response
	isSet bool
}

func (v NullableGETStockItemsStockItemId200Response) Get() *GETStockItemsStockItemId200Response {
	return v.value
}

func (v *NullableGETStockItemsStockItemId200Response) Set(val *GETStockItemsStockItemId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockItemsStockItemId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockItemsStockItemId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockItemsStockItemId200Response(val *GETStockItemsStockItemId200Response) *NullableGETStockItemsStockItemId200Response {
	return &NullableGETStockItemsStockItemId200Response{value: val, isSet: true}
}

func (v NullableGETStockItemsStockItemId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockItemsStockItemId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
