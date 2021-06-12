// server_test.go
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	// ルーティングの設定はそのままに
	// テスト用のサーバを立ち上げることができる
	ts := httptest.NewServer(Route())
	defer ts.Close()

	// 通常通りHTTPリクエストを送ることができる
	// テスト用サーバのURLは ts.URL で取得できる
	res, err := http.Get(ts.URL + "/greet?name=gopher")
	if err != nil {
		t.Fatalf("http.Get failed: %s", err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("read from HTTP Response Body failed: %s", err)
	}
	expected := "Hello, gopher"
	if string(greeting) != expected {
		t.Fatalf("response of /greet?name=gopher returns %s, want %s", string(greeting), expected)
	}
}

