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

// checks if the PATCHPriceListsPriceListId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPriceListsPriceListId200Response{}

// PATCHPriceListsPriceListId200Response struct for PATCHPriceListsPriceListId200Response
type PATCHPriceListsPriceListId200Response struct {
	Data *PATCHPriceListsPriceListId200ResponseData `json:"data,omitempty"`
}

// NewPATCHPriceListsPriceListId200Response instantiates a new PATCHPriceListsPriceListId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPriceListsPriceListId200Response() *PATCHPriceListsPriceListId200Response {
	this := PATCHPriceListsPriceListId200Response{}
	return &this
}

// NewPATCHPriceListsPriceListId200ResponseWithDefaults instantiates a new PATCHPriceListsPriceListId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPriceListsPriceListId200ResponseWithDefaults() *PATCHPriceListsPriceListId200Response {
	this := PATCHPriceListsPriceListId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHPriceListsPriceListId200Response) GetData() PATCHPriceListsPriceListId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHPriceListsPriceListId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPriceListsPriceListId200Response) GetDataOk() (*PATCHPriceListsPriceListId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHPriceListsPriceListId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHPriceListsPriceListId200ResponseData and assigns it to the Data field.
func (o *PATCHPriceListsPriceListId200Response) SetData(v PATCHPriceListsPriceListId200ResponseData) {
	o.Data = &v
}

func (o PATCHPriceListsPriceListId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPriceListsPriceListId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHPriceListsPriceListId200Response struct {
	value *PATCHPriceListsPriceListId200Response
	isSet bool
}

func (v NullablePATCHPriceListsPriceListId200Response) Get() *PATCHPriceListsPriceListId200Response {
	return v.value
}

func (v *NullablePATCHPriceListsPriceListId200Response) Set(val *PATCHPriceListsPriceListId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPriceListsPriceListId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPriceListsPriceListId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPriceListsPriceListId200Response(val *PATCHPriceListsPriceListId200Response) *NullablePATCHPriceListsPriceListId200Response {
	return &NullablePATCHPriceListsPriceListId200Response{value: val, isSet: true}
}

func (v NullablePATCHPriceListsPriceListId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPriceListsPriceListId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
