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

// checks if the GETShippingZonesShippingZoneId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETShippingZonesShippingZoneId200Response{}

// GETShippingZonesShippingZoneId200Response struct for GETShippingZonesShippingZoneId200Response
type GETShippingZonesShippingZoneId200Response struct {
	Data *GETShippingZonesShippingZoneId200ResponseData `json:"data,omitempty"`
}

// NewGETShippingZonesShippingZoneId200Response instantiates a new GETShippingZonesShippingZoneId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingZonesShippingZoneId200Response() *GETShippingZonesShippingZoneId200Response {
	this := GETShippingZonesShippingZoneId200Response{}
	return &this
}

// NewGETShippingZonesShippingZoneId200ResponseWithDefaults instantiates a new GETShippingZonesShippingZoneId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingZonesShippingZoneId200ResponseWithDefaults() *GETShippingZonesShippingZoneId200Response {
	this := GETShippingZonesShippingZoneId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETShippingZonesShippingZoneId200Response) GetData() GETShippingZonesShippingZoneId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETShippingZonesShippingZoneId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingZonesShippingZoneId200Response) GetDataOk() (*GETShippingZonesShippingZoneId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETShippingZonesShippingZoneId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETShippingZonesShippingZoneId200ResponseData and assigns it to the Data field.
func (o *GETShippingZonesShippingZoneId200Response) SetData(v GETShippingZonesShippingZoneId200ResponseData) {
	o.Data = &v
}

func (o GETShippingZonesShippingZoneId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETShippingZonesShippingZoneId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETShippingZonesShippingZoneId200Response struct {
	value *GETShippingZonesShippingZoneId200Response
	isSet bool
}

func (v NullableGETShippingZonesShippingZoneId200Response) Get() *GETShippingZonesShippingZoneId200Response {
	return v.value
}

func (v *NullableGETShippingZonesShippingZoneId200Response) Set(val *GETShippingZonesShippingZoneId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingZonesShippingZoneId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingZonesShippingZoneId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingZonesShippingZoneId200Response(val *GETShippingZonesShippingZoneId200Response) *NullableGETShippingZonesShippingZoneId200Response {
	return &NullableGETShippingZonesShippingZoneId200Response{value: val, isSet: true}
}

func (v NullableGETShippingZonesShippingZoneId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingZonesShippingZoneId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
