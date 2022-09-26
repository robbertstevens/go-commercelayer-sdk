/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETPackages200ResponseDataInnerRelationships struct for GETPackages200ResponseDataInnerRelationships
type GETPackages200ResponseDataInnerRelationships struct {
	StockLocation *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation `json:"stock_location,omitempty"`
	Parcels       *GETPackages200ResponseDataInnerRelationshipsParcels                `json:"parcels,omitempty"`
	Attachments   *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments     `json:"attachments,omitempty"`
}

// NewGETPackages200ResponseDataInnerRelationships instantiates a new GETPackages200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPackages200ResponseDataInnerRelationships() *GETPackages200ResponseDataInnerRelationships {
	this := GETPackages200ResponseDataInnerRelationships{}
	return &this
}

// NewGETPackages200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETPackages200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPackages200ResponseDataInnerRelationshipsWithDefaults() *GETPackages200ResponseDataInnerRelationships {
	this := GETPackages200ResponseDataInnerRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *GETPackages200ResponseDataInnerRelationships) GetStockLocation() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPackages200ResponseDataInnerRelationships) GetStockLocationOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *GETPackages200ResponseDataInnerRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *GETPackages200ResponseDataInnerRelationships) SetStockLocation(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetParcels returns the Parcels field value if set, zero value otherwise.
func (o *GETPackages200ResponseDataInnerRelationships) GetParcels() GETPackages200ResponseDataInnerRelationshipsParcels {
	if o == nil || o.Parcels == nil {
		var ret GETPackages200ResponseDataInnerRelationshipsParcels
		return ret
	}
	return *o.Parcels
}

// GetParcelsOk returns a tuple with the Parcels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPackages200ResponseDataInnerRelationships) GetParcelsOk() (*GETPackages200ResponseDataInnerRelationshipsParcels, bool) {
	if o == nil || o.Parcels == nil {
		return nil, false
	}
	return o.Parcels, true
}

// HasParcels returns a boolean if a field has been set.
func (o *GETPackages200ResponseDataInnerRelationships) HasParcels() bool {
	if o != nil && o.Parcels != nil {
		return true
	}

	return false
}

// SetParcels gets a reference to the given GETPackages200ResponseDataInnerRelationshipsParcels and assigns it to the Parcels field.
func (o *GETPackages200ResponseDataInnerRelationships) SetParcels(v GETPackages200ResponseDataInnerRelationshipsParcels) {
	o.Parcels = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETPackages200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPackages200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETPackages200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETPackages200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETPackages200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.Parcels != nil {
		toSerialize["parcels"] = o.Parcels
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETPackages200ResponseDataInnerRelationships struct {
	value *GETPackages200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETPackages200ResponseDataInnerRelationships) Get() *GETPackages200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETPackages200ResponseDataInnerRelationships) Set(val *GETPackages200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPackages200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPackages200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPackages200ResponseDataInnerRelationships(val *GETPackages200ResponseDataInnerRelationships) *NullableGETPackages200ResponseDataInnerRelationships {
	return &NullableGETPackages200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETPackages200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPackages200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
