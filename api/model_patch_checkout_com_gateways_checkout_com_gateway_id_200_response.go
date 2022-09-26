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

// PATCHCheckoutComGatewaysCheckoutComGatewayId200Response struct for PATCHCheckoutComGatewaysCheckoutComGatewayId200Response
type PATCHCheckoutComGatewaysCheckoutComGatewayId200Response struct {
	Data *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseData `json:"data,omitempty"`
}

// NewPATCHCheckoutComGatewaysCheckoutComGatewayId200Response instantiates a new PATCHCheckoutComGatewaysCheckoutComGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCheckoutComGatewaysCheckoutComGatewayId200Response() *PATCHCheckoutComGatewaysCheckoutComGatewayId200Response {
	this := PATCHCheckoutComGatewaysCheckoutComGatewayId200Response{}
	return &this
}

// NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseWithDefaults instantiates a new PATCHCheckoutComGatewaysCheckoutComGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseWithDefaults() *PATCHCheckoutComGatewaysCheckoutComGatewayId200Response {
	this := PATCHCheckoutComGatewaysCheckoutComGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200Response) GetData() PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200Response) GetDataOk() (*PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseData and assigns it to the Data field.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200Response) SetData(v PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseData) {
	o.Data = &v
}

func (o PATCHCheckoutComGatewaysCheckoutComGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200Response struct {
	value *PATCHCheckoutComGatewaysCheckoutComGatewayId200Response
	isSet bool
}

func (v NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200Response) Get() *PATCHCheckoutComGatewaysCheckoutComGatewayId200Response {
	return v.value
}

func (v *NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200Response) Set(val *PATCHCheckoutComGatewaysCheckoutComGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCheckoutComGatewaysCheckoutComGatewayId200Response(val *PATCHCheckoutComGatewaysCheckoutComGatewayId200Response) *NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200Response {
	return &NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200Response{value: val, isSet: true}
}

func (v NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
