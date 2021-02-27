package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

func (i *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	i.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (f *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	f.val = math.Float32frombits(bytes)
}

type ConstantLongInfo struct {
	val int64
}

func (l *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	l.val = int64(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}

func (d *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	d.val = math.Float64frombits(bytes)
}
