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

// checks if the PATCHTaxCategoriesTaxCategoryIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHTaxCategoriesTaxCategoryIdRequest{}

// PATCHTaxCategoriesTaxCategoryIdRequest struct for PATCHTaxCategoriesTaxCategoryIdRequest
type PATCHTaxCategoriesTaxCategoryIdRequest struct {
	Data PATCHTaxCategoriesTaxCategoryIdRequestData `json:"data"`
}

// NewPATCHTaxCategoriesTaxCategoryIdRequest instantiates a new PATCHTaxCategoriesTaxCategoryIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHTaxCategoriesTaxCategoryIdRequest(data PATCHTaxCategoriesTaxCategoryIdRequestData) *PATCHTaxCategoriesTaxCategoryIdRequest {
	this := PATCHTaxCategoriesTaxCategoryIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHTaxCategoriesTaxCategoryIdRequestWithDefaults instantiates a new PATCHTaxCategoriesTaxCategoryIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHTaxCategoriesTaxCategoryIdRequestWithDefaults() *PATCHTaxCategoriesTaxCategoryIdRequest {
	this := PATCHTaxCategoriesTaxCategoryIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHTaxCategoriesTaxCategoryIdRequest) GetData() PATCHTaxCategoriesTaxCategoryIdRequestData {
	if o == nil {
		var ret PATCHTaxCategoriesTaxCategoryIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHTaxCategoriesTaxCategoryIdRequest) GetDataOk() (*PATCHTaxCategoriesTaxCategoryIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHTaxCategoriesTaxCategoryIdRequest) SetData(v PATCHTaxCategoriesTaxCategoryIdRequestData) {
	o.Data = v
}

func (o PATCHTaxCategoriesTaxCategoryIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHTaxCategoriesTaxCategoryIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHTaxCategoriesTaxCategoryIdRequest struct {
	value *PATCHTaxCategoriesTaxCategoryIdRequest
	isSet bool
}

func (v NullablePATCHTaxCategoriesTaxCategoryIdRequest) Get() *PATCHTaxCategoriesTaxCategoryIdRequest {
	return v.value
}

func (v *NullablePATCHTaxCategoriesTaxCategoryIdRequest) Set(val *PATCHTaxCategoriesTaxCategoryIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHTaxCategoriesTaxCategoryIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHTaxCategoriesTaxCategoryIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHTaxCategoriesTaxCategoryIdRequest(val *PATCHTaxCategoriesTaxCategoryIdRequest) *NullablePATCHTaxCategoriesTaxCategoryIdRequest {
	return &NullablePATCHTaxCategoriesTaxCategoryIdRequest{value: val, isSet: true}
}

func (v NullablePATCHTaxCategoriesTaxCategoryIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHTaxCategoriesTaxCategoryIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
