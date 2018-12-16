/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateKeysChildrenGetParams creates a new WeaviateKeysChildrenGetParams object
// with the default values initialized.
func NewWeaviateKeysChildrenGetParams() *WeaviateKeysChildrenGetParams {
	var ()
	return &WeaviateKeysChildrenGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateKeysChildrenGetParamsWithTimeout creates a new WeaviateKeysChildrenGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateKeysChildrenGetParamsWithTimeout(timeout time.Duration) *WeaviateKeysChildrenGetParams {
	var ()
	return &WeaviateKeysChildrenGetParams{

		timeout: timeout,
	}
}

// NewWeaviateKeysChildrenGetParamsWithContext creates a new WeaviateKeysChildrenGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateKeysChildrenGetParamsWithContext(ctx context.Context) *WeaviateKeysChildrenGetParams {
	var ()
	return &WeaviateKeysChildrenGetParams{

		Context: ctx,
	}
}

// NewWeaviateKeysChildrenGetParamsWithHTTPClient creates a new WeaviateKeysChildrenGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateKeysChildrenGetParamsWithHTTPClient(client *http.Client) *WeaviateKeysChildrenGetParams {
	var ()
	return &WeaviateKeysChildrenGetParams{
		HTTPClient: client,
	}
}

/*WeaviateKeysChildrenGetParams contains all the parameters to send to the API endpoint
for the weaviate keys children get operation typically these are written to a http.Request
*/
type WeaviateKeysChildrenGetParams struct {

	/*KeyID
	  Unique ID of the key.

	*/
	KeyID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate keys children get params
func (o *WeaviateKeysChildrenGetParams) WithTimeout(timeout time.Duration) *WeaviateKeysChildrenGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate keys children get params
func (o *WeaviateKeysChildrenGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate keys children get params
func (o *WeaviateKeysChildrenGetParams) WithContext(ctx context.Context) *WeaviateKeysChildrenGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate keys children get params
func (o *WeaviateKeysChildrenGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate keys children get params
func (o *WeaviateKeysChildrenGetParams) WithHTTPClient(client *http.Client) *WeaviateKeysChildrenGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate keys children get params
func (o *WeaviateKeysChildrenGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyID adds the keyID to the weaviate keys children get params
func (o *WeaviateKeysChildrenGetParams) WithKeyID(keyID strfmt.UUID) *WeaviateKeysChildrenGetParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the weaviate keys children get params
func (o *WeaviateKeysChildrenGetParams) SetKeyID(keyID strfmt.UUID) {
	o.KeyID = keyID
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateKeysChildrenGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param keyId
	if err := r.SetPathParam("keyId", o.KeyID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
