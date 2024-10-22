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

// checks if the GoogleGeocoderCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleGeocoderCreate{}

// GoogleGeocoderCreate struct for GoogleGeocoderCreate
type GoogleGeocoderCreate struct {
	Data GoogleGeocoderCreateData `json:"data"`
}

// NewGoogleGeocoderCreate instantiates a new GoogleGeocoderCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleGeocoderCreate(data GoogleGeocoderCreateData) *GoogleGeocoderCreate {
	this := GoogleGeocoderCreate{}
	this.Data = data
	return &this
}

// NewGoogleGeocoderCreateWithDefaults instantiates a new GoogleGeocoderCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleGeocoderCreateWithDefaults() *GoogleGeocoderCreate {
	this := GoogleGeocoderCreate{}
	return &this
}

// GetData returns the Data field value
func (o *GoogleGeocoderCreate) GetData() GoogleGeocoderCreateData {
	if o == nil {
		var ret GoogleGeocoderCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GoogleGeocoderCreate) GetDataOk() (*GoogleGeocoderCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GoogleGeocoderCreate) SetData(v GoogleGeocoderCreateData) {
	o.Data = v
}

func (o GoogleGeocoderCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleGeocoderCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGoogleGeocoderCreate struct {
	value *GoogleGeocoderCreate
	isSet bool
}

func (v NullableGoogleGeocoderCreate) Get() *GoogleGeocoderCreate {
	return v.value
}

func (v *NullableGoogleGeocoderCreate) Set(val *GoogleGeocoderCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleGeocoderCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleGeocoderCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleGeocoderCreate(val *GoogleGeocoderCreate) *NullableGoogleGeocoderCreate {
	return &NullableGoogleGeocoderCreate{value: val, isSet: true}
}

func (v NullableGoogleGeocoderCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleGeocoderCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
