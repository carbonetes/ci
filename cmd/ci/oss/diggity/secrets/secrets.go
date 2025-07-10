package secrets

import (
	stream "github.com/carbonetes/diggity/cmd/diggity/grove"
	"github.com/carbonetes/diggity/pkg/cdx"
	"github.com/carbonetes/diggity/pkg/scanner/secret"
	diggityTypes "github.com/carbonetes/diggity/pkg/types"
)

func Analyze() []diggityTypes.Secret {

	addr, _ := diggityTypes.NewAddress()
	cdx.New(addr)
	// Secrets
	cdx.New(addr)
	secretAddr := *addr
	secretAddr.NID = "secret"
	secret.New(&secretAddr)

	// --- Retrieve secrets ---
	s, _ := stream.Get(secretAddr.String())
	secrets := s.([]diggityTypes.Secret)
	// -----------------------
	return secrets
}
