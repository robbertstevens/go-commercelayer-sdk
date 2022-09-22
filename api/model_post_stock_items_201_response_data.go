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

// POSTStockItems201ResponseData struct for POSTStockItems201ResponseData
type POSTStockItems201ResponseData struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                         `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks          `json:"links,omitempty"`
	Attributes    *POSTStockItems201ResponseDataAttributes        `json:"attributes,omitempty"`
	Relationships *GETStockItems200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewPOSTStockItems201ResponseData instantiates a new POSTStockItems201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStockItems201ResponseData() *POSTStockItems201ResponseData {
	this := POSTStockItems201ResponseData{}
	return &this
}

// NewPOSTStockItems201ResponseDataWithDefaults instantiates a new POSTStockItems201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStockItems201ResponseDataWithDefaults() *POSTStockItems201ResponseData {
	this := POSTStockItems201ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *POSTStockItems201ResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockItems201ResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTStockItems201ResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *POSTStockItems201ResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *POSTStockItems201ResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockItems201ResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTStockItems201ResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *POSTStockItems201ResponseData) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTStockItems201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockItems201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTStockItems201ResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *POSTStockItems201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *POSTStockItems201ResponseData) GetAttributes() POSTStockItems201ResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret POSTStockItems201ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockItems201ResponseData) GetAttributesOk() (*POSTStockItems201ResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *POSTStockItems201ResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given POSTStockItems201ResponseDataAttributes and assigns it to the Attributes field.
func (o *POSTStockItems201ResponseData) SetAttributes(v POSTStockItems201ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTStockItems201ResponseData) GetRelationships() GETStockItems200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETStockItems200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockItems201ResponseData) GetRelationshipsOk() (*GETStockItems200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTStockItems201ResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETStockItems200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *POSTStockItems201ResponseData) SetRelationships(v GETStockItems200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o POSTStockItems201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTStockItems201ResponseData struct {
	value *POSTStockItems201ResponseData
	isSet bool
}

func (v NullablePOSTStockItems201ResponseData) Get() *POSTStockItems201ResponseData {
	return v.value
}

func (v *NullablePOSTStockItems201ResponseData) Set(val *POSTStockItems201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStockItems201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStockItems201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStockItems201ResponseData(val *POSTStockItems201ResponseData) *NullablePOSTStockItems201ResponseData {
	return &NullablePOSTStockItems201ResponseData{value: val, isSet: true}
}

func (v NullablePOSTStockItems201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStockItems201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
