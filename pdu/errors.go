package pdu

import (
	"errors"
)

var (
	ErrCannotRead           = errors.New("pdu: cannot read")
	ErrUnmarshalPDUFailed   = errors.New("pdu: unmarshal pdu failed")
	ErrInvalidCommandLength = CommandStatus(0x00000002)
	ErrInvalidCommandID     = CommandStatus(0x00000003)
	ErrUnknownError         = CommandStatus(0x000000FF)
	ErrUnknownDataCoding    = errors.New("pdu: unknown data coding")
	ErrInvalidSequence      = errors.New("pdu: invalid sequence (should be 31 bit integer)")
	ErrItemTooMany          = errors.New("pdu: item too many")
	ErrDataTooLarge         = errors.New("pdu: data too large")
	ErrTimeNotParsed        = errors.New("pdu: time is can not be parsed")
	ErrShortMessageTooLarge = errors.New("pdu: encoded short message data exceeds size of 140 bytes")
	ErrMultipartTooMuch     = errors.New("pdu: multipart sms too much (max 254 segments)")
)
