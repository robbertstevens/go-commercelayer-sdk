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

// checks if the StockTransferData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockTransferData{}

// StockTransferData struct for StockTransferData
type StockTransferData struct {
	// The resource's type
	Type          interface{}                                               `json:"type"`
	Attributes    GETStockTransfersStockTransferId200ResponseDataAttributes `json:"attributes"`
	Relationships *StockTransferDataRelationships                           `json:"relationships,omitempty"`
}

// NewStockTransferData instantiates a new StockTransferData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockTransferData(type_ interface{}, attributes GETStockTransfersStockTransferId200ResponseDataAttributes) *StockTransferData {
	this := StockTransferData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewStockTransferDataWithDefaults instantiates a new StockTransferData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockTransferDataWithDefaults() *StockTransferData {
	this := StockTransferData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *StockTransferData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StockTransferData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StockTransferData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *StockTransferData) GetAttributes() GETStockTransfersStockTransferId200ResponseDataAttributes {
	if o == nil {
		var ret GETStockTransfersStockTransferId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StockTransferData) GetAttributesOk() (*GETStockTransfersStockTransferId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StockTransferData) SetAttributes(v GETStockTransfersStockTransferId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StockTransferData) GetRelationships() StockTransferDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret StockTransferDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferData) GetRelationshipsOk() (*StockTransferDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StockTransferData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given StockTransferDataRelationships and assigns it to the Relationships field.
func (o *StockTransferData) SetRelationships(v StockTransferDataRelationships) {
	o.Relationships = &v
}

func (o StockTransferData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockTransferData) ToMap() (map[string]interface{}, error) {
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

type NullableStockTransferData struct {
	value *StockTransferData
	isSet bool
}

func (v NullableStockTransferData) Get() *StockTransferData {
	return v.value
}

func (v *NullableStockTransferData) Set(val *StockTransferData) {
	v.value = val
	v.isSet = true
}

func (v NullableStockTransferData) IsSet() bool {
	return v.isSet
}

func (v *NullableStockTransferData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockTransferData(val *StockTransferData) *NullableStockTransferData {
	return &NullableStockTransferData{value: val, isSet: true}
}

func (v NullableStockTransferData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockTransferData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
