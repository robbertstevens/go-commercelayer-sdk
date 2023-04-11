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

// checks if the PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest{}

// PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest struct for PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest
type PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest struct {
	Data PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestData `json:"data"`
}

// NewPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest instantiates a new PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest(data PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestData) *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest {
	this := PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestWithDefaults instantiates a new PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestWithDefaults() *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest {
	this := PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) GetData() PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestData {
	if o == nil {
		var ret PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) GetDataOk() (*PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) SetData(v PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestData) {
	o.Data = v
}

func (o PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest struct {
	value *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest
	isSet bool
}

func (v NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) Get() *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest {
	return v.value
}

func (v *NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) Set(val *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest(val *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) *NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest {
	return &NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest{value: val, isSet: true}
}

func (v NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
