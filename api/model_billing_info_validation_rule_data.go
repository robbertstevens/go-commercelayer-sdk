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

// checks if the BillingInfoValidationRuleData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInfoValidationRuleData{}

// BillingInfoValidationRuleData struct for BillingInfoValidationRuleData
type BillingInfoValidationRuleData struct {
	// The resource's type
	Type          interface{}                                                                       `json:"type"`
	Attributes    GETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes `json:"attributes"`
	Relationships *BillingInfoValidationRuleDataRelationships                                       `json:"relationships,omitempty"`
}

// NewBillingInfoValidationRuleData instantiates a new BillingInfoValidationRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoValidationRuleData(type_ interface{}, attributes GETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes) *BillingInfoValidationRuleData {
	this := BillingInfoValidationRuleData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewBillingInfoValidationRuleDataWithDefaults instantiates a new BillingInfoValidationRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoValidationRuleDataWithDefaults() *BillingInfoValidationRuleData {
	this := BillingInfoValidationRuleData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *BillingInfoValidationRuleData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfoValidationRuleData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BillingInfoValidationRuleData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BillingInfoValidationRuleData) GetAttributes() GETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes {
	if o == nil {
		var ret GETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleData) GetAttributesOk() (*GETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BillingInfoValidationRuleData) SetAttributes(v GETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BillingInfoValidationRuleData) GetRelationships() BillingInfoValidationRuleDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BillingInfoValidationRuleDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleData) GetRelationshipsOk() (*BillingInfoValidationRuleDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BillingInfoValidationRuleData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BillingInfoValidationRuleDataRelationships and assigns it to the Relationships field.
func (o *BillingInfoValidationRuleData) SetRelationships(v BillingInfoValidationRuleDataRelationships) {
	o.Relationships = &v
}

func (o BillingInfoValidationRuleData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInfoValidationRuleData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableBillingInfoValidationRuleData struct {
	value *BillingInfoValidationRuleData
	isSet bool
}

func (v NullableBillingInfoValidationRuleData) Get() *BillingInfoValidationRuleData {
	return v.value
}

func (v *NullableBillingInfoValidationRuleData) Set(val *BillingInfoValidationRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoValidationRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoValidationRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoValidationRuleData(val *BillingInfoValidationRuleData) *NullableBillingInfoValidationRuleData {
	return &NullableBillingInfoValidationRuleData{value: val, isSet: true}
}

func (v NullableBillingInfoValidationRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoValidationRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
