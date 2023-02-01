/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GoogleGeocoder struct for GoogleGeocoder
type GoogleGeocoder struct {
	Data *GoogleGeocoderData `json:"data,omitempty"`
}

// NewGoogleGeocoder instantiates a new GoogleGeocoder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleGeocoder() *GoogleGeocoder {
	this := GoogleGeocoder{}
	return &this
}

// NewGoogleGeocoderWithDefaults instantiates a new GoogleGeocoder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleGeocoderWithDefaults() *GoogleGeocoder {
	this := GoogleGeocoder{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GoogleGeocoder) GetData() GoogleGeocoderData {
	if o == nil || o.Data == nil {
		var ret GoogleGeocoderData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleGeocoder) GetDataOk() (*GoogleGeocoderData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GoogleGeocoder) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GoogleGeocoderData and assigns it to the Data field.
func (o *GoogleGeocoder) SetData(v GoogleGeocoderData) {
	o.Data = &v
}

func (o GoogleGeocoder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleGeocoder struct {
	value *GoogleGeocoder
	isSet bool
}

func (v NullableGoogleGeocoder) Get() *GoogleGeocoder {
	return v.value
}

func (v *NullableGoogleGeocoder) Set(val *GoogleGeocoder) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleGeocoder) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleGeocoder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleGeocoder(val *GoogleGeocoder) *NullableGoogleGeocoder {
	return &NullableGoogleGeocoder{value: val, isSet: true}
}

func (v NullableGoogleGeocoder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleGeocoder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
