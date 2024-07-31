# ExportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**LinksResponse**](LinksResponse.md) |  | 
**AccountingEntryId** | **string** | This is the Pleo internal identifier of the export item | 
**AccountingEntryLines** | [**[]ExportItemLine**](ExportItemLine.md) | Accounting entry broken down in entry lines. There will always be an accounting entry line. If the accounting entry has been split into separate lines, then each line would represent the details of that accounting entry line, else there would only be one line present, representing the whole entry. | 
**Amount** | [**ExportItemAmount**](ExportItemAmount.md) |  | 
**CompanyId** | **string** | Pleo company identifier this export item belongs to | 
**Date** | **time.Time** | Date the accounting entry should be bookkept | 
**Files** | Pointer to [**[]ExportItemFile**](ExportItemFile.md) | Files that have been attached to this accounting entry | [optional] 
**Metadata** | [**ExportItemMetadata**](ExportItemMetadata.md) |  | 
**Note** | Pointer to **NullableString** | Additional comments potentially describing the accounting entry | [optional] 
**Supplier** | Pointer to [**NullableSupplier**](Supplier.md) |  | [optional] 
**TeamCode** | Pointer to **NullableString** | Team code is an identifier assigned to an expense to categorize it or associate it with a specific team | [optional] 
**Type** | **string** | This is the Pleo internal identifier of the accounting entry this export item is associated with | 
**User** | [**ExportItemUser**](ExportItemUser.md) |  | 

## Methods

### NewExportItem

`func NewExportItem(links LinksResponse, accountingEntryId string, accountingEntryLines []ExportItemLine, amount ExportItemAmount, companyId string, date time.Time, metadata ExportItemMetadata, type_ string, user ExportItemUser, ) *ExportItem`

NewExportItem instantiates a new ExportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportItemWithDefaults

`func NewExportItemWithDefaults() *ExportItem`

NewExportItemWithDefaults instantiates a new ExportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ExportItem) GetLinks() LinksResponse`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExportItem) GetLinksOk() (*LinksResponse, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExportItem) SetLinks(v LinksResponse)`

SetLinks sets Links field to given value.


### GetAccountingEntryId

`func (o *ExportItem) GetAccountingEntryId() string`

GetAccountingEntryId returns the AccountingEntryId field if non-nil, zero value otherwise.

### GetAccountingEntryIdOk

`func (o *ExportItem) GetAccountingEntryIdOk() (*string, bool)`

GetAccountingEntryIdOk returns a tuple with the AccountingEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryId

`func (o *ExportItem) SetAccountingEntryId(v string)`

SetAccountingEntryId sets AccountingEntryId field to given value.


### GetAccountingEntryLines

`func (o *ExportItem) GetAccountingEntryLines() []ExportItemLine`

GetAccountingEntryLines returns the AccountingEntryLines field if non-nil, zero value otherwise.

### GetAccountingEntryLinesOk

`func (o *ExportItem) GetAccountingEntryLinesOk() (*[]ExportItemLine, bool)`

GetAccountingEntryLinesOk returns a tuple with the AccountingEntryLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryLines

`func (o *ExportItem) SetAccountingEntryLines(v []ExportItemLine)`

SetAccountingEntryLines sets AccountingEntryLines field to given value.


### GetAmount

`func (o *ExportItem) GetAmount() ExportItemAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExportItem) GetAmountOk() (*ExportItemAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExportItem) SetAmount(v ExportItemAmount)`

SetAmount sets Amount field to given value.


### GetCompanyId

`func (o *ExportItem) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ExportItem) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ExportItem) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetDate

`func (o *ExportItem) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ExportItem) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ExportItem) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetFiles

`func (o *ExportItem) GetFiles() []ExportItemFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ExportItem) GetFilesOk() (*[]ExportItemFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ExportItem) SetFiles(v []ExportItemFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ExportItem) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *ExportItem) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *ExportItem) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetMetadata

`func (o *ExportItem) GetMetadata() ExportItemMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExportItem) GetMetadataOk() (*ExportItemMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExportItem) SetMetadata(v ExportItemMetadata)`

SetMetadata sets Metadata field to given value.


### GetNote

`func (o *ExportItem) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ExportItem) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ExportItem) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ExportItem) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *ExportItem) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *ExportItem) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetSupplier

`func (o *ExportItem) GetSupplier() Supplier`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *ExportItem) GetSupplierOk() (*Supplier, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *ExportItem) SetSupplier(v Supplier)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *ExportItem) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### SetSupplierNil

`func (o *ExportItem) SetSupplierNil(b bool)`

 SetSupplierNil sets the value for Supplier to be an explicit nil

### UnsetSupplier
`func (o *ExportItem) UnsetSupplier()`

UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
### GetTeamCode

`func (o *ExportItem) GetTeamCode() string`

GetTeamCode returns the TeamCode field if non-nil, zero value otherwise.

### GetTeamCodeOk

`func (o *ExportItem) GetTeamCodeOk() (*string, bool)`

GetTeamCodeOk returns a tuple with the TeamCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamCode

`func (o *ExportItem) SetTeamCode(v string)`

SetTeamCode sets TeamCode field to given value.

### HasTeamCode

`func (o *ExportItem) HasTeamCode() bool`

HasTeamCode returns a boolean if a field has been set.

### SetTeamCodeNil

`func (o *ExportItem) SetTeamCodeNil(b bool)`

 SetTeamCodeNil sets the value for TeamCode to be an explicit nil

### UnsetTeamCode
`func (o *ExportItem) UnsetTeamCode()`

UnsetTeamCode ensures that no value is present for TeamCode, not even an explicit nil
### GetType

`func (o *ExportItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportItem) SetType(v string)`

SetType sets Type field to given value.


### GetUser

`func (o *ExportItem) GetUser() ExportItemUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExportItem) GetUserOk() (*ExportItemUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExportItem) SetUser(v ExportItemUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


