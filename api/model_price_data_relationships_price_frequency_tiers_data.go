/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PriceDataRelationshipsPriceFrequencyTiersData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceDataRelationshipsPriceFrequencyTiersData{}

// PriceDataRelationshipsPriceFrequencyTiersData struct for PriceDataRelationshipsPriceFrequencyTiersData
type PriceDataRelationshipsPriceFrequencyTiersData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewPriceDataRelationshipsPriceFrequencyTiersData instantiates a new PriceDataRelationshipsPriceFrequencyTiersData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceDataRelationshipsPriceFrequencyTiersData() *PriceDataRelationshipsPriceFrequencyTiersData {
	this := PriceDataRelationshipsPriceFrequencyTiersData{}
	return &this
}

// NewPriceDataRelationshipsPriceFrequencyTiersDataWithDefaults instantiates a new PriceDataRelationshipsPriceFrequencyTiersData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceDataRelationshipsPriceFrequencyTiersDataWithDefaults() *PriceDataRelationshipsPriceFrequencyTiersData {
	this := PriceDataRelationshipsPriceFrequencyTiersData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceDataRelationshipsPriceFrequencyTiersData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceDataRelationshipsPriceFrequencyTiersData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PriceDataRelationshipsPriceFrequencyTiersData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *PriceDataRelationshipsPriceFrequencyTiersData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceDataRelationshipsPriceFrequencyTiersData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceDataRelationshipsPriceFrequencyTiersData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PriceDataRelationshipsPriceFrequencyTiersData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *PriceDataRelationshipsPriceFrequencyTiersData) SetId(v interface{}) {
	o.Id = v
}

func (o PriceDataRelationshipsPriceFrequencyTiersData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceDataRelationshipsPriceFrequencyTiersData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePriceDataRelationshipsPriceFrequencyTiersData struct {
	value *PriceDataRelationshipsPriceFrequencyTiersData
	isSet bool
}

func (v NullablePriceDataRelationshipsPriceFrequencyTiersData) Get() *PriceDataRelationshipsPriceFrequencyTiersData {
	return v.value
}

func (v *NullablePriceDataRelationshipsPriceFrequencyTiersData) Set(val *PriceDataRelationshipsPriceFrequencyTiersData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceDataRelationshipsPriceFrequencyTiersData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceDataRelationshipsPriceFrequencyTiersData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceDataRelationshipsPriceFrequencyTiersData(val *PriceDataRelationshipsPriceFrequencyTiersData) *NullablePriceDataRelationshipsPriceFrequencyTiersData {
	return &NullablePriceDataRelationshipsPriceFrequencyTiersData{value: val, isSet: true}
}

func (v NullablePriceDataRelationshipsPriceFrequencyTiersData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceDataRelationshipsPriceFrequencyTiersData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
