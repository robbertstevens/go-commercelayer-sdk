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

// checks if the POSTLineItemsRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTLineItemsRequestDataRelationships{}

// POSTLineItemsRequestDataRelationships struct for POSTLineItemsRequestDataRelationships
type POSTLineItemsRequestDataRelationships struct {
	Order POSTAdyenPaymentsRequestDataRelationshipsOrder `json:"order"`
	Item  *POSTLineItemsRequestDataRelationshipsItem     `json:"item,omitempty"`
}

// NewPOSTLineItemsRequestDataRelationships instantiates a new POSTLineItemsRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItemsRequestDataRelationships(order POSTAdyenPaymentsRequestDataRelationshipsOrder) *POSTLineItemsRequestDataRelationships {
	this := POSTLineItemsRequestDataRelationships{}
	this.Order = order
	return &this
}

// NewPOSTLineItemsRequestDataRelationshipsWithDefaults instantiates a new POSTLineItemsRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItemsRequestDataRelationshipsWithDefaults() *POSTLineItemsRequestDataRelationships {
	this := POSTLineItemsRequestDataRelationships{}
	return &this
}

// GetOrder returns the Order field value
func (o *POSTLineItemsRequestDataRelationships) GetOrder() POSTAdyenPaymentsRequestDataRelationshipsOrder {
	if o == nil {
		var ret POSTAdyenPaymentsRequestDataRelationshipsOrder
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *POSTLineItemsRequestDataRelationships) GetOrderOk() (*POSTAdyenPaymentsRequestDataRelationshipsOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *POSTLineItemsRequestDataRelationships) SetOrder(v POSTAdyenPaymentsRequestDataRelationshipsOrder) {
	o.Order = v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *POSTLineItemsRequestDataRelationships) GetItem() POSTLineItemsRequestDataRelationshipsItem {
	if o == nil || IsNil(o.Item) {
		var ret POSTLineItemsRequestDataRelationshipsItem
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItemsRequestDataRelationships) GetItemOk() (*POSTLineItemsRequestDataRelationshipsItem, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *POSTLineItemsRequestDataRelationships) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given POSTLineItemsRequestDataRelationshipsItem and assigns it to the Item field.
func (o *POSTLineItemsRequestDataRelationships) SetItem(v POSTLineItemsRequestDataRelationshipsItem) {
	o.Item = &v
}

func (o POSTLineItemsRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTLineItemsRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order"] = o.Order
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	return toSerialize, nil
}

type NullablePOSTLineItemsRequestDataRelationships struct {
	value *POSTLineItemsRequestDataRelationships
	isSet bool
}

func (v NullablePOSTLineItemsRequestDataRelationships) Get() *POSTLineItemsRequestDataRelationships {
	return v.value
}

func (v *NullablePOSTLineItemsRequestDataRelationships) Set(val *POSTLineItemsRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItemsRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItemsRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItemsRequestDataRelationships(val *POSTLineItemsRequestDataRelationships) *NullablePOSTLineItemsRequestDataRelationships {
	return &NullablePOSTLineItemsRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTLineItemsRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItemsRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
