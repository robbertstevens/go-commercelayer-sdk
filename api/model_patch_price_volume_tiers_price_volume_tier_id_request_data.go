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

// checks if the PATCHPriceVolumeTiersPriceVolumeTierIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPriceVolumeTiersPriceVolumeTierIdRequestData{}

// PATCHPriceVolumeTiersPriceVolumeTierIdRequestData struct for PATCHPriceVolumeTiersPriceVolumeTierIdRequestData
type PATCHPriceVolumeTiersPriceVolumeTierIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                           `json:"id"`
	Attributes    PATCHPriceVolumeTiersPriceVolumeTierIdRequestDataAttributes           `json:"attributes"`
	Relationships *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPATCHPriceVolumeTiersPriceVolumeTierIdRequestData instantiates a new PATCHPriceVolumeTiersPriceVolumeTierIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPriceVolumeTiersPriceVolumeTierIdRequestData(type_ interface{}, id interface{}, attributes PATCHPriceVolumeTiersPriceVolumeTierIdRequestDataAttributes) *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData {
	this := PATCHPriceVolumeTiersPriceVolumeTierIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHPriceVolumeTiersPriceVolumeTierIdRequestDataWithDefaults instantiates a new PATCHPriceVolumeTiersPriceVolumeTierIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPriceVolumeTiersPriceVolumeTierIdRequestDataWithDefaults() *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData {
	this := PATCHPriceVolumeTiersPriceVolumeTierIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) GetAttributes() PATCHPriceVolumeTiersPriceVolumeTierIdRequestDataAttributes {
	if o == nil {
		var ret PATCHPriceVolumeTiersPriceVolumeTierIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) GetAttributesOk() (*PATCHPriceVolumeTiersPriceVolumeTierIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) SetAttributes(v PATCHPriceVolumeTiersPriceVolumeTierIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) GetRelationships() PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) GetRelationshipsOk() (*PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataRelationships and assigns it to the Relationships field.
func (o *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) SetRelationships(v PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataRelationships) {
	o.Relationships = &v
}

func (o PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePATCHPriceVolumeTiersPriceVolumeTierIdRequestData struct {
	value *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData
	isSet bool
}

func (v NullablePATCHPriceVolumeTiersPriceVolumeTierIdRequestData) Get() *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData {
	return v.value
}

func (v *NullablePATCHPriceVolumeTiersPriceVolumeTierIdRequestData) Set(val *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPriceVolumeTiersPriceVolumeTierIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPriceVolumeTiersPriceVolumeTierIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPriceVolumeTiersPriceVolumeTierIdRequestData(val *PATCHPriceVolumeTiersPriceVolumeTierIdRequestData) *NullablePATCHPriceVolumeTiersPriceVolumeTierIdRequestData {
	return &NullablePATCHPriceVolumeTiersPriceVolumeTierIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHPriceVolumeTiersPriceVolumeTierIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPriceVolumeTiersPriceVolumeTierIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
