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

// AvalaraAccountCreateDataRelationships struct for AvalaraAccountCreateDataRelationships
type AvalaraAccountCreateDataRelationships struct {
	TaxCategories *AvalaraAccountCreateDataRelationshipsTaxCategories `json:"tax_categories,omitempty"`
}

// NewAvalaraAccountCreateDataRelationships instantiates a new AvalaraAccountCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccountCreateDataRelationships() *AvalaraAccountCreateDataRelationships {
	this := AvalaraAccountCreateDataRelationships{}
	return &this
}

// NewAvalaraAccountCreateDataRelationshipsWithDefaults instantiates a new AvalaraAccountCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountCreateDataRelationshipsWithDefaults() *AvalaraAccountCreateDataRelationships {
	this := AvalaraAccountCreateDataRelationships{}
	return &this
}

// GetTaxCategories returns the TaxCategories field value if set, zero value otherwise.
func (o *AvalaraAccountCreateDataRelationships) GetTaxCategories() AvalaraAccountCreateDataRelationshipsTaxCategories {
	if o == nil || o.TaxCategories == nil {
		var ret AvalaraAccountCreateDataRelationshipsTaxCategories
		return ret
	}
	return *o.TaxCategories
}

// GetTaxCategoriesOk returns a tuple with the TaxCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountCreateDataRelationships) GetTaxCategoriesOk() (*AvalaraAccountCreateDataRelationshipsTaxCategories, bool) {
	if o == nil || o.TaxCategories == nil {
		return nil, false
	}
	return o.TaxCategories, true
}

// HasTaxCategories returns a boolean if a field has been set.
func (o *AvalaraAccountCreateDataRelationships) HasTaxCategories() bool {
	if o != nil && o.TaxCategories != nil {
		return true
	}

	return false
}

// SetTaxCategories gets a reference to the given AvalaraAccountCreateDataRelationshipsTaxCategories and assigns it to the TaxCategories field.
func (o *AvalaraAccountCreateDataRelationships) SetTaxCategories(v AvalaraAccountCreateDataRelationshipsTaxCategories) {
	o.TaxCategories = &v
}

func (o AvalaraAccountCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaxCategories != nil {
		toSerialize["tax_categories"] = o.TaxCategories
	}
	return json.Marshal(toSerialize)
}

type NullableAvalaraAccountCreateDataRelationships struct {
	value *AvalaraAccountCreateDataRelationships
	isSet bool
}

func (v NullableAvalaraAccountCreateDataRelationships) Get() *AvalaraAccountCreateDataRelationships {
	return v.value
}

func (v *NullableAvalaraAccountCreateDataRelationships) Set(val *AvalaraAccountCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccountCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccountCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccountCreateDataRelationships(val *AvalaraAccountCreateDataRelationships) *NullableAvalaraAccountCreateDataRelationships {
	return &NullableAvalaraAccountCreateDataRelationships{value: val, isSet: true}
}

func (v NullableAvalaraAccountCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccountCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
