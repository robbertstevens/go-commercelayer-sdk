/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PATCHWebhooksWebhookId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHWebhooksWebhookId200Response{}

// PATCHWebhooksWebhookId200Response struct for PATCHWebhooksWebhookId200Response
type PATCHWebhooksWebhookId200Response struct {
	Data *PATCHWebhooksWebhookId200ResponseData `json:"data,omitempty"`
}

// NewPATCHWebhooksWebhookId200Response instantiates a new PATCHWebhooksWebhookId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHWebhooksWebhookId200Response() *PATCHWebhooksWebhookId200Response {
	this := PATCHWebhooksWebhookId200Response{}
	return &this
}

// NewPATCHWebhooksWebhookId200ResponseWithDefaults instantiates a new PATCHWebhooksWebhookId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHWebhooksWebhookId200ResponseWithDefaults() *PATCHWebhooksWebhookId200Response {
	this := PATCHWebhooksWebhookId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHWebhooksWebhookId200Response) GetData() PATCHWebhooksWebhookId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHWebhooksWebhookId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHWebhooksWebhookId200Response) GetDataOk() (*PATCHWebhooksWebhookId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHWebhooksWebhookId200ResponseData and assigns it to the Data field.
func (o *PATCHWebhooksWebhookId200Response) SetData(v PATCHWebhooksWebhookId200ResponseData) {
	o.Data = &v
}

func (o PATCHWebhooksWebhookId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHWebhooksWebhookId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHWebhooksWebhookId200Response struct {
	value *PATCHWebhooksWebhookId200Response
	isSet bool
}

func (v NullablePATCHWebhooksWebhookId200Response) Get() *PATCHWebhooksWebhookId200Response {
	return v.value
}

func (v *NullablePATCHWebhooksWebhookId200Response) Set(val *PATCHWebhooksWebhookId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHWebhooksWebhookId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHWebhooksWebhookId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHWebhooksWebhookId200Response(val *PATCHWebhooksWebhookId200Response) *NullablePATCHWebhooksWebhookId200Response {
	return &NullablePATCHWebhooksWebhookId200Response{value: val, isSet: true}
}

func (v NullablePATCHWebhooksWebhookId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHWebhooksWebhookId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
