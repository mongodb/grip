package slogger

type Level uint8

const (
	OFF Level = iota
	DEBUG
	INFO
	WARN
	ERROR
)
