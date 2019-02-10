/*
 * 匿名掲示板API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type CommentProperties struct {

	Id int32 `json:"id"`

	Content string `json:"content"`

	PostId int32 `json:"post_id"`

	CommentedAt time.Time `json:"commented_at"`
}
