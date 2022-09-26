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

// PATCHPaymentMethodsPaymentMethodId200Response struct for PATCHPaymentMethodsPaymentMethodId200Response
type PATCHPaymentMethodsPaymentMethodId200Response struct {
	Data *PATCHPaymentMethodsPaymentMethodId200ResponseData `json:"data,omitempty"`
}

// NewPATCHPaymentMethodsPaymentMethodId200Response instantiates a new PATCHPaymentMethodsPaymentMethodId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPaymentMethodsPaymentMethodId200Response() *PATCHPaymentMethodsPaymentMethodId200Response {
	this := PATCHPaymentMethodsPaymentMethodId200Response{}
	return &this
}

// NewPATCHPaymentMethodsPaymentMethodId200ResponseWithDefaults instantiates a new PATCHPaymentMethodsPaymentMethodId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPaymentMethodsPaymentMethodId200ResponseWithDefaults() *PATCHPaymentMethodsPaymentMethodId200Response {
	this := PATCHPaymentMethodsPaymentMethodId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHPaymentMethodsPaymentMethodId200Response) GetData() PATCHPaymentMethodsPaymentMethodId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHPaymentMethodsPaymentMethodId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200Response) GetDataOk() (*PATCHPaymentMethodsPaymentMethodId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHPaymentMethodsPaymentMethodId200ResponseData and assigns it to the Data field.
func (o *PATCHPaymentMethodsPaymentMethodId200Response) SetData(v PATCHPaymentMethodsPaymentMethodId200ResponseData) {
	o.Data = &v
}

func (o PATCHPaymentMethodsPaymentMethodId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHPaymentMethodsPaymentMethodId200Response struct {
	value *PATCHPaymentMethodsPaymentMethodId200Response
	isSet bool
}

func (v NullablePATCHPaymentMethodsPaymentMethodId200Response) Get() *PATCHPaymentMethodsPaymentMethodId200Response {
	return v.value
}

func (v *NullablePATCHPaymentMethodsPaymentMethodId200Response) Set(val *PATCHPaymentMethodsPaymentMethodId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPaymentMethodsPaymentMethodId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPaymentMethodsPaymentMethodId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPaymentMethodsPaymentMethodId200Response(val *PATCHPaymentMethodsPaymentMethodId200Response) *NullablePATCHPaymentMethodsPaymentMethodId200Response {
	return &NullablePATCHPaymentMethodsPaymentMethodId200Response{value: val, isSet: true}
}

func (v NullablePATCHPaymentMethodsPaymentMethodId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPaymentMethodsPaymentMethodId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
