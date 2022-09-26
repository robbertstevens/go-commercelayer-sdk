/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes struct for PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes
type PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name *string `json:"name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The gateway merchant account.
	MerchantAccount *string `json:"merchant_account,omitempty"`
	// The gateway API key.
	ApiKey *string `json:"api_key,omitempty"`
	// The public key linked to your API credential.
	PublicKey *string `json:"public_key,omitempty"`
	// The prefix of the endpoint used for live transactions.
	LiveUrlPrefix *string `json:"live_url_prefix,omitempty"`
}

// NewPATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes instantiates a new PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes() *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes {
	this := PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes{}
	return &this
}

// NewPATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributesWithDefaults() *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes {
	this := PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMerchantAccount returns the MerchantAccount field value if set, zero value otherwise.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetMerchantAccount() string {
	if o == nil || o.MerchantAccount == nil {
		var ret string
		return ret
	}
	return *o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetMerchantAccountOk() (*string, bool) {
	if o == nil || o.MerchantAccount == nil {
		return nil, false
	}
	return o.MerchantAccount, true
}

// HasMerchantAccount returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasMerchantAccount() bool {
	if o != nil && o.MerchantAccount != nil {
		return true
	}

	return false
}

// SetMerchantAccount gets a reference to the given string and assigns it to the MerchantAccount field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetMerchantAccount(v string) {
	o.MerchantAccount = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetLiveUrlPrefix returns the LiveUrlPrefix field value if set, zero value otherwise.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetLiveUrlPrefix() string {
	if o == nil || o.LiveUrlPrefix == nil {
		var ret string
		return ret
	}
	return *o.LiveUrlPrefix
}

// GetLiveUrlPrefixOk returns a tuple with the LiveUrlPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) GetLiveUrlPrefixOk() (*string, bool) {
	if o == nil || o.LiveUrlPrefix == nil {
		return nil, false
	}
	return o.LiveUrlPrefix, true
}

// HasLiveUrlPrefix returns a boolean if a field has been set.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) HasLiveUrlPrefix() bool {
	if o != nil && o.LiveUrlPrefix != nil {
		return true
	}

	return false
}

// SetLiveUrlPrefix gets a reference to the given string and assigns it to the LiveUrlPrefix field.
func (o *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) SetLiveUrlPrefix(v string) {
	o.LiveUrlPrefix = &v
}

func (o PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.MerchantAccount != nil {
		toSerialize["merchant_account"] = o.MerchantAccount
	}
	if o.ApiKey != nil {
		toSerialize["api_key"] = o.ApiKey
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	if o.LiveUrlPrefix != nil {
		toSerialize["live_url_prefix"] = o.LiveUrlPrefix
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes struct {
	value *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) Get() *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) Set(val *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes(val *PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) *NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes {
	return &NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
