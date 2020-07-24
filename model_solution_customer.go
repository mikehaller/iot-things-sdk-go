/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).  The Bosch IoT Things HTTP API uses response status codes     (see [RFC 7231](https://tools.ietf.org/html/rfc7231#section-6))     to indicate whether a specific request has been successfully completed, or not. However, the descriptions we provide additionally to the status code (e.g. in our API docs, or error codes like. \"solutions:transaction.count.exceeded\") might change without advance notice. These are not be considered as official API, and must therefore not be applied in your applications or tests.
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package iotthings
// SolutionCustomer struct for SolutionCustomer
type SolutionCustomer struct {
	// The name of the customer
	Name string `json:"name"`
	// The email address of the customer
	Email string `json:"email"`
	// Additional information about the customer
	Info string `json:"info"`
}
