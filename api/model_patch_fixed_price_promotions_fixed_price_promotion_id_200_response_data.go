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

// checks if the PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData{}

// PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData struct for PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData
type PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                                          `json:"type,omitempty"`
	Links         *POSTAddresses201ResponseDataLinks                                   `json:"links,omitempty"`
	Attributes    *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataAttributes `json:"attributes,omitempty"`
	Relationships *POSTFixedPricePromotions201ResponseDataRelationships                `json:"relationships,omitempty"`
}

// NewPATCHFixedPricePromotionsFixedPricePromotionId200ResponseData instantiates a new PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHFixedPricePromotionsFixedPricePromotionId200ResponseData() *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData {
	this := PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData{}
	return &this
}

// NewPATCHFixedPricePromotionsFixedPricePromotionId200ResponseDataWithDefaults instantiates a new PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHFixedPricePromotionsFixedPricePromotionId200ResponseDataWithDefaults() *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData {
	this := PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) GetLinks() POSTAddresses201ResponseDataLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) GetLinksOk() (*POSTAddresses201ResponseDataLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataLinks and assigns it to the Links field.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) SetLinks(v POSTAddresses201ResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) GetAttributes() PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) GetAttributesOk() (*PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataAttributes and assigns it to the Attributes field.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) SetAttributes(v PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) GetRelationships() POSTFixedPricePromotions201ResponseDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTFixedPricePromotions201ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) GetRelationshipsOk() (*POSTFixedPricePromotions201ResponseDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTFixedPricePromotions201ResponseDataRelationships and assigns it to the Relationships field.
func (o *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) SetRelationships(v POSTFixedPricePromotions201ResponseDataRelationships) {
	o.Relationships = &v
}

func (o PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePATCHFixedPricePromotionsFixedPricePromotionId200ResponseData struct {
	value *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData
	isSet bool
}

func (v NullablePATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) Get() *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData {
	return v.value
}

func (v *NullablePATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) Set(val *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHFixedPricePromotionsFixedPricePromotionId200ResponseData(val *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) *NullablePATCHFixedPricePromotionsFixedPricePromotionId200ResponseData {
	return &NullablePATCHFixedPricePromotionsFixedPricePromotionId200ResponseData{value: val, isSet: true}
}

func (v NullablePATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHFixedPricePromotionsFixedPricePromotionId200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
