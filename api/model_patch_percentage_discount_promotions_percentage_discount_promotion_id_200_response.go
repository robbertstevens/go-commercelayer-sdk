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

// PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response struct for PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response
type PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response struct {
	Data *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseData `json:"data,omitempty"`
}

// NewPATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response instantiates a new PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response() *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response {
	this := PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response{}
	return &this
}

// NewPATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseWithDefaults instantiates a new PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseWithDefaults() *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response {
	this := PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response) GetData() PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response) GetDataOk() (*PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseData and assigns it to the Data field.
func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response) SetData(v PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseData) {
	o.Data = &v
}

func (o PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response struct {
	value *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response
	isSet bool
}

func (v NullablePATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response) Get() *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response {
	return v.value
}

func (v *NullablePATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response) Set(val *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response(val *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response) *NullablePATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response {
	return &NullablePATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response{value: val, isSet: true}
}

func (v NullablePATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
