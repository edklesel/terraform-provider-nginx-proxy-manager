// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"mime/multipart"
	"reflect"
	"slices"
	"strings"
	"time"
)

func setZeroSlices(object interface{}) {

	currentObject := reflect.ValueOf(object).Elem()

	setZeroSliceFields(currentObject)
}

func setZeroSliceFields(value reflect.Value) {

	objType := value.Type()

	for i := 0; i < objType.NumField(); i++ {

		field := objType.Field(i)
		fieldValue := value.FieldByName(field.Name)

		// For struct arguments, then recurse
		if fieldValue.Kind() == reflect.Struct {
			setZeroSliceFields(fieldValue)
			continue

			// Otherwise, skip non-pointers
		} else if !slices.Contains([]reflect.Kind{reflect.Pointer, reflect.Slice}, fieldValue.Kind()) {
			continue
		}

		if fieldValue.IsNil() {

			// We want to set the pointer value to
			// the address of an ampty slice
			if fieldValue.Kind() == reflect.Pointer && field.Type.Elem().Kind() == reflect.Slice {
				// First set the pointer to an nil slice
				fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
				// Then make that slice len 0
				fieldValue.Elem().Set(reflect.MakeSlice(field.Type.Elem(), 0, 0))
				// If it's a nil slice, set it to an empty value
			} else if fieldValue.Kind() == reflect.Slice {
				fieldValue.Set(reflect.MakeSlice(field.Type, 0, 0))
			}

		}

	}

}

func returnPointer[S any](s S) *S {
	return &s
}

func newMultipartBody(fields map[string]string) (body bytes.Buffer, content_type string, err error) {

	var fieldWriter io.Writer
	var writer = multipart.NewWriter(&body)

	// Cycle through the fields and cretae a multipart body
	for field, value := range fields {
		if fieldWriter, err = writer.CreateFormFile(field, field); err != nil {
			return
		} else {
			_, _ = io.Copy(fieldWriter, strings.NewReader(value))
		}
	}

	// Get the content-type from the write object
	content_type = writer.FormDataContentType()

	// Tidy up
	_ = writer.Close()

	return
}

func NewCertWithKey() (cert string, key string) {

	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate private key: %v", err))
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"NPM Testing"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:  x509.KeyUsageEncipherOnly | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
		BasicConstraintsValid: true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privKey.PublicKey, privKey)
	if err != nil {
		panic(fmt.Sprintf("Failed to create certificate: %v", err))
	}

	cert = string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certBytes}))
	key = string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privKey)}))

	return

}
