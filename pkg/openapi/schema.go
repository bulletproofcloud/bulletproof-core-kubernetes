// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/unikorn-cloud/core/pkg/openapi"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"FwAA///D2MnBin4AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "https://raw.githubusercontent.com/unikorn-cloud/core/main/pkg/openapi/common.spec.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
