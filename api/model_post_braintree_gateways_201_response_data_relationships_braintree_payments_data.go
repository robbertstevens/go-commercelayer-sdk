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

// checks if the POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData{}

// POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData struct for POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData
type POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData instantiates a new POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData() *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData {
	this := POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData{}
	return &this
}

// NewPOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsDataWithDefaults instantiates a new POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsDataWithDefaults() *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData {
	this := POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData struct {
	value *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData
	isSet bool
}

func (v NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) Get() *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData {
	return v.value
}

func (v *NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) Set(val *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData(val *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) *NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData {
	return &NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData{value: val, isSet: true}
}

func (v NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
