package gather

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math/rand"
	math_rand "math/rand"
)

// init initializes a generic random seed for any
// functions that use random number generation
func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

// randomStringNoConsecutive creates a random string
// from upper case characters; identical consecutive
// letters are not allowed.
func randomStringNoConsecutive(l int) string {
	bytes := make([]byte, l)
	var temp int
	for i := 0; i < l; {
		if randInt(65, 90) != temp {
			temp = randInt('A', 'Z')
			bytes[i] = byte(temp)
			i++
		}
	}
	return string(bytes)
}

// randomStringNoConsecutive creates a random string
// from upper case characters
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt('A', 'Z'))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomFromPool(l int, pool string) string {
	return ""
}
