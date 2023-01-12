// Package binary handles embedding of the iPXE binaries.
package binary

// embed lib does the work of embedding the on disk iPXE binaries.
import _ "embed"

// IpxeEFI is the UEFI iPXE binary for x86 architectures.
//go:embed ipxe.efi
var IpxeEFI []byte

// Undionly is the BIOS iPXE binary for x86 architectures.
//go:embed undionly.kpxe
var Undionly []byte

// SNP is the UEFI iPXE binary for ARM architectures.
//go:embed snp.efi
var SNP []byte

// MagicString is included in each iPXE binary within the embedded script. It
// can be overwritten to change the behavior at startup.
var MagicString = []byte(`#a8b7e61f1075c37a793f2f92cee89f7bba00c4a8d7842ce3d40b5889032d8881
#ddd16a4fc4926ecefdfb6941e33c44ed3647133638f5e84021ea44d3152e7f97`)

// Files is the mapping to the embedded iPXE binaries.
var Files = map[string][]byte{
	"undionly.kpxe": Undionly,
	"ipxe.efi":      IpxeEFI,
	"snp.efi":       SNP,
}
