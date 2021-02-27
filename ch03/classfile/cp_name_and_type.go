package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (nat *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	nat.nameIndex = reader.readUint16()
	nat.descriptorIndex = reader.readUint16()
}
