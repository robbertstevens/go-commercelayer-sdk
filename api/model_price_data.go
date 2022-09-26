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

// PriceData struct for PriceData
type PriceData struct {
	// The resource's type
	Type          string                                  `json:"type"`
	Attributes    GETPrices200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *PriceDataRelationships                 `json:"relationships,omitempty"`
}

// NewPriceData instantiates a new PriceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceData(type_ string, attributes GETPrices200ResponseDataInnerAttributes) *PriceData {
	this := PriceData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPriceDataWithDefaults instantiates a new PriceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceDataWithDefaults() *PriceData {
	this := PriceData{}
	return &this
}

// GetType returns the Type field value
func (o *PriceData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PriceData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PriceData) GetAttributes() GETPrices200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETPrices200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceData) GetAttributesOk() (*GETPrices200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceData) SetAttributes(v GETPrices200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PriceData) GetRelationships() PriceDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PriceDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceData) GetRelationshipsOk() (*PriceDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PriceDataRelationships and assigns it to the Relationships field.
func (o *PriceData) SetRelationships(v PriceDataRelationships) {
	o.Relationships = &v
}

func (o PriceData) MarshalJSON() ([]byte, error) {
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

type NullablePriceData struct {
	value *PriceData
	isSet bool
}

func (v NullablePriceData) Get() *PriceData {
	return v.value
}

func (v *NullablePriceData) Set(val *PriceData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceData(val *PriceData) *NullablePriceData {
	return &NullablePriceData{value: val, isSet: true}
}

func (v NullablePriceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
