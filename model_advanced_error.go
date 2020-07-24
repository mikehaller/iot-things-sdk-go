/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).  The Bosch IoT Things HTTP API uses response status codes     (see [RFC 7231](https://tools.ietf.org/html/rfc7231#section-6))     to indicate whether a specific request has been successfully completed, or not. However, the descriptions we provide additionally to the status code (e.g. in our API docs, or error codes like. \"solutions:transaction.count.exceeded\") might change without advance notice. These are not be considered as official API, and must therefore not be applied in your applications or tests.
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package iotthings
// AdvancedError struct for AdvancedError
type AdvancedError struct {
	// The HTTP status of the error
	Status int32 `json:"status"`
	// The error code of the occurred exception
	Error string `json:"error"`
	// The message of the error - what went wrong
	Message string `json:"message"`
	// A description how to fix the error or more details
	Description string `json:"description,omitempty"`
	// A link to further information about the error and how to fix it
	Href string `json:"href,omitempty"`
}
