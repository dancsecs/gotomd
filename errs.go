package main

import "errors"

// Exported errors.
var (
	ErrInvalidRelativeDir  = errors.New("invalid relative directory")
	ErrInvalidDirectory    = errors.New("invalid directory")
	ErrMissingAction       = errors.New("missing action")
	ErrNoTestToRun         = errors.New("no tests to run")
	ErrNoPackageToRun      = errors.New("no package to run")
	ErrUnknownObject       = errors.New("unknown package object")
	ErrMissingHeaderLine   = errors.New("missing blank header line")
	ErrTagOutOfSequence    = errors.New("out of sequence: End before begin")
	ErrUnknownCommand      = errors.New("unknown command")
	ErrUnexpectedExtension = errors.New("unexpected file extension")
	ErrInvalidDefPerm      = errors.New("invalid default perm")
	ErrInvalidOptionsRC    = errors.New("invalid option mix replace and clean")
	ErrInvalidOutputDir    = errors.New("invalid output directory")
)
