package proxy

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	nginxServer := newNginxServer()
	appSatatusURL := "/app/status"
	createuserURL := "/create/user"

	httpCode, body := nginxServer.handleRequest(appSatatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appSatatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appSatatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appSatatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appSatatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appSatatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, httpCode, body)
}
