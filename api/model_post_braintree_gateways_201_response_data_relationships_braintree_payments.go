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

// checks if the POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments{}

// POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments struct for POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments
type POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                 `json:"links,omitempty"`
	Data  *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData `json:"data,omitempty"`
}

// NewPOSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments instantiates a new POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments() *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments {
	this := POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments{}
	return &this
}

// NewPOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsWithDefaults instantiates a new POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsWithDefaults() *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments {
	this := POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) GetData() POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) GetDataOk() (*POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData and assigns it to the Data field.
func (o *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) SetData(v POSTBraintreeGateways201ResponseDataRelationshipsBraintreePaymentsData) {
	o.Data = &v
}

func (o POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments struct {
	value *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments
	isSet bool
}

func (v NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) Get() *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments {
	return v.value
}

func (v *NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) Set(val *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments(val *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) *NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments {
	return &NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments{value: val, isSet: true}
}

func (v NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
