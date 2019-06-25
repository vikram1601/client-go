/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Request to create notes in batch.
type V1beta1BatchCreateNotesRequest struct {
	// The name of the project in the form of `projects/[PROJECT_ID]`, under which the notes are to be created.
	Parent string `json:"parent,omitempty"`
	// The notes to create. Max allowed length is 1000.
	Notes map[string]V1beta1Note `json:"notes,omitempty"`
}
