package middleware

//
//import (
//	"context"
//	"github.com/flow-lab/auxospore/internal/tenant"
//	"github.com/golang/mock/gomock"
//	"net/http"
//	"net/http/httptest"
//	"regexp"
//	"testing"
//)
//
//func TestAuthenticate(t *testing.T) {
//	t.Run("should not authenticate", func(t *testing.T) {
//		req, err := http.NewRequest("GET", "/", nil)
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		rec := httptest.NewRecorder()
//		h := Chain(func(w http.ResponseWriter, _ *http.Request) {
//			_, _ = w.Write([]byte("ok"))
//		}, Authenticate(tenantServiceMock{}))
//
//		h.ServeHTTP(rec, req)
//
//		if rec.Code != http.StatusUnauthorized {
//			t.Errorf("response code was %v instead of %d", rec.Code, http.StatusUnauthorized)
//		}
//	})
//
//	t.Run("should authenticate", func(t *testing.T) {
//		req, err := http.NewRequest("GET", "/", nil)
//		if err != nil {
//			t.Fatal(err)
//		}
//		req.Header.Set(ApiKeyKey, "1asd2sad8ajdjsadjk3")
//
//		ctrl := gomock.NewController(t)
//		defer ctrl.Finish()
//
//		rec := httptest.NewRecorder()
//		h := Chain(func(w http.ResponseWriter, _ *http.Request) {
//			_, _ = w.Write([]byte("ok"))
//		}, Authenticate(tenantServiceMock{}))
//
//		h.ServeHTTP(rec, req)
//
//		if rec.Code != http.StatusOK {
//			t.Errorf("response code was %v instead of %d", rec.Code, http.StatusOK)
//		}
//	})
//}
//
//func Test_matches(t *testing.T) {
//	type args struct {
//		re  *regexp.Regexp
//		key string
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		{
//			name: "should match",
//			args: struct {
//				re  *regexp.Regexp
//				key string
//			}{
//				re: regexp.MustCompile(apiKeyRegex), key: "1234asd213123ad4sff",
//			},
//			want: true,
//		},
//		{
//			name: "should match",
//			args: struct {
//				re  *regexp.Regexp
//				key string
//			}{
//				re: regexp.MustCompile(apiKeyRegex), key: "99999999999999",
//			},
//			want: true,
//		},
//		{
//			name: "should match",
//			args: struct {
//				re  *regexp.Regexp
//				key string
//			}{
//				re: regexp.MustCompile(apiKeyRegex), key: "A99999999999999",
//			},
//			want: false,
//		},
//		{
//			name: "should not match",
//			args: struct {
//				re  *regexp.Regexp
//				key string
//			}{
//				re: regexp.MustCompile(apiKeyRegex), key: " 99999999999999 ",
//			},
//			want: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := matches(tt.args.re, tt.args.key); got != tt.want {
//				t.Errorf("matches() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//type tenantServiceMock struct {
//}
//
//func (t tenantServiceMock) GetTenantByApiKey(_ context.Context, _ string) (*tenant.Tenant, error) {
//	return &tenant.Tenant{}, nil
//}
//
//func (t tenantServiceMock) CreateOrUpdate(ctx context.Context, cmd tenant.CreateOrUpdateTenantCmd) error {
//	panic("implement me")
//}
//
//func (t tenantServiceMock) Ping(ctx context.Context) error {
//	panic("implement me")
//}
