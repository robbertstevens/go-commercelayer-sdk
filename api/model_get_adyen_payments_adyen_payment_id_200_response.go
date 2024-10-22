/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETAdyenPaymentsAdyenPaymentId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETAdyenPaymentsAdyenPaymentId200Response{}

// GETAdyenPaymentsAdyenPaymentId200Response struct for GETAdyenPaymentsAdyenPaymentId200Response
type GETAdyenPaymentsAdyenPaymentId200Response struct {
	Data *GETAdyenPaymentsAdyenPaymentId200ResponseData `json:"data,omitempty"`
}

// NewGETAdyenPaymentsAdyenPaymentId200Response instantiates a new GETAdyenPaymentsAdyenPaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAdyenPaymentsAdyenPaymentId200Response() *GETAdyenPaymentsAdyenPaymentId200Response {
	this := GETAdyenPaymentsAdyenPaymentId200Response{}
	return &this
}

// NewGETAdyenPaymentsAdyenPaymentId200ResponseWithDefaults instantiates a new GETAdyenPaymentsAdyenPaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAdyenPaymentsAdyenPaymentId200ResponseWithDefaults() *GETAdyenPaymentsAdyenPaymentId200Response {
	this := GETAdyenPaymentsAdyenPaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAdyenPaymentsAdyenPaymentId200Response) GetData() GETAdyenPaymentsAdyenPaymentId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETAdyenPaymentsAdyenPaymentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdyenPaymentsAdyenPaymentId200Response) GetDataOk() (*GETAdyenPaymentsAdyenPaymentId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAdyenPaymentsAdyenPaymentId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAdyenPaymentsAdyenPaymentId200ResponseData and assigns it to the Data field.
func (o *GETAdyenPaymentsAdyenPaymentId200Response) SetData(v GETAdyenPaymentsAdyenPaymentId200ResponseData) {
	o.Data = &v
}

func (o GETAdyenPaymentsAdyenPaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETAdyenPaymentsAdyenPaymentId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETAdyenPaymentsAdyenPaymentId200Response struct {
	value *GETAdyenPaymentsAdyenPaymentId200Response
	isSet bool
}

func (v NullableGETAdyenPaymentsAdyenPaymentId200Response) Get() *GETAdyenPaymentsAdyenPaymentId200Response {
	return v.value
}

func (v *NullableGETAdyenPaymentsAdyenPaymentId200Response) Set(val *GETAdyenPaymentsAdyenPaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAdyenPaymentsAdyenPaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAdyenPaymentsAdyenPaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAdyenPaymentsAdyenPaymentId200Response(val *GETAdyenPaymentsAdyenPaymentId200Response) *NullableGETAdyenPaymentsAdyenPaymentId200Response {
	return &NullableGETAdyenPaymentsAdyenPaymentId200Response{value: val, isSet: true}
}

func (v NullableGETAdyenPaymentsAdyenPaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAdyenPaymentsAdyenPaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
