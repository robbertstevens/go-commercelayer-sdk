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

// checks if the CustomerUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerUpdateData{}

// CustomerUpdateData struct for CustomerUpdateData
type CustomerUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                       `json:"id"`
	Attributes    PATCHCustomersCustomerId200ResponseDataAttributes `json:"attributes"`
	Relationships *CustomerCreateDataRelationships                  `json:"relationships,omitempty"`
}

// NewCustomerUpdateData instantiates a new CustomerUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerUpdateData(type_ interface{}, id interface{}, attributes PATCHCustomersCustomerId200ResponseDataAttributes) *CustomerUpdateData {
	this := CustomerUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCustomerUpdateDataWithDefaults instantiates a new CustomerUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerUpdateDataWithDefaults() *CustomerUpdateData {
	this := CustomerUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CustomerUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CustomerUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerUpdateData) GetAttributes() PATCHCustomersCustomerId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHCustomersCustomerId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerUpdateData) GetAttributesOk() (*PATCHCustomersCustomerId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerUpdateData) SetAttributes(v PATCHCustomersCustomerId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerUpdateData) GetRelationships() CustomerCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CustomerCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateData) GetRelationshipsOk() (*CustomerCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerCreateDataRelationships and assigns it to the Relationships field.
func (o *CustomerUpdateData) SetRelationships(v CustomerCreateDataRelationships) {
	o.Relationships = &v
}

func (o CustomerUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerUpdateData) ToMap() (map[string]interface{}, error) {
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

type NullableCustomerUpdateData struct {
	value *CustomerUpdateData
	isSet bool
}

func (v NullableCustomerUpdateData) Get() *CustomerUpdateData {
	return v.value
}

func (v *NullableCustomerUpdateData) Set(val *CustomerUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerUpdateData(val *CustomerUpdateData) *NullableCustomerUpdateData {
	return &NullableCustomerUpdateData{value: val, isSet: true}
}

func (v NullableCustomerUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
