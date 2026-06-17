package api

import (
	"context"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func testClient(srv *httptest.Server) *Client {
	return &Client{
		httpClient: srv.Client(),
		baseURL:    srv.URL,
		realmID:    "123",
	}
}

// The /download endpoint returns the URL as a JSON-quoted string literal;
// callers must receive it without the surrounding quotes.
func TestFetchDownloadURLStripsQuotes(t *testing.T) {
	const want = "https://intuit-storage.example/attachments/receipt.pdf?Expires=123&Signature=abc"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/company/123/download/789" {
			t.Errorf("unexpected path %q", r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(want) // writes a quoted string literal
	}))
	defer srv.Close()

	got, err := testClient(srv).FetchDownloadURL(context.Background(), "789")
	if err != nil {
		t.Fatalf("FetchDownloadURL: %v", err)
	}
	if got != want {
		t.Errorf("got %q, want %q (surrounding quotes not stripped?)", got, want)
	}
}

func TestUploadMultipart(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/company/123/upload" {
			t.Errorf("unexpected path %q", r.URL.Path)
		}
		mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if err != nil || !strings.HasPrefix(mediaType, "multipart/") {
			t.Fatalf("content-type = %q", r.Header.Get("Content-Type"))
		}
		mr := multipart.NewReader(r.Body, params["boundary"])
		var names []string
		var meta map[string]any
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatalf("read part: %v", err)
			}
			names = append(names, p.FormName())
			if p.FormName() == "file_metadata_01" {
				b, _ := io.ReadAll(p)
				_ = json.Unmarshal(b, &meta)
			}
		}
		if len(names) != 2 || names[0] != "file_metadata_01" || names[1] != "file_content_01" {
			t.Errorf("part names = %v, want [file_metadata_01 file_content_01]", names)
		}
		if meta["FileName"] != "receipt.pdf" {
			t.Errorf("metadata FileName = %v, want receipt.pdf", meta["FileName"])
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"AttachableResponse":[{"Attachable":{"Id":"42","FileName":"receipt.pdf"}}]}`))
	}))
	defer srv.Close()

	meta := []byte(`{"FileName":"receipt.pdf","ContentType":"application/pdf"}`)
	att, err := testClient(srv).Upload(context.Background(), meta, "receipt.pdf", "application/pdf", strings.NewReader("PDFDATA"))
	if err != nil {
		t.Fatalf("Upload: %v", err)
	}
	if att["Id"] != "42" {
		t.Errorf("Attachable Id = %v, want 42", att["Id"])
	}
}

func TestDownloadStreamsBody(t *testing.T) {
	const payload = "file-bytes-here"
	mux := http.NewServeMux()
	srv := httptest.NewServer(mux)
	defer srv.Close()
	mux.HandleFunc("/v3/company/123/download/789", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(srv.URL + "/ext/receipt.pdf")
	})
	mux.HandleFunc("/ext/receipt.pdf", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(payload))
	})

	body, err := testClient(srv).Download(context.Background(), "789")
	if err != nil {
		t.Fatalf("Download: %v", err)
	}
	defer func() { _ = body.Close() }()
	got, _ := io.ReadAll(body)
	if string(got) != payload {
		t.Errorf("got %q, want %q", got, payload)
	}
}
