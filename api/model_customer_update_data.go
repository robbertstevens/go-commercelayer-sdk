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

// CustomerUpdateData struct for CustomerUpdateData
type CustomerUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
	Attributes PATCHCustomersCustomerId200ResponseDataAttributes `json:"attributes"`
	Relationships *CustomerCreateDataRelationships `json:"relationships,omitempty"`
}

// NewCustomerUpdateData instantiates a new CustomerUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerUpdateData(type_ string, id string, attributes PATCHCustomersCustomerId200ResponseDataAttributes) *CustomerUpdateData {
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
	var type_ string = "customers"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CustomerUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CustomerUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerUpdateData) SetId(v string) {
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
	if o == nil || o.Relationships == nil {
		var ret CustomerCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateData) GetRelationshipsOk() (*CustomerCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerCreateDataRelationships and assigns it to the Relationships field.
func (o *CustomerUpdateData) SetRelationships(v CustomerCreateDataRelationships) {
	o.Relationships = &v
}

func (o CustomerUpdateData) MarshalJSON() ([]byte, error) {
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


