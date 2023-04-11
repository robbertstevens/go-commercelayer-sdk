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

// checks if the PATCHBundlesBundleIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHBundlesBundleIdRequest{}

// PATCHBundlesBundleIdRequest struct for PATCHBundlesBundleIdRequest
type PATCHBundlesBundleIdRequest struct {
	Data PATCHBundlesBundleIdRequestData `json:"data"`
}

// NewPATCHBundlesBundleIdRequest instantiates a new PATCHBundlesBundleIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHBundlesBundleIdRequest(data PATCHBundlesBundleIdRequestData) *PATCHBundlesBundleIdRequest {
	this := PATCHBundlesBundleIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHBundlesBundleIdRequestWithDefaults instantiates a new PATCHBundlesBundleIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHBundlesBundleIdRequestWithDefaults() *PATCHBundlesBundleIdRequest {
	this := PATCHBundlesBundleIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHBundlesBundleIdRequest) GetData() PATCHBundlesBundleIdRequestData {
	if o == nil {
		var ret PATCHBundlesBundleIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHBundlesBundleIdRequest) GetDataOk() (*PATCHBundlesBundleIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHBundlesBundleIdRequest) SetData(v PATCHBundlesBundleIdRequestData) {
	o.Data = v
}

func (o PATCHBundlesBundleIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHBundlesBundleIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHBundlesBundleIdRequest struct {
	value *PATCHBundlesBundleIdRequest
	isSet bool
}

func (v NullablePATCHBundlesBundleIdRequest) Get() *PATCHBundlesBundleIdRequest {
	return v.value
}

func (v *NullablePATCHBundlesBundleIdRequest) Set(val *PATCHBundlesBundleIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHBundlesBundleIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHBundlesBundleIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHBundlesBundleIdRequest(val *PATCHBundlesBundleIdRequest) *NullablePATCHBundlesBundleIdRequest {
	return &NullablePATCHBundlesBundleIdRequest{value: val, isSet: true}
}

func (v NullablePATCHBundlesBundleIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHBundlesBundleIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
