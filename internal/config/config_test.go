package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/openlibrecommunity/olcrtc/internal/app/session"
)

func TestLoadAndApply(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "olcrtc.yaml")
	body := `
mode: srv
link: direct
auth:
  provider: wbstream
room:
  id: r1
  client_id: c1
crypto:
  key: deadbeef
net:
  transport: datachannel
  dns: 1.1.1.1:53
socks:
  host: 127.0.0.1
  port: 1080
  user: u
  pass: p
vp8:
  fps: 25
  batch_size: 4
gen:
  amount: 3
debug: true
`
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}

	f, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if f.Mode != "srv" || f.Auth.Provider != "wbstream" || f.Room.ID != "r1" || f.Crypto.Key != "deadbeef" {
		t.Fatalf("unexpected file: %+v", f)
	}

	got := Apply(session.Config{}, f)
	if got.Mode != "srv" || got.Link != "direct" || got.Auth != "wbstream" ||
		got.RoomID != "r1" || got.ClientID != "c1" || got.KeyHex != "deadbeef" ||
		got.Transport != "datachannel" || got.DNSServer != "1.1.1.1:53" ||
		got.SOCKSHost != "127.0.0.1" || got.SOCKSPort != 1080 ||
		got.SOCKSUser != "u" || got.SOCKSPass != "p" ||
		got.VP8FPS != 25 || got.VP8BatchSize != 4 || got.Amount != 3 {
		t.Fatalf("Apply produced wrong config: %+v", got)
	}
}

func TestApplyCLIWins(t *testing.T) {
	cli := session.Config{
		Mode:      "cnc",
		KeyHex:    "from-cli",
		SOCKSPort: 9999,
	}
	f := File{
		Mode:   "srv",
		Crypto: Crypto{Key: "from-yaml"},
		SOCKS:  SOCKS{Port: 1234, Host: "0.0.0.0"},
	}
	got := Apply(cli, f)
	if got.Mode != "cnc" {
		t.Errorf("Mode: got %q, want cnc (CLI wins)", got.Mode)
	}
	if got.KeyHex != "from-cli" {
		t.Errorf("KeyHex: got %q, want from-cli (CLI wins)", got.KeyHex)
	}
	if got.SOCKSPort != 9999 {
		t.Errorf("SOCKSPort: got %d, want 9999 (CLI wins)", got.SOCKSPort)
	}
	if got.SOCKSHost != "0.0.0.0" {
		t.Errorf("SOCKSHost: got %q, want 0.0.0.0 (YAML fills empty CLI)", got.SOCKSHost)
	}
}

func TestLoadMissing(t *testing.T) {
	_, err := Load(filepath.Join(t.TempDir(), "nope.yaml"))
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}
