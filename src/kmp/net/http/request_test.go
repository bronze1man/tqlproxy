
package http
import "testing"
import "net/http"
func TestNewRequestFromOrigin(T *testing.T){
	NewRequestFromOrigin(&http.Request{})
}

