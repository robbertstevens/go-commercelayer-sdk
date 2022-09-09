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

// PATCHShipmentsShipmentId200Response struct for PATCHShipmentsShipmentId200Response
type PATCHShipmentsShipmentId200Response struct {
	Data *PATCHShipmentsShipmentId200ResponseData `json:"data,omitempty"`
}

// NewPATCHShipmentsShipmentId200Response instantiates a new PATCHShipmentsShipmentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHShipmentsShipmentId200Response() *PATCHShipmentsShipmentId200Response {
	this := PATCHShipmentsShipmentId200Response{}
	return &this
}

// NewPATCHShipmentsShipmentId200ResponseWithDefaults instantiates a new PATCHShipmentsShipmentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHShipmentsShipmentId200ResponseWithDefaults() *PATCHShipmentsShipmentId200Response {
	this := PATCHShipmentsShipmentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHShipmentsShipmentId200Response) GetData() PATCHShipmentsShipmentId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHShipmentsShipmentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShipmentsShipmentId200Response) GetDataOk() (*PATCHShipmentsShipmentId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHShipmentsShipmentId200ResponseData and assigns it to the Data field.
func (o *PATCHShipmentsShipmentId200Response) SetData(v PATCHShipmentsShipmentId200ResponseData) {
	o.Data = &v
}

func (o PATCHShipmentsShipmentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHShipmentsShipmentId200Response struct {
	value *PATCHShipmentsShipmentId200Response
	isSet bool
}

func (v NullablePATCHShipmentsShipmentId200Response) Get() *PATCHShipmentsShipmentId200Response {
	return v.value
}

func (v *NullablePATCHShipmentsShipmentId200Response) Set(val *PATCHShipmentsShipmentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHShipmentsShipmentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHShipmentsShipmentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHShipmentsShipmentId200Response(val *PATCHShipmentsShipmentId200Response) *NullablePATCHShipmentsShipmentId200Response {
	return &NullablePATCHShipmentsShipmentId200Response{value: val, isSet: true}
}

func (v NullablePATCHShipmentsShipmentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHShipmentsShipmentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


