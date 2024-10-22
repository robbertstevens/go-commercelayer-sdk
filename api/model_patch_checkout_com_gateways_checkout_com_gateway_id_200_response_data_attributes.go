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

// checks if the PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes{}

// PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes struct for PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes
type PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name interface{} `json:"name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The gateway secret key.
	SecretKey interface{} `json:"secret_key,omitempty"`
	// The gateway public key.
	PublicKey interface{} `json:"public_key,omitempty"`
}

// NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes instantiates a new PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes() *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes {
	this := PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes{}
	return &this
}

// NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributesWithDefaults() *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes {
	this := PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetSecretKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetSecretKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return &o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasSecretKey() bool {
	if o != nil && IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given interface{} and assigns it to the SecretKey field.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetSecretKey(v interface{}) {
	o.SecretKey = v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetPublicKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetPublicKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return &o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasPublicKey() bool {
	if o != nil && IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given interface{} and assigns it to the PublicKey field.
func (o *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetPublicKey(v interface{}) {
	o.PublicKey = v
}

func (o PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.SecretKey != nil {
		toSerialize["secret_key"] = o.SecretKey
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	return toSerialize, nil
}

type NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes struct {
	value *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) Get() *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) Set(val *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes(val *PATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) *NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes {
	return &NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
