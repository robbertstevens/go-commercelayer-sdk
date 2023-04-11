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

// checks if the GETExternalPaymentsExternalPaymentId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETExternalPaymentsExternalPaymentId200Response{}

// GETExternalPaymentsExternalPaymentId200Response struct for GETExternalPaymentsExternalPaymentId200Response
type GETExternalPaymentsExternalPaymentId200Response struct {
	Data *GETExternalPaymentsExternalPaymentId200ResponseData `json:"data,omitempty"`
}

// NewGETExternalPaymentsExternalPaymentId200Response instantiates a new GETExternalPaymentsExternalPaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalPaymentsExternalPaymentId200Response() *GETExternalPaymentsExternalPaymentId200Response {
	this := GETExternalPaymentsExternalPaymentId200Response{}
	return &this
}

// NewGETExternalPaymentsExternalPaymentId200ResponseWithDefaults instantiates a new GETExternalPaymentsExternalPaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalPaymentsExternalPaymentId200ResponseWithDefaults() *GETExternalPaymentsExternalPaymentId200Response {
	this := GETExternalPaymentsExternalPaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETExternalPaymentsExternalPaymentId200Response) GetData() GETExternalPaymentsExternalPaymentId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETExternalPaymentsExternalPaymentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPaymentsExternalPaymentId200Response) GetDataOk() (*GETExternalPaymentsExternalPaymentId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETExternalPaymentsExternalPaymentId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETExternalPaymentsExternalPaymentId200ResponseData and assigns it to the Data field.
func (o *GETExternalPaymentsExternalPaymentId200Response) SetData(v GETExternalPaymentsExternalPaymentId200ResponseData) {
	o.Data = &v
}

func (o GETExternalPaymentsExternalPaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETExternalPaymentsExternalPaymentId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETExternalPaymentsExternalPaymentId200Response struct {
	value *GETExternalPaymentsExternalPaymentId200Response
	isSet bool
}

func (v NullableGETExternalPaymentsExternalPaymentId200Response) Get() *GETExternalPaymentsExternalPaymentId200Response {
	return v.value
}

func (v *NullableGETExternalPaymentsExternalPaymentId200Response) Set(val *GETExternalPaymentsExternalPaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalPaymentsExternalPaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalPaymentsExternalPaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalPaymentsExternalPaymentId200Response(val *GETExternalPaymentsExternalPaymentId200Response) *NullableGETExternalPaymentsExternalPaymentId200Response {
	return &NullableGETExternalPaymentsExternalPaymentId200Response{value: val, isSet: true}
}

func (v NullableGETExternalPaymentsExternalPaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalPaymentsExternalPaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
