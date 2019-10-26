// Copyright (c) 2015-2016 The Utopia developers
// Copyright (c) 2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chainhash

import (
	x19r "github.com/UtopiaCoinOrg/ucd/chaincfg/chainhash/x19rhash"
	"github.com/dchest/blake256"
	"sync"
)

var x19rHashMtx   sync.Mutex
// HashFunc calculates the hash of the supplied bytes.
// TODO(jcv) Should modify blake256 so it has the same interface as blake2
// and sha256 so these function can look more like btcsuite.  Then should
// try to get it to the upstream blake256 repo
func HashFunc(b []byte) [blake256.Size]byte {
	var outB [blake256.Size]byte
	copy(outB[:], HashB(b))
	return outB
}

// HashB calculates hash(b) and returns the resulting bytes.
func HashB(b []byte) []byte {
	a := blake256.New()
	a.Write(b)
	out := a.Sum(nil)
	return out
}

// HashH calculates hash(b) and returns the resulting bytes as a Hash.
func HashH(b []byte) Hash {
	return Hash(HashFunc(b))
}

func HashHx19(b []byte)Hash  {

	x19rHashMtx.Lock()
	defer x19rHashMtx.Unlock()

	//height := binary.LittleEndian.Uint32(b[36:])
	//fmt.Println(hex.EncodeToString(b))
	//hashTest := x19r.X19r_Sum256(string(b))
	//var testHash Hash
	//testHash.SetBytes(hashTest[:])
	//fmt.Println(testHash.String())
	//log.Infof( "height = %d,  %s", height, hex.EncodeToString(b))
	//log.Infof("height = %d,  %s", height, testHash.String())
	return x19r.X19r_Sum256(string(b))
}

// HashBlockSize is the block size of the hash algorithm in bytes.
const HashBlockSize = blake256.BlockSize
