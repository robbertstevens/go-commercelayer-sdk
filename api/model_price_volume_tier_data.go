/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PriceVolumeTierData struct for PriceVolumeTierData
type PriceVolumeTierData struct {
	// The resource's type
	Type          string                                      `json:"type"`
	Attributes    GETPriceTiers200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *PriceTierDataRelationships                 `json:"relationships,omitempty"`
}

// NewPriceVolumeTierData instantiates a new PriceVolumeTierData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceVolumeTierData(type_ string, attributes GETPriceTiers200ResponseDataInnerAttributes) *PriceVolumeTierData {
	this := PriceVolumeTierData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPriceVolumeTierDataWithDefaults instantiates a new PriceVolumeTierData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceVolumeTierDataWithDefaults() *PriceVolumeTierData {
	this := PriceVolumeTierData{}
	return &this
}

// GetType returns the Type field value
func (o *PriceVolumeTierData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceVolumeTierData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PriceVolumeTierData) GetAttributes() GETPriceTiers200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETPriceTiers200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierData) GetAttributesOk() (*GETPriceTiers200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceVolumeTierData) SetAttributes(v GETPriceTiers200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PriceVolumeTierData) GetRelationships() PriceTierDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PriceTierDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierData) GetRelationshipsOk() (*PriceTierDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceVolumeTierData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PriceTierDataRelationships and assigns it to the Relationships field.
func (o *PriceVolumeTierData) SetRelationships(v PriceTierDataRelationships) {
	o.Relationships = &v
}

func (o PriceVolumeTierData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullablePriceVolumeTierData struct {
	value *PriceVolumeTierData
	isSet bool
}

func (v NullablePriceVolumeTierData) Get() *PriceVolumeTierData {
	return v.value
}

func (v *NullablePriceVolumeTierData) Set(val *PriceVolumeTierData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceVolumeTierData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceVolumeTierData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceVolumeTierData(val *PriceVolumeTierData) *NullablePriceVolumeTierData {
	return &NullablePriceVolumeTierData{value: val, isSet: true}
}

func (v NullablePriceVolumeTierData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceVolumeTierData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
