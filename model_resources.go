/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).  The Bosch IoT Things HTTP API uses response status codes     (see [RFC 7231](https://tools.ietf.org/html/rfc7231#section-6))     to indicate whether a specific request has been successfully completed, or not. However, the descriptions we provide additionally to the status code (e.g. in our API docs, or error codes like. \"solutions:transaction.count.exceeded\") might change without advance notice. These are not be considered as official API, and must therefore not be applied in your applications or tests.
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package iotthings
// Resources (Authorization) Resources containing one ResourceEntry for each `type:path` key, `type` being one of the following `thing`, `policy`, `message` and `solution`   Note: The solution resources can be addressed ONLY within a special policy per solution.  See [policy of a solution](https://docs.bosch-iot-suite.com/things/basic-concepts/auth/auth-policy-solution/)
type Resources struct {
	Thing ResourceEntry `json:"thing:/,omitempty"`
	ThingAttributesSomePath ResourceEntryRestricted `json:"thing:/attributes/some/path,omitempty"`
	Policy ResourceEntry `json:"policy:/,omitempty"`
	Message ResourceEntry `json:"message:/,omitempty"`
	Solution ResourceEntry `json:"solution:/,omitempty"`
}
