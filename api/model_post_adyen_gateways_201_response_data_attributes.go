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

// POSTAdyenGateways201ResponseDataAttributes struct for POSTAdyenGateways201ResponseDataAttributes
type POSTAdyenGateways201ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name string `json:"name"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The gateway merchant account.
	MerchantAccount string `json:"merchant_account"`
	// The gateway API key.
	ApiKey string `json:"api_key"`
	// The public key linked to your API credential.
	PublicKey *string `json:"public_key,omitempty"`
	// The prefix of the endpoint used for live transactions.
	LiveUrlPrefix string `json:"live_url_prefix"`
}

// NewPOSTAdyenGateways201ResponseDataAttributes instantiates a new POSTAdyenGateways201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAdyenGateways201ResponseDataAttributes(name string, merchantAccount string, apiKey string, liveUrlPrefix string) *POSTAdyenGateways201ResponseDataAttributes {
	this := POSTAdyenGateways201ResponseDataAttributes{}
	this.Name = name
	this.MerchantAccount = merchantAccount
	this.ApiKey = apiKey
	this.LiveUrlPrefix = liveUrlPrefix
	return &this
}

// NewPOSTAdyenGateways201ResponseDataAttributesWithDefaults instantiates a new POSTAdyenGateways201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAdyenGateways201ResponseDataAttributesWithDefaults() *POSTAdyenGateways201ResponseDataAttributes {
	this := POSTAdyenGateways201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *POSTAdyenGateways201ResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTAdyenGateways201ResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTAdyenGateways201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTAdyenGateways201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTAdyenGateways201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTAdyenGateways201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTAdyenGateways201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTAdyenGateways201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTAdyenGateways201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTAdyenGateways201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTAdyenGateways201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *POSTAdyenGateways201ResponseDataAttributes) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseDataAttributes) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *POSTAdyenGateways201ResponseDataAttributes) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetApiKey returns the ApiKey field value
func (o *POSTAdyenGateways201ResponseDataAttributes) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseDataAttributes) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *POSTAdyenGateways201ResponseDataAttributes) SetApiKey(v string) {
	o.ApiKey = v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *POSTAdyenGateways201ResponseDataAttributes) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseDataAttributes) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *POSTAdyenGateways201ResponseDataAttributes) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *POSTAdyenGateways201ResponseDataAttributes) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetLiveUrlPrefix returns the LiveUrlPrefix field value
func (o *POSTAdyenGateways201ResponseDataAttributes) GetLiveUrlPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LiveUrlPrefix
}

// GetLiveUrlPrefixOk returns a tuple with the LiveUrlPrefix field value
// and a boolean to check if the value has been set.
func (o *POSTAdyenGateways201ResponseDataAttributes) GetLiveUrlPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LiveUrlPrefix, true
}

// SetLiveUrlPrefix sets field value
func (o *POSTAdyenGateways201ResponseDataAttributes) SetLiveUrlPrefix(v string) {
	o.LiveUrlPrefix = v
}

func (o POSTAdyenGateways201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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
	if true {
		toSerialize["merchant_account"] = o.MerchantAccount
	}
	if true {
		toSerialize["api_key"] = o.ApiKey
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	if true {
		toSerialize["live_url_prefix"] = o.LiveUrlPrefix
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTAdyenGateways201ResponseDataAttributes struct {
	value *POSTAdyenGateways201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTAdyenGateways201ResponseDataAttributes) Get() *POSTAdyenGateways201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTAdyenGateways201ResponseDataAttributes) Set(val *POSTAdyenGateways201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAdyenGateways201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAdyenGateways201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAdyenGateways201ResponseDataAttributes(val *POSTAdyenGateways201ResponseDataAttributes) *NullablePOSTAdyenGateways201ResponseDataAttributes {
	return &NullablePOSTAdyenGateways201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTAdyenGateways201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAdyenGateways201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
