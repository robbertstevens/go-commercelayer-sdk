/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTInventoryStockLocationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTInventoryStockLocationsRequest{}

// POSTInventoryStockLocationsRequest struct for POSTInventoryStockLocationsRequest
type POSTInventoryStockLocationsRequest struct {
	Data POSTInventoryStockLocationsRequestData `json:"data"`
}

// NewPOSTInventoryStockLocationsRequest instantiates a new POSTInventoryStockLocationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTInventoryStockLocationsRequest(data POSTInventoryStockLocationsRequestData) *POSTInventoryStockLocationsRequest {
	this := POSTInventoryStockLocationsRequest{}
	this.Data = data
	return &this
}

// NewPOSTInventoryStockLocationsRequestWithDefaults instantiates a new POSTInventoryStockLocationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTInventoryStockLocationsRequestWithDefaults() *POSTInventoryStockLocationsRequest {
	this := POSTInventoryStockLocationsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTInventoryStockLocationsRequest) GetData() POSTInventoryStockLocationsRequestData {
	if o == nil {
		var ret POSTInventoryStockLocationsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTInventoryStockLocationsRequest) GetDataOk() (*POSTInventoryStockLocationsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTInventoryStockLocationsRequest) SetData(v POSTInventoryStockLocationsRequestData) {
	o.Data = v
}

func (o POSTInventoryStockLocationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTInventoryStockLocationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTInventoryStockLocationsRequest struct {
	value *POSTInventoryStockLocationsRequest
	isSet bool
}

func (v NullablePOSTInventoryStockLocationsRequest) Get() *POSTInventoryStockLocationsRequest {
	return v.value
}

func (v *NullablePOSTInventoryStockLocationsRequest) Set(val *POSTInventoryStockLocationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTInventoryStockLocationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTInventoryStockLocationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTInventoryStockLocationsRequest(val *POSTInventoryStockLocationsRequest) *NullablePOSTInventoryStockLocationsRequest {
	return &NullablePOSTInventoryStockLocationsRequest{value: val, isSet: true}
}

func (v NullablePOSTInventoryStockLocationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTInventoryStockLocationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
