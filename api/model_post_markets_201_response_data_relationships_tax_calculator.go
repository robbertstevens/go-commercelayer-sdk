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

// checks if the POSTMarkets201ResponseDataRelationshipsTaxCalculator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMarkets201ResponseDataRelationshipsTaxCalculator{}

// POSTMarkets201ResponseDataRelationshipsTaxCalculator struct for POSTMarkets201ResponseDataRelationshipsTaxCalculator
type POSTMarkets201ResponseDataRelationshipsTaxCalculator struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData `json:"data,omitempty"`
}

// NewPOSTMarkets201ResponseDataRelationshipsTaxCalculator instantiates a new POSTMarkets201ResponseDataRelationshipsTaxCalculator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMarkets201ResponseDataRelationshipsTaxCalculator() *POSTMarkets201ResponseDataRelationshipsTaxCalculator {
	this := POSTMarkets201ResponseDataRelationshipsTaxCalculator{}
	return &this
}

// NewPOSTMarkets201ResponseDataRelationshipsTaxCalculatorWithDefaults instantiates a new POSTMarkets201ResponseDataRelationshipsTaxCalculator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMarkets201ResponseDataRelationshipsTaxCalculatorWithDefaults() *POSTMarkets201ResponseDataRelationshipsTaxCalculator {
	this := POSTMarkets201ResponseDataRelationshipsTaxCalculator{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculator) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculator) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculator) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculator) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculator) GetData() POSTMarkets201ResponseDataRelationshipsTaxCalculatorData {
	if o == nil || IsNil(o.Data) {
		var ret POSTMarkets201ResponseDataRelationshipsTaxCalculatorData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculator) GetDataOk() (*POSTMarkets201ResponseDataRelationshipsTaxCalculatorData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculator) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTMarkets201ResponseDataRelationshipsTaxCalculatorData and assigns it to the Data field.
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculator) SetData(v POSTMarkets201ResponseDataRelationshipsTaxCalculatorData) {
	o.Data = &v
}

func (o POSTMarkets201ResponseDataRelationshipsTaxCalculator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMarkets201ResponseDataRelationshipsTaxCalculator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculator struct {
	value *POSTMarkets201ResponseDataRelationshipsTaxCalculator
	isSet bool
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculator) Get() *POSTMarkets201ResponseDataRelationshipsTaxCalculator {
	return v.value
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculator) Set(val *POSTMarkets201ResponseDataRelationshipsTaxCalculator) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculator) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMarkets201ResponseDataRelationshipsTaxCalculator(val *POSTMarkets201ResponseDataRelationshipsTaxCalculator) *NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculator {
	return &NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculator{value: val, isSet: true}
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
