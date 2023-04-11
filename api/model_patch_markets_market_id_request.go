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

// checks if the PATCHMarketsMarketIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHMarketsMarketIdRequest{}

// PATCHMarketsMarketIdRequest struct for PATCHMarketsMarketIdRequest
type PATCHMarketsMarketIdRequest struct {
	Data PATCHMarketsMarketIdRequestData `json:"data"`
}

// NewPATCHMarketsMarketIdRequest instantiates a new PATCHMarketsMarketIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHMarketsMarketIdRequest(data PATCHMarketsMarketIdRequestData) *PATCHMarketsMarketIdRequest {
	this := PATCHMarketsMarketIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHMarketsMarketIdRequestWithDefaults instantiates a new PATCHMarketsMarketIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHMarketsMarketIdRequestWithDefaults() *PATCHMarketsMarketIdRequest {
	this := PATCHMarketsMarketIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHMarketsMarketIdRequest) GetData() PATCHMarketsMarketIdRequestData {
	if o == nil {
		var ret PATCHMarketsMarketIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHMarketsMarketIdRequest) GetDataOk() (*PATCHMarketsMarketIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHMarketsMarketIdRequest) SetData(v PATCHMarketsMarketIdRequestData) {
	o.Data = v
}

func (o PATCHMarketsMarketIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHMarketsMarketIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHMarketsMarketIdRequest struct {
	value *PATCHMarketsMarketIdRequest
	isSet bool
}

func (v NullablePATCHMarketsMarketIdRequest) Get() *PATCHMarketsMarketIdRequest {
	return v.value
}

func (v *NullablePATCHMarketsMarketIdRequest) Set(val *PATCHMarketsMarketIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHMarketsMarketIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHMarketsMarketIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHMarketsMarketIdRequest(val *PATCHMarketsMarketIdRequest) *NullablePATCHMarketsMarketIdRequest {
	return &NullablePATCHMarketsMarketIdRequest{value: val, isSet: true}
}

func (v NullablePATCHMarketsMarketIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHMarketsMarketIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
