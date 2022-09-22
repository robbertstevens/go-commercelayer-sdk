/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETCustomerPasswordResets200ResponseDataInnerRelationships struct for GETCustomerPasswordResets200ResponseDataInnerRelationships
type GETCustomerPasswordResets200ResponseDataInnerRelationships struct {
	Customer *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer `json:"customer,omitempty"`
	Events   *GETCustomerAddresses200ResponseDataInnerRelationshipsEvents  `json:"events,omitempty"`
}

// NewGETCustomerPasswordResets200ResponseDataInnerRelationships instantiates a new GETCustomerPasswordResets200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerPasswordResets200ResponseDataInnerRelationships() *GETCustomerPasswordResets200ResponseDataInnerRelationships {
	this := GETCustomerPasswordResets200ResponseDataInnerRelationships{}
	return &this
}

// NewGETCustomerPasswordResets200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETCustomerPasswordResets200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerPasswordResets200ResponseDataInnerRelationshipsWithDefaults() *GETCustomerPasswordResets200ResponseDataInnerRelationships {
	this := GETCustomerPasswordResets200ResponseDataInnerRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *GETCustomerPasswordResets200ResponseDataInnerRelationships) GetCustomer() GETCouponRecipients200ResponseDataInnerRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret GETCouponRecipients200ResponseDataInnerRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerPasswordResets200ResponseDataInnerRelationships) GetCustomerOk() (*GETCouponRecipients200ResponseDataInnerRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *GETCustomerPasswordResets200ResponseDataInnerRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given GETCouponRecipients200ResponseDataInnerRelationshipsCustomer and assigns it to the Customer field.
func (o *GETCustomerPasswordResets200ResponseDataInnerRelationships) SetCustomer(v GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) {
	o.Customer = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETCustomerPasswordResets200ResponseDataInnerRelationships) GetEvents() GETCustomerAddresses200ResponseDataInnerRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret GETCustomerAddresses200ResponseDataInnerRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerPasswordResets200ResponseDataInnerRelationships) GetEventsOk() (*GETCustomerAddresses200ResponseDataInnerRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETCustomerPasswordResets200ResponseDataInnerRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given GETCustomerAddresses200ResponseDataInnerRelationshipsEvents and assigns it to the Events field.
func (o *GETCustomerPasswordResets200ResponseDataInnerRelationships) SetEvents(v GETCustomerAddresses200ResponseDataInnerRelationshipsEvents) {
	o.Events = &v
}

func (o GETCustomerPasswordResets200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomerPasswordResets200ResponseDataInnerRelationships struct {
	value *GETCustomerPasswordResets200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETCustomerPasswordResets200ResponseDataInnerRelationships) Get() *GETCustomerPasswordResets200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETCustomerPasswordResets200ResponseDataInnerRelationships) Set(val *GETCustomerPasswordResets200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerPasswordResets200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerPasswordResets200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerPasswordResets200ResponseDataInnerRelationships(val *GETCustomerPasswordResets200ResponseDataInnerRelationships) *NullableGETCustomerPasswordResets200ResponseDataInnerRelationships {
	return &NullableGETCustomerPasswordResets200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETCustomerPasswordResets200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerPasswordResets200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
