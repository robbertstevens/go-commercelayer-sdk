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

// checks if the POSTBingGeocodersRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBingGeocodersRequestDataAttributes{}

// POSTBingGeocodersRequestDataAttributes struct for POSTBingGeocodersRequestDataAttributes
type POSTBingGeocodersRequestDataAttributes struct {
	// The geocoder's internal name
	Name interface{} `json:"name"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The Bing Virtualearth key
	Key interface{} `json:"key"`
}

// NewPOSTBingGeocodersRequestDataAttributes instantiates a new POSTBingGeocodersRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBingGeocodersRequestDataAttributes(name interface{}, key interface{}) *POSTBingGeocodersRequestDataAttributes {
	this := POSTBingGeocodersRequestDataAttributes{}
	this.Name = name
	this.Key = key
	return &this
}

// NewPOSTBingGeocodersRequestDataAttributesWithDefaults instantiates a new POSTBingGeocodersRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBingGeocodersRequestDataAttributesWithDefaults() *POSTBingGeocodersRequestDataAttributes {
	this := POSTBingGeocodersRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTBingGeocodersRequestDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTBingGeocodersRequestDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTBingGeocodersRequestDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTBingGeocodersRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTBingGeocodersRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTBingGeocodersRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTBingGeocodersRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTBingGeocodersRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTBingGeocodersRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTBingGeocodersRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTBingGeocodersRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTBingGeocodersRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTBingGeocodersRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTBingGeocodersRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTBingGeocodersRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTBingGeocodersRequestDataAttributes) GetKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTBingGeocodersRequestDataAttributes) GetKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *POSTBingGeocodersRequestDataAttributes) SetKey(v interface{}) {
	o.Key = v
}

func (o POSTBingGeocodersRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBingGeocodersRequestDataAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullablePOSTBingGeocodersRequestDataAttributes struct {
	value *POSTBingGeocodersRequestDataAttributes
	isSet bool
}

func (v NullablePOSTBingGeocodersRequestDataAttributes) Get() *POSTBingGeocodersRequestDataAttributes {
	return v.value
}

func (v *NullablePOSTBingGeocodersRequestDataAttributes) Set(val *POSTBingGeocodersRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBingGeocodersRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBingGeocodersRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBingGeocodersRequestDataAttributes(val *POSTBingGeocodersRequestDataAttributes) *NullablePOSTBingGeocodersRequestDataAttributes {
	return &NullablePOSTBingGeocodersRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTBingGeocodersRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBingGeocodersRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
