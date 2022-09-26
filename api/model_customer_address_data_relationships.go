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

// CustomerAddressDataRelationships struct for CustomerAddressDataRelationships
type CustomerAddressDataRelationships struct {
	Customer *CouponRecipientDataRelationshipsCustomer `json:"customer,omitempty"`
	Address  *BingGeocoderDataRelationshipsAddresses   `json:"address,omitempty"`
	Events   *CustomerAddressDataRelationshipsEvents   `json:"events,omitempty"`
}

// NewCustomerAddressDataRelationships instantiates a new CustomerAddressDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddressDataRelationships() *CustomerAddressDataRelationships {
	this := CustomerAddressDataRelationships{}
	return &this
}

// NewCustomerAddressDataRelationshipsWithDefaults instantiates a new CustomerAddressDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressDataRelationshipsWithDefaults() *CustomerAddressDataRelationships {
	this := CustomerAddressDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CustomerAddressDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *CustomerAddressDataRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientDataRelationshipsCustomer and assigns it to the Customer field.
func (o *CustomerAddressDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CustomerAddressDataRelationships) GetAddress() BingGeocoderDataRelationshipsAddresses {
	if o == nil || o.Address == nil {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressDataRelationships) GetAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CustomerAddressDataRelationships) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given BingGeocoderDataRelationshipsAddresses and assigns it to the Address field.
func (o *CustomerAddressDataRelationships) SetAddress(v BingGeocoderDataRelationshipsAddresses) {
	o.Address = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *CustomerAddressDataRelationships) GetEvents() CustomerAddressDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressDataRelationships) GetEventsOk() (*CustomerAddressDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *CustomerAddressDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressDataRelationshipsEvents and assigns it to the Events field.
func (o *CustomerAddressDataRelationships) SetEvents(v CustomerAddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o CustomerAddressDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerAddressDataRelationships struct {
	value *CustomerAddressDataRelationships
	isSet bool
}

func (v NullableCustomerAddressDataRelationships) Get() *CustomerAddressDataRelationships {
	return v.value
}

func (v *NullableCustomerAddressDataRelationships) Set(val *CustomerAddressDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressDataRelationships(val *CustomerAddressDataRelationships) *NullableCustomerAddressDataRelationships {
	return &NullableCustomerAddressDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerAddressDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
