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

// EventCallbackDataRelationshipsWebhook struct for EventCallbackDataRelationshipsWebhook
type EventCallbackDataRelationshipsWebhook struct {
	Data *EventCallbackDataRelationshipsWebhookData `json:"data,omitempty"`
}

// NewEventCallbackDataRelationshipsWebhook instantiates a new EventCallbackDataRelationshipsWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventCallbackDataRelationshipsWebhook() *EventCallbackDataRelationshipsWebhook {
	this := EventCallbackDataRelationshipsWebhook{}
	return &this
}

// NewEventCallbackDataRelationshipsWebhookWithDefaults instantiates a new EventCallbackDataRelationshipsWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventCallbackDataRelationshipsWebhookWithDefaults() *EventCallbackDataRelationshipsWebhook {
	this := EventCallbackDataRelationshipsWebhook{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EventCallbackDataRelationshipsWebhook) GetData() EventCallbackDataRelationshipsWebhookData {
	if o == nil || o.Data == nil {
		var ret EventCallbackDataRelationshipsWebhookData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCallbackDataRelationshipsWebhook) GetDataOk() (*EventCallbackDataRelationshipsWebhookData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EventCallbackDataRelationshipsWebhook) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given EventCallbackDataRelationshipsWebhookData and assigns it to the Data field.
func (o *EventCallbackDataRelationshipsWebhook) SetData(v EventCallbackDataRelationshipsWebhookData) {
	o.Data = &v
}

func (o EventCallbackDataRelationshipsWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableEventCallbackDataRelationshipsWebhook struct {
	value *EventCallbackDataRelationshipsWebhook
	isSet bool
}

func (v NullableEventCallbackDataRelationshipsWebhook) Get() *EventCallbackDataRelationshipsWebhook {
	return v.value
}

func (v *NullableEventCallbackDataRelationshipsWebhook) Set(val *EventCallbackDataRelationshipsWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCallbackDataRelationshipsWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCallbackDataRelationshipsWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCallbackDataRelationshipsWebhook(val *EventCallbackDataRelationshipsWebhook) *NullableEventCallbackDataRelationshipsWebhook {
	return &NullableEventCallbackDataRelationshipsWebhook{value: val, isSet: true}
}

func (v NullableEventCallbackDataRelationshipsWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCallbackDataRelationshipsWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
