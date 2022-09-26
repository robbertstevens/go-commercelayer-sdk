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

// GETCarrierAccountsCarrierAccountId200Response struct for GETCarrierAccountsCarrierAccountId200Response
type GETCarrierAccountsCarrierAccountId200Response struct {
	Data *GETCarrierAccounts200ResponseDataInner `json:"data,omitempty"`
}

// NewGETCarrierAccountsCarrierAccountId200Response instantiates a new GETCarrierAccountsCarrierAccountId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCarrierAccountsCarrierAccountId200Response() *GETCarrierAccountsCarrierAccountId200Response {
	this := GETCarrierAccountsCarrierAccountId200Response{}
	return &this
}

// NewGETCarrierAccountsCarrierAccountId200ResponseWithDefaults instantiates a new GETCarrierAccountsCarrierAccountId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCarrierAccountsCarrierAccountId200ResponseWithDefaults() *GETCarrierAccountsCarrierAccountId200Response {
	this := GETCarrierAccountsCarrierAccountId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCarrierAccountsCarrierAccountId200Response) GetData() GETCarrierAccounts200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETCarrierAccounts200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCarrierAccountsCarrierAccountId200Response) GetDataOk() (*GETCarrierAccounts200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCarrierAccountsCarrierAccountId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCarrierAccounts200ResponseDataInner and assigns it to the Data field.
func (o *GETCarrierAccountsCarrierAccountId200Response) SetData(v GETCarrierAccounts200ResponseDataInner) {
	o.Data = &v
}

func (o GETCarrierAccountsCarrierAccountId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCarrierAccountsCarrierAccountId200Response struct {
	value *GETCarrierAccountsCarrierAccountId200Response
	isSet bool
}

func (v NullableGETCarrierAccountsCarrierAccountId200Response) Get() *GETCarrierAccountsCarrierAccountId200Response {
	return v.value
}

func (v *NullableGETCarrierAccountsCarrierAccountId200Response) Set(val *GETCarrierAccountsCarrierAccountId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCarrierAccountsCarrierAccountId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCarrierAccountsCarrierAccountId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCarrierAccountsCarrierAccountId200Response(val *GETCarrierAccountsCarrierAccountId200Response) *NullableGETCarrierAccountsCarrierAccountId200Response {
	return &NullableGETCarrierAccountsCarrierAccountId200Response{value: val, isSet: true}
}

func (v NullableGETCarrierAccountsCarrierAccountId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCarrierAccountsCarrierAccountId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
