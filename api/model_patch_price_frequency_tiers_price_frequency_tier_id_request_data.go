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

// checks if the PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData{}

// PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData struct for PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData
type PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                           `json:"id"`
	Attributes    PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes     `json:"attributes"`
	Relationships *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData instantiates a new PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData(type_ interface{}, id interface{}, attributes PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData {
	this := PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataWithDefaults instantiates a new PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataWithDefaults() *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData {
	this := PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) GetAttributes() PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes {
	if o == nil {
		var ret PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) GetAttributesOk() (*PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) SetAttributes(v PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) GetRelationships() PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) GetRelationshipsOk() (*PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataRelationships and assigns it to the Relationships field.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) SetRelationships(v PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataRelationships) {
	o.Relationships = &v
}

func (o PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData struct {
	value *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData
	isSet bool
}

func (v NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) Get() *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData {
	return v.value
}

func (v *NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) Set(val *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData(val *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) *NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData {
	return &NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
