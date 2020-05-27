// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewDownloadClusterISOParams creates a new DownloadClusterISOParams object
// no default values defined in spec.
func NewDownloadClusterISOParams() DownloadClusterISOParams {

	return DownloadClusterISOParams{}
}

// DownloadClusterISOParams contains all the bound params for the download cluster i s o operation
// typically these are obtained from a http.Request
//
// swagger:parameters DownloadClusterISO
type DownloadClusterISOParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ClusterID strfmt.UUID
	/*
	  Required: true
	  In: query
	*/
	ImageID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDownloadClusterISOParams() beforehand.
func (o *DownloadClusterISOParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rClusterID, rhkClusterID, _ := route.Params.GetOK("cluster_id")
	if err := o.bindClusterID(rClusterID, rhkClusterID, route.Formats); err != nil {
		res = append(res, err)
	}

	qImageID, qhkImageID, _ := qs.GetOK("image_id")
	if err := o.bindImageID(qImageID, qhkImageID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClusterID binds and validates parameter ClusterID from path.
func (o *DownloadClusterISOParams) bindClusterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("cluster_id", "path", "strfmt.UUID", raw)
	}
	o.ClusterID = *(value.(*strfmt.UUID))

	if err := o.validateClusterID(formats); err != nil {
		return err
	}

	return nil
}

// validateClusterID carries on validations for parameter ClusterID
func (o *DownloadClusterISOParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.FormatOf("cluster_id", "path", "uuid", o.ClusterID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindImageID binds and validates parameter ImageID from query.
func (o *DownloadClusterISOParams) bindImageID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("image_id", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("image_id", "query", raw); err != nil {
		return err
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("image_id", "query", "strfmt.UUID", raw)
	}
	o.ImageID = *(value.(*strfmt.UUID))

	if err := o.validateImageID(formats); err != nil {
		return err
	}

	return nil
}

// validateImageID carries on validations for parameter ImageID
func (o *DownloadClusterISOParams) validateImageID(formats strfmt.Registry) error {

	if err := validate.FormatOf("image_id", "query", "uuid", o.ImageID.String(), formats); err != nil {
		return err
	}
	return nil
}
