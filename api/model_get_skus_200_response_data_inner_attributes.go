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

// GETSkus200ResponseDataInnerAttributes struct for GETSkus200ResponseDataInnerAttributes
type GETSkus200ResponseDataInnerAttributes struct {
	// The SKU code, that uniquely identifies the SKU within the organization.
	Code *string `json:"code,omitempty"`
	// The internal name of the SKU.
	Name *string `json:"name,omitempty"`
	// An internal description of the SKU.
	Description *string `json:"description,omitempty"`
	// The URL of an image that represents the SKU.
	ImageUrl *string `json:"image_url,omitempty"`
	// The number of pieces that compose the SKU. This is useful to describe sets and bundles.
	PiecesPerPack *int32 `json:"pieces_per_pack,omitempty"`
	// The weight of the SKU. If present, it will be used to calculate the shipping rates.
	Weight *float32 `json:"weight,omitempty"`
	// Can be one of 'gr', 'lb', or 'oz'
	UnitOfWeight *string `json:"unit_of_weight,omitempty"`
	// The Harmonized System Code used by customs to identify the products shipped across international borders.
	HsTariffNumber *string `json:"hs_tariff_number,omitempty"`
	// Indicates if the SKU doesn't generate shipments.
	DoNotShip *bool `json:"do_not_ship,omitempty"`
	// Indicates if the SKU doesn't track the stock inventory.
	DoNotTrack *bool `json:"do_not_track,omitempty"`
	// Aggregated information about the SKU's inventory. Returned only when retrieving a single SKU.
	Inventory map[string]interface{} `json:"inventory,omitempty"`
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

// NewGETSkus200ResponseDataInnerAttributes instantiates a new GETSkus200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkus200ResponseDataInnerAttributes() *GETSkus200ResponseDataInnerAttributes {
	this := GETSkus200ResponseDataInnerAttributes{}
	return &this
}

// NewGETSkus200ResponseDataInnerAttributesWithDefaults instantiates a new GETSkus200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkus200ResponseDataInnerAttributesWithDefaults() *GETSkus200ResponseDataInnerAttributes {
	this := GETSkus200ResponseDataInnerAttributes{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GETSkus200ResponseDataInnerAttributes) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETSkus200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GETSkus200ResponseDataInnerAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *GETSkus200ResponseDataInnerAttributes) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetPiecesPerPack returns the PiecesPerPack field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetPiecesPerPack() int32 {
	if o == nil || o.PiecesPerPack == nil {
		var ret int32
		return ret
	}
	return *o.PiecesPerPack
}

// GetPiecesPerPackOk returns a tuple with the PiecesPerPack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetPiecesPerPackOk() (*int32, bool) {
	if o == nil || o.PiecesPerPack == nil {
		return nil, false
	}
	return o.PiecesPerPack, true
}

// HasPiecesPerPack returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasPiecesPerPack() bool {
	if o != nil && o.PiecesPerPack != nil {
		return true
	}

	return false
}

// SetPiecesPerPack gets a reference to the given int32 and assigns it to the PiecesPerPack field.
func (o *GETSkus200ResponseDataInnerAttributes) SetPiecesPerPack(v int32) {
	o.PiecesPerPack = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetWeight() float32 {
	if o == nil || o.Weight == nil {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetWeightOk() (*float32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *GETSkus200ResponseDataInnerAttributes) SetWeight(v float32) {
	o.Weight = &v
}

// GetUnitOfWeight returns the UnitOfWeight field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetUnitOfWeight() string {
	if o == nil || o.UnitOfWeight == nil {
		var ret string
		return ret
	}
	return *o.UnitOfWeight
}

// GetUnitOfWeightOk returns a tuple with the UnitOfWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetUnitOfWeightOk() (*string, bool) {
	if o == nil || o.UnitOfWeight == nil {
		return nil, false
	}
	return o.UnitOfWeight, true
}

// HasUnitOfWeight returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasUnitOfWeight() bool {
	if o != nil && o.UnitOfWeight != nil {
		return true
	}

	return false
}

// SetUnitOfWeight gets a reference to the given string and assigns it to the UnitOfWeight field.
func (o *GETSkus200ResponseDataInnerAttributes) SetUnitOfWeight(v string) {
	o.UnitOfWeight = &v
}

// GetHsTariffNumber returns the HsTariffNumber field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetHsTariffNumber() string {
	if o == nil || o.HsTariffNumber == nil {
		var ret string
		return ret
	}
	return *o.HsTariffNumber
}

// GetHsTariffNumberOk returns a tuple with the HsTariffNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetHsTariffNumberOk() (*string, bool) {
	if o == nil || o.HsTariffNumber == nil {
		return nil, false
	}
	return o.HsTariffNumber, true
}

// HasHsTariffNumber returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasHsTariffNumber() bool {
	if o != nil && o.HsTariffNumber != nil {
		return true
	}

	return false
}

// SetHsTariffNumber gets a reference to the given string and assigns it to the HsTariffNumber field.
func (o *GETSkus200ResponseDataInnerAttributes) SetHsTariffNumber(v string) {
	o.HsTariffNumber = &v
}

// GetDoNotShip returns the DoNotShip field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetDoNotShip() bool {
	if o == nil || o.DoNotShip == nil {
		var ret bool
		return ret
	}
	return *o.DoNotShip
}

// GetDoNotShipOk returns a tuple with the DoNotShip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetDoNotShipOk() (*bool, bool) {
	if o == nil || o.DoNotShip == nil {
		return nil, false
	}
	return o.DoNotShip, true
}

// HasDoNotShip returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasDoNotShip() bool {
	if o != nil && o.DoNotShip != nil {
		return true
	}

	return false
}

// SetDoNotShip gets a reference to the given bool and assigns it to the DoNotShip field.
func (o *GETSkus200ResponseDataInnerAttributes) SetDoNotShip(v bool) {
	o.DoNotShip = &v
}

// GetDoNotTrack returns the DoNotTrack field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetDoNotTrack() bool {
	if o == nil || o.DoNotTrack == nil {
		var ret bool
		return ret
	}
	return *o.DoNotTrack
}

// GetDoNotTrackOk returns a tuple with the DoNotTrack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetDoNotTrackOk() (*bool, bool) {
	if o == nil || o.DoNotTrack == nil {
		return nil, false
	}
	return o.DoNotTrack, true
}

// HasDoNotTrack returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasDoNotTrack() bool {
	if o != nil && o.DoNotTrack != nil {
		return true
	}

	return false
}

// SetDoNotTrack gets a reference to the given bool and assigns it to the DoNotTrack field.
func (o *GETSkus200ResponseDataInnerAttributes) SetDoNotTrack(v bool) {
	o.DoNotTrack = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetInventory() map[string]interface{} {
	if o == nil || o.Inventory == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetInventoryOk() (map[string]interface{}, bool) {
	if o == nil || o.Inventory == nil {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasInventory() bool {
	if o != nil && o.Inventory != nil {
		return true
	}

	return false
}

// SetInventory gets a reference to the given map[string]interface{} and assigns it to the Inventory field.
func (o *GETSkus200ResponseDataInnerAttributes) SetInventory(v map[string]interface{}) {
	o.Inventory = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETSkus200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETSkus200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETSkus200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETSkus200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETSkus200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETSkus200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.PiecesPerPack != nil {
		toSerialize["pieces_per_pack"] = o.PiecesPerPack
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.UnitOfWeight != nil {
		toSerialize["unit_of_weight"] = o.UnitOfWeight
	}
	if o.HsTariffNumber != nil {
		toSerialize["hs_tariff_number"] = o.HsTariffNumber
	}
	if o.DoNotShip != nil {
		toSerialize["do_not_ship"] = o.DoNotShip
	}
	if o.DoNotTrack != nil {
		toSerialize["do_not_track"] = o.DoNotTrack
	}
	if o.Inventory != nil {
		toSerialize["inventory"] = o.Inventory
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

type NullableGETSkus200ResponseDataInnerAttributes struct {
	value *GETSkus200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETSkus200ResponseDataInnerAttributes) Get() *GETSkus200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETSkus200ResponseDataInnerAttributes) Set(val *GETSkus200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkus200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkus200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkus200ResponseDataInnerAttributes(val *GETSkus200ResponseDataInnerAttributes) *NullableGETSkus200ResponseDataInnerAttributes {
	return &NullableGETSkus200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETSkus200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkus200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
