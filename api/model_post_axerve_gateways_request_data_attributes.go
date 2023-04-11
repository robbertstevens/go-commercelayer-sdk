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

// checks if the POSTAxerveGatewaysRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAxerveGatewaysRequestDataAttributes{}

// POSTAxerveGatewaysRequestDataAttributes struct for POSTAxerveGatewaysRequestDataAttributes
type POSTAxerveGatewaysRequestDataAttributes struct {
	// The payment gateway's internal name.
	Name interface{} `json:"name"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The merchant login code.
	Login interface{} `json:"login"`
	// The gateway API key.
	ApiKey interface{} `json:"api_key"`
}

// NewPOSTAxerveGatewaysRequestDataAttributes instantiates a new POSTAxerveGatewaysRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAxerveGatewaysRequestDataAttributes(name interface{}, login interface{}, apiKey interface{}) *POSTAxerveGatewaysRequestDataAttributes {
	this := POSTAxerveGatewaysRequestDataAttributes{}
	this.Name = name
	this.Login = login
	this.ApiKey = apiKey
	return &this
}

// NewPOSTAxerveGatewaysRequestDataAttributesWithDefaults instantiates a new POSTAxerveGatewaysRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAxerveGatewaysRequestDataAttributesWithDefaults() *POSTAxerveGatewaysRequestDataAttributes {
	this := POSTAxerveGatewaysRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTAxerveGatewaysRequestDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxerveGatewaysRequestDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTAxerveGatewaysRequestDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAxerveGatewaysRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxerveGatewaysRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTAxerveGatewaysRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTAxerveGatewaysRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAxerveGatewaysRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxerveGatewaysRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTAxerveGatewaysRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTAxerveGatewaysRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAxerveGatewaysRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxerveGatewaysRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTAxerveGatewaysRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTAxerveGatewaysRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetLogin returns the Login field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTAxerveGatewaysRequestDataAttributes) GetLogin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxerveGatewaysRequestDataAttributes) GetLoginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *POSTAxerveGatewaysRequestDataAttributes) SetLogin(v interface{}) {
	o.Login = v
}

// GetApiKey returns the ApiKey field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTAxerveGatewaysRequestDataAttributes) GetApiKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxerveGatewaysRequestDataAttributes) GetApiKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *POSTAxerveGatewaysRequestDataAttributes) SetApiKey(v interface{}) {
	o.ApiKey = v
}

func (o POSTAxerveGatewaysRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAxerveGatewaysRequestDataAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.ApiKey != nil {
		toSerialize["api_key"] = o.ApiKey
	}
	return toSerialize, nil
}

type NullablePOSTAxerveGatewaysRequestDataAttributes struct {
	value *POSTAxerveGatewaysRequestDataAttributes
	isSet bool
}

func (v NullablePOSTAxerveGatewaysRequestDataAttributes) Get() *POSTAxerveGatewaysRequestDataAttributes {
	return v.value
}

func (v *NullablePOSTAxerveGatewaysRequestDataAttributes) Set(val *POSTAxerveGatewaysRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAxerveGatewaysRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAxerveGatewaysRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAxerveGatewaysRequestDataAttributes(val *POSTAxerveGatewaysRequestDataAttributes) *NullablePOSTAxerveGatewaysRequestDataAttributes {
	return &NullablePOSTAxerveGatewaysRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTAxerveGatewaysRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAxerveGatewaysRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
