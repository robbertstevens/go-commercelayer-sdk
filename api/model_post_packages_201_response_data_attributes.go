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

// checks if the POSTPackages201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPackages201ResponseDataAttributes{}

// POSTPackages201ResponseDataAttributes struct for POSTPackages201ResponseDataAttributes
type POSTPackages201ResponseDataAttributes struct {
	// Unique name for the package.
	Name interface{} `json:"name"`
	// The package identifying code.
	Code interface{} `json:"code,omitempty"`
	// The package length, used to automatically calculate the tax rates from the available carrier accounts.
	Length interface{} `json:"length"`
	// The package width, used to automatically calculate the tax rates from the available carrier accounts.
	Width interface{} `json:"width"`
	// The package height, used to automatically calculate the tax rates from the available carrier accounts.
	Height interface{} `json:"height"`
	// The unit of length. Can be one of 'cm', or 'in'.
	UnitOfLength interface{} `json:"unit_of_length"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTPackages201ResponseDataAttributes instantiates a new POSTPackages201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPackages201ResponseDataAttributes(name interface{}, length interface{}, width interface{}, height interface{}, unitOfLength interface{}) *POSTPackages201ResponseDataAttributes {
	this := POSTPackages201ResponseDataAttributes{}
	this.Name = name
	this.Length = length
	this.Width = width
	this.Height = height
	this.UnitOfLength = unitOfLength
	return &this
}

// NewPOSTPackages201ResponseDataAttributesWithDefaults instantiates a new POSTPackages201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPackages201ResponseDataAttributesWithDefaults() *POSTPackages201ResponseDataAttributes {
	this := POSTPackages201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPackages201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPackages201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTPackages201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPackages201ResponseDataAttributes) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPackages201ResponseDataAttributes) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *POSTPackages201ResponseDataAttributes) HasCode() bool {
	if o != nil && IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *POSTPackages201ResponseDataAttributes) SetCode(v interface{}) {
	o.Code = v
}

// GetLength returns the Length field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPackages201ResponseDataAttributes) GetLength() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPackages201ResponseDataAttributes) GetLengthOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *POSTPackages201ResponseDataAttributes) SetLength(v interface{}) {
	o.Length = v
}

// GetWidth returns the Width field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPackages201ResponseDataAttributes) GetWidth() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPackages201ResponseDataAttributes) GetWidthOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *POSTPackages201ResponseDataAttributes) SetWidth(v interface{}) {
	o.Width = v
}

// GetHeight returns the Height field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPackages201ResponseDataAttributes) GetHeight() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPackages201ResponseDataAttributes) GetHeightOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *POSTPackages201ResponseDataAttributes) SetHeight(v interface{}) {
	o.Height = v
}

// GetUnitOfLength returns the UnitOfLength field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPackages201ResponseDataAttributes) GetUnitOfLength() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.UnitOfLength
}

// GetUnitOfLengthOk returns a tuple with the UnitOfLength field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPackages201ResponseDataAttributes) GetUnitOfLengthOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UnitOfLength) {
		return nil, false
	}
	return &o.UnitOfLength, true
}

// SetUnitOfLength sets field value
func (o *POSTPackages201ResponseDataAttributes) SetUnitOfLength(v interface{}) {
	o.UnitOfLength = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPackages201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPackages201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTPackages201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTPackages201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPackages201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPackages201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTPackages201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTPackages201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPackages201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPackages201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTPackages201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTPackages201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTPackages201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPackages201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.UnitOfLength != nil {
		toSerialize["unit_of_length"] = o.UnitOfLength
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullablePOSTPackages201ResponseDataAttributes struct {
	value *POSTPackages201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTPackages201ResponseDataAttributes) Get() *POSTPackages201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTPackages201ResponseDataAttributes) Set(val *POSTPackages201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPackages201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPackages201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPackages201ResponseDataAttributes(val *POSTPackages201ResponseDataAttributes) *NullablePOSTPackages201ResponseDataAttributes {
	return &NullablePOSTPackages201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTPackages201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPackages201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
