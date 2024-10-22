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

// checks if the LineItemDataRelationshipsStockTransfers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemDataRelationshipsStockTransfers{}

// LineItemDataRelationshipsStockTransfers struct for LineItemDataRelationshipsStockTransfers
type LineItemDataRelationshipsStockTransfers struct {
	Data *LineItemDataRelationshipsStockTransfersData `json:"data,omitempty"`
}

// NewLineItemDataRelationshipsStockTransfers instantiates a new LineItemDataRelationshipsStockTransfers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemDataRelationshipsStockTransfers() *LineItemDataRelationshipsStockTransfers {
	this := LineItemDataRelationshipsStockTransfers{}
	return &this
}

// NewLineItemDataRelationshipsStockTransfersWithDefaults instantiates a new LineItemDataRelationshipsStockTransfers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemDataRelationshipsStockTransfersWithDefaults() *LineItemDataRelationshipsStockTransfers {
	this := LineItemDataRelationshipsStockTransfers{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LineItemDataRelationshipsStockTransfers) GetData() LineItemDataRelationshipsStockTransfersData {
	if o == nil || IsNil(o.Data) {
		var ret LineItemDataRelationshipsStockTransfersData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemDataRelationshipsStockTransfers) GetDataOk() (*LineItemDataRelationshipsStockTransfersData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LineItemDataRelationshipsStockTransfers) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LineItemDataRelationshipsStockTransfersData and assigns it to the Data field.
func (o *LineItemDataRelationshipsStockTransfers) SetData(v LineItemDataRelationshipsStockTransfersData) {
	o.Data = &v
}

func (o LineItemDataRelationshipsStockTransfers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemDataRelationshipsStockTransfers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLineItemDataRelationshipsStockTransfers struct {
	value *LineItemDataRelationshipsStockTransfers
	isSet bool
}

func (v NullableLineItemDataRelationshipsStockTransfers) Get() *LineItemDataRelationshipsStockTransfers {
	return v.value
}

func (v *NullableLineItemDataRelationshipsStockTransfers) Set(val *LineItemDataRelationshipsStockTransfers) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemDataRelationshipsStockTransfers) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemDataRelationshipsStockTransfers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemDataRelationshipsStockTransfers(val *LineItemDataRelationshipsStockTransfers) *NullableLineItemDataRelationshipsStockTransfers {
	return &NullableLineItemDataRelationshipsStockTransfers{value: val, isSet: true}
}

func (v NullableLineItemDataRelationshipsStockTransfers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemDataRelationshipsStockTransfers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
