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

// checks if the POSTDeliveryLeadTimesRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTDeliveryLeadTimesRequestDataRelationships{}

// POSTDeliveryLeadTimesRequestDataRelationships struct for POSTDeliveryLeadTimesRequestDataRelationships
type POSTDeliveryLeadTimesRequestDataRelationships struct {
	StockLocation  POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation  `json:"stock_location"`
	ShippingMethod POSTDeliveryLeadTimesRequestDataRelationshipsShippingMethod `json:"shipping_method"`
}

// NewPOSTDeliveryLeadTimesRequestDataRelationships instantiates a new POSTDeliveryLeadTimesRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTDeliveryLeadTimesRequestDataRelationships(stockLocation POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation, shippingMethod POSTDeliveryLeadTimesRequestDataRelationshipsShippingMethod) *POSTDeliveryLeadTimesRequestDataRelationships {
	this := POSTDeliveryLeadTimesRequestDataRelationships{}
	this.StockLocation = stockLocation
	this.ShippingMethod = shippingMethod
	return &this
}

// NewPOSTDeliveryLeadTimesRequestDataRelationshipsWithDefaults instantiates a new POSTDeliveryLeadTimesRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTDeliveryLeadTimesRequestDataRelationshipsWithDefaults() *POSTDeliveryLeadTimesRequestDataRelationships {
	this := POSTDeliveryLeadTimesRequestDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value
func (o *POSTDeliveryLeadTimesRequestDataRelationships) GetStockLocation() POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation {
	if o == nil {
		var ret POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation
		return ret
	}

	return o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value
// and a boolean to check if the value has been set.
func (o *POSTDeliveryLeadTimesRequestDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StockLocation, true
}

// SetStockLocation sets field value
func (o *POSTDeliveryLeadTimesRequestDataRelationships) SetStockLocation(v POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation) {
	o.StockLocation = v
}

// GetShippingMethod returns the ShippingMethod field value
func (o *POSTDeliveryLeadTimesRequestDataRelationships) GetShippingMethod() POSTDeliveryLeadTimesRequestDataRelationshipsShippingMethod {
	if o == nil {
		var ret POSTDeliveryLeadTimesRequestDataRelationshipsShippingMethod
		return ret
	}

	return o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value
// and a boolean to check if the value has been set.
func (o *POSTDeliveryLeadTimesRequestDataRelationships) GetShippingMethodOk() (*POSTDeliveryLeadTimesRequestDataRelationshipsShippingMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingMethod, true
}

// SetShippingMethod sets field value
func (o *POSTDeliveryLeadTimesRequestDataRelationships) SetShippingMethod(v POSTDeliveryLeadTimesRequestDataRelationshipsShippingMethod) {
	o.ShippingMethod = v
}

func (o POSTDeliveryLeadTimesRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTDeliveryLeadTimesRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stock_location"] = o.StockLocation
	toSerialize["shipping_method"] = o.ShippingMethod
	return toSerialize, nil
}

type NullablePOSTDeliveryLeadTimesRequestDataRelationships struct {
	value *POSTDeliveryLeadTimesRequestDataRelationships
	isSet bool
}

func (v NullablePOSTDeliveryLeadTimesRequestDataRelationships) Get() *POSTDeliveryLeadTimesRequestDataRelationships {
	return v.value
}

func (v *NullablePOSTDeliveryLeadTimesRequestDataRelationships) Set(val *POSTDeliveryLeadTimesRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTDeliveryLeadTimesRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTDeliveryLeadTimesRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTDeliveryLeadTimesRequestDataRelationships(val *POSTDeliveryLeadTimesRequestDataRelationships) *NullablePOSTDeliveryLeadTimesRequestDataRelationships {
	return &NullablePOSTDeliveryLeadTimesRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTDeliveryLeadTimesRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTDeliveryLeadTimesRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
