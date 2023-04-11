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

// checks if the PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships{}

// PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships struct for PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships
type PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships struct {
	Sku *POSTInStockSubscriptionsRequestDataRelationshipsSku `json:"sku,omitempty"`
}

// NewPATCHTaxCategoriesTaxCategoryIdRequestDataRelationships instantiates a new PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHTaxCategoriesTaxCategoryIdRequestDataRelationships() *PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships {
	this := PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships{}
	return &this
}

// NewPATCHTaxCategoriesTaxCategoryIdRequestDataRelationshipsWithDefaults instantiates a new PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHTaxCategoriesTaxCategoryIdRequestDataRelationshipsWithDefaults() *PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships {
	this := PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) GetSku() POSTInStockSubscriptionsRequestDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret POSTInStockSubscriptionsRequestDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) GetSkuOk() (*POSTInStockSubscriptionsRequestDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given POSTInStockSubscriptionsRequestDataRelationshipsSku and assigns it to the Sku field.
func (o *PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) SetSku(v POSTInStockSubscriptionsRequestDataRelationshipsSku) {
	o.Sku = &v
}

func (o PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullablePATCHTaxCategoriesTaxCategoryIdRequestDataRelationships struct {
	value *PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships
	isSet bool
}

func (v NullablePATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) Get() *PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships {
	return v.value
}

func (v *NullablePATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) Set(val *PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHTaxCategoriesTaxCategoryIdRequestDataRelationships(val *PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) *NullablePATCHTaxCategoriesTaxCategoryIdRequestDataRelationships {
	return &NullablePATCHTaxCategoriesTaxCategoryIdRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
