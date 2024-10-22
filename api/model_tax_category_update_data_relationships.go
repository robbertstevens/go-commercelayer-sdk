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

// checks if the TaxCategoryUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxCategoryUpdateDataRelationships{}

// TaxCategoryUpdateDataRelationships struct for TaxCategoryUpdateDataRelationships
type TaxCategoryUpdateDataRelationships struct {
	Sku *InStockSubscriptionCreateDataRelationshipsSku `json:"sku,omitempty"`
}

// NewTaxCategoryUpdateDataRelationships instantiates a new TaxCategoryUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCategoryUpdateDataRelationships() *TaxCategoryUpdateDataRelationships {
	this := TaxCategoryUpdateDataRelationships{}
	return &this
}

// NewTaxCategoryUpdateDataRelationshipsWithDefaults instantiates a new TaxCategoryUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCategoryUpdateDataRelationshipsWithDefaults() *TaxCategoryUpdateDataRelationships {
	this := TaxCategoryUpdateDataRelationships{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *TaxCategoryUpdateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret InStockSubscriptionCreateDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCategoryUpdateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *TaxCategoryUpdateDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given InStockSubscriptionCreateDataRelationshipsSku and assigns it to the Sku field.
func (o *TaxCategoryUpdateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku) {
	o.Sku = &v
}

func (o TaxCategoryUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxCategoryUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullableTaxCategoryUpdateDataRelationships struct {
	value *TaxCategoryUpdateDataRelationships
	isSet bool
}

func (v NullableTaxCategoryUpdateDataRelationships) Get() *TaxCategoryUpdateDataRelationships {
	return v.value
}

func (v *NullableTaxCategoryUpdateDataRelationships) Set(val *TaxCategoryUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCategoryUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCategoryUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCategoryUpdateDataRelationships(val *TaxCategoryUpdateDataRelationships) *NullableTaxCategoryUpdateDataRelationships {
	return &NullableTaxCategoryUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableTaxCategoryUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCategoryUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
