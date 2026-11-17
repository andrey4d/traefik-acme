package resolver

import (
	"bytes"
	"encoding/base64"
	"os"
	"path/filepath"
	"testing"
)

const sampleACME = `{
  "my-cert-resolver": {
    "Account": {
      "Email": "test@example.com",
      "Registration": {"body": {"status": "valid"}, "uri": "https://acme-v02.api.letsencrypt.org/directory/1"},
      "PrivateKey": "dGVzdA==",
      "KeyType": "ECDSA-P256"
    },
    "Certificates": [
      {
        "domain": {"main": "example.com"},
        "certificate": "Y2VydC1kYXRh",
        "key": "a2V5LWRhdGE=",
        "Store": "default"
      }
    ]
  }
}`

func TestLoad_WhenValidFile_ReturnsResolvers(t *testing.T) {
	t.Parallel()

	file := filepath.Join(t.TempDir(), "acme.json")
	if err := os.WriteFile(file, []byte(sampleACME), 0o644); err != nil {
		t.Fatalf("write sample file: %v", err)
	}

	resolvers, err := Load(file)
	if err != nil {
		t.Fatalf("Load() error = %v, want nil", err)
	}

	r, ok := resolvers["my-cert-resolver"]
	if !ok {
		t.Fatalf("resolver %q not found in result: %+v", "my-cert-resolver", resolvers)
	}
	if r.Account.Email != "test@example.com" {
		t.Errorf("Account.Email = %q, want %q", r.Account.Email, "test@example.com")
	}

	certs := r.GetCerts()
	if len(certs) != 1 || certs[0].Domain != "example.com" {
		t.Fatalf("GetCerts() = %+v, want one cert for example.com", certs)
	}
}

func TestLoad_WhenFileMissingOrInvalid_ReturnsError(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	invalid := filepath.Join(dir, "invalid.json")
	if err := os.WriteFile(invalid, []byte("{not-json"), 0o644); err != nil {
		t.Fatalf("write invalid file: %v", err)
	}

	tests := []struct {
		name string
		file string
	}{
		{"missing file", filepath.Join(dir, "does-not-exist.json")},
		{"invalid json", invalid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Load(tt.file); err == nil {
				t.Fatalf("Load(%q) error = nil, want error", tt.file)
			}
		})
	}
}

func TestSave_WhenDirMissing_CreatesFilesWithPermissions(t *testing.T) {
	t.Parallel()

	out := filepath.Join(t.TempDir(), "nested", "certs")
	crt := DomainCert{
		Domain:      "example.com",
		Certificate: base64.StdEncoding.EncodeToString([]byte("CERT-DATA")),
		Key:         base64.StdEncoding.EncodeToString([]byte("KEY-DATA")),
	}

	if err := Save(crt, out); err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	checkFile := func(name string, want []byte, wantMode os.FileMode) {
		data, err := os.ReadFile(filepath.Join(out, name))
		if err != nil {
			t.Fatalf("read %q: %v", name, err)
		}
		if !bytes.Equal(data, want) {
			t.Errorf("%s content = %q, want %q", name, data, want)
		}
		info, err := os.Stat(filepath.Join(out, name))
		if err != nil {
			t.Fatalf("stat %q: %v", name, err)
		}
		if info.Mode().Perm() != wantMode {
			t.Errorf("%s mode = %o, want %o", name, info.Mode().Perm(), wantMode)
		}
	}

	checkFile("example.com.crt", []byte("CERT-DATA"), 0o644)
	checkFile("example.com.key", []byte("KEY-DATA"), 0o600)
}

func TestSave_WhenInvalidBase64_ReturnsError(t *testing.T) {
	t.Parallel()

	if err := Save(DomainCert{Domain: "example.com", Certificate: "!!not-base64!!"}, t.TempDir()); err == nil {
		t.Fatal("Save() with invalid base64 error = nil, want error")
	}
}
