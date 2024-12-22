# Silent Errors in Go Maps

This example demonstrates a common issue in Go where accessing a map with a non-existent key doesn't result in an error, but instead returns the zero value of the map's value type. This can lead to subtle bugs that are hard to track down.

The `bug.go` file shows the problem: attempting to access a key that doesn't exist in the map returns 0, which might be a valid value in certain contexts, masking the error.

The `bugSolution.go` file provides a solution by explicitly checking if the key exists using the comma ok idiom.
