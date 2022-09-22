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

// PATCHReturnLineItemsReturnLineItemId200Response struct for PATCHReturnLineItemsReturnLineItemId200Response
type PATCHReturnLineItemsReturnLineItemId200Response struct {
	Data *PATCHReturnLineItemsReturnLineItemId200ResponseData `json:"data,omitempty"`
}

// NewPATCHReturnLineItemsReturnLineItemId200Response instantiates a new PATCHReturnLineItemsReturnLineItemId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHReturnLineItemsReturnLineItemId200Response() *PATCHReturnLineItemsReturnLineItemId200Response {
	this := PATCHReturnLineItemsReturnLineItemId200Response{}
	return &this
}

// NewPATCHReturnLineItemsReturnLineItemId200ResponseWithDefaults instantiates a new PATCHReturnLineItemsReturnLineItemId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHReturnLineItemsReturnLineItemId200ResponseWithDefaults() *PATCHReturnLineItemsReturnLineItemId200Response {
	this := PATCHReturnLineItemsReturnLineItemId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHReturnLineItemsReturnLineItemId200Response) GetData() PATCHReturnLineItemsReturnLineItemId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHReturnLineItemsReturnLineItemId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200Response) GetDataOk() (*PATCHReturnLineItemsReturnLineItemId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHReturnLineItemsReturnLineItemId200ResponseData and assigns it to the Data field.
func (o *PATCHReturnLineItemsReturnLineItemId200Response) SetData(v PATCHReturnLineItemsReturnLineItemId200ResponseData) {
	o.Data = &v
}

func (o PATCHReturnLineItemsReturnLineItemId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHReturnLineItemsReturnLineItemId200Response struct {
	value *PATCHReturnLineItemsReturnLineItemId200Response
	isSet bool
}

func (v NullablePATCHReturnLineItemsReturnLineItemId200Response) Get() *PATCHReturnLineItemsReturnLineItemId200Response {
	return v.value
}

func (v *NullablePATCHReturnLineItemsReturnLineItemId200Response) Set(val *PATCHReturnLineItemsReturnLineItemId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHReturnLineItemsReturnLineItemId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHReturnLineItemsReturnLineItemId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHReturnLineItemsReturnLineItemId200Response(val *PATCHReturnLineItemsReturnLineItemId200Response) *NullablePATCHReturnLineItemsReturnLineItemId200Response {
	return &NullablePATCHReturnLineItemsReturnLineItemId200Response{value: val, isSet: true}
}

func (v NullablePATCHReturnLineItemsReturnLineItemId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHReturnLineItemsReturnLineItemId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
