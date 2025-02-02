# PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The tax calculator&#39;s internal name. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**TaxCalculatorUrl** | Pointer to **interface{}** | The URL to the service that will compute the taxes. | [optional] 
**ResetCircuit** | Pointer to **interface{}** | Send this attribute if you want to reset the circuit breaker associated to this resource to &#39;closed&#39; state and zero failures count. Cannot be passed by sales channels. | [optional] 

## Methods

### NewPATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes

`func NewPATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes() *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes`

NewPATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes instantiates a new PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributesWithDefaults

`func NewPATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributesWithDefaults() *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes`

NewPATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributesWithDefaults instantiates a new PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetTaxCalculatorUrl

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) GetTaxCalculatorUrl() interface{}`

GetTaxCalculatorUrl returns the TaxCalculatorUrl field if non-nil, zero value otherwise.

### GetTaxCalculatorUrlOk

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) GetTaxCalculatorUrlOk() (*interface{}, bool)`

GetTaxCalculatorUrlOk returns a tuple with the TaxCalculatorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculatorUrl

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) SetTaxCalculatorUrl(v interface{})`

SetTaxCalculatorUrl sets TaxCalculatorUrl field to given value.

### HasTaxCalculatorUrl

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) HasTaxCalculatorUrl() bool`

HasTaxCalculatorUrl returns a boolean if a field has been set.

### SetTaxCalculatorUrlNil

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) SetTaxCalculatorUrlNil(b bool)`

 SetTaxCalculatorUrlNil sets the value for TaxCalculatorUrl to be an explicit nil

### UnsetTaxCalculatorUrl
`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) UnsetTaxCalculatorUrl()`

UnsetTaxCalculatorUrl ensures that no value is present for TaxCalculatorUrl, not even an explicit nil
### GetResetCircuit

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) GetResetCircuit() interface{}`

GetResetCircuit returns the ResetCircuit field if non-nil, zero value otherwise.

### GetResetCircuitOk

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) GetResetCircuitOk() (*interface{}, bool)`

GetResetCircuitOk returns a tuple with the ResetCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCircuit

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) SetResetCircuit(v interface{})`

SetResetCircuit sets ResetCircuit field to given value.

### HasResetCircuit

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) HasResetCircuit() bool`

HasResetCircuit returns a boolean if a field has been set.

### SetResetCircuitNil

`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) SetResetCircuitNil(b bool)`

 SetResetCircuitNil sets the value for ResetCircuit to be an explicit nil

### UnsetResetCircuit
`func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes) UnsetResetCircuit()`

UnsetResetCircuit ensures that no value is present for ResetCircuit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


