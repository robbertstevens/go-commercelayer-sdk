/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the OrderSubscriptionItemData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSubscriptionItemData{}

// OrderSubscriptionItemData struct for OrderSubscriptionItemData
type OrderSubscriptionItemData struct {
	// The resource's type
	Type          interface{}                                                               `json:"type"`
	Attributes    GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes `json:"attributes"`
	Relationships *OrderSubscriptionItemDataRelationships                                   `json:"relationships,omitempty"`
}

// NewOrderSubscriptionItemData instantiates a new OrderSubscriptionItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionItemData(type_ interface{}, attributes GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) *OrderSubscriptionItemData {
	this := OrderSubscriptionItemData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewOrderSubscriptionItemDataWithDefaults instantiates a new OrderSubscriptionItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionItemDataWithDefaults() *OrderSubscriptionItemData {
	this := OrderSubscriptionItemData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OrderSubscriptionItemData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderSubscriptionItemData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderSubscriptionItemData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *OrderSubscriptionItemData) GetAttributes() GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes {
	if o == nil {
		var ret GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionItemData) GetAttributesOk() (*GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrderSubscriptionItemData) SetAttributes(v GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OrderSubscriptionItemData) GetRelationships() OrderSubscriptionItemDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret OrderSubscriptionItemDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionItemData) GetRelationshipsOk() (*OrderSubscriptionItemDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrderSubscriptionItemData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given OrderSubscriptionItemDataRelationships and assigns it to the Relationships field.
func (o *OrderSubscriptionItemData) SetRelationships(v OrderSubscriptionItemDataRelationships) {
	o.Relationships = &v
}

func (o OrderSubscriptionItemData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSubscriptionItemData) ToMap() (map[string]interface{}, error) {
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

type NullableOrderSubscriptionItemData struct {
	value *OrderSubscriptionItemData
	isSet bool
}

func (v NullableOrderSubscriptionItemData) Get() *OrderSubscriptionItemData {
	return v.value
}

func (v *NullableOrderSubscriptionItemData) Set(val *OrderSubscriptionItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionItemData(val *OrderSubscriptionItemData) *NullableOrderSubscriptionItemData {
	return &NullableOrderSubscriptionItemData{value: val, isSet: true}
}

func (v NullableOrderSubscriptionItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
