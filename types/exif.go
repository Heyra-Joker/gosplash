/*
 * Copyright (c) 2024. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package types

type Exif struct {
	Make         string `json:"make"`
	Model        string `json:"model"`
	Name         string `json:"name"`
	ExposureTime string `json:"exposure_time"`
	Aperture     string `json:"aperture"`
	FocalLength  string `json:"focal_length"`
	Iso          int    `json:"iso"`
}
