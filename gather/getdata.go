package gather

import (
	"math/rand"
	"time"
)

type gather struct{}

// init initializes a generic random seed for any
// functions that do not supply their own
func init() { rand.Seed(time.Now().UTC().UnixNano()) }

func get() (gather, error) {

}
