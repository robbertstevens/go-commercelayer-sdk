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

// checks if the SkuCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuCreateDataRelationships{}

// SkuCreateDataRelationships struct for SkuCreateDataRelationships
type SkuCreateDataRelationships struct {
	ShippingCategory ShipmentCreateDataRelationshipsShippingCategory `json:"shipping_category"`
	Tags             *AddressCreateDataRelationshipsTags             `json:"tags,omitempty"`
}

// NewSkuCreateDataRelationships instantiates a new SkuCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuCreateDataRelationships(shippingCategory ShipmentCreateDataRelationshipsShippingCategory) *SkuCreateDataRelationships {
	this := SkuCreateDataRelationships{}
	this.ShippingCategory = shippingCategory
	return &this
}

// NewSkuCreateDataRelationshipsWithDefaults instantiates a new SkuCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuCreateDataRelationshipsWithDefaults() *SkuCreateDataRelationships {
	this := SkuCreateDataRelationships{}
	return &this
}

// GetShippingCategory returns the ShippingCategory field value
func (o *SkuCreateDataRelationships) GetShippingCategory() ShipmentCreateDataRelationshipsShippingCategory {
	if o == nil {
		var ret ShipmentCreateDataRelationshipsShippingCategory
		return ret
	}

	return o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value
// and a boolean to check if the value has been set.
func (o *SkuCreateDataRelationships) GetShippingCategoryOk() (*ShipmentCreateDataRelationshipsShippingCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingCategory, true
}

// SetShippingCategory sets field value
func (o *SkuCreateDataRelationships) SetShippingCategory(v ShipmentCreateDataRelationshipsShippingCategory) {
	o.ShippingCategory = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SkuCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressCreateDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SkuCreateDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressCreateDataRelationshipsTags and assigns it to the Tags field.
func (o *SkuCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags) {
	o.Tags = &v
}

func (o SkuCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipping_category"] = o.ShippingCategory
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableSkuCreateDataRelationships struct {
	value *SkuCreateDataRelationships
	isSet bool
}

func (v NullableSkuCreateDataRelationships) Get() *SkuCreateDataRelationships {
	return v.value
}

func (v *NullableSkuCreateDataRelationships) Set(val *SkuCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuCreateDataRelationships(val *SkuCreateDataRelationships) *NullableSkuCreateDataRelationships {
	return &NullableSkuCreateDataRelationships{value: val, isSet: true}
}

func (v NullableSkuCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
