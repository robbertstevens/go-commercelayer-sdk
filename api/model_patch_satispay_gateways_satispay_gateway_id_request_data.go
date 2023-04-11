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

// checks if the PATCHSatispayGatewaysSatispayGatewayIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHSatispayGatewaysSatispayGatewayIdRequestData{}

// PATCHSatispayGatewaysSatispayGatewayIdRequestData struct for PATCHSatispayGatewaysSatispayGatewayIdRequestData
type PATCHSatispayGatewaysSatispayGatewayIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                             `json:"id"`
	Attributes    PATCHManualGatewaysManualGatewayIdRequestDataAttributes `json:"attributes"`
	Relationships *POSTSatispayGatewaysRequestDataRelationships           `json:"relationships,omitempty"`
}

// NewPATCHSatispayGatewaysSatispayGatewayIdRequestData instantiates a new PATCHSatispayGatewaysSatispayGatewayIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHSatispayGatewaysSatispayGatewayIdRequestData(type_ interface{}, id interface{}, attributes PATCHManualGatewaysManualGatewayIdRequestDataAttributes) *PATCHSatispayGatewaysSatispayGatewayIdRequestData {
	this := PATCHSatispayGatewaysSatispayGatewayIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHSatispayGatewaysSatispayGatewayIdRequestDataWithDefaults instantiates a new PATCHSatispayGatewaysSatispayGatewayIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHSatispayGatewaysSatispayGatewayIdRequestDataWithDefaults() *PATCHSatispayGatewaysSatispayGatewayIdRequestData {
	this := PATCHSatispayGatewaysSatispayGatewayIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHSatispayGatewaysSatispayGatewayIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSatispayGatewaysSatispayGatewayIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHSatispayGatewaysSatispayGatewayIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHSatispayGatewaysSatispayGatewayIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSatispayGatewaysSatispayGatewayIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHSatispayGatewaysSatispayGatewayIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHSatispayGatewaysSatispayGatewayIdRequestData) GetAttributes() PATCHManualGatewaysManualGatewayIdRequestDataAttributes {
	if o == nil {
		var ret PATCHManualGatewaysManualGatewayIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHSatispayGatewaysSatispayGatewayIdRequestData) GetAttributesOk() (*PATCHManualGatewaysManualGatewayIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHSatispayGatewaysSatispayGatewayIdRequestData) SetAttributes(v PATCHManualGatewaysManualGatewayIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHSatispayGatewaysSatispayGatewayIdRequestData) GetRelationships() POSTSatispayGatewaysRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTSatispayGatewaysRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSatispayGatewaysSatispayGatewayIdRequestData) GetRelationshipsOk() (*POSTSatispayGatewaysRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHSatispayGatewaysSatispayGatewayIdRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTSatispayGatewaysRequestDataRelationships and assigns it to the Relationships field.
func (o *PATCHSatispayGatewaysSatispayGatewayIdRequestData) SetRelationships(v POSTSatispayGatewaysRequestDataRelationships) {
	o.Relationships = &v
}

func (o PATCHSatispayGatewaysSatispayGatewayIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHSatispayGatewaysSatispayGatewayIdRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePATCHSatispayGatewaysSatispayGatewayIdRequestData struct {
	value *PATCHSatispayGatewaysSatispayGatewayIdRequestData
	isSet bool
}

func (v NullablePATCHSatispayGatewaysSatispayGatewayIdRequestData) Get() *PATCHSatispayGatewaysSatispayGatewayIdRequestData {
	return v.value
}

func (v *NullablePATCHSatispayGatewaysSatispayGatewayIdRequestData) Set(val *PATCHSatispayGatewaysSatispayGatewayIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHSatispayGatewaysSatispayGatewayIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHSatispayGatewaysSatispayGatewayIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHSatispayGatewaysSatispayGatewayIdRequestData(val *PATCHSatispayGatewaysSatispayGatewayIdRequestData) *NullablePATCHSatispayGatewaysSatispayGatewayIdRequestData {
	return &NullablePATCHSatispayGatewaysSatispayGatewayIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHSatispayGatewaysSatispayGatewayIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHSatispayGatewaysSatispayGatewayIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
