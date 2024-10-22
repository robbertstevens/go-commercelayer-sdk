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

// checks if the GETManualGatewaysManualGatewayId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETManualGatewaysManualGatewayId200Response{}

// GETManualGatewaysManualGatewayId200Response struct for GETManualGatewaysManualGatewayId200Response
type GETManualGatewaysManualGatewayId200Response struct {
	Data *GETManualGatewaysManualGatewayId200ResponseData `json:"data,omitempty"`
}

// NewGETManualGatewaysManualGatewayId200Response instantiates a new GETManualGatewaysManualGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETManualGatewaysManualGatewayId200Response() *GETManualGatewaysManualGatewayId200Response {
	this := GETManualGatewaysManualGatewayId200Response{}
	return &this
}

// NewGETManualGatewaysManualGatewayId200ResponseWithDefaults instantiates a new GETManualGatewaysManualGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETManualGatewaysManualGatewayId200ResponseWithDefaults() *GETManualGatewaysManualGatewayId200Response {
	this := GETManualGatewaysManualGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETManualGatewaysManualGatewayId200Response) GetData() GETManualGatewaysManualGatewayId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETManualGatewaysManualGatewayId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETManualGatewaysManualGatewayId200Response) GetDataOk() (*GETManualGatewaysManualGatewayId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETManualGatewaysManualGatewayId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETManualGatewaysManualGatewayId200ResponseData and assigns it to the Data field.
func (o *GETManualGatewaysManualGatewayId200Response) SetData(v GETManualGatewaysManualGatewayId200ResponseData) {
	o.Data = &v
}

func (o GETManualGatewaysManualGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETManualGatewaysManualGatewayId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETManualGatewaysManualGatewayId200Response struct {
	value *GETManualGatewaysManualGatewayId200Response
	isSet bool
}

func (v NullableGETManualGatewaysManualGatewayId200Response) Get() *GETManualGatewaysManualGatewayId200Response {
	return v.value
}

func (v *NullableGETManualGatewaysManualGatewayId200Response) Set(val *GETManualGatewaysManualGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETManualGatewaysManualGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETManualGatewaysManualGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETManualGatewaysManualGatewayId200Response(val *GETManualGatewaysManualGatewayId200Response) *NullableGETManualGatewaysManualGatewayId200Response {
	return &NullableGETManualGatewaysManualGatewayId200Response{value: val, isSet: true}
}

func (v NullableGETManualGatewaysManualGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETManualGatewaysManualGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
