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

// POSTShippingZones201ResponseDataAttributes struct for POSTShippingZones201ResponseDataAttributes
type POSTShippingZones201ResponseDataAttributes struct {
	// The shipping zone's internal name.
	Name string `json:"name"`
	// The regex that will be evaluated to match the shipping address country code.
	CountryCodeRegex *string `json:"country_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping address country code.
	NotCountryCodeRegex *string `json:"not_country_code_regex,omitempty"`
	// The regex that will be evaluated to match the shipping address state code.
	StateCodeRegex *string `json:"state_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping address state code.
	NotStateCodeRegex *string `json:"not_state_code_regex,omitempty"`
	// The regex that will be evaluated to match the shipping address zip code.
	ZipCodeRegex *string `json:"zip_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping zip country code.
	NotZipCodeRegex *string `json:"not_zip_code_regex,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTShippingZones201ResponseDataAttributes instantiates a new POSTShippingZones201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingZones201ResponseDataAttributes(name string) *POSTShippingZones201ResponseDataAttributes {
	this := POSTShippingZones201ResponseDataAttributes{}
	this.Name = name
	return &this
}

// NewPOSTShippingZones201ResponseDataAttributesWithDefaults instantiates a new POSTShippingZones201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingZones201ResponseDataAttributesWithDefaults() *POSTShippingZones201ResponseDataAttributes {
	this := POSTShippingZones201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *POSTShippingZones201ResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *POSTShippingZones201ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTShippingZones201ResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetCountryCodeRegex returns the CountryCodeRegex field value if set, zero value otherwise.
func (o *POSTShippingZones201ResponseDataAttributes) GetCountryCodeRegex() string {
	if o == nil || o.CountryCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.CountryCodeRegex
}

// GetCountryCodeRegexOk returns a tuple with the CountryCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingZones201ResponseDataAttributes) GetCountryCodeRegexOk() (*string, bool) {
	if o == nil || o.CountryCodeRegex == nil {
		return nil, false
	}
	return o.CountryCodeRegex, true
}

// HasCountryCodeRegex returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasCountryCodeRegex() bool {
	if o != nil && o.CountryCodeRegex != nil {
		return true
	}

	return false
}

// SetCountryCodeRegex gets a reference to the given string and assigns it to the CountryCodeRegex field.
func (o *POSTShippingZones201ResponseDataAttributes) SetCountryCodeRegex(v string) {
	o.CountryCodeRegex = &v
}

// GetNotCountryCodeRegex returns the NotCountryCodeRegex field value if set, zero value otherwise.
func (o *POSTShippingZones201ResponseDataAttributes) GetNotCountryCodeRegex() string {
	if o == nil || o.NotCountryCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.NotCountryCodeRegex
}

// GetNotCountryCodeRegexOk returns a tuple with the NotCountryCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingZones201ResponseDataAttributes) GetNotCountryCodeRegexOk() (*string, bool) {
	if o == nil || o.NotCountryCodeRegex == nil {
		return nil, false
	}
	return o.NotCountryCodeRegex, true
}

// HasNotCountryCodeRegex returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasNotCountryCodeRegex() bool {
	if o != nil && o.NotCountryCodeRegex != nil {
		return true
	}

	return false
}

// SetNotCountryCodeRegex gets a reference to the given string and assigns it to the NotCountryCodeRegex field.
func (o *POSTShippingZones201ResponseDataAttributes) SetNotCountryCodeRegex(v string) {
	o.NotCountryCodeRegex = &v
}

// GetStateCodeRegex returns the StateCodeRegex field value if set, zero value otherwise.
func (o *POSTShippingZones201ResponseDataAttributes) GetStateCodeRegex() string {
	if o == nil || o.StateCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.StateCodeRegex
}

// GetStateCodeRegexOk returns a tuple with the StateCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingZones201ResponseDataAttributes) GetStateCodeRegexOk() (*string, bool) {
	if o == nil || o.StateCodeRegex == nil {
		return nil, false
	}
	return o.StateCodeRegex, true
}

// HasStateCodeRegex returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasStateCodeRegex() bool {
	if o != nil && o.StateCodeRegex != nil {
		return true
	}

	return false
}

// SetStateCodeRegex gets a reference to the given string and assigns it to the StateCodeRegex field.
func (o *POSTShippingZones201ResponseDataAttributes) SetStateCodeRegex(v string) {
	o.StateCodeRegex = &v
}

// GetNotStateCodeRegex returns the NotStateCodeRegex field value if set, zero value otherwise.
func (o *POSTShippingZones201ResponseDataAttributes) GetNotStateCodeRegex() string {
	if o == nil || o.NotStateCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.NotStateCodeRegex
}

// GetNotStateCodeRegexOk returns a tuple with the NotStateCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingZones201ResponseDataAttributes) GetNotStateCodeRegexOk() (*string, bool) {
	if o == nil || o.NotStateCodeRegex == nil {
		return nil, false
	}
	return o.NotStateCodeRegex, true
}

// HasNotStateCodeRegex returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasNotStateCodeRegex() bool {
	if o != nil && o.NotStateCodeRegex != nil {
		return true
	}

	return false
}

// SetNotStateCodeRegex gets a reference to the given string and assigns it to the NotStateCodeRegex field.
func (o *POSTShippingZones201ResponseDataAttributes) SetNotStateCodeRegex(v string) {
	o.NotStateCodeRegex = &v
}

// GetZipCodeRegex returns the ZipCodeRegex field value if set, zero value otherwise.
func (o *POSTShippingZones201ResponseDataAttributes) GetZipCodeRegex() string {
	if o == nil || o.ZipCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.ZipCodeRegex
}

// GetZipCodeRegexOk returns a tuple with the ZipCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingZones201ResponseDataAttributes) GetZipCodeRegexOk() (*string, bool) {
	if o == nil || o.ZipCodeRegex == nil {
		return nil, false
	}
	return o.ZipCodeRegex, true
}

// HasZipCodeRegex returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasZipCodeRegex() bool {
	if o != nil && o.ZipCodeRegex != nil {
		return true
	}

	return false
}

// SetZipCodeRegex gets a reference to the given string and assigns it to the ZipCodeRegex field.
func (o *POSTShippingZones201ResponseDataAttributes) SetZipCodeRegex(v string) {
	o.ZipCodeRegex = &v
}

// GetNotZipCodeRegex returns the NotZipCodeRegex field value if set, zero value otherwise.
func (o *POSTShippingZones201ResponseDataAttributes) GetNotZipCodeRegex() string {
	if o == nil || o.NotZipCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.NotZipCodeRegex
}

// GetNotZipCodeRegexOk returns a tuple with the NotZipCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingZones201ResponseDataAttributes) GetNotZipCodeRegexOk() (*string, bool) {
	if o == nil || o.NotZipCodeRegex == nil {
		return nil, false
	}
	return o.NotZipCodeRegex, true
}

// HasNotZipCodeRegex returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasNotZipCodeRegex() bool {
	if o != nil && o.NotZipCodeRegex != nil {
		return true
	}

	return false
}

// SetNotZipCodeRegex gets a reference to the given string and assigns it to the NotZipCodeRegex field.
func (o *POSTShippingZones201ResponseDataAttributes) SetNotZipCodeRegex(v string) {
	o.NotZipCodeRegex = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTShippingZones201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingZones201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTShippingZones201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTShippingZones201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingZones201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTShippingZones201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTShippingZones201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingZones201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTShippingZones201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTShippingZones201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.CountryCodeRegex != nil {
		toSerialize["country_code_regex"] = o.CountryCodeRegex
	}
	if o.NotCountryCodeRegex != nil {
		toSerialize["not_country_code_regex"] = o.NotCountryCodeRegex
	}
	if o.StateCodeRegex != nil {
		toSerialize["state_code_regex"] = o.StateCodeRegex
	}
	if o.NotStateCodeRegex != nil {
		toSerialize["not_state_code_regex"] = o.NotStateCodeRegex
	}
	if o.ZipCodeRegex != nil {
		toSerialize["zip_code_regex"] = o.ZipCodeRegex
	}
	if o.NotZipCodeRegex != nil {
		toSerialize["not_zip_code_regex"] = o.NotZipCodeRegex
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
	return json.Marshal(toSerialize)
}

type NullablePOSTShippingZones201ResponseDataAttributes struct {
	value *POSTShippingZones201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTShippingZones201ResponseDataAttributes) Get() *POSTShippingZones201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTShippingZones201ResponseDataAttributes) Set(val *POSTShippingZones201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingZones201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingZones201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingZones201ResponseDataAttributes(val *POSTShippingZones201ResponseDataAttributes) *NullablePOSTShippingZones201ResponseDataAttributes {
	return &NullablePOSTShippingZones201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTShippingZones201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingZones201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
