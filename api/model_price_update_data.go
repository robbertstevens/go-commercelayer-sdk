/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PriceUpdateData struct for PriceUpdateData
type PriceUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                      `json:"id"`
	Attributes    PATCHPricesPriceId200ResponseDataAttributes `json:"attributes"`
	Relationships *PriceUpdateDataRelationships               `json:"relationships,omitempty"`
}

// NewPriceUpdateData instantiates a new PriceUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceUpdateData(type_ string, id string, attributes PATCHPricesPriceId200ResponseDataAttributes) *PriceUpdateData {
	this := PriceUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPriceUpdateDataWithDefaults instantiates a new PriceUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceUpdateDataWithDefaults() *PriceUpdateData {
	this := PriceUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *PriceUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PriceUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PriceUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PriceUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PriceUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PriceUpdateData) GetAttributes() PATCHPricesPriceId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHPricesPriceId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceUpdateData) GetAttributesOk() (*PATCHPricesPriceId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceUpdateData) SetAttributes(v PATCHPricesPriceId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PriceUpdateData) GetRelationships() PriceUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PriceUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceUpdateData) GetRelationshipsOk() (*PriceUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PriceUpdateDataRelationships and assigns it to the Relationships field.
func (o *PriceUpdateData) SetRelationships(v PriceUpdateDataRelationships) {
	o.Relationships = &v
}

func (o PriceUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullablePriceUpdateData struct {
	value *PriceUpdateData
	isSet bool
}

func (v NullablePriceUpdateData) Get() *PriceUpdateData {
	return v.value
}

func (v *NullablePriceUpdateData) Set(val *PriceUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceUpdateData(val *PriceUpdateData) *NullablePriceUpdateData {
	return &NullablePriceUpdateData{value: val, isSet: true}
}

func (v NullablePriceUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
