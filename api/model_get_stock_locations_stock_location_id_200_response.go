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

// checks if the GETStockLocationsStockLocationId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETStockLocationsStockLocationId200Response{}

// GETStockLocationsStockLocationId200Response struct for GETStockLocationsStockLocationId200Response
type GETStockLocationsStockLocationId200Response struct {
	Data *GETStockLocationsStockLocationId200ResponseData `json:"data,omitempty"`
}

// NewGETStockLocationsStockLocationId200Response instantiates a new GETStockLocationsStockLocationId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockLocationsStockLocationId200Response() *GETStockLocationsStockLocationId200Response {
	this := GETStockLocationsStockLocationId200Response{}
	return &this
}

// NewGETStockLocationsStockLocationId200ResponseWithDefaults instantiates a new GETStockLocationsStockLocationId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockLocationsStockLocationId200ResponseWithDefaults() *GETStockLocationsStockLocationId200Response {
	this := GETStockLocationsStockLocationId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETStockLocationsStockLocationId200Response) GetData() GETStockLocationsStockLocationId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETStockLocationsStockLocationId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocationsStockLocationId200Response) GetDataOk() (*GETStockLocationsStockLocationId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETStockLocationsStockLocationId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETStockLocationsStockLocationId200ResponseData and assigns it to the Data field.
func (o *GETStockLocationsStockLocationId200Response) SetData(v GETStockLocationsStockLocationId200ResponseData) {
	o.Data = &v
}

func (o GETStockLocationsStockLocationId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETStockLocationsStockLocationId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETStockLocationsStockLocationId200Response struct {
	value *GETStockLocationsStockLocationId200Response
	isSet bool
}

func (v NullableGETStockLocationsStockLocationId200Response) Get() *GETStockLocationsStockLocationId200Response {
	return v.value
}

func (v *NullableGETStockLocationsStockLocationId200Response) Set(val *GETStockLocationsStockLocationId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockLocationsStockLocationId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockLocationsStockLocationId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockLocationsStockLocationId200Response(val *GETStockLocationsStockLocationId200Response) *NullableGETStockLocationsStockLocationId200Response {
	return &NullableGETStockLocationsStockLocationId200Response{value: val, isSet: true}
}

func (v NullableGETStockLocationsStockLocationId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockLocationsStockLocationId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
