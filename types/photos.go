/*
 * Copyright (c) 2024. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package types

type Photos struct {
	ID             string `json:"id"`
	Slug           string `json:"slug"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	BlurHash       string `json:"blur_hash"`
	AssetType      string `json:"asset_type"`
	PromotedAt     string `json:"promoted_at"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Color          string `json:"color"`
	Description    string `json:"description"`
	AltDescription string `json:"alt_description"`
	Likes          int    `json:"likes"`
	LikedByUser    bool   `json:"liked_by_user"`
	Sponsorship    string `json:"sponsorship"`
	PublicDomain   bool   `json:"public_domain"`
	Views          int    `json:"views"`
	Downloads      int    `json:"downloads"`
}
