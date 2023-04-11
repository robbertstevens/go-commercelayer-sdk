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

// checks if the PATCHAxerveGatewaysAxerveGatewayIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHAxerveGatewaysAxerveGatewayIdRequestData{}

// PATCHAxerveGatewaysAxerveGatewayIdRequestData struct for PATCHAxerveGatewaysAxerveGatewayIdRequestData
type PATCHAxerveGatewaysAxerveGatewayIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                             `json:"id"`
	Attributes    PATCHAxerveGatewaysAxerveGatewayIdRequestDataAttributes `json:"attributes"`
	Relationships *POSTAxerveGatewaysRequestDataRelationships             `json:"relationships,omitempty"`
}

// NewPATCHAxerveGatewaysAxerveGatewayIdRequestData instantiates a new PATCHAxerveGatewaysAxerveGatewayIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAxerveGatewaysAxerveGatewayIdRequestData(type_ interface{}, id interface{}, attributes PATCHAxerveGatewaysAxerveGatewayIdRequestDataAttributes) *PATCHAxerveGatewaysAxerveGatewayIdRequestData {
	this := PATCHAxerveGatewaysAxerveGatewayIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHAxerveGatewaysAxerveGatewayIdRequestDataWithDefaults instantiates a new PATCHAxerveGatewaysAxerveGatewayIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAxerveGatewaysAxerveGatewayIdRequestDataWithDefaults() *PATCHAxerveGatewaysAxerveGatewayIdRequestData {
	this := PATCHAxerveGatewaysAxerveGatewayIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHAxerveGatewaysAxerveGatewayIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAxerveGatewaysAxerveGatewayIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHAxerveGatewaysAxerveGatewayIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHAxerveGatewaysAxerveGatewayIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAxerveGatewaysAxerveGatewayIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHAxerveGatewaysAxerveGatewayIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHAxerveGatewaysAxerveGatewayIdRequestData) GetAttributes() PATCHAxerveGatewaysAxerveGatewayIdRequestDataAttributes {
	if o == nil {
		var ret PATCHAxerveGatewaysAxerveGatewayIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHAxerveGatewaysAxerveGatewayIdRequestData) GetAttributesOk() (*PATCHAxerveGatewaysAxerveGatewayIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHAxerveGatewaysAxerveGatewayIdRequestData) SetAttributes(v PATCHAxerveGatewaysAxerveGatewayIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHAxerveGatewaysAxerveGatewayIdRequestData) GetRelationships() POSTAxerveGatewaysRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTAxerveGatewaysRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAxerveGatewaysAxerveGatewayIdRequestData) GetRelationshipsOk() (*POSTAxerveGatewaysRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHAxerveGatewaysAxerveGatewayIdRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTAxerveGatewaysRequestDataRelationships and assigns it to the Relationships field.
func (o *PATCHAxerveGatewaysAxerveGatewayIdRequestData) SetRelationships(v POSTAxerveGatewaysRequestDataRelationships) {
	o.Relationships = &v
}

func (o PATCHAxerveGatewaysAxerveGatewayIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHAxerveGatewaysAxerveGatewayIdRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePATCHAxerveGatewaysAxerveGatewayIdRequestData struct {
	value *PATCHAxerveGatewaysAxerveGatewayIdRequestData
	isSet bool
}

func (v NullablePATCHAxerveGatewaysAxerveGatewayIdRequestData) Get() *PATCHAxerveGatewaysAxerveGatewayIdRequestData {
	return v.value
}

func (v *NullablePATCHAxerveGatewaysAxerveGatewayIdRequestData) Set(val *PATCHAxerveGatewaysAxerveGatewayIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAxerveGatewaysAxerveGatewayIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAxerveGatewaysAxerveGatewayIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAxerveGatewaysAxerveGatewayIdRequestData(val *PATCHAxerveGatewaysAxerveGatewayIdRequestData) *NullablePATCHAxerveGatewaysAxerveGatewayIdRequestData {
	return &NullablePATCHAxerveGatewaysAxerveGatewayIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHAxerveGatewaysAxerveGatewayIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAxerveGatewaysAxerveGatewayIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
