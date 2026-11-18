package accountingentries_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/dieagenturverwaltung/pleo-api/accountingentries"
	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func newTestClient(handler http.HandlerFunc) *accountingentries.Client {
	transport := roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if err := r.Context().Err(); err != nil {
			return nil, err
		}
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		return w.Result(), nil
	})
	return accountingentries.New(&shared.Config{
		BaseURL:    "https://example.test",
		HttpClient: &http.Client{Transport: transport},
	})
}

func checkRequest(t *testing.T, r *http.Request, method, path string, query url.Values, body string) {
	t.Helper()
	if r.Method != method || r.URL.EscapedPath() != path {
		t.Errorf("request = %s %s, want %s %s", r.Method, r.URL.EscapedPath(), method, path)
	}
	if got := r.URL.Query(); !reflect.DeepEqual(got, query) {
		t.Errorf("query = %v, want %v", got, query)
	}
	if body == "" {
		if r.ContentLength != 0 {
			t.Errorf("unexpected request body of length %d", r.ContentLength)
		}
		return
	}
	var got, want map[string]any
	if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
		t.Errorf("decode request: %v", err)
		return
	}
	if err := json.Unmarshal([]byte(body), &want); err != nil {
		t.Errorf("decode expected body: %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("body = %#v, want %#v", got, want)
	}
}

func TestGetAccountingEntry(t *testing.T) {
	client := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		checkRequest(t, r, http.MethodGet, "/v1/accounting-entries/entry%2Fwith%20%3F%23", url.Values{
			"include_deleted":                  {"false"},
			"include_per_diem_related_entries": {"true"},
		}, "")
		w.Header().Set("ratelimit-remaining", "599")
		_, _ = w.Write([]byte(`{"data":{
			"id":"entry-id","companyId":"company-id","family":"CARD_PURCHASE",
			"subFamily":"ONLINE","status":"COMPLETED","exportStatus":"EXPORTED","reviewStatus":"OK",
			"createdAt":"2026-11-01T12:00:00Z","updatedAt":"2026-11-02T12:00:00Z",
			"performedAt":"2026-11-01T12:00:00Z","settledAt":"2026-11-02T12:00:00Z",
			"bookkeepingDate":"2026-11-02T00:00:00Z","employeeId":"employee-id","teamId":"team-id",
			"foreignExtensionId":"external-id","accountId":"account-id","accountCode":"4000",
			"taxCodeId":"tax-id","note":"Travel","receiptIds":["receipt-id"],
			"transactionValue":{"currency":"EUR","minors":-9007199254740993},
			"totalBillValue":{"currency":"EUR","minors":1234},
			"attendees":[{"id":"attendee-id","attendableId":"contact-id","attendableType":"CONTACT"}],
			"tags":[{"tagGroupId":"group-id","tagId":"tag-id"}],
			"splitItems":[{"id":"split-id","accountId":"split-account-id","tags":[],
				"transactionValue":{"currency":"EUR","minors":1234},
				"totalBillValue":{"currency":"EUR","minors":1234}}],
			"merchant":{"name":"Merchant","address":{"country":"DE"}},
			"supplier":{"cif":"supplier-cif","documentNumber":"invoice-1","documentType":"FACTURA"},
			"relatedAccountingEntries":[{"id":"related-id"}]
		}}`))
	})
	entry, extra, err := client.GetAccountingEntry(context.Background(), "entry/with ?#").
		WithIncludeDeleted(false).WithIncludePerDiemRelatedEntries(true).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if entry.ID != "entry-id" || entry.Family != accountingentries.FamilyCardPurchase ||
		entry.SubFamily == nil || *entry.SubFamily != accountingentries.SubFamilyOnline ||
		entry.Status != accountingentries.StatusCompleted || entry.ExportStatus != accountingentries.ExportStatusExported ||
		entry.ReviewStatus != accountingentries.ReviewStatusOK {
		t.Fatalf("unexpected entry: %+v", entry)
	}
	if entry.TransactionValue.Minors != -9007199254740993 || entry.TotalBillValue == nil || entry.TotalBillValue.Minors != 1234 {
		t.Fatalf("unexpected amounts: %+v, %+v", entry.TransactionValue, entry.TotalBillValue)
	}
	if len(entry.SplitItems) != 1 || entry.SplitItems[0].TransactionValue == nil || entry.SplitItems[0].TransactionValue.Minors != 1234 ||
		len(entry.Tags) != 1 || entry.Tags[0].TagID != "tag-id" || len(entry.Attendees) != 1 ||
		entry.Merchant == nil || entry.Merchant.Address["country"] != "DE" ||
		entry.Supplier == nil || entry.Supplier.DocumentNumber != "invoice-1" || len(entry.RelatedAccountingEntries) != 1 {
		t.Fatalf("unexpected nested data: %+v", entry)
	}
	if extra == nil || extra.RateLimit == nil || extra.RateLimit.Remaining == nil || *extra.RateLimit.Remaining != 599 {
		t.Fatalf("unexpected rate limit: %+v", extra)
	}
}

func TestSearchAccountingEntries(t *testing.T) {
	client := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		checkRequest(t, r, http.MethodPost, "/v1/accounting-entries:search", url.Values{
			"company_id": {"company-id"}, "organization_id": {"organization-id"},
			"after": {"cursor+/?="}, "limit": {"50"},
			"sorting_keys": {"performedAt", "id"}, "sorting_order": {"DESC", "ASC"},
		}, `{
			"includeDeleted":true,"accountIds":["account-1","account-2"],"employeeIds":["employee-id"],
			"exportStatus":["CREATED","QUEUED"],"families":["CARD_PURCHASE"],"foreignExtensionIds":["foreign-id"],
			"hasAccount":false,"hasBookkeepingDate":false,"hasNote":false,"settled":false,
			"performedAtStart":"2026-11-01T00:00:00Z","performedAtEnd":"2026-11-02T00:00:00Z",
			"reviewStatus":["OK"],"status":["COMPLETED"],"subFamilies":["ONLINE"],
			"tagGroupIdsUnset":["group-id"],"tagRowIds":["tag-id"],"taxCodeIds":["tax-id"],
			"teamIds":["team-id"],"transactionValueCurrencies":["EUR","GBP"]
		}`)
		_, _ = w.Write([]byte(`{"data":[
			{"id":"entry-1","subFamily":null,"transactionValue":{"currency":"EUR","minors":1}},
			{"id":"entry-2","subFamily":"FUTURE_SUB_FAMILY"}
		],"pagination":{"hasNextPage":true,"endCursor":"next-cursor","total":2,
			"currentRequestPagination":{"parameters":{"company_id":["company-id"]},"limit":50}}}`))
	})
	limit, after := 50, "cursor+/?="
	page, _, err := client.SearchAccountingEntries(context.Background()).
		WithBody(accountingentries.SearchAccountingEntriesRequest{IncludeDeleted: true}).
		WithCompanyID("company-id").WithOrganizationID("organization-id").
		WithPagingInfo(shared.PagingInfo{After: &after, Limit: &limit, SortingKeys: []string{"performedAt", "id"}, SortingOrder: []string{"DESC", "ASC"}}).
		WithAccountIDs("account-1").WithAccountIDs("account-2").WithEmployeeIDs("employee-id").
		WithExportStatus(accountingentries.ExportStatusCreated, accountingentries.ExportStatusQueued).
		WithFamilies(accountingentries.FamilyCardPurchase).WithForeignExtensionIDs("foreign-id").
		WithHasAccount(false).WithHasBookkeepingDate(false).WithHasNote(false).WithSettled(false).
		WithPerformedAtStart("2026-11-01T00:00:00Z").WithPerformedAtEnd("2026-11-02T00:00:00Z").
		WithReviewStatus(accountingentries.ReviewStatusOK).WithStatus(accountingentries.StatusCompleted).
		WithSubFamilies(accountingentries.SubFamilyOnline).WithTagGroupIDsUnset("group-id").
		WithTagRowIDs("tag-id").WithTaxCodeIDs("tax-id").WithTeamIDs("team-id").
		WithTransactionValueCurrencies("EUR", "GBP").Execute()
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Data) != 2 || page.Data[0].SubFamily != nil || page.Data[1].SubFamily == nil ||
		*page.Data[1].SubFamily != "FUTURE_SUB_FAMILY" || !page.Pagination.HasNextPage ||
		page.Pagination.EndCursor != "next-cursor" || page.Pagination.Total != 2 ||
		page.Pagination.CurrentRequestPagination.Limit != 50 {
		t.Fatalf("unexpected search page: %+v", page)
	}
}

const receiptJSON = `{"id":"receipt-id","accountingEntryId":"entry-id","createdAt":"2026-11-01T12:00:00Z",
	"fileType":"PREVIEW","mimeType":"application/pdf","sizeInBytes":123456,
	"source":"USER","url":"https://example.test/receipt?signature=test","ocrDocumentId":"ocr-id"}`

func checkReceipt(t *testing.T, receipt accountingentries.Receipt) {
	t.Helper()
	if receipt.ID != "receipt-id" || receipt.AccountingEntryID != "entry-id" ||
		receipt.FileType != accountingentries.ReceiptFileTypePreview || receipt.MIMEType != accountingentries.ReceiptMIMETypePDF ||
		receipt.Source != accountingentries.ReceiptSourceUser || receipt.SizeInBytes != 123456 ||
		receipt.URL != "https://example.test/receipt?signature=test" || receipt.OCRDocumentID != "ocr-id" {
		t.Errorf("unexpected receipt: %+v", receipt)
	}
}

func TestGetReceipts(t *testing.T) {
	client := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		checkRequest(t, r, http.MethodGet, "/v1/accounting-entries/entry%2Fid/receipts", url.Values{
			"source": {"USER", "CAMERA"}, "file_type": {"PREVIEW"},
			"before": {"previous-cursor"}, "limit": {"10"}, "sorting_keys": {"createdAt"}, "sorting_order": {"ASC"},
		}, "")
		_, _ = w.Write([]byte(`{"data":[` + receiptJSON + `],"pagination":{"hasPreviousPage":true,"startCursor":"start"}}`))
	})
	limit, before := 10, "previous-cursor"
	page, _, err := client.GetReceipts(context.Background(), "entry/id").
		WithSource(accountingentries.ReceiptSourceUser).WithSource(accountingentries.ReceiptSourceCamera).
		WithFileType(accountingentries.ReceiptFileTypePreview).
		WithPagingInfo(shared.PagingInfo{Before: &before, Limit: &limit, SortingKeys: []string{"createdAt"}, SortingOrder: []string{"ASC"}}).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Data) != 1 || !page.Pagination.HasPreviousPage || page.Pagination.StartCursor != "start" {
		t.Fatalf("unexpected receipts page: %+v", page)
	}
	checkReceipt(t, page.Data[0])
}

func TestGetReceipt(t *testing.T) {
	client := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		checkRequest(t, r, http.MethodGet, "/v1/accounting-entries/entry%2Fid/receipts/receipt%2Fid", url.Values{
			"file_type": {"PREVIEW"},
		}, "")
		_, _ = w.Write([]byte(`{"data":` + receiptJSON + `}`))
	})
	receipt, _, err := client.GetReceipt(context.Background(), "entry/id", "receipt/id").
		WithFileType(accountingentries.ReceiptFileTypePreview).Execute()
	if err != nil {
		t.Fatal(err)
	}
	checkReceipt(t, *receipt)
}

func TestSearchReceipts(t *testing.T) {
	client := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		checkRequest(t, r, http.MethodPost, "/v1/accounting-entries/receipts:search", url.Values{
			"company_id": {"company-id"}, "partner_id": {"partner-id"}, "organization_id": {"organization-id"},
			"after": {"cursor"}, "limit": {"25"},
		}, `{"includeDeleted":true,"accountingEntryIds":["entry-1","entry-2"],"employeeIds":["employee-id"],
			"fileType":"PREVIEW","mimeType":["application/pdf","image/jpeg"],"source":["USER","CAMERA"]}`)
		_, _ = w.Write([]byte(`{"data":[` + receiptJSON + `],"pagination":{"hasNextPage":true,"endCursor":"next"}}`))
	})
	limit, after := 25, "cursor"
	page, _, err := client.SearchReceipts(context.Background()).
		WithBody(accountingentries.SearchReceiptsRequest{AccountingEntryIDs: []string{"entry-1"}}).
		WithIncludeDeleted(true).WithAccountingEntryIDs("entry-2").WithEmployeeIDs("employee-id").
		WithCompanyID("company-id").WithPartnerID("partner-id").WithOrganizationID("organization-id").
		WithPagingInfo(shared.PagingInfo{After: &after, Limit: &limit}).WithFileType(accountingentries.ReceiptFileTypePreview).
		WithMIMEType(accountingentries.ReceiptMIMETypePDF, accountingentries.ReceiptMIMETypeJPEG).
		WithSource(accountingentries.ReceiptSourceUser, accountingentries.ReceiptSourceCamera).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Data) != 1 || !page.Pagination.HasNextPage || page.Pagination.EndCursor != "next" {
		t.Fatalf("unexpected receipts search page: %+v", page)
	}
	checkReceipt(t, page.Data[0])
}

func TestOptionalFiltersAreOmitted(t *testing.T) {
	client := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Query()) != 0 {
			t.Errorf("unexpected query: %v", r.URL.Query())
		}
		if r.Method == http.MethodPost {
			checkRequest(t, r, http.MethodPost, r.URL.EscapedPath(), url.Values{}, `{"includeDeleted":false}`)
		}
		if r.Method == http.MethodPost || r.URL.Path == "/v1/accounting-entries/entry/receipts" {
			_, _ = w.Write([]byte(`{"data":[],"pagination":{}}`))
		} else {
			_, _ = w.Write([]byte(`{"data":{}}`))
		}
	})
	ctx := context.Background()
	_, _, err := client.GetAccountingEntry(ctx, "entry").Execute()
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = client.SearchAccountingEntries(ctx).Execute()
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = client.GetReceipts(ctx, "entry").Execute()
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = client.GetReceipt(ctx, "entry", "receipt").Execute()
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = client.SearchReceipts(ctx).Execute()
	if err != nil {
		t.Fatal(err)
	}
}

func TestOperationsPropagateFailures(t *testing.T) {
	operations := map[string]func(*accountingentries.Client, context.Context) (bool, *shared.ResponseExtra, error){
		"get entry": func(c *accountingentries.Client, ctx context.Context) (bool, *shared.ResponseExtra, error) {
			out, extra, err := c.GetAccountingEntry(context.Background(), "entry").WithContext(ctx).Execute()
			return out == nil, extra, err
		},
		"search entries": func(c *accountingentries.Client, ctx context.Context) (bool, *shared.ResponseExtra, error) {
			out, extra, err := c.SearchAccountingEntries(context.Background()).WithContext(ctx).Execute()
			return out == nil, extra, err
		},
		"get receipts": func(c *accountingentries.Client, ctx context.Context) (bool, *shared.ResponseExtra, error) {
			out, extra, err := c.GetReceipts(context.Background(), "entry").WithContext(ctx).Execute()
			return out == nil, extra, err
		},
		"get receipt": func(c *accountingentries.Client, ctx context.Context) (bool, *shared.ResponseExtra, error) {
			out, extra, err := c.GetReceipt(context.Background(), "entry", "receipt").WithContext(ctx).Execute()
			return out == nil, extra, err
		},
		"search receipts": func(c *accountingentries.Client, ctx context.Context) (bool, *shared.ResponseExtra, error) {
			out, extra, err := c.SearchReceipts(context.Background()).WithContext(ctx).Execute()
			return out == nil, extra, err
		},
	}
	client := newTestClient(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "12")
		w.WriteHeader(http.StatusTooManyRequests)
	})
	for name, execute := range operations {
		t.Run(name, func(t *testing.T) {
			isNil, extra, err := execute(client, context.Background())
			var requestErr *shared.RequestError
			if !isNil || !errors.As(err, &requestErr) || requestErr.StatusCode != http.StatusTooManyRequests {
				t.Fatalf("expected nil result and HTTP 429, got nil=%v, err=%v", isNil, err)
			}
			if extra == nil || extra.RateLimit == nil || extra.RateLimit.RetryAfterSeconds == nil || *extra.RateLimit.RetryAfterSeconds != 12 {
				t.Fatalf("missing retry metadata: %+v", extra)
			}

			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			isNil, extra, err = execute(client, ctx)
			if !isNil || extra != nil || !errors.Is(err, context.Canceled) {
				t.Fatalf("expected cancellation without a result or metadata, got nil=%v, extra=%+v, err=%v", isNil, extra, err)
			}
		})
	}
}
