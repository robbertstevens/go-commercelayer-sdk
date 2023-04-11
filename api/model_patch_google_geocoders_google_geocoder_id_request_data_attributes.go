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

// checks if the PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes{}

// PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes struct for PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes
type PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes struct {
	// The geocoder's internal name
	Name interface{} `json:"name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The Google Map API key
	ApiKey interface{} `json:"api_key,omitempty"`
}

// NewPATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes instantiates a new PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes() *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes {
	this := PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes{}
	return &this
}

// NewPATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributesWithDefaults instantiates a new PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributesWithDefaults() *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes {
	this := PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) GetApiKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) GetApiKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return &o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) HasApiKey() bool {
	if o != nil && IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given interface{} and assigns it to the ApiKey field.
func (o *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) SetApiKey(v interface{}) {
	o.ApiKey = v
}

func (o PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.ApiKey != nil {
		toSerialize["api_key"] = o.ApiKey
	}
	return toSerialize, nil
}

type NullablePATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes struct {
	value *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes
	isSet bool
}

func (v NullablePATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) Get() *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes {
	return v.value
}

func (v *NullablePATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) Set(val *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes(val *PATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) *NullablePATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes {
	return &NullablePATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHGoogleGeocodersGoogleGeocoderIdRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
