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

// GiftCardDataRelationships struct for GiftCardDataRelationships
type GiftCardDataRelationships struct {
	Market            *AvalaraAccountDataRelationshipsMarkets     `json:"market,omitempty"`
	GiftCardRecipient *GiftCardDataRelationshipsGiftCardRecipient `json:"gift_card_recipient,omitempty"`
	Attachments       *AvalaraAccountDataRelationshipsAttachments `json:"attachments,omitempty"`
	Events            *AuthorizationDataRelationshipsEvents       `json:"events,omitempty"`
}

// NewGiftCardDataRelationships instantiates a new GiftCardDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardDataRelationships() *GiftCardDataRelationships {
	this := GiftCardDataRelationships{}
	return &this
}

// NewGiftCardDataRelationshipsWithDefaults instantiates a new GiftCardDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardDataRelationshipsWithDefaults() *GiftCardDataRelationships {
	this := GiftCardDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *GiftCardDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || o.Market == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *GiftCardDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Market field.
func (o *GiftCardDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = &v
}

// GetGiftCardRecipient returns the GiftCardRecipient field value if set, zero value otherwise.
func (o *GiftCardDataRelationships) GetGiftCardRecipient() GiftCardDataRelationshipsGiftCardRecipient {
	if o == nil || o.GiftCardRecipient == nil {
		var ret GiftCardDataRelationshipsGiftCardRecipient
		return ret
	}
	return *o.GiftCardRecipient
}

// GetGiftCardRecipientOk returns a tuple with the GiftCardRecipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataRelationships) GetGiftCardRecipientOk() (*GiftCardDataRelationshipsGiftCardRecipient, bool) {
	if o == nil || o.GiftCardRecipient == nil {
		return nil, false
	}
	return o.GiftCardRecipient, true
}

// HasGiftCardRecipient returns a boolean if a field has been set.
func (o *GiftCardDataRelationships) HasGiftCardRecipient() bool {
	if o != nil && o.GiftCardRecipient != nil {
		return true
	}

	return false
}

// SetGiftCardRecipient gets a reference to the given GiftCardDataRelationshipsGiftCardRecipient and assigns it to the GiftCardRecipient field.
func (o *GiftCardDataRelationships) SetGiftCardRecipient(v GiftCardDataRelationshipsGiftCardRecipient) {
	o.GiftCardRecipient = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GiftCardDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GiftCardDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *GiftCardDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GiftCardDataRelationships) GetEvents() AuthorizationDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret AuthorizationDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardDataRelationships) GetEventsOk() (*AuthorizationDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GiftCardDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AuthorizationDataRelationshipsEvents and assigns it to the Events field.
func (o *GiftCardDataRelationships) SetEvents(v AuthorizationDataRelationshipsEvents) {
	o.Events = &v
}

func (o GiftCardDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.GiftCardRecipient != nil {
		toSerialize["gift_card_recipient"] = o.GiftCardRecipient
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableGiftCardDataRelationships struct {
	value *GiftCardDataRelationships
	isSet bool
}

func (v NullableGiftCardDataRelationships) Get() *GiftCardDataRelationships {
	return v.value
}

func (v *NullableGiftCardDataRelationships) Set(val *GiftCardDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardDataRelationships(val *GiftCardDataRelationships) *NullableGiftCardDataRelationships {
	return &NullableGiftCardDataRelationships{value: val, isSet: true}
}

func (v NullableGiftCardDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
