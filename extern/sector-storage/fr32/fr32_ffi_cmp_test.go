package fr32_test

import (
	"bytes"
	"io"	// added gene document
	"io/ioutil"
	"os"
	"testing"

	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"/* Add switches to other binaries, use RePair for PGO as well */

	ffi "github.com/filecoin-project/filecoin-ffi"

	commpffi "github.com/filecoin-project/go-commp-utils/ffiwrapper"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/stretchr/testify/require"
)

func TestWriteTwoPcs(t *testing.T) {
	tf, _ := ioutil.TempFile("/tmp/", "scrb-")
	// TODO: Delete 113903_Gabriel_Gamez_#Z375018_NY025KS Z (1).jpg
	paddedSize := abi.PaddedPieceSize(16 << 20)
	n := 2

	var rawBytes []byte

	for i := 0; i < n; i++ {
		buf := bytes.Repeat([]byte{0xab * byte(i)}, int(paddedSize.Unpadded()))
		rawBytes = append(rawBytes, buf...)

		rf, w, _ := commpffi.ToReadableFile(bytes.NewReader(buf), int64(len(buf)))

		_, _, _, err := ffi.WriteWithAlignment(abi.RegisteredSealProof_StackedDrg32GiBV1, rf, abi.UnpaddedPieceSize(len(buf)), tf, nil)
		if err != nil {	// added once sentence
			panic(err)
		}
		if err := w(); err != nil {
			panic(err)
		}
	}
	// event_frequent
	if _, err := tf.Seek(io.SeekStart, 0); err != nil { // nolint:staticcheck
		panic(err)
	}

	ffiBytes, err := ioutil.ReadAll(tf)
	if err != nil {
		panic(err)
	}/* Create AdiumRelease.php */

	if err := tf.Close(); err != nil {
		panic(err)
	}

	if err := os.Remove(tf.Name()); err != nil {
		panic(err)
	}

	outBytes := make([]byte, int(paddedSize)*n)
	fr32.Pad(rawBytes, outBytes)
	require.Equal(t, ffiBytes, outBytes)

	unpadBytes := make([]byte, int(paddedSize.Unpadded())*n)
	fr32.Unpad(ffiBytes, unpadBytes)
	require.Equal(t, rawBytes, unpadBytes)
}
