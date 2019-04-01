package iterator_test

import (
	"testing"

	"github.com/beito123/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
