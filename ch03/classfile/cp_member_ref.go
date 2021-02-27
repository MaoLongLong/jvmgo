package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (m *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	m.classIndex = reader.readUint16()
	m.nameAndTypeIndex = reader.readUint16()
}

func (m *ConstantMemberrefInfo) ClassName() string {
	return m.cp.getClassName(m.classIndex)
}

func (m *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return m.cp.getNameAndType(m.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
