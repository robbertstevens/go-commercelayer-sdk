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

// PATCHGiftCardRecipientsGiftCardRecipientId200Response struct for PATCHGiftCardRecipientsGiftCardRecipientId200Response
type PATCHGiftCardRecipientsGiftCardRecipientId200Response struct {
	Data *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData `json:"data,omitempty"`
}

// NewPATCHGiftCardRecipientsGiftCardRecipientId200Response instantiates a new PATCHGiftCardRecipientsGiftCardRecipientId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHGiftCardRecipientsGiftCardRecipientId200Response() *PATCHGiftCardRecipientsGiftCardRecipientId200Response {
	this := PATCHGiftCardRecipientsGiftCardRecipientId200Response{}
	return &this
}

// NewPATCHGiftCardRecipientsGiftCardRecipientId200ResponseWithDefaults instantiates a new PATCHGiftCardRecipientsGiftCardRecipientId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHGiftCardRecipientsGiftCardRecipientId200ResponseWithDefaults() *PATCHGiftCardRecipientsGiftCardRecipientId200Response {
	this := PATCHGiftCardRecipientsGiftCardRecipientId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200Response) GetData() PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200Response) GetDataOk() (*PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData and assigns it to the Data field.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200Response) SetData(v PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) {
	o.Data = &v
}

func (o PATCHGiftCardRecipientsGiftCardRecipientId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHGiftCardRecipientsGiftCardRecipientId200Response struct {
	value *PATCHGiftCardRecipientsGiftCardRecipientId200Response
	isSet bool
}

func (v NullablePATCHGiftCardRecipientsGiftCardRecipientId200Response) Get() *PATCHGiftCardRecipientsGiftCardRecipientId200Response {
	return v.value
}

func (v *NullablePATCHGiftCardRecipientsGiftCardRecipientId200Response) Set(val *PATCHGiftCardRecipientsGiftCardRecipientId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHGiftCardRecipientsGiftCardRecipientId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHGiftCardRecipientsGiftCardRecipientId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHGiftCardRecipientsGiftCardRecipientId200Response(val *PATCHGiftCardRecipientsGiftCardRecipientId200Response) *NullablePATCHGiftCardRecipientsGiftCardRecipientId200Response {
	return &NullablePATCHGiftCardRecipientsGiftCardRecipientId200Response{value: val, isSet: true}
}

func (v NullablePATCHGiftCardRecipientsGiftCardRecipientId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHGiftCardRecipientsGiftCardRecipientId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
