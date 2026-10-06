// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCertificatesCRD(t *testing.T) {

	var client = newTestClient(t)

	// I'm not sure why we have to specify the types for this
	// function but not testCRUD for the other test files, but
	// I've done it to satisfy the compiler.
	// To investigate further.
	testCRD[Certificate, CertificateBase](t, client.Certificates)
}

func TestCertificatesCustom(t *testing.T) {

	var err error
	var client = newTestClient(t)

	var randInt = rand.IntN(1e5)

	t.Run("Test validating an invalid certificate", func(t *testing.T) {

		invalidCert := CertificateUpload{
			CertificateKey: returnPointer("dummyKey"),
			Certificate:    returnPointer("dummyCert"),
		}
		err = client.Certificates.Validate(invalidCert)
		assert.Error(t, err)
	})

	// Generate a certificate and key
	rawCert, rawKey := NewCertWithKey()

	_certBody := CertificateUpload{
		Certificate:    &rawCert,
		CertificateKey: &rawKey,
	}

	t.Run("Test validating a valid cert", func(t *testing.T) {
		err = client.Certificates.Validate(_certBody)
		assert.NoError(t, err)
	})

	t.Run("Test uploading a certificate", func(t *testing.T) {

		// Validate it for good measure
		err = client.Certificates.Validate(_certBody)
		assert.NoError(t, err)

		// Create the Certificate object in NPM
		_cert := CertificateBase{
			Provider:    returnPointer("other"),
			NiceName:    returnPointer(fmt.Sprintf("test_%d", randInt)),
			DomainNames: returnPointer([]string{fmt.Sprintf("domain_%d.com", randInt)}),
		}
		cert, err := client.Certificates.Create(_cert)
		assert.NoError(t, err)

		// Upload the certificate and check that the body matches
		newCert, err := client.Certificates.Upload(cert.ID, _certBody)
		assert.NoError(t, err)
		assert.Equal(t, newCert.Certificate, _certBody.Certificate)

		// Fetch the certificate from NPM and check the body matches
		certGet, err := client.Certificates.Get(cert.ID)
		assert.NoError(t, err)

		assert.Equal(t, certGet.Meta.Certificate, _certBody.Certificate)
		assert.Equal(t, certGet.Meta.CertificateKey, _certBody.CertificateKey)

		err = client.Certificates.Delete(cert.ID)
		if err != nil {
			t.Errorf("Unable to tidy up certificate: %v", err)
		}
	})

}

func TestCertificatesLetsencrypt(t *testing.T) {

	var err error
	var client = newTestClient(t)

	var randInt = rand.IntN(1e5)

	var _certMeta = CertificateMeta{
		DNSProvider:            returnPointer("powerdns"),
		DNSProviderCredentials: returnPointer("dns_powerdns_api_url = http://pdns:8081\r\ndns_powerdns_api_key = npm"),
		DNSChallenge:           returnPointer(true),
		PropagationSeconds:     returnPointer(10),
	}

	var _certCreate = CertificateBase{
		Provider:    returnPointer("letsencrypt"),
		NiceName:    returnPointer(fmt.Sprintf("test%d", randInt)),
		DomainNames: returnPointer([]string{"www.example.com", "2.example.com"}),
		Meta:        &_certMeta,
	}

	var cert Certificate
	cert, err = client.Certificates.Create(_certCreate)
	assert.NoError(t, err)

	assert.ElementsMatch(t, *cert.DomainNames, *_certCreate.DomainNames)
	assert.NotNil(t, cert.Meta.LetsencryptCertificate.Dates.To)
	assert.NotNil(t, cert.Meta.LetsencryptCertificate.Dates.From)

	// Don't check NiceName as it the returned NiceName is just the domain names
	// assert.Equal(t, *cert.NiceName, *_certCreate.NiceName)

}
