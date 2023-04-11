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

// checks if the POSTOrdersRequestDataRelationshipsPaymentMethodData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrdersRequestDataRelationshipsPaymentMethodData{}

// POSTOrdersRequestDataRelationshipsPaymentMethodData struct for POSTOrdersRequestDataRelationshipsPaymentMethodData
type POSTOrdersRequestDataRelationshipsPaymentMethodData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTOrdersRequestDataRelationshipsPaymentMethodData instantiates a new POSTOrdersRequestDataRelationshipsPaymentMethodData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrdersRequestDataRelationshipsPaymentMethodData() *POSTOrdersRequestDataRelationshipsPaymentMethodData {
	this := POSTOrdersRequestDataRelationshipsPaymentMethodData{}
	return &this
}

// NewPOSTOrdersRequestDataRelationshipsPaymentMethodDataWithDefaults instantiates a new POSTOrdersRequestDataRelationshipsPaymentMethodData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrdersRequestDataRelationshipsPaymentMethodDataWithDefaults() *POSTOrdersRequestDataRelationshipsPaymentMethodData {
	this := POSTOrdersRequestDataRelationshipsPaymentMethodData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrdersRequestDataRelationshipsPaymentMethodData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrdersRequestDataRelationshipsPaymentMethodData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTOrdersRequestDataRelationshipsPaymentMethodData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTOrdersRequestDataRelationshipsPaymentMethodData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrdersRequestDataRelationshipsPaymentMethodData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrdersRequestDataRelationshipsPaymentMethodData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTOrdersRequestDataRelationshipsPaymentMethodData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTOrdersRequestDataRelationshipsPaymentMethodData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTOrdersRequestDataRelationshipsPaymentMethodData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrdersRequestDataRelationshipsPaymentMethodData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTOrdersRequestDataRelationshipsPaymentMethodData struct {
	value *POSTOrdersRequestDataRelationshipsPaymentMethodData
	isSet bool
}

func (v NullablePOSTOrdersRequestDataRelationshipsPaymentMethodData) Get() *POSTOrdersRequestDataRelationshipsPaymentMethodData {
	return v.value
}

func (v *NullablePOSTOrdersRequestDataRelationshipsPaymentMethodData) Set(val *POSTOrdersRequestDataRelationshipsPaymentMethodData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrdersRequestDataRelationshipsPaymentMethodData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrdersRequestDataRelationshipsPaymentMethodData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrdersRequestDataRelationshipsPaymentMethodData(val *POSTOrdersRequestDataRelationshipsPaymentMethodData) *NullablePOSTOrdersRequestDataRelationshipsPaymentMethodData {
	return &NullablePOSTOrdersRequestDataRelationshipsPaymentMethodData{value: val, isSet: true}
}

func (v NullablePOSTOrdersRequestDataRelationshipsPaymentMethodData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrdersRequestDataRelationshipsPaymentMethodData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
