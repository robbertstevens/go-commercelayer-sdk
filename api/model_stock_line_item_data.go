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

// StockLineItemData struct for StockLineItemData
type StockLineItemData struct {
	// The resource's type
	Type          string                                          `json:"type"`
	Attributes    GETStockLineItems200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *StockLineItemDataRelationships                 `json:"relationships,omitempty"`
}

// NewStockLineItemData instantiates a new StockLineItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockLineItemData(type_ string, attributes GETStockLineItems200ResponseDataInnerAttributes) *StockLineItemData {
	this := StockLineItemData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewStockLineItemDataWithDefaults instantiates a new StockLineItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockLineItemDataWithDefaults() *StockLineItemData {
	this := StockLineItemData{}
	return &this
}

// GetType returns the Type field value
func (o *StockLineItemData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StockLineItemData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StockLineItemData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *StockLineItemData) GetAttributes() GETStockLineItems200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETStockLineItems200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StockLineItemData) GetAttributesOk() (*GETStockLineItems200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StockLineItemData) SetAttributes(v GETStockLineItems200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StockLineItemData) GetRelationships() StockLineItemDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret StockLineItemDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLineItemData) GetRelationshipsOk() (*StockLineItemDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StockLineItemData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given StockLineItemDataRelationships and assigns it to the Relationships field.
func (o *StockLineItemData) SetRelationships(v StockLineItemDataRelationships) {
	o.Relationships = &v
}

func (o StockLineItemData) MarshalJSON() ([]byte, error) {
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

type NullableStockLineItemData struct {
	value *StockLineItemData
	isSet bool
}

func (v NullableStockLineItemData) Get() *StockLineItemData {
	return v.value
}

func (v *NullableStockLineItemData) Set(val *StockLineItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableStockLineItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableStockLineItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockLineItemData(val *StockLineItemData) *NullableStockLineItemData {
	return &NullableStockLineItemData{value: val, isSet: true}
}

func (v NullableStockLineItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockLineItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
