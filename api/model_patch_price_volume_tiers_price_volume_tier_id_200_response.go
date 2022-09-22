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

// PATCHPriceVolumeTiersPriceVolumeTierId200Response struct for PATCHPriceVolumeTiersPriceVolumeTierId200Response
type PATCHPriceVolumeTiersPriceVolumeTierId200Response struct {
	Data *PATCHPriceVolumeTiersPriceVolumeTierId200ResponseData `json:"data,omitempty"`
}

// NewPATCHPriceVolumeTiersPriceVolumeTierId200Response instantiates a new PATCHPriceVolumeTiersPriceVolumeTierId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPriceVolumeTiersPriceVolumeTierId200Response() *PATCHPriceVolumeTiersPriceVolumeTierId200Response {
	this := PATCHPriceVolumeTiersPriceVolumeTierId200Response{}
	return &this
}

// NewPATCHPriceVolumeTiersPriceVolumeTierId200ResponseWithDefaults instantiates a new PATCHPriceVolumeTiersPriceVolumeTierId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPriceVolumeTiersPriceVolumeTierId200ResponseWithDefaults() *PATCHPriceVolumeTiersPriceVolumeTierId200Response {
	this := PATCHPriceVolumeTiersPriceVolumeTierId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHPriceVolumeTiersPriceVolumeTierId200Response) GetData() PATCHPriceVolumeTiersPriceVolumeTierId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHPriceVolumeTiersPriceVolumeTierId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPriceVolumeTiersPriceVolumeTierId200Response) GetDataOk() (*PATCHPriceVolumeTiersPriceVolumeTierId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHPriceVolumeTiersPriceVolumeTierId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHPriceVolumeTiersPriceVolumeTierId200ResponseData and assigns it to the Data field.
func (o *PATCHPriceVolumeTiersPriceVolumeTierId200Response) SetData(v PATCHPriceVolumeTiersPriceVolumeTierId200ResponseData) {
	o.Data = &v
}

func (o PATCHPriceVolumeTiersPriceVolumeTierId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHPriceVolumeTiersPriceVolumeTierId200Response struct {
	value *PATCHPriceVolumeTiersPriceVolumeTierId200Response
	isSet bool
}

func (v NullablePATCHPriceVolumeTiersPriceVolumeTierId200Response) Get() *PATCHPriceVolumeTiersPriceVolumeTierId200Response {
	return v.value
}

func (v *NullablePATCHPriceVolumeTiersPriceVolumeTierId200Response) Set(val *PATCHPriceVolumeTiersPriceVolumeTierId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPriceVolumeTiersPriceVolumeTierId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPriceVolumeTiersPriceVolumeTierId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPriceVolumeTiersPriceVolumeTierId200Response(val *PATCHPriceVolumeTiersPriceVolumeTierId200Response) *NullablePATCHPriceVolumeTiersPriceVolumeTierId200Response {
	return &NullablePATCHPriceVolumeTiersPriceVolumeTierId200Response{value: val, isSet: true}
}

func (v NullablePATCHPriceVolumeTiersPriceVolumeTierId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPriceVolumeTiersPriceVolumeTierId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
