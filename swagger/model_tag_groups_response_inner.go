/*
 * Pleo.io API
 *
 * Pleo.io API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TagGroupsResponseInner struct {
	// This is the internal pleo identifier of the tag group
	Id string `json:"id"`
	// This is the tag group name supplied on create
	Name string `json:"name"`
	// These are the attributes that define a tag in this tag group, i.e. the tag columns. Each tag column is included only once.
	TagAttributes []interface{} `json:"tagAttributes"`
	// This flag specifies whether the tag group has been archived in the Pleo system or not
	Archived bool `json:"archived"`
}
