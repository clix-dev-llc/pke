/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline
import (
	"time"
)
// TokenCreateRequest struct for TokenCreateRequest
type TokenCreateRequest struct {
	Name string `json:"name,omitempty"`
	VirtualUser string `json:"virtualUser,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}
