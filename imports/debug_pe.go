// this file was generated by gomacro command: import "debug/pe"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"debug/pe"
)

func init() {
	Packages["debug/pe"] = Package{
	Binds: map[string]Value{
		"COFFSymbolSize":	ValueOf(pe.COFFSymbolSize),
		"IMAGE_FILE_MACHINE_AM33":	ValueOf(pe.IMAGE_FILE_MACHINE_AM33),
		"IMAGE_FILE_MACHINE_AMD64":	ValueOf(pe.IMAGE_FILE_MACHINE_AMD64),
		"IMAGE_FILE_MACHINE_ARM":	ValueOf(pe.IMAGE_FILE_MACHINE_ARM),
		"IMAGE_FILE_MACHINE_EBC":	ValueOf(pe.IMAGE_FILE_MACHINE_EBC),
		"IMAGE_FILE_MACHINE_I386":	ValueOf(pe.IMAGE_FILE_MACHINE_I386),
		"IMAGE_FILE_MACHINE_IA64":	ValueOf(pe.IMAGE_FILE_MACHINE_IA64),
		"IMAGE_FILE_MACHINE_M32R":	ValueOf(pe.IMAGE_FILE_MACHINE_M32R),
		"IMAGE_FILE_MACHINE_MIPS16":	ValueOf(pe.IMAGE_FILE_MACHINE_MIPS16),
		"IMAGE_FILE_MACHINE_MIPSFPU":	ValueOf(pe.IMAGE_FILE_MACHINE_MIPSFPU),
		"IMAGE_FILE_MACHINE_MIPSFPU16":	ValueOf(pe.IMAGE_FILE_MACHINE_MIPSFPU16),
		"IMAGE_FILE_MACHINE_POWERPC":	ValueOf(pe.IMAGE_FILE_MACHINE_POWERPC),
		"IMAGE_FILE_MACHINE_POWERPCFP":	ValueOf(pe.IMAGE_FILE_MACHINE_POWERPCFP),
		"IMAGE_FILE_MACHINE_R4000":	ValueOf(pe.IMAGE_FILE_MACHINE_R4000),
		"IMAGE_FILE_MACHINE_SH3":	ValueOf(pe.IMAGE_FILE_MACHINE_SH3),
		"IMAGE_FILE_MACHINE_SH3DSP":	ValueOf(pe.IMAGE_FILE_MACHINE_SH3DSP),
		"IMAGE_FILE_MACHINE_SH4":	ValueOf(pe.IMAGE_FILE_MACHINE_SH4),
		"IMAGE_FILE_MACHINE_SH5":	ValueOf(pe.IMAGE_FILE_MACHINE_SH5),
		"IMAGE_FILE_MACHINE_THUMB":	ValueOf(pe.IMAGE_FILE_MACHINE_THUMB),
		"IMAGE_FILE_MACHINE_UNKNOWN":	ValueOf(pe.IMAGE_FILE_MACHINE_UNKNOWN),
		"IMAGE_FILE_MACHINE_WCEMIPSV2":	ValueOf(pe.IMAGE_FILE_MACHINE_WCEMIPSV2),
		"NewFile":	ValueOf(pe.NewFile),
		"Open":	ValueOf(pe.Open),
	},
	Types: map[string]Type{
		"COFFSymbol":	TypeOf((*pe.COFFSymbol)(nil)).Elem(),
		"DataDirectory":	TypeOf((*pe.DataDirectory)(nil)).Elem(),
		"File":	TypeOf((*pe.File)(nil)).Elem(),
		"FileHeader":	TypeOf((*pe.FileHeader)(nil)).Elem(),
		"FormatError":	TypeOf((*pe.FormatError)(nil)).Elem(),
		"ImportDirectory":	TypeOf((*pe.ImportDirectory)(nil)).Elem(),
		"OptionalHeader32":	TypeOf((*pe.OptionalHeader32)(nil)).Elem(),
		"OptionalHeader64":	TypeOf((*pe.OptionalHeader64)(nil)).Elem(),
		"Reloc":	TypeOf((*pe.Reloc)(nil)).Elem(),
		"Section":	TypeOf((*pe.Section)(nil)).Elem(),
		"SectionHeader":	TypeOf((*pe.SectionHeader)(nil)).Elem(),
		"SectionHeader32":	TypeOf((*pe.SectionHeader32)(nil)).Elem(),
		"StringTable":	TypeOf((*pe.StringTable)(nil)).Elem(),
		"Symbol":	TypeOf((*pe.Symbol)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	} }
}
