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

// checks if the PriceFrequencyTierData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceFrequencyTierData{}

// PriceFrequencyTierData struct for PriceFrequencyTierData
type PriceFrequencyTierData struct {
	// The resource's type
	Type          interface{}                                                         `json:"type"`
	Attributes    GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes `json:"attributes"`
	Relationships *PriceFrequencyTierDataRelationships                                `json:"relationships,omitempty"`
}

// NewPriceFrequencyTierData instantiates a new PriceFrequencyTierData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceFrequencyTierData(type_ interface{}, attributes GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) *PriceFrequencyTierData {
	this := PriceFrequencyTierData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPriceFrequencyTierDataWithDefaults instantiates a new PriceFrequencyTierData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceFrequencyTierDataWithDefaults() *PriceFrequencyTierData {
	this := PriceFrequencyTierData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PriceFrequencyTierData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceFrequencyTierData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceFrequencyTierData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PriceFrequencyTierData) GetAttributes() GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes {
	if o == nil {
		var ret GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceFrequencyTierData) GetAttributesOk() (*GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceFrequencyTierData) SetAttributes(v GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PriceFrequencyTierData) GetRelationships() PriceFrequencyTierDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PriceFrequencyTierDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceFrequencyTierData) GetRelationshipsOk() (*PriceFrequencyTierDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceFrequencyTierData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PriceFrequencyTierDataRelationships and assigns it to the Relationships field.
func (o *PriceFrequencyTierData) SetRelationships(v PriceFrequencyTierDataRelationships) {
	o.Relationships = &v
}

func (o PriceFrequencyTierData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceFrequencyTierData) ToMap() (map[string]interface{}, error) {
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

type NullablePriceFrequencyTierData struct {
	value *PriceFrequencyTierData
	isSet bool
}

func (v NullablePriceFrequencyTierData) Get() *PriceFrequencyTierData {
	return v.value
}

func (v *NullablePriceFrequencyTierData) Set(val *PriceFrequencyTierData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceFrequencyTierData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceFrequencyTierData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceFrequencyTierData(val *PriceFrequencyTierData) *NullablePriceFrequencyTierData {
	return &NullablePriceFrequencyTierData{value: val, isSet: true}
}

func (v NullablePriceFrequencyTierData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceFrequencyTierData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
