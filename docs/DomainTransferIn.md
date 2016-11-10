# DomainTransferIn

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | **string** | Authorization code from registrar for transferring a domain | [default to null]
**Period** | **int32** | Can be more than 1 but no more than 10 years total including current registration length | [optional] [default to null]
**RenewAuto** | **bool** | Whether or not the domain should be configured to automatically renew | [optional] [default to null]
**Privacy** | **bool** | Whether or not privacy has been requested | [optional] [default to null]
**Consent** | [**Consent**](Consent.md) | Required agreements can be retrieved via the GET ./domains/agreements endpoint | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


