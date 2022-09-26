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

// EventData struct for EventData
type EventData struct {
	// The resource's type
	Type          string                                  `json:"type"`
	Attributes    GETEvents200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *EventDataRelationships                 `json:"relationships,omitempty"`
}

// NewEventData instantiates a new EventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventData(type_ string, attributes GETEvents200ResponseDataInnerAttributes) *EventData {
	this := EventData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewEventDataWithDefaults instantiates a new EventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventDataWithDefaults() *EventData {
	this := EventData{}
	return &this
}

// GetType returns the Type field value
func (o *EventData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *EventData) GetAttributes() GETEvents200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETEvents200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *EventData) GetAttributesOk() (*GETEvents200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *EventData) SetAttributes(v GETEvents200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *EventData) GetRelationships() EventDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret EventDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventData) GetRelationshipsOk() (*EventDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *EventData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given EventDataRelationships and assigns it to the Relationships field.
func (o *EventData) SetRelationships(v EventDataRelationships) {
	o.Relationships = &v
}

func (o EventData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableEventData struct {
	value *EventData
	isSet bool
}

func (v NullableEventData) Get() *EventData {
	return v.value
}

func (v *NullableEventData) Set(val *EventData) {
	v.value = val
	v.isSet = true
}

func (v NullableEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventData(val *EventData) *NullableEventData {
	return &NullableEventData{value: val, isSet: true}
}

func (v NullableEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
