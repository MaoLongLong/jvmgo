package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (s *ConstantStringInfo) readInfo(reader *ClassReader) {
	s.stringIndex = reader.readUint16()
}

func (s *ConstantStringInfo) String() string {
	return s.cp.getUtf8(s.stringIndex)
}
