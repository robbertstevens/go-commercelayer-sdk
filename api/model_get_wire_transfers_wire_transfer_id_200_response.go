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

// GETWireTransfersWireTransferId200Response struct for GETWireTransfersWireTransferId200Response
type GETWireTransfersWireTransferId200Response struct {
	Data *GETWireTransfers200ResponseDataInner `json:"data,omitempty"`
}

// NewGETWireTransfersWireTransferId200Response instantiates a new GETWireTransfersWireTransferId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETWireTransfersWireTransferId200Response() *GETWireTransfersWireTransferId200Response {
	this := GETWireTransfersWireTransferId200Response{}
	return &this
}

// NewGETWireTransfersWireTransferId200ResponseWithDefaults instantiates a new GETWireTransfersWireTransferId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETWireTransfersWireTransferId200ResponseWithDefaults() *GETWireTransfersWireTransferId200Response {
	this := GETWireTransfersWireTransferId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETWireTransfersWireTransferId200Response) GetData() GETWireTransfers200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETWireTransfers200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWireTransfersWireTransferId200Response) GetDataOk() (*GETWireTransfers200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETWireTransfersWireTransferId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETWireTransfers200ResponseDataInner and assigns it to the Data field.
func (o *GETWireTransfersWireTransferId200Response) SetData(v GETWireTransfers200ResponseDataInner) {
	o.Data = &v
}

func (o GETWireTransfersWireTransferId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETWireTransfersWireTransferId200Response struct {
	value *GETWireTransfersWireTransferId200Response
	isSet bool
}

func (v NullableGETWireTransfersWireTransferId200Response) Get() *GETWireTransfersWireTransferId200Response {
	return v.value
}

func (v *NullableGETWireTransfersWireTransferId200Response) Set(val *GETWireTransfersWireTransferId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETWireTransfersWireTransferId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETWireTransfersWireTransferId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETWireTransfersWireTransferId200Response(val *GETWireTransfersWireTransferId200Response) *NullableGETWireTransfersWireTransferId200Response {
	return &NullableGETWireTransfersWireTransferId200Response{value: val, isSet: true}
}

func (v NullableGETWireTransfersWireTransferId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETWireTransfersWireTransferId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


