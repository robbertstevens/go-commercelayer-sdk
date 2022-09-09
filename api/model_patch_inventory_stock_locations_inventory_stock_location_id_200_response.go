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

// PATCHInventoryStockLocationsInventoryStockLocationId200Response struct for PATCHInventoryStockLocationsInventoryStockLocationId200Response
type PATCHInventoryStockLocationsInventoryStockLocationId200Response struct {
	Data *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseData `json:"data,omitempty"`
}

// NewPATCHInventoryStockLocationsInventoryStockLocationId200Response instantiates a new PATCHInventoryStockLocationsInventoryStockLocationId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHInventoryStockLocationsInventoryStockLocationId200Response() *PATCHInventoryStockLocationsInventoryStockLocationId200Response {
	this := PATCHInventoryStockLocationsInventoryStockLocationId200Response{}
	return &this
}

// NewPATCHInventoryStockLocationsInventoryStockLocationId200ResponseWithDefaults instantiates a new PATCHInventoryStockLocationsInventoryStockLocationId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHInventoryStockLocationsInventoryStockLocationId200ResponseWithDefaults() *PATCHInventoryStockLocationsInventoryStockLocationId200Response {
	this := PATCHInventoryStockLocationsInventoryStockLocationId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200Response) GetData() PATCHInventoryStockLocationsInventoryStockLocationId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHInventoryStockLocationsInventoryStockLocationId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200Response) GetDataOk() (*PATCHInventoryStockLocationsInventoryStockLocationId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHInventoryStockLocationsInventoryStockLocationId200ResponseData and assigns it to the Data field.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200Response) SetData(v PATCHInventoryStockLocationsInventoryStockLocationId200ResponseData) {
	o.Data = &v
}

func (o PATCHInventoryStockLocationsInventoryStockLocationId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHInventoryStockLocationsInventoryStockLocationId200Response struct {
	value *PATCHInventoryStockLocationsInventoryStockLocationId200Response
	isSet bool
}

func (v NullablePATCHInventoryStockLocationsInventoryStockLocationId200Response) Get() *PATCHInventoryStockLocationsInventoryStockLocationId200Response {
	return v.value
}

func (v *NullablePATCHInventoryStockLocationsInventoryStockLocationId200Response) Set(val *PATCHInventoryStockLocationsInventoryStockLocationId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHInventoryStockLocationsInventoryStockLocationId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHInventoryStockLocationsInventoryStockLocationId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHInventoryStockLocationsInventoryStockLocationId200Response(val *PATCHInventoryStockLocationsInventoryStockLocationId200Response) *NullablePATCHInventoryStockLocationsInventoryStockLocationId200Response {
	return &NullablePATCHInventoryStockLocationsInventoryStockLocationId200Response{value: val, isSet: true}
}

func (v NullablePATCHInventoryStockLocationsInventoryStockLocationId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHInventoryStockLocationsInventoryStockLocationId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


