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

// GETLineItems200ResponseDataInnerRelationshipsStockLineItems struct for GETLineItems200ResponseDataInnerRelationshipsStockLineItems
type GETLineItems200ResponseDataInnerRelationshipsStockLineItems struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks            `json:"links,omitempty"`
	Data  []GETLineItems200ResponseDataInnerRelationshipsStockLineItemsDataInner `json:"data,omitempty"`
}

// NewGETLineItems200ResponseDataInnerRelationshipsStockLineItems instantiates a new GETLineItems200ResponseDataInnerRelationshipsStockLineItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItems200ResponseDataInnerRelationshipsStockLineItems() *GETLineItems200ResponseDataInnerRelationshipsStockLineItems {
	this := GETLineItems200ResponseDataInnerRelationshipsStockLineItems{}
	return &this
}

// NewGETLineItems200ResponseDataInnerRelationshipsStockLineItemsWithDefaults instantiates a new GETLineItems200ResponseDataInnerRelationshipsStockLineItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItems200ResponseDataInnerRelationshipsStockLineItemsWithDefaults() *GETLineItems200ResponseDataInnerRelationshipsStockLineItems {
	this := GETLineItems200ResponseDataInnerRelationshipsStockLineItems{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockLineItems) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockLineItems) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockLineItems) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockLineItems) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockLineItems) GetData() []GETLineItems200ResponseDataInnerRelationshipsStockLineItemsDataInner {
	if o == nil || o.Data == nil {
		var ret []GETLineItems200ResponseDataInnerRelationshipsStockLineItemsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockLineItems) GetDataOk() ([]GETLineItems200ResponseDataInnerRelationshipsStockLineItemsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockLineItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETLineItems200ResponseDataInnerRelationshipsStockLineItemsDataInner and assigns it to the Data field.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockLineItems) SetData(v []GETLineItems200ResponseDataInnerRelationshipsStockLineItemsDataInner) {
	o.Data = v
}

func (o GETLineItems200ResponseDataInnerRelationshipsStockLineItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItems200ResponseDataInnerRelationshipsStockLineItems struct {
	value *GETLineItems200ResponseDataInnerRelationshipsStockLineItems
	isSet bool
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsStockLineItems) Get() *GETLineItems200ResponseDataInnerRelationshipsStockLineItems {
	return v.value
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsStockLineItems) Set(val *GETLineItems200ResponseDataInnerRelationshipsStockLineItems) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsStockLineItems) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsStockLineItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItems200ResponseDataInnerRelationshipsStockLineItems(val *GETLineItems200ResponseDataInnerRelationshipsStockLineItems) *NullableGETLineItems200ResponseDataInnerRelationshipsStockLineItems {
	return &NullableGETLineItems200ResponseDataInnerRelationshipsStockLineItems{value: val, isSet: true}
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsStockLineItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsStockLineItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
