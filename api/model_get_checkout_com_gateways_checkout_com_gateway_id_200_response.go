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

// checks if the GETCheckoutComGatewaysCheckoutComGatewayId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCheckoutComGatewaysCheckoutComGatewayId200Response{}

// GETCheckoutComGatewaysCheckoutComGatewayId200Response struct for GETCheckoutComGatewaysCheckoutComGatewayId200Response
type GETCheckoutComGatewaysCheckoutComGatewayId200Response struct {
	Data *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseData `json:"data,omitempty"`
}

// NewGETCheckoutComGatewaysCheckoutComGatewayId200Response instantiates a new GETCheckoutComGatewaysCheckoutComGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCheckoutComGatewaysCheckoutComGatewayId200Response() *GETCheckoutComGatewaysCheckoutComGatewayId200Response {
	this := GETCheckoutComGatewaysCheckoutComGatewayId200Response{}
	return &this
}

// NewGETCheckoutComGatewaysCheckoutComGatewayId200ResponseWithDefaults instantiates a new GETCheckoutComGatewaysCheckoutComGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCheckoutComGatewaysCheckoutComGatewayId200ResponseWithDefaults() *GETCheckoutComGatewaysCheckoutComGatewayId200Response {
	this := GETCheckoutComGatewaysCheckoutComGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200Response) GetData() GETCheckoutComGatewaysCheckoutComGatewayId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETCheckoutComGatewaysCheckoutComGatewayId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200Response) GetDataOk() (*GETCheckoutComGatewaysCheckoutComGatewayId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCheckoutComGatewaysCheckoutComGatewayId200ResponseData and assigns it to the Data field.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200Response) SetData(v GETCheckoutComGatewaysCheckoutComGatewayId200ResponseData) {
	o.Data = &v
}

func (o GETCheckoutComGatewaysCheckoutComGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCheckoutComGatewaysCheckoutComGatewayId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response struct {
	value *GETCheckoutComGatewaysCheckoutComGatewayId200Response
	isSet bool
}

func (v NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response) Get() *GETCheckoutComGatewaysCheckoutComGatewayId200Response {
	return v.value
}

func (v *NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response) Set(val *GETCheckoutComGatewaysCheckoutComGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCheckoutComGatewaysCheckoutComGatewayId200Response(val *GETCheckoutComGatewaysCheckoutComGatewayId200Response) *NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response {
	return &NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response{value: val, isSet: true}
}

func (v NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCheckoutComGatewaysCheckoutComGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
