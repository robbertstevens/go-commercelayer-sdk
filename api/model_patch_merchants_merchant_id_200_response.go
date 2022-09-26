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

// PATCHMerchantsMerchantId200Response struct for PATCHMerchantsMerchantId200Response
type PATCHMerchantsMerchantId200Response struct {
	Data *PATCHMerchantsMerchantId200ResponseData `json:"data,omitempty"`
}

// NewPATCHMerchantsMerchantId200Response instantiates a new PATCHMerchantsMerchantId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHMerchantsMerchantId200Response() *PATCHMerchantsMerchantId200Response {
	this := PATCHMerchantsMerchantId200Response{}
	return &this
}

// NewPATCHMerchantsMerchantId200ResponseWithDefaults instantiates a new PATCHMerchantsMerchantId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHMerchantsMerchantId200ResponseWithDefaults() *PATCHMerchantsMerchantId200Response {
	this := PATCHMerchantsMerchantId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHMerchantsMerchantId200Response) GetData() PATCHMerchantsMerchantId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHMerchantsMerchantId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHMerchantsMerchantId200Response) GetDataOk() (*PATCHMerchantsMerchantId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHMerchantsMerchantId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHMerchantsMerchantId200ResponseData and assigns it to the Data field.
func (o *PATCHMerchantsMerchantId200Response) SetData(v PATCHMerchantsMerchantId200ResponseData) {
	o.Data = &v
}

func (o PATCHMerchantsMerchantId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHMerchantsMerchantId200Response struct {
	value *PATCHMerchantsMerchantId200Response
	isSet bool
}

func (v NullablePATCHMerchantsMerchantId200Response) Get() *PATCHMerchantsMerchantId200Response {
	return v.value
}

func (v *NullablePATCHMerchantsMerchantId200Response) Set(val *PATCHMerchantsMerchantId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHMerchantsMerchantId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHMerchantsMerchantId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHMerchantsMerchantId200Response(val *PATCHMerchantsMerchantId200Response) *NullablePATCHMerchantsMerchantId200Response {
	return &NullablePATCHMerchantsMerchantId200Response{value: val, isSet: true}
}

func (v NullablePATCHMerchantsMerchantId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHMerchantsMerchantId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
