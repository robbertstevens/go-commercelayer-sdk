/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the WebhookCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookCreate{}

// WebhookCreate struct for WebhookCreate
type WebhookCreate struct {
	Data POSTWebhooksRequestData `json:"data"`
}

// NewWebhookCreate instantiates a new WebhookCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookCreate(data POSTWebhooksRequestData) *WebhookCreate {
	this := WebhookCreate{}
	this.Data = data
	return &this
}

// NewWebhookCreateWithDefaults instantiates a new WebhookCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookCreateWithDefaults() *WebhookCreate {
	this := WebhookCreate{}
	return &this
}

// GetData returns the Data field value
func (o *WebhookCreate) GetData() POSTWebhooksRequestData {
	if o == nil {
		var ret POSTWebhooksRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WebhookCreate) GetDataOk() (*POSTWebhooksRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *WebhookCreate) SetData(v POSTWebhooksRequestData) {
	o.Data = v
}

func (o WebhookCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableWebhookCreate struct {
	value *WebhookCreate
	isSet bool
}

func (v NullableWebhookCreate) Get() *WebhookCreate {
	return v.value
}

func (v *NullableWebhookCreate) Set(val *WebhookCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookCreate(val *WebhookCreate) *NullableWebhookCreate {
	return &NullableWebhookCreate{value: val, isSet: true}
}

func (v NullableWebhookCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
