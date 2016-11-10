# \VdomainsApi

All URIs are relative to *https://api.ote-godaddy.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Available**](VdomainsApi.md#Available) | **Get** /v1/domains/available | Determine whether or not the specified domain is available for purchase
[**AvailableBulk**](VdomainsApi.md#AvailableBulk) | **Post** /v1/domains/available | Determine whether or not the specified domains are available for purchase
[**Cancel**](VdomainsApi.md#Cancel) | **Delete** /v1/domains/{domain} | Cancel a purchased domain
[**CancelPrivacy**](VdomainsApi.md#CancelPrivacy) | **Delete** /v1/domains/{domain}/privacy | Submit a privacy cancellation request for the given domain
[**Get**](VdomainsApi.md#Get) | **Get** /v1/domains/{domain} | Retrieve details for the specified Domain
[**GetAgreement**](VdomainsApi.md#GetAgreement) | **Get** /v1/domains/agreements | Retrieve the legal agreement(s) required to purchase the specified TLD and add-ons
[**List**](VdomainsApi.md#List) | **Get** /v1/domains | Retrieve a list of Domains for the specified Shopper
[**Purchase**](VdomainsApi.md#Purchase) | **Post** /v1/domains/purchase | Purchase and register the specified Domain
[**PurchasePrivacy**](VdomainsApi.md#PurchasePrivacy) | **Post** /v1/domains/{domain}/privacy/purchase | Purchase privacy for a specified domain
[**RecordAdd**](VdomainsApi.md#RecordAdd) | **Patch** /v1/domains/{domain}/records | Add the specified DNS Records to the specified Domain
[**RecordGet**](VdomainsApi.md#RecordGet) | **Get** /v1/domains/{domain}/records/{type?}/{name?} | Retrieve DNS Records for the specified Domain, optionally with the specified Type and/or Name
[**RecordReplace**](VdomainsApi.md#RecordReplace) | **Put** /v1/domains/{domain}/records | Replace all DNS Records for the specified Domain
[**RecordReplaceType**](VdomainsApi.md#RecordReplaceType) | **Put** /v1/domains/{domain}/records/{type} | Replace all DNS Records for the specified Domain with the specified Type
[**RecordReplaceTypeName**](VdomainsApi.md#RecordReplaceTypeName) | **Put** /v1/domains/{domain}/records/{type}/{name} | Replace all DNS Records for the specified Domain with the specified Type and Name
[**Renew**](VdomainsApi.md#Renew) | **Post** /v1/domains/{domain}/renew | Renew the specified Domain
[**Schema**](VdomainsApi.md#Schema) | **Get** /v1/domains/purchase/schema/{tld} | Retrieve the schema to be submitted when registering a Domain for the specified TLD
[**Suggest**](VdomainsApi.md#Suggest) | **Get** /v1/domains/suggest | Suggest alternate Domain names based on a seed Domain, a set of keywords, or the shopper&#39;s purchase history
[**Tlds**](VdomainsApi.md#Tlds) | **Get** /v1/domains/tlds | Retrieves a list of TLDs supported and enabled for sale
[**TransferIn**](VdomainsApi.md#TransferIn) | **Post** /v1/domains/{domain}/transfer | Purchase and start or restart transfer process
[**Update**](VdomainsApi.md#Update) | **Patch** /v1/domains/{domain} | Update details for the specified Domain
[**UpdateContacts**](VdomainsApi.md#UpdateContacts) | **Patch** /v1/domains/{domain}/contacts | Update domain
[**Validate**](VdomainsApi.md#Validate) | **Post** /v1/domains/purchase/validate | Validate the request body using the Domain Purchase Schema for the specified TLD
[**VerifyEmail**](VdomainsApi.md#VerifyEmail) | **Post** /v1/domains/{domain}/verifyRegistrantEmail | Re-send Contact E-mail Verification for specified Domain


# **Available**
> DomainAvailableResponse Available($domain, $checkType, $forTransfer)

Determine whether or not the specified domain is available for purchase


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain name whose availability is to be checked | 
 **checkType** | **string**| Optimize for time (&#39;FAST&#39;) or accuracy (&#39;FULL&#39;) | [optional] [default to FAST]
 **forTransfer** | **bool**| Whether or not to include domains available for transfer. If set to True, checkType is ignored | [optional] [default to false]

### Return type

[**DomainAvailableResponse**](DomainAvailableResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AvailableBulk**
> DomainAvailableBulk AvailableBulk($domains, $checkType)

Determine whether or not the specified domains are available for purchase


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domains** | **[]string**| Domain names for which to check availability | 
 **checkType** | **string**| Optimize for time (&#39;FAST&#39;) or accuracy (&#39;FULL&#39;) | [optional] [default to FAST]

### Return type

[**DomainAvailableBulk**](DomainAvailableBulk.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Cancel**
> Cancel($domain)

Cancel a purchased domain


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain to cancel | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelPrivacy**
> CancelPrivacy($domain, $xShopperId)

Submit a privacy cancellation request for the given domain


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain whose privacy is to be cancelled | 
 **xShopperId** | **string**| Shopper ID of the owner of the domain | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> DomainDetail Get($domain, $xShopperId)

Retrieve details for the specified Domain


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain name whose details are to be retrieved | 
 **xShopperId** | **string**| Shopper ID expected to own the specified domain | [optional] 

### Return type

[**DomainDetail**](DomainDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgreement**
> []LegalAgreement GetAgreement($tlds, $privacy, $xMarketId, $forTransfer)

Retrieve the legal agreement(s) required to purchase the specified TLD and add-ons


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlds** | [**[]string**](string.md)| list of TLDs whose legal agreements are to be retrieved | 
 **privacy** | **bool**| Whether or not privacy has been requested | 
 **xMarketId** | **string**| Unique identifier of the Market used to retrieve/translate Legal Agreements | [optional] [default to en-US]
 **forTransfer** | **bool**| Whether or not domain tranfer has been requested | [optional] 

### Return type

[**[]LegalAgreement**](LegalAgreement.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []DomainSummary List($xShopperId, $statuses, $statusGroups, $limit, $marker, $includes, $modifiedDate)

Retrieve a list of Domains for the specified Shopper


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xShopperId** | **string**| Shopper ID whose domains are to be retrieved | [optional] 
 **statuses** | [**[]string**](string.md)| Only include results with &#x60;status&#x60; value in the specified set | [optional] 
 **statusGroups** | [**[]string**](string.md)| Only include results with &#x60;status&#x60; value in any of the specified groups | [optional] 
 **limit** | **int32**| Maximum number of domains to return | [optional] 
 **marker** | **string**| Marker Domain to use as the offset in results | [optional] 
 **includes** | [**[]string**](string.md)| Optional details to be included in the response | [optional] 
 **modifiedDate** | **string**| Only include results that have been modified since the specified date | [optional] 

### Return type

[**[]DomainSummary**](DomainSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Purchase**
> DomainPurchaseResponse Purchase($body, $xShopperId)

Purchase and register the specified Domain


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DomainPurchase**](DomainPurchase.md)| An instance document expected to match the JSON schema returned by &#x60;./schema/{tld}&#x60; | 
 **xShopperId** | **string**| The Shopper for whom the domain should be purchased | [optional] 

### Return type

[**DomainPurchaseResponse**](DomainPurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PurchasePrivacy**
> DomainPurchaseResponse PurchasePrivacy($domain, $body, $xShopperId)

Purchase privacy for a specified domain


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain for which to purchase privacy | 
 **body** | [**PrivacyPurchase**](PrivacyPurchase.md)| Options for purchasing privacy | 
 **xShopperId** | **string**| Shopper ID of the owner of the domain | [optional] 

### Return type

[**DomainPurchaseResponse**](DomainPurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordAdd**
> RecordAdd($domain, $records, $xShopperId)

Add the specified DNS Records to the specified Domain


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain whose DNS Records are to be augmented | 
 **records** | [**DnsRecord**](DnsRecord.md)| DNS Records to add to whatever currently exists | 
 **xShopperId** | **string**| Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordGet**
> []DnsRecord RecordGet($domain, $type_, $name, $xShopperId, $offset, $limit)

Retrieve DNS Records for the specified Domain, optionally with the specified Type and/or Name


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain whose DNS Records are to be retrieved | 
 **type_** | **string**| DNS Record Type for which DNS Records are to be retrieved | 
 **name** | **string**| DNS Record Name for which DNS Records are to be retrieved | 
 **xShopperId** | **string**| Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | [optional] 
 **offset** | **int32**| Number of results to skip for pagination | [optional] 
 **limit** | **int32**| Maximum number of items to return | [optional] 

### Return type

[**[]DnsRecord**](DNSRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordReplace**
> RecordReplace($domain, $records, $xShopperId)

Replace all DNS Records for the specified Domain


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain whose DNS Records are to be replaced | 
 **records** | [**[]DnsRecord**](DNSRecord.md)| DNS Records to replace whatever currently exists | 
 **xShopperId** | **string**| Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordReplaceType**
> RecordReplaceType($domain, $type_, $records, $xShopperId)

Replace all DNS Records for the specified Domain with the specified Type


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain whose DNS Records are to be replaced | 
 **type_** | **string**| DNS Record Type for which DNS Records are to be replaced | 
 **records** | [**[]DnsRecordCreateType**](DNSRecordCreateType.md)| DNS Records to replace whatever currently exists | 
 **xShopperId** | **string**| Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordReplaceTypeName**
> RecordReplaceTypeName($domain, $type_, $name, $records, $xShopperId)

Replace all DNS Records for the specified Domain with the specified Type and Name


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain whose DNS Records are to be replaced | 
 **type_** | **string**| DNS Record Type for which DNS Records are to be replaced | 
 **name** | **string**| DNS Record Name for which DNS Records are to be replaced | 
 **records** | [**[]DnsRecordCreateTypeName**](DNSRecordCreateTypeName.md)| DNS Records to replace whatever currently exists | 
 **xShopperId** | **string**| Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Renew**
> DomainPurchaseResponse Renew($domain, $xShopperId, $body)

Renew the specified Domain


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain to renew | 
 **xShopperId** | **string**| Shopper for whom Domain is to be renewed. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | [optional] 
 **body** | [**DomainRenew**](DomainRenew.md)| Options for renewing existing Domain | [optional] 

### Return type

[**DomainPurchaseResponse**](DomainPurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Schema**
> JsonSchema Schema($tld)

Retrieve the schema to be submitted when registering a Domain for the specified TLD




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tld** | **string**| The Top-Level Domain whose schema should be retrieved | 

### Return type

[**JsonSchema**](JsonSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Suggest**
> []DomainSuggestion Suggest($xShopperId, $query, $country, $city, $sources, $tlds, $lengthMax, $lengthMin, $limit, $waitMs)

Suggest alternate Domain names based on a seed Domain, a set of keywords, or the shopper's purchase history


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xShopperId** | **string**| Shopper ID for which the suggestions are being generated | [optional] 
 **query** | **string**| Domain name or set of keywords for which alternative domain names will be suggested | [optional] 
 **country** | **string**| Two-letter ISO country code to be used as a hint for target region&lt;br/&gt;&lt;br/&gt; NOTE: These are sample values, there are many &lt;a href&#x3D;\&quot;http://www.iso.org/iso/country_codes.htm\&quot;&gt;more&lt;/a&gt; | [optional] 
 **city** | **string**| Name of city to be used as a hint for target region | [optional] 
 **sources** | [**[]string**](string.md)| Sources to be queried&lt;br/&gt;&lt;br/&gt;&lt;ul&gt; &lt;li&gt;&lt;strong&gt;CC_TLD&lt;/strong&gt; - Varies the TLD using Country Codes&lt;/li&gt; &lt;li&gt;&lt;strong&gt;EXTENSION&lt;/strong&gt; - Varies the TLD&lt;/li&gt; &lt;li&gt;&lt;strong&gt;KEYWORD_SPIN&lt;/strong&gt; - Identifies keywords and then rotates each one&lt;/li&gt; &lt;li&gt;&lt;strong&gt;PREMIUM&lt;/strong&gt; - Includes variations with premium prices&lt;/li&gt;&lt;/ul&gt; | [optional] 
 **tlds** | [**[]string**](string.md)| Top-level domains to be included in suggestions&lt;br/&gt;&lt;br/&gt; NOTE: These are sample values, there are many &lt;a href&#x3D;\&quot;http://www.godaddy.com/tlds/gtld.aspx#domain_search_form\&quot;&gt;more&lt;/a&gt; | [optional] 
 **lengthMax** | **int32**| Maximum length of second-level domain | [optional] 
 **lengthMin** | **int32**| Minimum length of second-level domain | [optional] 
 **limit** | **int32**| Maximum number of suggestions to return | [optional] 
 **waitMs** | **int32**| Maximum amount of time, in milliseconds, to wait for responses If elapses, return the results compiled up to that point | [optional] [default to 1000]

### Return type

[**[]DomainSuggestion**](DomainSuggestion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Tlds**
> []TldSummary Tlds()

Retrieves a list of TLDs supported and enabled for sale


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]TldSummary**](TldSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferIn**
> DomainPurchaseResponse TransferIn($domain, $body, $xShopperId)

Purchase and start or restart transfer process


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain to transfer in | 
 **body** | [**DomainTransferIn**](DomainTransferIn.md)| Details for domain transfer purchase | 
 **xShopperId** | **string**| The Shopper to whom the domain should be transfered | [optional] 

### Return type

[**DomainPurchaseResponse**](DomainPurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> Update($domain, $body, $xShopperId)

Update details for the specified Domain


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain whose details are to be updated | 
 **body** | [**DomainUpdate**](DomainUpdate.md)| Changes to apply to existing Domain | 
 **xShopperId** | **string**| Shopper for whom Domain is to be updated. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContacts**
> UpdateContacts($domain, $contacts, $xShopperId)

Update domain


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain whose Contacts are to be updated. | 
 **contacts** | [**DomainContacts**](DomainContacts.md)| Changes to apply to existing Contacts | 
 **xShopperId** | **string**| Shopper for whom domain contacts are to be updated. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Validate**
> Validate($body)

Validate the request body using the Domain Purchase Schema for the specified TLD




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DomainPurchase**](DomainPurchase.md)| An instance document expected to match the JSON schema returned by &#x60;./schema/{tld}&#x60; | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyEmail**
> VerifyEmail($domain, $xShopperId)

Re-send Contact E-mail Verification for specified Domain


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Domain whose Contact E-mail should be verified. | 
 **xShopperId** | **string**| Shopper for whom domain contact e-mail should be verified. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

