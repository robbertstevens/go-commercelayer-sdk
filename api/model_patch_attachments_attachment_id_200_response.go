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

// PATCHAttachmentsAttachmentId200Response struct for PATCHAttachmentsAttachmentId200Response
type PATCHAttachmentsAttachmentId200Response struct {
	Data *PATCHAttachmentsAttachmentId200ResponseData `json:"data,omitempty"`
}

// NewPATCHAttachmentsAttachmentId200Response instantiates a new PATCHAttachmentsAttachmentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAttachmentsAttachmentId200Response() *PATCHAttachmentsAttachmentId200Response {
	this := PATCHAttachmentsAttachmentId200Response{}
	return &this
}

// NewPATCHAttachmentsAttachmentId200ResponseWithDefaults instantiates a new PATCHAttachmentsAttachmentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAttachmentsAttachmentId200ResponseWithDefaults() *PATCHAttachmentsAttachmentId200Response {
	this := PATCHAttachmentsAttachmentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHAttachmentsAttachmentId200Response) GetData() PATCHAttachmentsAttachmentId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHAttachmentsAttachmentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAttachmentsAttachmentId200Response) GetDataOk() (*PATCHAttachmentsAttachmentId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHAttachmentsAttachmentId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHAttachmentsAttachmentId200ResponseData and assigns it to the Data field.
func (o *PATCHAttachmentsAttachmentId200Response) SetData(v PATCHAttachmentsAttachmentId200ResponseData) {
	o.Data = &v
}

func (o PATCHAttachmentsAttachmentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHAttachmentsAttachmentId200Response struct {
	value *PATCHAttachmentsAttachmentId200Response
	isSet bool
}

func (v NullablePATCHAttachmentsAttachmentId200Response) Get() *PATCHAttachmentsAttachmentId200Response {
	return v.value
}

func (v *NullablePATCHAttachmentsAttachmentId200Response) Set(val *PATCHAttachmentsAttachmentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAttachmentsAttachmentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAttachmentsAttachmentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAttachmentsAttachmentId200Response(val *PATCHAttachmentsAttachmentId200Response) *NullablePATCHAttachmentsAttachmentId200Response {
	return &NullablePATCHAttachmentsAttachmentId200Response{value: val, isSet: true}
}

func (v NullablePATCHAttachmentsAttachmentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAttachmentsAttachmentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
