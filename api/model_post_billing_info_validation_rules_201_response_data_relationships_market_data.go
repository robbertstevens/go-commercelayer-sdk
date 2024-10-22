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

// checks if the POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData{}

// POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData struct for POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData
type POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData instantiates a new POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData() *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData {
	this := POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData{}
	return &this
}

// NewPOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketDataWithDefaults instantiates a new POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketDataWithDefaults() *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData {
	this := POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData struct {
	value *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData
	isSet bool
}

func (v NullablePOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) Get() *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData {
	return v.value
}

func (v *NullablePOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) Set(val *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData(val *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) *NullablePOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData {
	return &NullablePOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData{value: val, isSet: true}
}

func (v NullablePOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBillingInfoValidationRules201ResponseDataRelationshipsMarketData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
