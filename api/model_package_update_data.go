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

// PackageUpdateData struct for PackageUpdateData
type PackageUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                          `json:"id"`
	Attributes    PATCHPackagesPackageId200ResponseDataAttributes `json:"attributes"`
	Relationships *PackageUpdateDataRelationships                 `json:"relationships,omitempty"`
}

// NewPackageUpdateData instantiates a new PackageUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageUpdateData(type_ string, id string, attributes PATCHPackagesPackageId200ResponseDataAttributes) *PackageUpdateData {
	this := PackageUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPackageUpdateDataWithDefaults instantiates a new PackageUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageUpdateDataWithDefaults() *PackageUpdateData {
	this := PackageUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *PackageUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PackageUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PackageUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PackageUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PackageUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PackageUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PackageUpdateData) GetAttributes() PATCHPackagesPackageId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHPackagesPackageId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PackageUpdateData) GetAttributesOk() (*PATCHPackagesPackageId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PackageUpdateData) SetAttributes(v PATCHPackagesPackageId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PackageUpdateData) GetRelationships() PackageUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PackageUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageUpdateData) GetRelationshipsOk() (*PackageUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PackageUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PackageUpdateDataRelationships and assigns it to the Relationships field.
func (o *PackageUpdateData) SetRelationships(v PackageUpdateDataRelationships) {
	o.Relationships = &v
}

func (o PackageUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullablePackageUpdateData struct {
	value *PackageUpdateData
	isSet bool
}

func (v NullablePackageUpdateData) Get() *PackageUpdateData {
	return v.value
}

func (v *NullablePackageUpdateData) Set(val *PackageUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageUpdateData(val *PackageUpdateData) *NullablePackageUpdateData {
	return &NullablePackageUpdateData{value: val, isSet: true}
}

func (v NullablePackageUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
