# {{classname}}

All URIs are relative to *https://api-inmobile-pl.easypack24.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AuthenticatePost**](DefaultApi.md#V1AuthenticatePost) | **Post** /v1/authenticate | RegistrationApi.refreshToken
[**V1CollectCompartmentCloseSessionUuidGet**](DefaultApi.md#V1CollectCompartmentCloseSessionUuidGet) | **Get** /v1/collect/compartment/close/{sessionUuid} | CompartmentApi.closeCompartment
[**V1CollectCompartmentOpenSessionUuidPost**](DefaultApi.md#V1CollectCompartmentOpenSessionUuidPost) | **Post** /v1/collect/compartment/open/{sessionUuid} | CompartmentApi.openCompartment
[**V1CollectCompartmentReopenSessionUuidPost**](DefaultApi.md#V1CollectCompartmentReopenSessionUuidPost) | **Post** /v1/collect/compartment/reopen/{sessionUuid} | CompartmentApi.reopenCompartment
[**V1CollectCompartmentStatusSessionUuidGet**](DefaultApi.md#V1CollectCompartmentStatusSessionUuidGet) | **Get** /v1/collect/compartment/status/{sessionUuid} | CompartmentStatusApi.statusCompartment
[**V1CollectTerminateSessionUuidPost**](DefaultApi.md#V1CollectTerminateSessionUuidPost) | **Post** /v1/collect/terminate/{sessionUuid} | CompartmentApi.terminateCompartment
[**V1CollectValidatePost**](DefaultApi.md#V1CollectValidatePost) | **Post** /v1/collect/validate | CompartmentApi.validateCompartment
[**V1ConfirmSMSCodePhoneNumberSmsCodePost**](DefaultApi.md#V1ConfirmSMSCodePhoneNumberSmsCodePost) | **Post** /v1/confirmSMSCode/{phoneNumber}/{smsCode} | RegistrationApi.confirmSMSCode
[**V1LogoutPost**](DefaultApi.md#V1LogoutPost) | **Post** /v1/logout | UserApi.logout
[**V1NotificationsGet**](DefaultApi.md#V1NotificationsGet) | **Get** /v1/notifications | NotificationCenterApi.news
[**V1NotificationsReadAllPost**](DefaultApi.md#V1NotificationsReadAllPost) | **Post** /v1/notifications/readAll | NotificationCenterApi.markAsReadAll
[**V1NotificationsReadNotificationIdPost**](DefaultApi.md#V1NotificationsReadNotificationIdPost) | **Post** /v1/notifications/read/{notificationId} | NotificationCenterApi.markAsRead
[**V1ObservedParcelPost**](DefaultApi.md#V1ObservedParcelPost) | **Post** /v1/observedParcel | ParcelApi.subscribeParcel
[**V1ObservedParcelShipmentNumberDelete**](DefaultApi.md#V1ObservedParcelShipmentNumberDelete) | **Delete** /v1/observedParcel/{shipmentNumber} | ParcelApi.removeObservedParcel
[**V1ParcelGet**](DefaultApi.md#V1ParcelGet) | **Get** /v1/parcel | ParcelApi.parcelsWithDate
[**V1ParcelShipmentNumberGet**](DefaultApi.md#V1ParcelShipmentNumberGet) | **Get** /v1/parcel/{shipmentNumber} | ParcelApi.oneParcel
[**V1ParcelsGet**](DefaultApi.md#V1ParcelsGet) | **Get** /v1/parcels | ParcelApi.parcels
[**V1PointsGet**](DefaultApi.md#V1PointsGet) | **Get** /v1/points | MapApi.pointsFor
[**V1SendSMSCodePhoneNumberGet**](DefaultApi.md#V1SendSMSCodePhoneNumberGet) | **Get** /v1/sendSMSCode/{phoneNumber} | RegistrationApi.sendSMSCode
[**V2AgreementGet**](DefaultApi.md#V2AgreementGet) | **Get** /v2/agreement | AgreementApi.agreement
[**V2AgreementPost**](DefaultApi.md#V2AgreementPost) | **Post** /v2/agreement | AgreementApi.agreement
[**V2CollectCompartmentClaimSessionUuidPost**](DefaultApi.md#V2CollectCompartmentClaimSessionUuidPost) | **Post** /v2/collect/compartment/claim/{sessionUuid} | CompartmentApi.claimCompartment
[**V2SetPushIdPost**](DefaultApi.md#V2SetPushIdPost) | **Post** /v2/setPushId | CloudMessagingApi.setPushId

# **V1AuthenticatePost**
> RefreshTokenResponse V1AuthenticatePost(ctx, optional)
RegistrationApi.refreshToken

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiV1AuthenticatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AuthenticatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AuthenticateRequest**](AuthenticateRequest.md)|  | 

### Return type

[**RefreshTokenResponse**](RefreshTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectCompartmentCloseSessionUuidGet**
> CompartmentCloseResponse V1CollectCompartmentCloseSessionUuidGet(ctx, sessionUuid)
CompartmentApi.closeCompartment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionUuid** | **string**| CompartmentApi.closeCompartment.str | 

### Return type

[**CompartmentCloseResponse**](CompartmentCloseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectCompartmentOpenSessionUuidPost**
> CompartmentOpenResponse V1CollectCompartmentOpenSessionUuidPost(ctx, sessionUuid)
CompartmentApi.openCompartment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionUuid** | **string**| CompartmentApi.openCompartment.str | 

### Return type

[**CompartmentOpenResponse**](CompartmentOpenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectCompartmentReopenSessionUuidPost**
> CompartmentOpenResponse V1CollectCompartmentReopenSessionUuidPost(ctx, sessionUuid)
CompartmentApi.reopenCompartment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionUuid** | **string**| CompartmentApi.reopenCompartment.str | 

### Return type

[**CompartmentOpenResponse**](CompartmentOpenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectCompartmentStatusSessionUuidGet**
> CompartmentStatusResponse V1CollectCompartmentStatusSessionUuidGet(ctx, sessionUuid, optional)
CompartmentStatusApi.statusCompartment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionUuid** | **string**| CompartmentStatusApi.statusCompartment.str | 
 **optional** | ***DefaultApiV1CollectCompartmentStatusSessionUuidGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1CollectCompartmentStatusSessionUuidGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expected** | **optional.String**| CompartmentStatusApi.statusCompartment.str2 | 

### Return type

[**CompartmentStatusResponse**](CompartmentStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectTerminateSessionUuidPost**
> Completable V1CollectTerminateSessionUuidPost(ctx, sessionUuid)
CompartmentApi.terminateCompartment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionUuid** | **string**| CompartmentApi.terminateCompartment.str | 

### Return type

[**Completable**](Completable.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectValidatePost**
> CompartmentValidateResponse V1CollectValidatePost(ctx, optional)
CompartmentApi.validateCompartment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiV1CollectValidatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1CollectValidatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ValidationRequest**](ValidationRequest.md)|  | 

### Return type

[**CompartmentValidateResponse**](CompartmentValidateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ConfirmSMSCodePhoneNumberSmsCodePost**
> ConfirmSmsResponse V1ConfirmSMSCodePhoneNumberSmsCodePost(ctx, phoneNumber, smsCode, optional)
RegistrationApi.confirmSMSCode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phoneNumber** | **string**| RegistrationApi.confirmSMSCode.str | 
  **smsCode** | **string**| RegistrationApi.confirmSMSCode.str2 | 
 **optional** | ***DefaultApiV1ConfirmSMSCodePhoneNumberSmsCodePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1ConfirmSMSCodePhoneNumberSmsCodePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PhoneOsRequest**](PhoneOsRequest.md)|  | 

### Return type

[**ConfirmSmsResponse**](ConfirmSMSResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1LogoutPost**
> Completable V1LogoutPost(ctx, )
UserApi.logout

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Completable**](Completable.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1NotificationsGet**
> NotificationResponse V1NotificationsGet(ctx, optional)
NotificationCenterApi.news

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiV1NotificationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1NotificationsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **optional.String**| NotificationCenterApi.news.str | 

### Return type

[**NotificationResponse**](NotificationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1NotificationsReadAllPost**
> NotificationData V1NotificationsReadAllPost(ctx, optional)
NotificationCenterApi.markAsReadAll

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiV1NotificationsReadAllPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1NotificationsReadAllPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of time.Time**](time.Time.md)|  | 

### Return type

[**NotificationData**](NotificationData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1NotificationsReadNotificationIdPost**
> NotificationData V1NotificationsReadNotificationIdPost(ctx, notificationId)
NotificationCenterApi.markAsRead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationId** | **string**| NotificationCenterApi.markAsRead.str | 

### Return type

[**NotificationData**](NotificationData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObservedParcelPost**
> Completable V1ObservedParcelPost(ctx, optional)
ParcelApi.subscribeParcel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiV1ObservedParcelPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1ObservedParcelPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SubscribeRequest**](SubscribeRequest.md)|  | 

### Return type

[**Completable**](Completable.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObservedParcelShipmentNumberDelete**
> Completable V1ObservedParcelShipmentNumberDelete(ctx, shipmentNumber)
ParcelApi.removeObservedParcel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentNumber** | **string**| ParcelApi.removeObservedParcel.str | 

### Return type

[**Completable**](Completable.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ParcelGet**
> []Parcel V1ParcelGet(ctx, optional)
ParcelApi.parcelsWithDate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiV1ParcelGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1ParcelGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **optional.String**| ParcelApi.parcelsWithDate.str | 

### Return type

[**[]Parcel**](Parcel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ParcelShipmentNumberGet**
> Parcel V1ParcelShipmentNumberGet(ctx, shipmentNumber)
ParcelApi.oneParcel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipmentNumber** | **string**| ParcelApi.oneParcel.str | 

### Return type

[**Parcel**](Parcel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ParcelsGet**
> []Parcel V1ParcelsGet(ctx, optional)
ParcelApi.parcels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiV1ParcelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1ParcelsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipmentNumbers** | **optional.String**| ParcelApi.parcels.str | 

### Return type

[**[]Parcel**](Parcel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PointsGet**
> DeliveryPointsResponse V1PointsGet(ctx, optional)
MapApi.pointsFor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiV1PointsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1PointsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relativePoint** | **optional.String**| MapApi.pointsFor.str | 
 **maxDistance** | **optional.Float64**| MapApi.pointsFor.d | 
 **fields** | **optional.String**| MapApi.pointsFor.str2 | 
 **sortBy** | **optional.String**| MapApi.pointsFor.str3 | 
 **sortOrder** | **optional.String**| MapApi.pointsFor.str4 | 
 **perPage** | **optional.Int32**| MapApi.pointsFor.i | 

### Return type

[**DeliveryPointsResponse**](DeliveryPointsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SendSMSCodePhoneNumberGet**
> Completable V1SendSMSCodePhoneNumberGet(ctx, phoneNumber)
RegistrationApi.sendSMSCode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phoneNumber** | **string**| RegistrationApi.sendSMSCode.str | 

### Return type

[**Completable**](Completable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2AgreementGet**
> []AgreementGrant V2AgreementGet(ctx, )
AgreementApi.agreement

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AgreementGrant**](AgreementGrant.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2AgreementPost**
> Completable V2AgreementPost(ctx, optional)
AgreementApi.agreement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiV2AgreementPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV2AgreementPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AgreementGrant**](AgreementGrant.md)|  | 

### Return type

[**Completable**](Completable.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2CollectCompartmentClaimSessionUuidPost**
> CompartmentClaimResponse V2CollectCompartmentClaimSessionUuidPost(ctx, sessionUuid, optional)
CompartmentApi.claimCompartment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionUuid** | **string**| CompartmentApi.claimCompartment.str | 
 **optional** | ***DefaultApiV2CollectCompartmentClaimSessionUuidPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV2CollectCompartmentClaimSessionUuidPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CompartmentClaimRequest**](CompartmentClaimRequest.md)|  | 

### Return type

[**CompartmentClaimResponse**](CompartmentClaimResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2SetPushIdPost**
> Completable V2SetPushIdPost(ctx, optional)
CloudMessagingApi.setPushId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiV2SetPushIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV2SetPushIdPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PushIdBody**](PushIdBody.md)|  | 

### Return type

[**Completable**](Completable.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

