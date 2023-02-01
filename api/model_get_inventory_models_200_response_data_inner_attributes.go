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

// GETInventoryModels200ResponseDataInnerAttributes struct for GETInventoryModels200ResponseDataInnerAttributes
type GETInventoryModels200ResponseDataInnerAttributes struct {
	// The inventory model's internal name.
	Name *string `json:"name,omitempty"`
	// The inventory model's shipping strategy: one between 'no_split' (default), 'split_shipments', 'ship_from_primary' and 'ship_from_first_available_or_primary'.
	Strategy *string `json:"strategy,omitempty"`
	// The maximum number of stock locations used for inventory computation
	StockLocationsCutoff *int32 `json:"stock_locations_cutoff,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewGETInventoryModels200ResponseDataInnerAttributes instantiates a new GETInventoryModels200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryModels200ResponseDataInnerAttributes() *GETInventoryModels200ResponseDataInnerAttributes {
	this := GETInventoryModels200ResponseDataInnerAttributes{}
	return &this
}

// NewGETInventoryModels200ResponseDataInnerAttributesWithDefaults instantiates a new GETInventoryModels200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryModels200ResponseDataInnerAttributesWithDefaults() *GETInventoryModels200ResponseDataInnerAttributes {
	this := GETInventoryModels200ResponseDataInnerAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETInventoryModels200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetStrategy() string {
	if o == nil || o.Strategy == nil {
		var ret string
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetStrategyOk() (*string, bool) {
	if o == nil || o.Strategy == nil {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) HasStrategy() bool {
	if o != nil && o.Strategy != nil {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given string and assigns it to the Strategy field.
func (o *GETInventoryModels200ResponseDataInnerAttributes) SetStrategy(v string) {
	o.Strategy = &v
}

// GetStockLocationsCutoff returns the StockLocationsCutoff field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetStockLocationsCutoff() int32 {
	if o == nil || o.StockLocationsCutoff == nil {
		var ret int32
		return ret
	}
	return *o.StockLocationsCutoff
}

// GetStockLocationsCutoffOk returns a tuple with the StockLocationsCutoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetStockLocationsCutoffOk() (*int32, bool) {
	if o == nil || o.StockLocationsCutoff == nil {
		return nil, false
	}
	return o.StockLocationsCutoff, true
}

// HasStockLocationsCutoff returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) HasStockLocationsCutoff() bool {
	if o != nil && o.StockLocationsCutoff != nil {
		return true
	}

	return false
}

// SetStockLocationsCutoff gets a reference to the given int32 and assigns it to the StockLocationsCutoff field.
func (o *GETInventoryModels200ResponseDataInnerAttributes) SetStockLocationsCutoff(v int32) {
	o.StockLocationsCutoff = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETInventoryModels200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETInventoryModels200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETInventoryModels200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETInventoryModels200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETInventoryModels200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETInventoryModels200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Strategy != nil {
		toSerialize["strategy"] = o.Strategy
	}
	if o.StockLocationsCutoff != nil {
		toSerialize["stock_locations_cutoff"] = o.StockLocationsCutoff
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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

type NullableGETInventoryModels200ResponseDataInnerAttributes struct {
	value *GETInventoryModels200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETInventoryModels200ResponseDataInnerAttributes) Get() *GETInventoryModels200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETInventoryModels200ResponseDataInnerAttributes) Set(val *GETInventoryModels200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryModels200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryModels200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryModels200ResponseDataInnerAttributes(val *GETInventoryModels200ResponseDataInnerAttributes) *NullableGETInventoryModels200ResponseDataInnerAttributes {
	return &NullableGETInventoryModels200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETInventoryModels200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryModels200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
