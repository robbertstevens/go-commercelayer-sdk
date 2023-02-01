/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETCustomerGroups200ResponseDataInnerRelationships struct for GETCustomerGroups200ResponseDataInnerRelationships
type GETCustomerGroups200ResponseDataInnerRelationships struct {
	Customers   *GETCustomerGroups200ResponseDataInnerRelationshipsCustomers    `json:"customers,omitempty"`
	Markets     *GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets     `json:"markets,omitempty"`
	Attachments *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewGETCustomerGroups200ResponseDataInnerRelationships instantiates a new GETCustomerGroups200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerGroups200ResponseDataInnerRelationships() *GETCustomerGroups200ResponseDataInnerRelationships {
	this := GETCustomerGroups200ResponseDataInnerRelationships{}
	return &this
}

// NewGETCustomerGroups200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETCustomerGroups200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerGroups200ResponseDataInnerRelationshipsWithDefaults() *GETCustomerGroups200ResponseDataInnerRelationships {
	this := GETCustomerGroups200ResponseDataInnerRelationships{}
	return &this
}

// GetCustomers returns the Customers field value if set, zero value otherwise.
func (o *GETCustomerGroups200ResponseDataInnerRelationships) GetCustomers() GETCustomerGroups200ResponseDataInnerRelationshipsCustomers {
	if o == nil || o.Customers == nil {
		var ret GETCustomerGroups200ResponseDataInnerRelationshipsCustomers
		return ret
	}
	return *o.Customers
}

// GetCustomersOk returns a tuple with the Customers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerGroups200ResponseDataInnerRelationships) GetCustomersOk() (*GETCustomerGroups200ResponseDataInnerRelationshipsCustomers, bool) {
	if o == nil || o.Customers == nil {
		return nil, false
	}
	return o.Customers, true
}

// HasCustomers returns a boolean if a field has been set.
func (o *GETCustomerGroups200ResponseDataInnerRelationships) HasCustomers() bool {
	if o != nil && o.Customers != nil {
		return true
	}

	return false
}

// SetCustomers gets a reference to the given GETCustomerGroups200ResponseDataInnerRelationshipsCustomers and assigns it to the Customers field.
func (o *GETCustomerGroups200ResponseDataInnerRelationships) SetCustomers(v GETCustomerGroups200ResponseDataInnerRelationshipsCustomers) {
	o.Customers = &v
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *GETCustomerGroups200ResponseDataInnerRelationships) GetMarkets() GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets {
	if o == nil || o.Markets == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets
		return ret
	}
	return *o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerGroups200ResponseDataInnerRelationships) GetMarketsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets, bool) {
	if o == nil || o.Markets == nil {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *GETCustomerGroups200ResponseDataInnerRelationships) HasMarkets() bool {
	if o != nil && o.Markets != nil {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets and assigns it to the Markets field.
func (o *GETCustomerGroups200ResponseDataInnerRelationships) SetMarkets(v GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets) {
	o.Markets = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETCustomerGroups200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerGroups200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETCustomerGroups200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETCustomerGroups200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETCustomerGroups200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customers != nil {
		toSerialize["customers"] = o.Customers
	}
	if o.Markets != nil {
		toSerialize["markets"] = o.Markets
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomerGroups200ResponseDataInnerRelationships struct {
	value *GETCustomerGroups200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETCustomerGroups200ResponseDataInnerRelationships) Get() *GETCustomerGroups200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETCustomerGroups200ResponseDataInnerRelationships) Set(val *GETCustomerGroups200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerGroups200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerGroups200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerGroups200ResponseDataInnerRelationships(val *GETCustomerGroups200ResponseDataInnerRelationships) *NullableGETCustomerGroups200ResponseDataInnerRelationships {
	return &NullableGETCustomerGroups200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETCustomerGroups200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerGroups200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
