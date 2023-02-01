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

// PATCHInventoryModelsInventoryModelId200Response struct for PATCHInventoryModelsInventoryModelId200Response
type PATCHInventoryModelsInventoryModelId200Response struct {
	Data *PATCHInventoryModelsInventoryModelId200ResponseData `json:"data,omitempty"`
}

// NewPATCHInventoryModelsInventoryModelId200Response instantiates a new PATCHInventoryModelsInventoryModelId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHInventoryModelsInventoryModelId200Response() *PATCHInventoryModelsInventoryModelId200Response {
	this := PATCHInventoryModelsInventoryModelId200Response{}
	return &this
}

// NewPATCHInventoryModelsInventoryModelId200ResponseWithDefaults instantiates a new PATCHInventoryModelsInventoryModelId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHInventoryModelsInventoryModelId200ResponseWithDefaults() *PATCHInventoryModelsInventoryModelId200Response {
	this := PATCHInventoryModelsInventoryModelId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHInventoryModelsInventoryModelId200Response) GetData() PATCHInventoryModelsInventoryModelId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHInventoryModelsInventoryModelId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHInventoryModelsInventoryModelId200Response) GetDataOk() (*PATCHInventoryModelsInventoryModelId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHInventoryModelsInventoryModelId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHInventoryModelsInventoryModelId200ResponseData and assigns it to the Data field.
func (o *PATCHInventoryModelsInventoryModelId200Response) SetData(v PATCHInventoryModelsInventoryModelId200ResponseData) {
	o.Data = &v
}

func (o PATCHInventoryModelsInventoryModelId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHInventoryModelsInventoryModelId200Response struct {
	value *PATCHInventoryModelsInventoryModelId200Response
	isSet bool
}

func (v NullablePATCHInventoryModelsInventoryModelId200Response) Get() *PATCHInventoryModelsInventoryModelId200Response {
	return v.value
}

func (v *NullablePATCHInventoryModelsInventoryModelId200Response) Set(val *PATCHInventoryModelsInventoryModelId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHInventoryModelsInventoryModelId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHInventoryModelsInventoryModelId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHInventoryModelsInventoryModelId200Response(val *PATCHInventoryModelsInventoryModelId200Response) *NullablePATCHInventoryModelsInventoryModelId200Response {
	return &NullablePATCHInventoryModelsInventoryModelId200Response{value: val, isSet: true}
}

func (v NullablePATCHInventoryModelsInventoryModelId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHInventoryModelsInventoryModelId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
