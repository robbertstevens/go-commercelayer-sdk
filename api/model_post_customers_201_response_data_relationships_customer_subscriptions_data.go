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

// checks if the POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData{}

// POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData struct for POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData
type POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData instantiates a new POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData() *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData {
	this := POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData{}
	return &this
}

// NewPOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsDataWithDefaults instantiates a new POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsDataWithDefaults() *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData {
	this := POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData struct {
	value *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData
	isSet bool
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) Get() *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData {
	return v.value
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) Set(val *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData(val *POSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) *NullablePOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData {
	return &NullablePOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData{value: val, isSet: true}
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsCustomerSubscriptionsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
