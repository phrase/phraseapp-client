package phraseapp

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/peterbourgon/diskv"
	"github.com/pkg/errors"
)

type httpCacheClient struct {
	contentCache *diskv.Diskv
	etagCache    *diskv.Diskv
	debug        bool
}

type cacheRecord struct {
	URL      string
	Response *httpResponse
	Payload  []byte
}

// httpResponse is a serializable copy of a http.Response
type httpResponse struct {
	Status           string
	StatusCode       int
	Proto            string
	ProtoMajor       int
	ProtoMinor       int
	Header           http.Header
	ContentLength    int64
	TransferEncoding []string
	Uncompressed     bool
	Trailer          http.Header
}

// newHTTPCacheClient returns a client to interact with the PhraseApp API and is caching the results
// This is experimental and should be used with care
func newHTTPCacheClient(debug bool) (*httpCacheClient, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}

	cachePath := filepath.Join(cacheDir, "phraseapp")
	var cacheSizeMax uint64 = 1024 * 1024 * 100 // 100MB
	cache := &httpCacheClient{
		contentCache: diskv.New(diskv.Options{
			BasePath:     cachePath,
			CacheSizeMax: cacheSizeMax,
		}),
		etagCache: diskv.New(diskv.Options{
			BasePath:     cachePath,
			CacheSizeMax: cacheSizeMax,
		}),
	}
	cache.debug = debug
	return cache, nil
}

func (client *httpCacheClient) RoundTrip(req *http.Request) (*http.Response, error) {
	url := req.URL.String()
	cachedResponse := client.getCache(req, url)
	rsp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	if client.debug {
		log.Printf("real status=%d", rsp.StatusCode)
	}
	if rsp.StatusCode == http.StatusNotModified {
		if client.debug {
			log.Printf("found in cache returning cached body")
		}
		cachedResponse.setCachedResponse(rsp)
		return rsp, nil
	} else if rsp.Status[0] != '2' {
		return nil, errors.Errorf("got status %s but expected 2x. body=%s", rsp.Status, string(body))
	}

	etag := rsp.Header.Get("Etag")
	etagCacheKey := md5sum(url)
	err = client.etagCache.Write(etagCacheKey, []byte(etag))
	if err != nil {
		return nil, err
	}

	contentCacheKey := md5sum(etag + url)
	encodedCache := cachedResponse.encode(rsp, url, body)
	err = client.contentCache.Write(contentCacheKey, encodedCache)
	if err != nil {
		return nil, err
	}

	rsp.Body = ioutil.NopCloser(bytes.NewReader(body))
	return rsp, nil
}

func (client *httpCacheClient) getCache(req *http.Request, url string) *cacheRecord {
	var cachedResponse *cacheRecord
	etagResult, err := client.etagCache.Read(md5sum(url))
	if err != nil {
		if client.debug {
			log.Println("doing request without etag")
		}
	} else {
		etag := string(etagResult)
		if client.debug {
			log.Printf("using etag %s in request", etag)
		}
		cache, err := client.contentCache.Read(md5sum(etag + url))
		if err != nil {
			if client.debug {
				log.Println("found etag but no cached response")
			}
		} else {
			req.Header.Set("If-None-Match", etag)
			var buf bytes.Buffer
			buf.Write(cache)
			decoder := gob.NewDecoder(&buf)
			err = decoder.Decode(&cachedResponse)
		}
	}

	return cachedResponse
}

func (record *cacheRecord) encode(rsp *http.Response, url string, body []byte) []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(cacheRecord{URL: url, Payload: body, Response: &httpResponse{
		Status:           rsp.Status,
		StatusCode:       rsp.StatusCode,
		Proto:            rsp.Proto,
		ProtoMajor:       rsp.ProtoMajor,
		ProtoMinor:       rsp.ProtoMinor,
		Header:           rsp.Header,
		ContentLength:    rsp.ContentLength,
		TransferEncoding: rsp.TransferEncoding,
		Trailer:          rsp.Header,
	}})

	return buf.Bytes()
}

func (record *cacheRecord) setCachedResponse(rsp *http.Response) {
	rsp.Status = record.Response.Status
	rsp.StatusCode = record.Response.StatusCode
	rsp.Proto = record.Response.Proto
	rsp.ProtoMajor = record.Response.ProtoMajor
	rsp.ProtoMinor = record.Response.ProtoMinor
	rsp.Header = record.Response.Header
	rsp.ContentLength = record.Response.ContentLength
	rsp.TransferEncoding = record.Response.TransferEncoding
	rsp.Trailer = record.Response.Header
	rsp.Body = ioutil.NopCloser(bytes.NewReader(record.Payload))
}

func md5sum(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
