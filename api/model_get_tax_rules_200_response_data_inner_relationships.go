/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETTaxRules200ResponseDataInnerRelationships struct for GETTaxRules200ResponseDataInnerRelationships
type GETTaxRules200ResponseDataInnerRelationships struct {
	ManualTaxCalculator *GETAddresses200ResponseDataInnerRelationshipsGeocoder `json:"manual_tax_calculator,omitempty"`
}

// NewGETTaxRules200ResponseDataInnerRelationships instantiates a new GETTaxRules200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETTaxRules200ResponseDataInnerRelationships() *GETTaxRules200ResponseDataInnerRelationships {
	this := GETTaxRules200ResponseDataInnerRelationships{}
	return &this
}

// NewGETTaxRules200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETTaxRules200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETTaxRules200ResponseDataInnerRelationshipsWithDefaults() *GETTaxRules200ResponseDataInnerRelationships {
	this := GETTaxRules200ResponseDataInnerRelationships{}
	return &this
}

// GetManualTaxCalculator returns the ManualTaxCalculator field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerRelationships) GetManualTaxCalculator() GETAddresses200ResponseDataInnerRelationshipsGeocoder {
	if o == nil || o.ManualTaxCalculator == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoder
		return ret
	}
	return *o.ManualTaxCalculator
}

// GetManualTaxCalculatorOk returns a tuple with the ManualTaxCalculator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerRelationships) GetManualTaxCalculatorOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool) {
	if o == nil || o.ManualTaxCalculator == nil {
		return nil, false
	}
	return o.ManualTaxCalculator, true
}

// HasManualTaxCalculator returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerRelationships) HasManualTaxCalculator() bool {
	if o != nil && o.ManualTaxCalculator != nil {
		return true
	}

	return false
}

// SetManualTaxCalculator gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoder and assigns it to the ManualTaxCalculator field.
func (o *GETTaxRules200ResponseDataInnerRelationships) SetManualTaxCalculator(v GETAddresses200ResponseDataInnerRelationshipsGeocoder) {
	o.ManualTaxCalculator = &v
}

func (o GETTaxRules200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ManualTaxCalculator != nil {
		toSerialize["manual_tax_calculator"] = o.ManualTaxCalculator
	}
	return json.Marshal(toSerialize)
}

type NullableGETTaxRules200ResponseDataInnerRelationships struct {
	value *GETTaxRules200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETTaxRules200ResponseDataInnerRelationships) Get() *GETTaxRules200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETTaxRules200ResponseDataInnerRelationships) Set(val *GETTaxRules200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETTaxRules200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETTaxRules200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETTaxRules200ResponseDataInnerRelationships(val *GETTaxRules200ResponseDataInnerRelationships) *NullableGETTaxRules200ResponseDataInnerRelationships {
	return &NullableGETTaxRules200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETTaxRules200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETTaxRules200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


