package middleware

//
//import (
//	"context"
//	"github.com/flow-lab/auxospore/internal/tenant"
//	"net/http"
//	"regexp"
//)
//
//// just for demo
//const (
//	ApiKeyKey = "api-key"
//	AuthKey   = "auth-key"
//
//	apiKeyRegex = "^[0-9a-z]+$"
//)
//
//// Auth represents user.
//type Auth struct {
//	AccountID string
//}
//
//// Authenticate authenticates and set token in the context.
//func Authenticate(srv tenant.TenantSrv) Middleware {
//	re := regexp.MustCompile(apiKeyRegex)
//
//	return func(f http.HandlerFunc) http.HandlerFunc {
//		return func(w http.ResponseWriter, r *http.Request) {
//			at, ok := r.URL.Query()[ApiKeyKey]
//			var key string
//			if ok && len(at) == 1 {
//				key = at[0]
//			} else {
//				key = r.Header.Get(ApiKeyKey)
//			}
//
//			if !matches(re, key) {
//				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
//				return
//			}
//
//			t, err := srv.GetTenantByApiKey(r.Context(), key)
//			if err != nil {
//				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//				return
//			}
//
//			if t == nil {
//				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
//				return
//			}
//
//			// TODO: [grokrz] hit the Cache and get Auth
//			// If not in cash, hit db and populate cache
//			// For now i will just populate with empty Auth
//			ctx := context.WithValue(r.Context(), AuthKey, Auth{AccountID: "1"})
//			ctx = context.WithValue(ctx, ApiKeyKey, key)
//			f(w, r.WithContext(ctx))
//		}
//	}
//}
//
//func matches(re *regexp.Regexp, key string) bool {
//	return re.MatchString(key)
//}
//
//// ApiKey gets the api key from the context.
//func ApiKey(ctx context.Context) (string, bool) {
//	tokenStr, ok := ctx.Value(ApiKeyKey).(string)
//	return tokenStr, ok
//}
