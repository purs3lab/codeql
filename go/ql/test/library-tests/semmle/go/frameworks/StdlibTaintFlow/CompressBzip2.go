// Code generated by https://github.com/gagliardetto/codebox. DO NOT EDIT.

package main

import (
	"compress/bzip2"
	"io"
)

func TaintStepTest_CompressBzip2NewReader_B0I0O0(sourceCQL interface{}) interface{} {
	fromReader656 := sourceCQL.(io.Reader)
	intoReader414 := bzip2.NewReader(fromReader656)
	return intoReader414
}

func RunAllTaints_CompressBzip2() {
	{
		source := newSource(0)
		out := TaintStepTest_CompressBzip2NewReader_B0I0O0(source)
		sink(0, out)
	}
}