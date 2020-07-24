/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).  The Bosch IoT Things HTTP API uses response status codes     (see [RFC 7231](https://tools.ietf.org/html/rfc7231#section-6))     to indicate whether a specific request has been successfully completed, or not. However, the descriptions we provide additionally to the status code (e.g. in our API docs, or error codes like. \"solutions:transaction.count.exceeded\") might change without advance notice. These are not be considered as official API, and must therefore not be applied in your applications or tests.
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package iotthings
// NewThing struct for NewThing
type NewThing struct {
	// The initial policy to create for this thing. This will create a separate policy entity managed by resource `/policies/{thingId}`.   Use the placeholder `{{ request:subjectId }}` in order to let the backend insert the authenticated subjectId of the HTTP request.
	Policy Policy `json:"_policy,omitempty"`
	// This field may contain the policy ID of an existing policy. The policy is copied and used for this newly created thing. This field may also contain a placeholder reference to a thing in the format `{{ ref:things/[thingId]/policyId }}` where you need to replace `[thingId]` with a valid thing ID. The newly created thing will then obtain a copy of the policy of the referenced thing. In the case of using a reference, the caller needs to have READ access to both the thing and the policy of the thing. In the case of using an explicit policy id to copy from, the caller needs to have READ access to the policy. If you want to specify a policy ID for the copied policy, use the policyId field. This field must not be used together with the field `_policy`. If you specify both `_policy` and `_copyPolicyFrom` this will lead to an error response.
	CopyPolicyFrom string `json:"_copyPolicyFrom,omitempty"`
	// The policy ID used for controlling access to this thing. Managed by resource `/policies/{policyId}`.
	PolicyId string `json:"policyId,omitempty"`
	// A single fully qualified identifier of a definition in the form 'namespace:name:version'
	Definition string `json:"definition,omitempty"`
	// An arbitrary JSON object describing the attributes of a thing.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// List of features where the key represents the `featureId` of each feature. The `featureId` key must be unique in the list.
	Features map[string]Feature `json:"features,omitempty"`
}
