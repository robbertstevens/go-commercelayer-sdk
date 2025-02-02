/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PATCHAddressesAddressId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHAddressesAddressId200Response{}

// PATCHAddressesAddressId200Response struct for PATCHAddressesAddressId200Response
type PATCHAddressesAddressId200Response struct {
	Data *PATCHAddressesAddressId200ResponseData `json:"data,omitempty"`
}

// NewPATCHAddressesAddressId200Response instantiates a new PATCHAddressesAddressId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAddressesAddressId200Response() *PATCHAddressesAddressId200Response {
	this := PATCHAddressesAddressId200Response{}
	return &this
}

// NewPATCHAddressesAddressId200ResponseWithDefaults instantiates a new PATCHAddressesAddressId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAddressesAddressId200ResponseWithDefaults() *PATCHAddressesAddressId200Response {
	this := PATCHAddressesAddressId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200Response) GetData() PATCHAddressesAddressId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHAddressesAddressId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200Response) GetDataOk() (*PATCHAddressesAddressId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHAddressesAddressId200ResponseData and assigns it to the Data field.
func (o *PATCHAddressesAddressId200Response) SetData(v PATCHAddressesAddressId200ResponseData) {
	o.Data = &v
}

func (o PATCHAddressesAddressId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHAddressesAddressId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHAddressesAddressId200Response struct {
	value *PATCHAddressesAddressId200Response
	isSet bool
}

func (v NullablePATCHAddressesAddressId200Response) Get() *PATCHAddressesAddressId200Response {
	return v.value
}

func (v *NullablePATCHAddressesAddressId200Response) Set(val *PATCHAddressesAddressId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAddressesAddressId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAddressesAddressId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAddressesAddressId200Response(val *PATCHAddressesAddressId200Response) *NullablePATCHAddressesAddressId200Response {
	return &NullablePATCHAddressesAddressId200Response{value: val, isSet: true}
}

func (v NullablePATCHAddressesAddressId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAddressesAddressId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
