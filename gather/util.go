package gather

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

// init initializes a generic random seed for any
// functions that use random number generation
func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

// RandomStringNoConsecutive creates a random string
// from upper case characters; identical consecutive
// letters are not allowed.
func RandomStringNoConsecutive(l int) string {
	bytes := make([]byte, l)
	var temp int
	for i := 0; i < l; {
		if RandInt(65, 90) != temp {
			temp = RandInt('A', 'Z')
			bytes[i] = byte(temp)
			i++
		}
	}
	return string(bytes)
}

// randomStringNoConsecutive creates a random string
// from upper case characters
func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt('A', 'Z'))
	}
	return string(bytes)
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func RandomFromPool(l int, pool string) string {
	return ""
}
