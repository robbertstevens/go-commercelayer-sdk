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

// GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources struct for GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources
type GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                          `json:"links,omitempty"`
	Data  []GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesDataInner `json:"data,omitempty"`
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources() *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources {
	this := GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources{}
	return &this
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesWithDefaults() *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources {
	this := GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) GetData() []GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesDataInner {
	if o == nil || o.Data == nil {
		var ret []GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) GetDataOk() ([]GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesDataInner and assigns it to the Data field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) SetData(v []GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesDataInner) {
	o.Data = v
}

func (o GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources struct {
	value *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources
	isSet bool
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) Get() *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources {
	return v.value
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) Set(val *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources(val *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) *NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources {
	return &NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources{value: val, isSet: true}
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
