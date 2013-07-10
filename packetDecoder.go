package kafka

type packetDecoder interface {
	getInt16() (int16, error)
	getInt32() (int32, error)
	getError() (KError, error)
	getString() (*string, error)
	getBytes() (*[]byte, error)
	getArrayCount() (int, error)
}
