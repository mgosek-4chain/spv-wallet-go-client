package httpx

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	bip32 "github.com/bitcoin-sv/go-sdk/compat/bip32"
	bsm "github.com/bitcoin-sv/go-sdk/compat/bsm"
	ec "github.com/bitcoin-sv/go-sdk/primitives/ec"
	"github.com/bitcoin-sv/spv-wallet-go-client/internal/cryptoutil"
	"github.com/bitcoin-sv/spv-wallet/models"
)

type NoopAuthHeaders struct{}

func (*NoopAuthHeaders) Append(req *http.Header) error { return nil }

type AuthHeaders struct {
	ExtendedKey *bip32.ExtendedKey
	AccessKey   *ec.PrivateKey
	Sign        bool
}

func (a *AuthHeaders) Append(h *http.Header) error {
	switch {
	case a.AccessKey != nil:
		return a.AppendForAccessKey(h)
	case a.Sign:
		return a.AppendForXPriv(h)
	default:
		return a.AppendForXPub(h)
	}
}

func (a *AuthHeaders) AppendForAccessKey(h *http.Header) error {
	b := AuthPayloadBuilder{
		PrivateHexKey: hex.EncodeToString(a.AccessKey.Serialize()),
		ExtendedKey:   a.ExtendedKey,
	}
	p, err := b.BuildForAccessKey()
	if err != nil {
		return fmt.Errorf("auth payload builder - build auth for access key op failure: %w", err)
	}

	h.Set(models.AuthAccessKey, p.AccessKey)
	h.Set(models.AuthHeaderHash, p.AuthHash)
	h.Set(models.AuthHeaderNonce, p.AuthNonce)
	h.Set(models.AuthHeaderTime, fmt.Sprintf("%d", p.AuthTime))
	h.Set(models.AuthSignature, p.Signature)
	return nil
}

func (a *AuthHeaders) AppendForXPub(h *http.Header) error {
	xPub, err := bip32.GetExtendedPublicKey(a.ExtendedKey)
	if err != nil {
		return fmt.Errorf("bip32 - get exended public key op failure: %w", err)
	}

	h.Set(models.AuthHeader, xPub)
	return nil
}

func (a *AuthHeaders) AppendForXPriv(h *http.Header) error {
	b := AuthPayloadBuilder{ExtendedKey: a.ExtendedKey}
	p, err := b.BuildForXPriv()
	if err != nil {
		return fmt.Errorf("auth payload builder - build auth for xpriv op failure: %w", err)
	}

	h.Set(models.AuthHeader, p.XPub)
	h.Set(models.AuthHeaderHash, p.AuthHash)
	h.Set(models.AuthHeaderNonce, p.AuthNonce)
	h.Set(models.AuthHeaderTime, fmt.Sprintf("%d", p.AuthTime))
	h.Set(models.AuthSignature, p.Signature)
	return nil
}

type AuthPayloadBuilder struct {
	PrivateHexKey string
	ExtendedKey   *bip32.ExtendedKey
	Body          string
}

func (b *AuthPayloadBuilder) BuildForAccessKey() (*models.AuthPayload, error) {
	privKey, err := ec.PrivateKeyFromHex(b.PrivateHexKey)
	if err != nil {
		return nil, fmt.Errorf("ec - private key from hex op failure: %w", err)
	}
	nonce, err := cryptoutil.RandomHex(32)
	if err != nil {
		return nil, fmt.Errorf("cryptoutil - random hex op failure: %w", err)
	}
	p, err := NewAuthPayloadAccessKey(privKey, b.Body, nonce)
	if err != nil {
		return nil, fmt.Errorf("new auth payload access key op failure: %w", err)
	}
	return p, nil
}

func (b *AuthPayloadBuilder) BuildForXPriv() (*models.AuthPayload, error) {
	xPub, err := bip32.GetExtendedPublicKey(b.ExtendedKey)
	if err != nil {
		return nil, fmt.Errorf("bip32 - get extended public key op failure: %w", err)
	}
	nonce, err := cryptoutil.RandomHex(32)
	if err != nil {
		return nil, fmt.Errorf("cryptoutil - random hex op failure: %w", err)
	}
	cKey, err := cryptoutil.DeriveChildKeyFromHex(b.ExtendedKey, nonce)
	if err != nil {
		return nil, fmt.Errorf("cryptoutil - derive child key op failure: %w", err)
	}
	pKey, err := bip32.GetPrivateKeyFromHDKey(cKey)
	if err != nil {
		return nil, fmt.Errorf("cryptoutil - get private key from HD key failure: %w", err)
	}
	p, err := NewAuthPayloadXPriv(pKey, xPub, b.Body, nonce)
	if err != nil {
		return nil, fmt.Errorf("new auth payload xPriv op failure: %w", err)
	}
	return p, nil
}

func PrivateKeyFromHexOrWIF(s string) (*ec.PrivateKey, error) {
	pk, err1 := ec.PrivateKeyFromWif(s)
	if err1 == nil {
		return pk, nil
	}
	pk, err2 := ec.PrivateKeyFromHex(s)
	if err2 != nil {
		return nil, errors.Join(err1, err2)
	}
	return pk, nil
}

func NewAuthPayloadAccessKey(privKey *ec.PrivateKey, body, nonce string) (*models.AuthPayload, error) {
	hash := cryptoutil.Hash(body)
	ts := time.Now().UnixMilli()
	hex := hex.EncodeToString(privKey.PubKey().SerializeCompressed())
	msg := []byte(hex + hash + nonce + strconv.FormatInt(ts, 10))
	bb, err := bsm.SignMessage(privKey, msg)
	if err != nil {
		return nil, fmt.Errorf("bsm - sign message op failure: %w", err)
	}

	p := models.AuthPayload{
		AuthHash:  hash,
		AuthTime:  ts,
		AuthNonce: nonce,
		AccessKey: hex,
		Signature: base64.StdEncoding.EncodeToString(bb),
	}
	return &p, nil
}

func NewAuthPayloadXPriv(pKey *ec.PrivateKey, xPub, body, nonce string) (*models.AuthPayload, error) {
	hash := cryptoutil.Hash(body)
	ts := time.Now().UnixMilli()
	msg := []byte(xPub + hash + nonce + strconv.FormatInt(ts, 10))
	sign, err := bsm.SignMessage(pKey, msg)
	if err != nil {
		return nil, fmt.Errorf("bsm - sign message op failure: %w", err)
	}

	p := models.AuthPayload{
		XPub:      xPub,
		AuthHash:  hash,
		AuthTime:  ts,
		AuthNonce: nonce,
		Signature: base64.StdEncoding.EncodeToString(sign),
	}
	return &p, nil
}
