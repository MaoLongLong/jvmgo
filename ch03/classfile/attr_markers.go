package classfile

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {
}

func (attr *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
