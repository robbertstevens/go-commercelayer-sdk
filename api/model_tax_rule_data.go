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

// checks if the TaxRuleData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxRuleData{}

// TaxRuleData struct for TaxRuleData
type TaxRuleData struct {
	// The resource's type
	Type          interface{}                                   `json:"type"`
	Attributes    GETTaxRulesTaxRuleId200ResponseDataAttributes `json:"attributes"`
	Relationships *TaxRuleDataRelationships                     `json:"relationships,omitempty"`
}

// NewTaxRuleData instantiates a new TaxRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRuleData(type_ interface{}, attributes GETTaxRulesTaxRuleId200ResponseDataAttributes) *TaxRuleData {
	this := TaxRuleData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTaxRuleDataWithDefaults instantiates a new TaxRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleDataWithDefaults() *TaxRuleData {
	this := TaxRuleData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *TaxRuleData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxRuleData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxRuleData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TaxRuleData) GetAttributes() GETTaxRulesTaxRuleId200ResponseDataAttributes {
	if o == nil {
		var ret GETTaxRulesTaxRuleId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaxRuleData) GetAttributesOk() (*GETTaxRulesTaxRuleId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TaxRuleData) SetAttributes(v GETTaxRulesTaxRuleId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxRuleData) GetRelationships() TaxRuleDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret TaxRuleDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRuleData) GetRelationshipsOk() (*TaxRuleDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxRuleData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given TaxRuleDataRelationships and assigns it to the Relationships field.
func (o *TaxRuleData) SetRelationships(v TaxRuleDataRelationships) {
	o.Relationships = &v
}

func (o TaxRuleData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxRuleData) ToMap() (map[string]interface{}, error) {
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

type NullableTaxRuleData struct {
	value *TaxRuleData
	isSet bool
}

func (v NullableTaxRuleData) Get() *TaxRuleData {
	return v.value
}

func (v *NullableTaxRuleData) Set(val *TaxRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRuleData(val *TaxRuleData) *NullableTaxRuleData {
	return &NullableTaxRuleData{value: val, isSet: true}
}

func (v NullableTaxRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
