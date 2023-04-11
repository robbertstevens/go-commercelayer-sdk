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

// checks if the DeliveryLeadTimeData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryLeadTimeData{}

// DeliveryLeadTimeData struct for DeliveryLeadTimeData
type DeliveryLeadTimeData struct {
	// The resource's type
	Type          interface{}                                                     `json:"type"`
	Attributes    GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes `json:"attributes"`
	Relationships *DeliveryLeadTimeDataRelationships                              `json:"relationships,omitempty"`
}

// NewDeliveryLeadTimeData instantiates a new DeliveryLeadTimeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeData(type_ interface{}, attributes GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) *DeliveryLeadTimeData {
	this := DeliveryLeadTimeData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewDeliveryLeadTimeDataWithDefaults instantiates a new DeliveryLeadTimeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeDataWithDefaults() *DeliveryLeadTimeData {
	this := DeliveryLeadTimeData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *DeliveryLeadTimeData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryLeadTimeData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeliveryLeadTimeData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *DeliveryLeadTimeData) GetAttributes() GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes {
	if o == nil {
		var ret GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeData) GetAttributesOk() (*GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *DeliveryLeadTimeData) SetAttributes(v GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *DeliveryLeadTimeData) GetRelationships() DeliveryLeadTimeDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret DeliveryLeadTimeDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeData) GetRelationshipsOk() (*DeliveryLeadTimeDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *DeliveryLeadTimeData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given DeliveryLeadTimeDataRelationships and assigns it to the Relationships field.
func (o *DeliveryLeadTimeData) SetRelationships(v DeliveryLeadTimeDataRelationships) {
	o.Relationships = &v
}

func (o DeliveryLeadTimeData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryLeadTimeData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableDeliveryLeadTimeData struct {
	value *DeliveryLeadTimeData
	isSet bool
}

func (v NullableDeliveryLeadTimeData) Get() *DeliveryLeadTimeData {
	return v.value
}

func (v *NullableDeliveryLeadTimeData) Set(val *DeliveryLeadTimeData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeData(val *DeliveryLeadTimeData) *NullableDeliveryLeadTimeData {
	return &NullableDeliveryLeadTimeData{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
