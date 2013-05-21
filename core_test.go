package curl

import (
	"testing"
)

func TestVersionInfo(t *testing.T) {
	info := VersionInfo(VERSION_FIRST)
	expectedProtocols := []string{"dict", "file", "ftp", "ftps", "gopher", "http", "https", "imap", "imaps", "ldap", "ldaps", "pop3", "pop3s", "rtmp", "rtsp", "smtp", "smtps", "telnet", "tftp"}
	protocols := info.Protocols
	for index, protocol := range protocols {
		expectedProtocol := expectedProtocols[index]
		if protocol != expectedProtocol {
			t.Errorf("protocol should be %v and is %v.", expectedProtocol, protocol)
		}
	}
}
