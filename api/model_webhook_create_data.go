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

// WebhookCreateData struct for WebhookCreateData
type WebhookCreateData struct {
	// The resource's type
	Type          string                                `json:"type"`
	Attributes    POSTWebhooks201ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                `json:"relationships,omitempty"`
}

// NewWebhookCreateData instantiates a new WebhookCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookCreateData(type_ string, attributes POSTWebhooks201ResponseDataAttributes) *WebhookCreateData {
	this := WebhookCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewWebhookCreateDataWithDefaults instantiates a new WebhookCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookCreateDataWithDefaults() *WebhookCreateData {
	this := WebhookCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *WebhookCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebhookCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebhookCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *WebhookCreateData) GetAttributes() POSTWebhooks201ResponseDataAttributes {
	if o == nil {
		var ret POSTWebhooks201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WebhookCreateData) GetAttributesOk() (*POSTWebhooks201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *WebhookCreateData) SetAttributes(v POSTWebhooks201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WebhookCreateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookCreateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WebhookCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *WebhookCreateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o WebhookCreateData) MarshalJSON() ([]byte, error) {
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

type NullableWebhookCreateData struct {
	value *WebhookCreateData
	isSet bool
}

func (v NullableWebhookCreateData) Get() *WebhookCreateData {
	return v.value
}

func (v *NullableWebhookCreateData) Set(val *WebhookCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookCreateData(val *WebhookCreateData) *NullableWebhookCreateData {
	return &NullableWebhookCreateData{value: val, isSet: true}
}

func (v NullableWebhookCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
