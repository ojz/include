package README

import (
	"encoding/base64"
)

// Bytes returns the contents of the file.
func Bytes() []byte {
	decoded, err := base64.StdEncoding.DecodeString("SWYgeW91IHdhbnQgdG8gaW5jbHVkZSBzb21lIGZpbGVzIGluIGEgcGFja2FnZSwgdXNlIHRoaXMgc3RhbnphOgoKICAgIC8vZ286Z2VuZXJhdGUgaW5jbHVkZSBSRUFETUUubWQKCk5vdyB5b3UgY2FuIGluY2x1ZGUgdGhlIGZpbGUgbGlrZSB0aGlzOgoKICAgIGltcG9ydCAiPHlvdXIgcGF0aD4vaW50ZXJuYWwvYXNzZXRzL1JFQURNRSIKCkFuZCB5b3UgY2FuIHVzZSB0aGUgZmlsZSBsaWtlIHRoaXM6CgogICAgUkVBRE1FLkJ5dGVzKCkK")
	if err != nil {
		panic(err)
	}
	return decoded
}
