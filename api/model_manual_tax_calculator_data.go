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

// checks if the ManualTaxCalculatorData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualTaxCalculatorData{}

// ManualTaxCalculatorData struct for ManualTaxCalculatorData
type ManualTaxCalculatorData struct {
	// The resource's type
	Type          interface{}                                                           `json:"type"`
	Attributes    GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes `json:"attributes"`
	Relationships *ManualTaxCalculatorDataRelationships                                 `json:"relationships,omitempty"`
}

// NewManualTaxCalculatorData instantiates a new ManualTaxCalculatorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualTaxCalculatorData(type_ interface{}, attributes GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) *ManualTaxCalculatorData {
	this := ManualTaxCalculatorData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewManualTaxCalculatorDataWithDefaults instantiates a new ManualTaxCalculatorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualTaxCalculatorDataWithDefaults() *ManualTaxCalculatorData {
	this := ManualTaxCalculatorData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ManualTaxCalculatorData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualTaxCalculatorData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ManualTaxCalculatorData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ManualTaxCalculatorData) GetAttributes() GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes {
	if o == nil {
		var ret GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorData) GetAttributesOk() (*GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ManualTaxCalculatorData) SetAttributes(v GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ManualTaxCalculatorData) GetRelationships() ManualTaxCalculatorDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ManualTaxCalculatorDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorData) GetRelationshipsOk() (*ManualTaxCalculatorDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ManualTaxCalculatorData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ManualTaxCalculatorDataRelationships and assigns it to the Relationships field.
func (o *ManualTaxCalculatorData) SetRelationships(v ManualTaxCalculatorDataRelationships) {
	o.Relationships = &v
}

func (o ManualTaxCalculatorData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualTaxCalculatorData) ToMap() (map[string]interface{}, error) {
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

type NullableManualTaxCalculatorData struct {
	value *ManualTaxCalculatorData
	isSet bool
}

func (v NullableManualTaxCalculatorData) Get() *ManualTaxCalculatorData {
	return v.value
}

func (v *NullableManualTaxCalculatorData) Set(val *ManualTaxCalculatorData) {
	v.value = val
	v.isSet = true
}

func (v NullableManualTaxCalculatorData) IsSet() bool {
	return v.isSet
}

func (v *NullableManualTaxCalculatorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualTaxCalculatorData(val *ManualTaxCalculatorData) *NullableManualTaxCalculatorData {
	return &NullableManualTaxCalculatorData{value: val, isSet: true}
}

func (v NullableManualTaxCalculatorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualTaxCalculatorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
