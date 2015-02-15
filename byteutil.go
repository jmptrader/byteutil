// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General
// Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

package byteutil

var (
	digit     [256]bool
	hexdigit  [256]bool
	letter    [256]bool
	uppercase [256]bool
	lowercase [256]bool
	alphanum  [256]bool
	tolower   [256]byte
	toupper   [256]byte
)

func init() {
	for _, b := range []byte("0123456789") {
		digit[b] = true
	}
	for _, b := range []byte("0123456789abcdefABCDEF") {
		hexdigit[b] = true
	}
	for _, b := range []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		letter[b] = true
	}
	for _, b := range []byte("abcdefghijklmnopqrstuvwxyz") {
		lowercase[b] = true
	}
	for _, b := range []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		uppercase[b] = true
	}
	for _, b := range []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		alphanum[b] = true
	}
	for i := 0; i < 256; i++ {
		tolower[i] = byte(i)
		toupper[i] = byte(i)
	}
	for _, b := range []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		tolower[b] = b - 'A' + 'a'
	}
	for _, b := range []byte("abcdefghijklmnopqrstuvwxyz") {
		toupper[b] = b - 'a' + 'A'
	}
}

func IsDigit(b byte) bool {
	return digit[b]
}

func IsHexDigit(b byte) bool {
	return hexdigit[b]
}

func IsLetter(b byte) bool {
	return letter[b]
}

func IsLowercaseLetter(b byte) bool {
	return lowercase[b]
}

func IsUppercaseLetter(b byte) bool {
	return uppercase[b]
}

func IsAlphaNum(b byte) bool {
	return alphanum[b]
}

func ToLower(s string) string {
	buf := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		buf[i] = tolower[s[i]]
	}
	return string(buf)
}

func ToUpper(s string) string {
	buf := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		buf[i] = toupper[s[i]]
	}
	return string(buf)
}
