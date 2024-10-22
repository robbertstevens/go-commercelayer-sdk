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

// checks if the ParcelUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParcelUpdateDataRelationships{}

// ParcelUpdateDataRelationships struct for ParcelUpdateDataRelationships
type ParcelUpdateDataRelationships struct {
	Shipment *ParcelCreateDataRelationshipsShipment `json:"shipment,omitempty"`
	Package  *ParcelCreateDataRelationshipsPackage  `json:"package,omitempty"`
}

// NewParcelUpdateDataRelationships instantiates a new ParcelUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelUpdateDataRelationships() *ParcelUpdateDataRelationships {
	this := ParcelUpdateDataRelationships{}
	return &this
}

// NewParcelUpdateDataRelationshipsWithDefaults instantiates a new ParcelUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelUpdateDataRelationshipsWithDefaults() *ParcelUpdateDataRelationships {
	this := ParcelUpdateDataRelationships{}
	return &this
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *ParcelUpdateDataRelationships) GetShipment() ParcelCreateDataRelationshipsShipment {
	if o == nil || IsNil(o.Shipment) {
		var ret ParcelCreateDataRelationshipsShipment
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelUpdateDataRelationships) GetShipmentOk() (*ParcelCreateDataRelationshipsShipment, bool) {
	if o == nil || IsNil(o.Shipment) {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *ParcelUpdateDataRelationships) HasShipment() bool {
	if o != nil && !IsNil(o.Shipment) {
		return true
	}

	return false
}

// SetShipment gets a reference to the given ParcelCreateDataRelationshipsShipment and assigns it to the Shipment field.
func (o *ParcelUpdateDataRelationships) SetShipment(v ParcelCreateDataRelationshipsShipment) {
	o.Shipment = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *ParcelUpdateDataRelationships) GetPackage() ParcelCreateDataRelationshipsPackage {
	if o == nil || IsNil(o.Package) {
		var ret ParcelCreateDataRelationshipsPackage
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelUpdateDataRelationships) GetPackageOk() (*ParcelCreateDataRelationshipsPackage, bool) {
	if o == nil || IsNil(o.Package) {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *ParcelUpdateDataRelationships) HasPackage() bool {
	if o != nil && !IsNil(o.Package) {
		return true
	}

	return false
}

// SetPackage gets a reference to the given ParcelCreateDataRelationshipsPackage and assigns it to the Package field.
func (o *ParcelUpdateDataRelationships) SetPackage(v ParcelCreateDataRelationshipsPackage) {
	o.Package = &v
}

func (o ParcelUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParcelUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Shipment) {
		toSerialize["shipment"] = o.Shipment
	}
	if !IsNil(o.Package) {
		toSerialize["package"] = o.Package
	}
	return toSerialize, nil
}

type NullableParcelUpdateDataRelationships struct {
	value *ParcelUpdateDataRelationships
	isSet bool
}

func (v NullableParcelUpdateDataRelationships) Get() *ParcelUpdateDataRelationships {
	return v.value
}

func (v *NullableParcelUpdateDataRelationships) Set(val *ParcelUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelUpdateDataRelationships(val *ParcelUpdateDataRelationships) *NullableParcelUpdateDataRelationships {
	return &NullableParcelUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableParcelUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
