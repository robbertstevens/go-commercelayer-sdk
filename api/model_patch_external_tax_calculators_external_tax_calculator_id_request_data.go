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

// checks if the PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData{}

// PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData struct for PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData
type PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                             `json:"id"`
	Attributes    PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestDataAttributes `json:"attributes"`
	Relationships interface{}                                                             `json:"relationships,omitempty"`
}

// NewPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData instantiates a new PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData(type_ interface{}, id interface{}, attributes PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestDataAttributes) *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData {
	this := PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestDataWithDefaults instantiates a new PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestDataWithDefaults() *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData {
	this := PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) GetAttributes() PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestDataAttributes {
	if o == nil {
		var ret PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) GetAttributesOk() (*PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) SetAttributes(v PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) HasRelationships() bool {
	if o != nil && IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData struct {
	value *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData
	isSet bool
}

func (v NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) Get() *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData {
	return v.value
}

func (v *NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) Set(val *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData(val *PATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) *NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData {
	return &NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
