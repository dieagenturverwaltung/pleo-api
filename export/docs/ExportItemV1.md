# ExportItemV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**LinksResponse**](LinksResponse.md) |  | 
**AccountingEntryId** | **string** | This is the Pleo internal identifier of the export item | 
**AccountingEntryLines** | [**[]ExportItemLineV1**](ExportItemLineV1.md) | Accounting entry broken down in entry lines. There will always be an accounting entry line. If the accounting entry has been split into separate lines, then each line would represent the details of that accounting entry line, else there would only be one line present, representing the whole entry. | 
**AdditionalInformation** | [**ExportItemAdditionalInformationV1**](ExportItemAdditionalInformationV1.md) |  | 
**Amount** | [**ExportItemAmount**](ExportItemAmount.md) |  | 
**CompanyId** | **string** | Pleo company identifier this export item belongs to | 
**Date** | **time.Time** | Date the accounting entry should be bookkept | 
**Files** | Pointer to [**[]ExportItemFile**](ExportItemFile.md) | Files that have been attached to this accounting entry | [optional] 
**Note** | Pointer to **NullableString** | Additional comments potentially describing the accounting entry | [optional] 
**Supplier** | Pointer to [**NullableSupplierV1**](SupplierV1.md) |  | [optional] 
**Team** | Pointer to [**NullableTeam**](Team.md) |  | [optional] 
**TeamCode** | Pointer to **NullableString** | Team code is an identifier assigned to an expense to categorize it or associate it with a specific team | [optional] 
**Type** | **string** | This is the Pleo internal identifier of the accounting entry this export item is associated with | 
**User** | [**ExportItemUserV1**](ExportItemUserV1.md) |  | 

## Methods

### NewExportItemV1

`func NewExportItemV1(links LinksResponse, accountingEntryId string, accountingEntryLines []ExportItemLineV1, additionalInformation ExportItemAdditionalInformationV1, amount ExportItemAmount, companyId string, date time.Time, type_ string, user ExportItemUserV1, ) *ExportItemV1`

NewExportItemV1 instantiates a new ExportItemV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportItemV1WithDefaults

`func NewExportItemV1WithDefaults() *ExportItemV1`

NewExportItemV1WithDefaults instantiates a new ExportItemV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ExportItemV1) GetLinks() LinksResponse`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExportItemV1) GetLinksOk() (*LinksResponse, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExportItemV1) SetLinks(v LinksResponse)`

SetLinks sets Links field to given value.


### GetAccountingEntryId

`func (o *ExportItemV1) GetAccountingEntryId() string`

GetAccountingEntryId returns the AccountingEntryId field if non-nil, zero value otherwise.

### GetAccountingEntryIdOk

`func (o *ExportItemV1) GetAccountingEntryIdOk() (*string, bool)`

GetAccountingEntryIdOk returns a tuple with the AccountingEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryId

`func (o *ExportItemV1) SetAccountingEntryId(v string)`

SetAccountingEntryId sets AccountingEntryId field to given value.


### GetAccountingEntryLines

`func (o *ExportItemV1) GetAccountingEntryLines() []ExportItemLineV1`

GetAccountingEntryLines returns the AccountingEntryLines field if non-nil, zero value otherwise.

### GetAccountingEntryLinesOk

`func (o *ExportItemV1) GetAccountingEntryLinesOk() (*[]ExportItemLineV1, bool)`

GetAccountingEntryLinesOk returns a tuple with the AccountingEntryLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryLines

`func (o *ExportItemV1) SetAccountingEntryLines(v []ExportItemLineV1)`

SetAccountingEntryLines sets AccountingEntryLines field to given value.


### GetAdditionalInformation

`func (o *ExportItemV1) GetAdditionalInformation() ExportItemAdditionalInformationV1`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *ExportItemV1) GetAdditionalInformationOk() (*ExportItemAdditionalInformationV1, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *ExportItemV1) SetAdditionalInformation(v ExportItemAdditionalInformationV1)`

SetAdditionalInformation sets AdditionalInformation field to given value.


### GetAmount

`func (o *ExportItemV1) GetAmount() ExportItemAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExportItemV1) GetAmountOk() (*ExportItemAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExportItemV1) SetAmount(v ExportItemAmount)`

SetAmount sets Amount field to given value.


### GetCompanyId

`func (o *ExportItemV1) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ExportItemV1) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ExportItemV1) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetDate

`func (o *ExportItemV1) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ExportItemV1) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ExportItemV1) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetFiles

`func (o *ExportItemV1) GetFiles() []ExportItemFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ExportItemV1) GetFilesOk() (*[]ExportItemFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ExportItemV1) SetFiles(v []ExportItemFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ExportItemV1) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *ExportItemV1) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *ExportItemV1) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetNote

`func (o *ExportItemV1) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ExportItemV1) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ExportItemV1) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ExportItemV1) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *ExportItemV1) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *ExportItemV1) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetSupplier

`func (o *ExportItemV1) GetSupplier() SupplierV1`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *ExportItemV1) GetSupplierOk() (*SupplierV1, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *ExportItemV1) SetSupplier(v SupplierV1)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *ExportItemV1) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### SetSupplierNil

`func (o *ExportItemV1) SetSupplierNil(b bool)`

 SetSupplierNil sets the value for Supplier to be an explicit nil

### UnsetSupplier
`func (o *ExportItemV1) UnsetSupplier()`

UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
### GetTeam

`func (o *ExportItemV1) GetTeam() Team`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ExportItemV1) GetTeamOk() (*Team, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ExportItemV1) SetTeam(v Team)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ExportItemV1) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *ExportItemV1) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *ExportItemV1) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetTeamCode

`func (o *ExportItemV1) GetTeamCode() string`

GetTeamCode returns the TeamCode field if non-nil, zero value otherwise.

### GetTeamCodeOk

`func (o *ExportItemV1) GetTeamCodeOk() (*string, bool)`

GetTeamCodeOk returns a tuple with the TeamCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamCode

`func (o *ExportItemV1) SetTeamCode(v string)`

SetTeamCode sets TeamCode field to given value.

### HasTeamCode

`func (o *ExportItemV1) HasTeamCode() bool`

HasTeamCode returns a boolean if a field has been set.

### SetTeamCodeNil

`func (o *ExportItemV1) SetTeamCodeNil(b bool)`

 SetTeamCodeNil sets the value for TeamCode to be an explicit nil

### UnsetTeamCode
`func (o *ExportItemV1) UnsetTeamCode()`

UnsetTeamCode ensures that no value is present for TeamCode, not even an explicit nil
### GetType

`func (o *ExportItemV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportItemV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportItemV1) SetType(v string)`

SetType sets Type field to given value.


### GetUser

`func (o *ExportItemV1) GetUser() ExportItemUserV1`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExportItemV1) GetUserOk() (*ExportItemUserV1, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExportItemV1) SetUser(v ExportItemUserV1)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


