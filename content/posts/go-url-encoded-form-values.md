---
title: Working with URL encoded form values
date: 2023-09-19
categories:
-  articles
---

In the book, we discussed how to send and received multipart form data. That is,
forms containing both text fields and file uploads, typically using the content
type as `multipart/form-data`. When we are working with only text fields, it is
sufficient to instead use the content type, `application/x-www-form-urlencoded`
and send the data as a _URL encoded set of key value pairs_. For example, as
described by the [MDN article](https://developer.mozilla.org/en-US/docs/Learn/Forms/Sending_and_retrieving_form_data#on_the_client_side_defining_how_to_send_the_data),
`say=Hi&to=Rheo` will be the request payload for a form with two text fields, `say` and `to` with the values,
`Hi` and `Rheo` respectively.

In this blog post, we will write a client and a server to send and process form data.

Let's get started!

- [Prerequisites](#prerequisites)
- [Reading URL encoded form data in a Go server](#reading-url-encoded-form-data-in-a-go-server)
- [Using a browser as the client](#using-a-browser-as-the-client)
- [Sending URL encoded form data from a Go client](#sending-url-encoded-form-data-from-a-go-client)
- [Learn more](#learn-more)


## Prerequisites

- An installation of Go
- A tool to edit your code. Any text editor you have will work fine
- A command terminal to run the `go` commands from

## Reading URL encoded form data in a Go server

Reading URL encoded form data in a server's handler function involves two key steps:

1. Call [`ParseForm()`](https://pkg.go.dev/net/http#Request.ParseForm) method of the request object
1. Then, retrieve the values of the fields using [`FormValue()`](https://pkg.go.dev/net/http#Request.FormValue)

Example code snippet:

```go
func handleForm(w http.ResponseWriter, req *http.Request) {
	// this will look for application/x-www-form-urlencoded
	// content type
	err := req.ParseForm()
	if err != nil {
		// if there was an error here, we return the error
		// as response along with a 400 http response code
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// once the form has been parsed correctly
	// we can obtain the form fields  using
	// req.FormValue passing the field name as the key
	// see ./post-method.html for the field names

	// it's worth noting here that FormValue() will return
	// the first value if there are multiple values specified
	// in the request
	// if you want to access the multiple values specified,
	// use req.Form directly (see https://pkg.go.dev/net/http#Request)
	say := req.FormValue("say")
	to := req.FormValue("to")

    // do something with the values read
    ..
}
```
You can find the complete code for the form handler in 
[form_handler.go](https://github.com/practicalgo/form-url-encoded-demo).

To run the application:

```
go build
./go-form-url-encoded.md # or .\go-form-url-encoded.exe 
```

## Using a browser as the client

First, let's make a request to the server from a browser by visiting http://localhost:8000. 

[Here](https://youtu.be/UX-ZWlq36EQ) is a video demonstration of using the application from a 
browser.

Key points to note from the above video are:

1. The "raw" request data (payload) is: `say=Hi&to=Rheo`
2. The content-type header is set as, `Content-Type: application/x-www-form-urlencoded`

Both are automatically done for us when we click on the "Send my greetings" button.

## Sending URL encoded form data from a Go client

Now, let's say you want to write a test for the `handleForm()` handler function.
You will need to send the relevant data as URL encoded form fields. Here is 
a code snippet showing you how:

```go
func TestHandleForm(t *testing.T) {
	w := httptest.NewRecorder()

    // set the key value pairs
	formBody := url.Values{
		"say": []string{"hi"}, // we only have a single element, but these need to be a slice
		"to":  []string{"rheo"}, // same as above
	}
	dataReader := formBody.Encode() // this perform the URL encoding and returns us a string

	req := httptest.NewRequest(
		http.MethodPost,
		"http://localhost",
		strings.NewReader(dataReader), // we need a io.Reader
	)

    // set the content-type header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	handleForm(w, req)

    // rest of the test
}
```

You can find the complete code for the form handler test function in 
[form_handler_test.go](https://github.com/practicalgo/form-url-encoded-demo).

Outside of test functions, you can use the same approach with `http.NewRequest` or 
`http.NewRequestWithContext`.

Key references:

- [`url.Values`](https://pkg.go.dev/net/url#Values) and [`Encode()`](https://pkg.go.dev/net/url#Values.Encode)
  method
- [`strings.NewReader()`](https://pkg.go.dev/strings#NewReader)

## Learn more

- [Sending and retrieving form data](https://developer.mozilla.org/en-US/docs/Learn/Forms/Sending_and_retrieving_form_data)

