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

// GETVoids200ResponseDataInnerRelationships struct for GETVoids200ResponseDataInnerRelationships
type GETVoids200ResponseDataInnerRelationships struct {
	Order                  *GETAdyenPayments200ResponseDataInnerRelationshipsOrder             `json:"order,omitempty"`
	ReferenceAuthorization *GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization `json:"reference_authorization,omitempty"`
	Events                 *GETAuthorizations200ResponseDataInnerRelationshipsEvents           `json:"events,omitempty"`
}

// NewGETVoids200ResponseDataInnerRelationships instantiates a new GETVoids200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETVoids200ResponseDataInnerRelationships() *GETVoids200ResponseDataInnerRelationships {
	this := GETVoids200ResponseDataInnerRelationships{}
	return &this
}

// NewGETVoids200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETVoids200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETVoids200ResponseDataInnerRelationshipsWithDefaults() *GETVoids200ResponseDataInnerRelationships {
	this := GETVoids200ResponseDataInnerRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GETVoids200ResponseDataInnerRelationships) GetOrder() GETAdyenPayments200ResponseDataInnerRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret GETAdyenPayments200ResponseDataInnerRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETVoids200ResponseDataInnerRelationships) GetOrderOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GETVoids200ResponseDataInnerRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given GETAdyenPayments200ResponseDataInnerRelationshipsOrder and assigns it to the Order field.
func (o *GETVoids200ResponseDataInnerRelationships) SetOrder(v GETAdyenPayments200ResponseDataInnerRelationshipsOrder) {
	o.Order = &v
}

// GetReferenceAuthorization returns the ReferenceAuthorization field value if set, zero value otherwise.
func (o *GETVoids200ResponseDataInnerRelationships) GetReferenceAuthorization() GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization {
	if o == nil || o.ReferenceAuthorization == nil {
		var ret GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization
		return ret
	}
	return *o.ReferenceAuthorization
}

// GetReferenceAuthorizationOk returns a tuple with the ReferenceAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETVoids200ResponseDataInnerRelationships) GetReferenceAuthorizationOk() (*GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization, bool) {
	if o == nil || o.ReferenceAuthorization == nil {
		return nil, false
	}
	return o.ReferenceAuthorization, true
}

// HasReferenceAuthorization returns a boolean if a field has been set.
func (o *GETVoids200ResponseDataInnerRelationships) HasReferenceAuthorization() bool {
	if o != nil && o.ReferenceAuthorization != nil {
		return true
	}

	return false
}

// SetReferenceAuthorization gets a reference to the given GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization and assigns it to the ReferenceAuthorization field.
func (o *GETVoids200ResponseDataInnerRelationships) SetReferenceAuthorization(v GETCaptures200ResponseDataInnerRelationshipsReferenceAuthorization) {
	o.ReferenceAuthorization = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETVoids200ResponseDataInnerRelationships) GetEvents() GETAuthorizations200ResponseDataInnerRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret GETAuthorizations200ResponseDataInnerRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETVoids200ResponseDataInnerRelationships) GetEventsOk() (*GETAuthorizations200ResponseDataInnerRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETVoids200ResponseDataInnerRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given GETAuthorizations200ResponseDataInnerRelationshipsEvents and assigns it to the Events field.
func (o *GETVoids200ResponseDataInnerRelationships) SetEvents(v GETAuthorizations200ResponseDataInnerRelationshipsEvents) {
	o.Events = &v
}

func (o GETVoids200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.ReferenceAuthorization != nil {
		toSerialize["reference_authorization"] = o.ReferenceAuthorization
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableGETVoids200ResponseDataInnerRelationships struct {
	value *GETVoids200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETVoids200ResponseDataInnerRelationships) Get() *GETVoids200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETVoids200ResponseDataInnerRelationships) Set(val *GETVoids200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETVoids200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETVoids200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETVoids200ResponseDataInnerRelationships(val *GETVoids200ResponseDataInnerRelationships) *NullableGETVoids200ResponseDataInnerRelationships {
	return &NullableGETVoids200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETVoids200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETVoids200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
