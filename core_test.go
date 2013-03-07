
package curl

import (
	"testing"
	"fmt"
)


func TestVersionInfo(t *testing.T) {
	info := VersionInfo(VERSION_FIRST)
	fmt.Printf("%#v\n", info.Protocols)
}
