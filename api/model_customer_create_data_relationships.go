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

// checks if the CustomerCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerCreateDataRelationships{}

// CustomerCreateDataRelationships struct for CustomerCreateDataRelationships
type CustomerCreateDataRelationships struct {
	CustomerGroup *CustomerCreateDataRelationshipsCustomerGroup `json:"customer_group,omitempty"`
	Tags          *AddressCreateDataRelationshipsTags           `json:"tags,omitempty"`
}

// NewCustomerCreateDataRelationships instantiates a new CustomerCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerCreateDataRelationships() *CustomerCreateDataRelationships {
	this := CustomerCreateDataRelationships{}
	return &this
}

// NewCustomerCreateDataRelationshipsWithDefaults instantiates a new CustomerCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCreateDataRelationshipsWithDefaults() *CustomerCreateDataRelationships {
	this := CustomerCreateDataRelationships{}
	return &this
}

// GetCustomerGroup returns the CustomerGroup field value if set, zero value otherwise.
func (o *CustomerCreateDataRelationships) GetCustomerGroup() CustomerCreateDataRelationshipsCustomerGroup {
	if o == nil || IsNil(o.CustomerGroup) {
		var ret CustomerCreateDataRelationshipsCustomerGroup
		return ret
	}
	return *o.CustomerGroup
}

// GetCustomerGroupOk returns a tuple with the CustomerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateDataRelationships) GetCustomerGroupOk() (*CustomerCreateDataRelationshipsCustomerGroup, bool) {
	if o == nil || IsNil(o.CustomerGroup) {
		return nil, false
	}
	return o.CustomerGroup, true
}

// HasCustomerGroup returns a boolean if a field has been set.
func (o *CustomerCreateDataRelationships) HasCustomerGroup() bool {
	if o != nil && !IsNil(o.CustomerGroup) {
		return true
	}

	return false
}

// SetCustomerGroup gets a reference to the given CustomerCreateDataRelationshipsCustomerGroup and assigns it to the CustomerGroup field.
func (o *CustomerCreateDataRelationships) SetCustomerGroup(v CustomerCreateDataRelationshipsCustomerGroup) {
	o.CustomerGroup = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CustomerCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressCreateDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CustomerCreateDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressCreateDataRelationshipsTags and assigns it to the Tags field.
func (o *CustomerCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags) {
	o.Tags = &v
}

func (o CustomerCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerGroup) {
		toSerialize["customer_group"] = o.CustomerGroup
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableCustomerCreateDataRelationships struct {
	value *CustomerCreateDataRelationships
	isSet bool
}

func (v NullableCustomerCreateDataRelationships) Get() *CustomerCreateDataRelationships {
	return v.value
}

func (v *NullableCustomerCreateDataRelationships) Set(val *CustomerCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCreateDataRelationships(val *CustomerCreateDataRelationships) *NullableCustomerCreateDataRelationships {
	return &NullableCustomerCreateDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
