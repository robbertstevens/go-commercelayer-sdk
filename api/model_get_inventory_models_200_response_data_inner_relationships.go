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

// GETInventoryModels200ResponseDataInnerRelationships struct for GETInventoryModels200ResponseDataInnerRelationships
type GETInventoryModels200ResponseDataInnerRelationships struct {
	InventoryStockLocations  *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations  `json:"inventory_stock_locations,omitempty"`
	InventoryReturnLocations *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations `json:"inventory_return_locations,omitempty"`
	Attachments              *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments              `json:"attachments,omitempty"`
}

// NewGETInventoryModels200ResponseDataInnerRelationships instantiates a new GETInventoryModels200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryModels200ResponseDataInnerRelationships() *GETInventoryModels200ResponseDataInnerRelationships {
	this := GETInventoryModels200ResponseDataInnerRelationships{}
	return &this
}

// NewGETInventoryModels200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETInventoryModels200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryModels200ResponseDataInnerRelationshipsWithDefaults() *GETInventoryModels200ResponseDataInnerRelationships {
	this := GETInventoryModels200ResponseDataInnerRelationships{}
	return &this
}

// GetInventoryStockLocations returns the InventoryStockLocations field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerRelationships) GetInventoryStockLocations() GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations {
	if o == nil || o.InventoryStockLocations == nil {
		var ret GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations
		return ret
	}
	return *o.InventoryStockLocations
}

// GetInventoryStockLocationsOk returns a tuple with the InventoryStockLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationships) GetInventoryStockLocationsOk() (*GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations, bool) {
	if o == nil || o.InventoryStockLocations == nil {
		return nil, false
	}
	return o.InventoryStockLocations, true
}

// HasInventoryStockLocations returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationships) HasInventoryStockLocations() bool {
	if o != nil && o.InventoryStockLocations != nil {
		return true
	}

	return false
}

// SetInventoryStockLocations gets a reference to the given GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations and assigns it to the InventoryStockLocations field.
func (o *GETInventoryModels200ResponseDataInnerRelationships) SetInventoryStockLocations(v GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) {
	o.InventoryStockLocations = &v
}

// GetInventoryReturnLocations returns the InventoryReturnLocations field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerRelationships) GetInventoryReturnLocations() GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations {
	if o == nil || o.InventoryReturnLocations == nil {
		var ret GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations
		return ret
	}
	return *o.InventoryReturnLocations
}

// GetInventoryReturnLocationsOk returns a tuple with the InventoryReturnLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationships) GetInventoryReturnLocationsOk() (*GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations, bool) {
	if o == nil || o.InventoryReturnLocations == nil {
		return nil, false
	}
	return o.InventoryReturnLocations, true
}

// HasInventoryReturnLocations returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationships) HasInventoryReturnLocations() bool {
	if o != nil && o.InventoryReturnLocations != nil {
		return true
	}

	return false
}

// SetInventoryReturnLocations gets a reference to the given GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations and assigns it to the InventoryReturnLocations field.
func (o *GETInventoryModels200ResponseDataInnerRelationships) SetInventoryReturnLocations(v GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) {
	o.InventoryReturnLocations = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETInventoryModels200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETInventoryModels200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InventoryStockLocations != nil {
		toSerialize["inventory_stock_locations"] = o.InventoryStockLocations
	}
	if o.InventoryReturnLocations != nil {
		toSerialize["inventory_return_locations"] = o.InventoryReturnLocations
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETInventoryModels200ResponseDataInnerRelationships struct {
	value *GETInventoryModels200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETInventoryModels200ResponseDataInnerRelationships) Get() *GETInventoryModels200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETInventoryModels200ResponseDataInnerRelationships) Set(val *GETInventoryModels200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryModels200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryModels200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryModels200ResponseDataInnerRelationships(val *GETInventoryModels200ResponseDataInnerRelationships) *NullableGETInventoryModels200ResponseDataInnerRelationships {
	return &NullableGETInventoryModels200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETInventoryModels200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryModels200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
