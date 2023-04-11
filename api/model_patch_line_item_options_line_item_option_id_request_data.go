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

// checks if the PATCHLineItemOptionsLineItemOptionIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHLineItemOptionsLineItemOptionIdRequestData{}

// PATCHLineItemOptionsLineItemOptionIdRequestData struct for PATCHLineItemOptionsLineItemOptionIdRequestData
type PATCHLineItemOptionsLineItemOptionIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                   `json:"id"`
	Attributes    PATCHLineItemOptionsLineItemOptionIdRequestDataAttributes     `json:"attributes"`
	Relationships *PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPATCHLineItemOptionsLineItemOptionIdRequestData instantiates a new PATCHLineItemOptionsLineItemOptionIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHLineItemOptionsLineItemOptionIdRequestData(type_ interface{}, id interface{}, attributes PATCHLineItemOptionsLineItemOptionIdRequestDataAttributes) *PATCHLineItemOptionsLineItemOptionIdRequestData {
	this := PATCHLineItemOptionsLineItemOptionIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHLineItemOptionsLineItemOptionIdRequestDataWithDefaults instantiates a new PATCHLineItemOptionsLineItemOptionIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHLineItemOptionsLineItemOptionIdRequestDataWithDefaults() *PATCHLineItemOptionsLineItemOptionIdRequestData {
	this := PATCHLineItemOptionsLineItemOptionIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHLineItemOptionsLineItemOptionIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemOptionsLineItemOptionIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHLineItemOptionsLineItemOptionIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHLineItemOptionsLineItemOptionIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemOptionsLineItemOptionIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHLineItemOptionsLineItemOptionIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHLineItemOptionsLineItemOptionIdRequestData) GetAttributes() PATCHLineItemOptionsLineItemOptionIdRequestDataAttributes {
	if o == nil {
		var ret PATCHLineItemOptionsLineItemOptionIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHLineItemOptionsLineItemOptionIdRequestData) GetAttributesOk() (*PATCHLineItemOptionsLineItemOptionIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHLineItemOptionsLineItemOptionIdRequestData) SetAttributes(v PATCHLineItemOptionsLineItemOptionIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHLineItemOptionsLineItemOptionIdRequestData) GetRelationships() PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHLineItemOptionsLineItemOptionIdRequestData) GetRelationshipsOk() (*PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionIdRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships and assigns it to the Relationships field.
func (o *PATCHLineItemOptionsLineItemOptionIdRequestData) SetRelationships(v PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) {
	o.Relationships = &v
}

func (o PATCHLineItemOptionsLineItemOptionIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHLineItemOptionsLineItemOptionIdRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePATCHLineItemOptionsLineItemOptionIdRequestData struct {
	value *PATCHLineItemOptionsLineItemOptionIdRequestData
	isSet bool
}

func (v NullablePATCHLineItemOptionsLineItemOptionIdRequestData) Get() *PATCHLineItemOptionsLineItemOptionIdRequestData {
	return v.value
}

func (v *NullablePATCHLineItemOptionsLineItemOptionIdRequestData) Set(val *PATCHLineItemOptionsLineItemOptionIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHLineItemOptionsLineItemOptionIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHLineItemOptionsLineItemOptionIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHLineItemOptionsLineItemOptionIdRequestData(val *PATCHLineItemOptionsLineItemOptionIdRequestData) *NullablePATCHLineItemOptionsLineItemOptionIdRequestData {
	return &NullablePATCHLineItemOptionsLineItemOptionIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHLineItemOptionsLineItemOptionIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHLineItemOptionsLineItemOptionIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
