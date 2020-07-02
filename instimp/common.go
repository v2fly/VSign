package instimp

import "github.com/xiaokangwang/VSign/instructions"

func NewVersionIns(version string) instructions.Instruction {
	return NewSimpleFilenameKeyValueInst("", "version", version, true)
}

func NewProjectIns(project string) instructions.Instruction {
	return NewSimpleFilenameKeyValueInst("", "project", project, true)
}
