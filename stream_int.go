package jsoniter

var digits []uint32

func init() {
	digits = make([]uint32, 1000)
	for i := uint32(0); i < 1000; i++ {
		digits[i] = (((i / 100) + '0') << 16) + ((((i / 10) % 10) + '0') << 8) + i%10 + '0'
		if i < 10 {
			digits[i] += 2 << 24
		} else if i < 100 {
			digits[i] += 1 << 24
		}
	}
}

func (stream *Stream) writeFirstBuf(v uint32) {
	start := v >> 24
	if start == 0 {
		stream.writeTwoBytes(byte(v>>16), byte(v>>8))
	} else if start == 1 {
		stream.writeByte(byte(v >> 8))
	}
	stream.writeByte(byte(v))
}

func (stream *Stream) writeBuf(v uint32) {
	stream.writeThreeBytes(byte(v>>16), byte(v>>8), byte(v))
}

// WriteUint8 write uint8 to stream
func (stream *Stream) WriteUint8(val uint8) {
	stream.writeFirstBuf(digits[val])
}

// WriteInt8 write int8 to stream
func (stream *Stream) WriteInt8(nval int8) {
	var val uint8
	if nval < 0 {
		val = uint8(-nval)
		stream.writeByte('-')
	} else {
		val = uint8(nval)
	}
	stream.writeFirstBuf(digits[val])
}

// WriteUint16 write uint16 to stream
func (stream *Stream) WriteUint16(val uint16) {
	q1 := val / 1000
	if q1 == 0 {
		stream.writeFirstBuf(digits[val])
		return
	}
	r1 := val - q1*1000
	stream.writeFirstBuf(digits[q1])
	stream.writeBuf(digits[r1])
	return
}

// WriteInt16 write int16 to stream
func (stream *Stream) WriteInt16(nval int16) {
	var val uint16
	if nval < 0 {
		val = uint16(-nval)
		stream.writeByte('-')
	} else {
		val = uint16(nval)
	}
	stream.WriteUint16(val)
}

// WriteUint32 write uint32 to stream
func (stream *Stream) WriteUint32(val uint32) {
	q1 := val / 1000
	if q1 == 0 {
		stream.writeFirstBuf(digits[val])
		return
	}
	r1 := val - q1*1000
	q2 := q1 / 1000
	if q2 == 0 {
		stream.writeFirstBuf(digits[q1])
		stream.writeBuf(digits[r1])
		return
	}
	r2 := q1 - q2*1000
	q3 := q2 / 1000
	if q3 == 0 {
		stream.writeFirstBuf(digits[q2])
	} else {
		r3 := q2 - q3*1000
		stream.writeByte(byte(q2 + '0'))
		stream.writeBuf(digits[r3])
	}
	stream.writeBuf(digits[r2])
	stream.writeBuf(digits[r1])
}

// WriteInt32 write int32 to stream
func (stream *Stream) WriteInt32(nval int32) {
	var val uint32
	if nval < 0 {
		val = uint32(-nval)
		stream.writeByte('-')
	} else {
		val = uint32(nval)
	}
	stream.WriteUint32(val)
}

// WriteUint64 write uint64 to stream
func (stream *Stream) WriteUint64(val uint64) {
	q1 := val / 1000
	if q1 == 0 {
		stream.writeFirstBuf(digits[val])
		return
	}
	r1 := val - q1*1000
	q2 := q1 / 1000
	if q2 == 0 {
		stream.writeFirstBuf(digits[q1])
		stream.writeBuf(digits[r1])
		return
	}
	r2 := q1 - q2*1000
	q3 := q2 / 1000
	if q3 == 0 {
		stream.writeFirstBuf(digits[q2])
		stream.writeBuf(digits[r2])
		stream.writeBuf(digits[r1])
		return
	}
	r3 := q2 - q3*1000
	q4 := q3 / 1000
	if q4 == 0 {
		stream.writeFirstBuf(digits[q3])
		stream.writeBuf(digits[r3])
		stream.writeBuf(digits[r2])
		stream.writeBuf(digits[r1])
		return
	}
	r4 := q3 - q4*1000
	q5 := q4 / 1000
	if q5 == 0 {
		stream.writeFirstBuf(digits[q4])
		stream.writeBuf(digits[r4])
		stream.writeBuf(digits[r3])
		stream.writeBuf(digits[r2])
		stream.writeBuf(digits[r1])
		return
	}
	r5 := q4 - q5*1000
	q6 := q5 / 1000
	if q6 == 0 {
		stream.writeFirstBuf(digits[q5])
	} else {
		stream.writeFirstBuf(digits[q6])
		r6 := q5 - q6*1000
		stream.writeBuf(digits[r6])
	}
	stream.writeBuf(digits[r5])
	stream.writeBuf(digits[r4])
	stream.writeBuf(digits[r3])
	stream.writeBuf(digits[r2])
	stream.writeBuf(digits[r1])
}

// WriteInt64 write int64 to stream
func (stream *Stream) WriteInt64(nval int64) {
	var val uint64
	if nval < 0 {
		val = uint64(-nval)
		stream.writeByte('-')
	} else {
		val = uint64(nval)
	}
	stream.WriteUint64(val)
}

// WriteInt write int to stream
func (stream *Stream) WriteInt(val int) {
	stream.WriteInt64(int64(val))
}

// WriteUint write uint to stream
func (stream *Stream) WriteUint(val uint) {
	stream.WriteUint64(uint64(val))
}
