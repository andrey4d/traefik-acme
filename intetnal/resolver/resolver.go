package resolver

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

/*
 *   Copyright (c) 2026 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */

type CertResolver struct {
	Account      Account        `json:"Account"`
	Certificates []Certificates `json:"Certificates"`
}

// Account struct
type Account struct {
	Email        string       `json:"Email"`
	Registration Registration `json:"Registration"`
	PrivateKey   string       `json:"PrivateKey"`
	KeyType      string       `json:"KeyType"`
}
type Registration struct {
	Body Body   `json:"body"`
	URI  string `json:"uri"`
}
type Body struct {
	Status string `json:"status"`
}

// Certificates struct
type Certificates struct {
	Domain      Domain `json:"domain"`
	Certificate string `json:"certificate"`
	Key         string `json:"key"`
	Store       string `json:"Store"`
}
type Domain struct {
	Main string `json:"main"`
}

type DomainCert struct {
	Domain      string
	Certificate string
	Key         string
}

func (r *CertResolver) GetCerts() (domainCrts []DomainCert) {
	for _, v := range r.Certificates {
		domainCrts = append(domainCrts, DomainCert{
			Domain:      v.Domain.Main,
			Certificate: v.Certificate,
			Key:         v.Key,
		})
	}
	return domainCrts
}

func Load(filename string) (map[string]*CertResolver, error) {
	// Open the file for reading
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	resolvers := make(map[string]*CertResolver)
	if err = json.Unmarshal(data, &resolvers); err != nil {
		return nil, err
	}

	return resolvers, nil
}

func DecodeBase64(input string) ([]byte, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %s", err)
	}
	return decodedBytes, nil
}

func Save(domainCertificate DomainCert, path string) error {
	filename := filepath.Join(path, domainCertificate.Domain)
	certificate, err := DecodeBase64(domainCertificate.Certificate)
	if err != nil {
		return err
	}
	key, err := DecodeBase64(domainCertificate.Key)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filename+".crt", certificate, 0644); err != nil {
		return err
	}
	if err := os.WriteFile(filename+".key", key, 0644); err != nil {
		return err
	}
	return nil
}
