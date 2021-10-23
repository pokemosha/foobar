package main

import (
	"bytes"
	"testing"
)

func TestFooBar_CorrectAnswer(t *testing.T) {
	ibuf := bytes.NewBufferString("4\n")
	obuf := bytes.NewBufferString("")
	err := foobar(obuf, ibuf)
	if err != nil {
		t.Error(err)
	}
	if obuf.String() != "1 Foo\n2 Foo\n3 Foo Bar\n4 Foo\n" {
		t.Errorf("wrong output: %s", obuf.String())
	}
}

func TestFooBar_NotString(t *testing.T) {
	ibuf := bytes.NewBufferString("not a number\n")
	obuf := bytes.NewBufferString("")
	err := foobar(obuf, ibuf)
	if err == nil {
		t.Error(err)
	}
}

func TestFooBar_NotZero(t *testing.T) {
	ibuf := bytes.NewBufferString("0\n")
	obuf := bytes.NewBufferString("")
	err := foobar(obuf, ibuf)
	if err == nil {
		t.Error(err)
	}
}
