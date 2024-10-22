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

// checks if the TaxjarAccountUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxjarAccountUpdateData{}

// TaxjarAccountUpdateData struct for TaxjarAccountUpdateData
type TaxjarAccountUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                 `json:"id"`
	Attributes    PATCHTaxjarAccountsTaxjarAccountId200ResponseDataAttributes `json:"attributes"`
	Relationships *AvalaraAccountCreateDataRelationships                      `json:"relationships,omitempty"`
}

// NewTaxjarAccountUpdateData instantiates a new TaxjarAccountUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxjarAccountUpdateData(type_ interface{}, id interface{}, attributes PATCHTaxjarAccountsTaxjarAccountId200ResponseDataAttributes) *TaxjarAccountUpdateData {
	this := TaxjarAccountUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewTaxjarAccountUpdateDataWithDefaults instantiates a new TaxjarAccountUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxjarAccountUpdateDataWithDefaults() *TaxjarAccountUpdateData {
	this := TaxjarAccountUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *TaxjarAccountUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxjarAccountUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxjarAccountUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *TaxjarAccountUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxjarAccountUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaxjarAccountUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *TaxjarAccountUpdateData) GetAttributes() PATCHTaxjarAccountsTaxjarAccountId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHTaxjarAccountsTaxjarAccountId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaxjarAccountUpdateData) GetAttributesOk() (*PATCHTaxjarAccountsTaxjarAccountId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TaxjarAccountUpdateData) SetAttributes(v PATCHTaxjarAccountsTaxjarAccountId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxjarAccountUpdateData) GetRelationships() AvalaraAccountCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AvalaraAccountCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxjarAccountUpdateData) GetRelationshipsOk() (*AvalaraAccountCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxjarAccountUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AvalaraAccountCreateDataRelationships and assigns it to the Relationships field.
func (o *TaxjarAccountUpdateData) SetRelationships(v AvalaraAccountCreateDataRelationships) {
	o.Relationships = &v
}

func (o TaxjarAccountUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxjarAccountUpdateData) ToMap() (map[string]interface{}, error) {
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

type NullableTaxjarAccountUpdateData struct {
	value *TaxjarAccountUpdateData
	isSet bool
}

func (v NullableTaxjarAccountUpdateData) Get() *TaxjarAccountUpdateData {
	return v.value
}

func (v *NullableTaxjarAccountUpdateData) Set(val *TaxjarAccountUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxjarAccountUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxjarAccountUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxjarAccountUpdateData(val *TaxjarAccountUpdateData) *NullableTaxjarAccountUpdateData {
	return &NullableTaxjarAccountUpdateData{value: val, isSet: true}
}

func (v NullableTaxjarAccountUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxjarAccountUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
