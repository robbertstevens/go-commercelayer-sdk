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

// checks if the PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData{}

// PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData struct for PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData
type PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                         `json:"id"`
	Attributes    PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestDataAttributes `json:"attributes"`
	Relationships *POSTManualTaxCalculatorsRequestDataRelationships                   `json:"relationships,omitempty"`
}

// NewPATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData instantiates a new PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData(type_ interface{}, id interface{}, attributes PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestDataAttributes) *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData {
	this := PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHManualTaxCalculatorsManualTaxCalculatorIdRequestDataWithDefaults instantiates a new PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHManualTaxCalculatorsManualTaxCalculatorIdRequestDataWithDefaults() *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData {
	this := PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) GetAttributes() PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestDataAttributes {
	if o == nil {
		var ret PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) GetAttributesOk() (*PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) SetAttributes(v PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) GetRelationships() POSTManualTaxCalculatorsRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTManualTaxCalculatorsRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) GetRelationshipsOk() (*POSTManualTaxCalculatorsRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTManualTaxCalculatorsRequestDataRelationships and assigns it to the Relationships field.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) SetRelationships(v POSTManualTaxCalculatorsRequestDataRelationships) {
	o.Relationships = &v
}

func (o PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData struct {
	value *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData
	isSet bool
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) Get() *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData {
	return v.value
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) Set(val *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData(val *PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) *NullablePATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData {
	return &NullablePATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
