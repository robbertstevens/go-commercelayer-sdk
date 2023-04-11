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

// checks if the PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest{}

// PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest struct for PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest
type PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest struct {
	Data PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestData `json:"data"`
}

// NewPATCHFreeGiftPromotionsFreeGiftPromotionIdRequest instantiates a new PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHFreeGiftPromotionsFreeGiftPromotionIdRequest(data PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestData) *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest {
	this := PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHFreeGiftPromotionsFreeGiftPromotionIdRequestWithDefaults instantiates a new PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHFreeGiftPromotionsFreeGiftPromotionIdRequestWithDefaults() *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest {
	this := PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) GetData() PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestData {
	if o == nil {
		var ret PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) GetDataOk() (*PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) SetData(v PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestData) {
	o.Data = v
}

func (o PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHFreeGiftPromotionsFreeGiftPromotionIdRequest struct {
	value *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest
	isSet bool
}

func (v NullablePATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) Get() *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest {
	return v.value
}

func (v *NullablePATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) Set(val *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHFreeGiftPromotionsFreeGiftPromotionIdRequest(val *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) *NullablePATCHFreeGiftPromotionsFreeGiftPromotionIdRequest {
	return &NullablePATCHFreeGiftPromotionsFreeGiftPromotionIdRequest{value: val, isSet: true}
}

func (v NullablePATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHFreeGiftPromotionsFreeGiftPromotionIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
