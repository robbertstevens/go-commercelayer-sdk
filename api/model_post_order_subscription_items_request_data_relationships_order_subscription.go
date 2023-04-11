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

// checks if the POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription{}

// POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription struct for POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription
type POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription struct {
	Data POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscriptionData `json:"data"`
}

// NewPOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription instantiates a new POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription(data POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscriptionData) *POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription {
	this := POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription{}
	this.Data = data
	return &this
}

// NewPOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscriptionWithDefaults instantiates a new POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscriptionWithDefaults() *POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription {
	this := POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription{}
	return &this
}

// GetData returns the Data field value
func (o *POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription) GetData() POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscriptionData {
	if o == nil {
		var ret POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscriptionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription) GetDataOk() (*POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscriptionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription) SetData(v POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscriptionData) {
	o.Data = v
}

func (o POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription struct {
	value *POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription
	isSet bool
}

func (v NullablePOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription) Get() *POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription {
	return v.value
}

func (v *NullablePOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription) Set(val *POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription(val *POSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription) *NullablePOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription {
	return &NullablePOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription{value: val, isSet: true}
}

func (v NullablePOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderSubscriptionItemsRequestDataRelationshipsOrderSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
