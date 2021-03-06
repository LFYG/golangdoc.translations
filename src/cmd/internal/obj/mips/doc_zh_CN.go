// +build ingore

package mips

import (
	"cmd/internal/obj"
	"cmd/internal/sys"
	"fmt"
	"log"
	"math"
	"sort"
)

const (
	AABSD = obj.ABaseMIPS64 + obj.A_ARCHSPECIFIC + iota
	AABSF
	AABSW
	AADD
	AADDD
	AADDF
	AADDU
	AADDW
	AAND
	ABEQ
	ABFPF
	ABFPT
	ABGEZ
	ABGEZAL
	ABGTZ
	ABLEZ
	ABLTZ
	ABLTZAL
	ABNE
	ABREAK
	ACMPEQD
	ACMPEQF
	ACMPGED
	ACMPGEF
	ACMPGTD
	ACMPGTF
	ADIV
	ADIVD
	ADIVF
	ADIVU
	ADIVW
	AGOK
	ALUI
	AMOVB
	AMOVBU
	AMOVD
	AMOVDF
	AMOVDW
	AMOVF
	AMOVFD
	AMOVFW
	AMOVH
	AMOVHU
	AMOVW
	AMOVWD
	AMOVWF
	AMOVWL
	AMOVWR
	AMUL
	AMULD
	AMULF
	AMULU
	AMULW
	ANEGD
	ANEGF
	ANEGW
	ANOR
	AOR
	AREM
	AREMU
	ARFE
	ASGT
	ASGTU
	ASLL
	ASRA
	ASRL
	ASUB
	ASUBD
	ASUBF
	ASUBU
	ASUBW
	ASYSCALL
	ATLBP
	ATLBR
	ATLBWI
	ATLBWR
	AWORD
	AXOR

	//  64-bit
	AMOVV
	AMOVVL
	AMOVVR
	ASLLV
	ASRAV
	ASRLV
	ADIVV
	ADIVVU
	AREMV
	AREMVU
	AMULV
	AMULVU
	AADDV
	AADDVU
	ASUBV
	ASUBVU

	//  64-bit FP
	ATRUNCFV
	ATRUNCDV
	ATRUNCFW
	ATRUNCDW
	AMOVWU
	AMOVFV
	AMOVDV
	AMOVVF
	AMOVVD
	ALAST

	// aliases
	AJMP = obj.AJMP
	AJAL = obj.ACALL
	ARET = obj.ARET
)

const (
	BIG = 32766
)

const (
	C_NONE = iota
	C_REG
	C_FREG
	C_FCREG
	C_MREG /* special processor register*/
	C_HI
	C_LO
	C_ZCON
	C_SCON /* 16 bit signed*/
	C_UCON /* 32 bit signed, low 16 bits 0*/
	C_ADD0CON
	C_AND0CON
	C_ADDCON /* -0x8000 <= v < 0*/
	C_ANDCON /* 0 < v <= 0xFFFF*/
	C_LCON   /* other 32*/
	C_DCON   /* other 64 (could subdivide further)*/
	C_SACON  /* $n(REG) where n <= int16*/
	C_SECON
	C_LACON /* $n(REG) where int16 < n <= int32*/
	C_LECON
	C_DACON /* $n(REG) where int32 < n*/
	C_STCON /* $tlsvar*/
	C_SBRA
	C_LBRA
	C_SAUTO
	C_LAUTO
	C_SEXT
	C_LEXT
	C_ZOREG
	C_SOREG
	C_LOREG
	C_GOK
	C_ADDR
	C_TLS
	C_TEXTSIZE
	C_NCLASS /* must be the last*/
)

const (
	E_HILO  = 1 << 0
	E_FCR   = 1 << 1
	E_MCR   = 1 << 2
	E_MEM   = 1 << 3
	E_MEMSP = 1 << 4 /* uses offset and size*/
	E_MEMSB = 1 << 5 /* uses offset and size*/
	ANYMEM  = E_MEM | E_MEMSP | E_MEMSB

	// DELAY = LOAD|BRANCH|FCMP
	DELAY = BRANCH /* only schedule branch*/
)

const (
	//  mark flags
	FOLL    = 1 << 0
	LABEL   = 1 << 1
	LEAF    = 1 << 2
	SYNC    = 1 << 3
	BRANCH  = 1 << 4
	LOAD    = 1 << 5
	FCMP    = 1 << 6
	NOSCHED = 1 << 7
	NSCHED  = 20
)

const (
	FuncAlign = 8
)

//  * mips 64

// * mips 64
const (
	NSNAME = 8
	NSYM   = 50
	NREG   = 32 /* number of general registers*/
	NFREG  = 32 /* number of floating point registers*/
)

const (
	REG_R0 = obj.RBaseMIPS64 + iota
	REG_R1
	REG_R2
	REG_R3
	REG_R4
	REG_R5
	REG_R6
	REG_R7
	REG_R8
	REG_R9
	REG_R10
	REG_R11
	REG_R12
	REG_R13
	REG_R14
	REG_R15
	REG_R16
	REG_R17
	REG_R18
	REG_R19
	REG_R20
	REG_R21
	REG_R22
	REG_R23
	REG_R24
	REG_R25
	REG_R26
	REG_R27
	REG_R28
	REG_R29
	REG_R30
	REG_R31
	REG_F0
	REG_F1
	REG_F2
	REG_F3
	REG_F4
	REG_F5
	REG_F6
	REG_F7
	REG_F8
	REG_F9
	REG_F10
	REG_F11
	REG_F12
	REG_F13
	REG_F14
	REG_F15
	REG_F16
	REG_F17
	REG_F18
	REG_F19
	REG_F20
	REG_F21
	REG_F22
	REG_F23
	REG_F24
	REG_F25
	REG_F26
	REG_F27
	REG_F28
	REG_F29
	REG_F30
	REG_F31
	REG_HI
	REG_LO

	// co-processor 0 control registers
	REG_M0
	REG_M1
	REG_M2
	REG_M3
	REG_M4
	REG_M5
	REG_M6
	REG_M7
	REG_M8
	REG_M9
	REG_M10
	REG_M11
	REG_M12
	REG_M13
	REG_M14
	REG_M15
	REG_M16
	REG_M17
	REG_M18
	REG_M19
	REG_M20
	REG_M21
	REG_M22
	REG_M23
	REG_M24
	REG_M25
	REG_M26
	REG_M27
	REG_M28
	REG_M29
	REG_M30
	REG_M31

	// FPU control registers
	REG_FCR0
	REG_FCR1
	REG_FCR2
	REG_FCR3
	REG_FCR4
	REG_FCR5
	REG_FCR6
	REG_FCR7
	REG_FCR8
	REG_FCR9
	REG_FCR10
	REG_FCR11
	REG_FCR12
	REG_FCR13
	REG_FCR14
	REG_FCR15
	REG_FCR16
	REG_FCR17
	REG_FCR18
	REG_FCR19
	REG_FCR20
	REG_FCR21
	REG_FCR22
	REG_FCR23
	REG_FCR24
	REG_FCR25
	REG_FCR26
	REG_FCR27
	REG_FCR28
	REG_FCR29
	REG_FCR30
	REG_FCR31
	REG_LAST    = REG_FCR31 // the last defined register
	REG_SPECIAL = REG_M0
	REGZERO     = REG_R0 /* set to zero*/
	REGSP       = REG_R29
	REGSB       = REG_R28
	REGLINK     = REG_R31
	REGRET      = REG_R1
	REGARG      = -1      /* -1 disables passing the first argument in register*/
	REGRT1      = REG_R1  /* reserved for runtime, duffzero and duffcopy*/
	REGRT2      = REG_R2  /* reserved for runtime, duffcopy*/
	REGCTXT     = REG_R22 /* context for closures*/
	REGG        = REG_R30 /* G*/
	REGTMP      = REG_R23 /* used by the linker*/
	FREGRET     = REG_F0
	FREGZERO    = REG_F24 /* both float and double*/
	FREGHALF    = REG_F26 /* double*/
	FREGONE     = REG_F28 /* double*/
	FREGTWO     = REG_F30 /* double*/
)

var Anames = []string{obj.A_ARCHSPECIFIC: "ABSD", "ABSF", "ABSW", "ADD", "ADDD", "ADDF", "ADDU", "ADDW", "AND", "BEQ", "BFPF", "BFPT", "BGEZ", "BGEZAL", "BGTZ", "BLEZ", "BLTZ", "BLTZAL", "BNE", "BREAK", "CMPEQD", "CMPEQF", "CMPGED", "CMPGEF", "CMPGTD", "CMPGTF", "DIV", "DIVD", "DIVF", "DIVU", "DIVW", "GOK", "LUI", "MOVB", "MOVBU", "MOVD", "MOVDF", "MOVDW", "MOVF", "MOVFD", "MOVFW", "MOVH", "MOVHU", "MOVW", "MOVWD", "MOVWF", "MOVWL", "MOVWR", "MUL", "MULD", "MULF", "MULU", "MULW", "NEGD", "NEGF", "NEGW", "NOR", "OR", "REM", "REMU", "RFE", "SGT", "SGTU", "SLL", "SRA", "SRL", "SUB", "SUBD", "SUBF", "SUBU", "SUBW", "SYSCALL", "TLBP", "TLBR", "TLBWI", "TLBWR", "WORD", "XOR", "MOVV", "MOVVL", "MOVVR", "SLLV", "SRAV", "SRLV", "DIVV", "DIVVU", "REMV", "REMVU", "MULV", "MULVU", "ADDV", "ADDVU", "SUBV", "SUBVU", "TRUNCFV", "TRUNCDV", "TRUNCFW", "TRUNCDW", "MOVWU", "MOVFV", "MOVDV", "MOVVF", "MOVVD", "LAST"}

var Linkmips64 = obj.LinkArch{Arch: sys.ArchMIPS64, Preprocess: preprocess, Assemble: span0, Follow: follow, Progedit: progedit}

var Linkmips64le = obj.LinkArch{Arch: sys.ArchMIPS64LE, Preprocess: preprocess, Assemble: span0, Follow: follow, Progedit: progedit}

type Dep struct {
}

type Optab struct {
}

type Sch struct {
}

func BCOND(x uint32, y uint32) uint32

func DRconv(a int) string

func FPD(x uint32, y uint32) uint32

func FPF(x uint32, y uint32) uint32

func FPV(x uint32, y uint32) uint32

func FPW(x uint32, y uint32) uint32

func MMU(x uint32, y uint32) uint32

func OP(x uint32, y uint32) uint32

func OP_FRRR(op uint32, r1 uint32, r2 uint32, r3 uint32) uint32

func OP_IRR(op uint32, i uint32, r2 uint32, r3 uint32) uint32

func OP_JMP(op uint32, i uint32) uint32

func OP_RRR(op uint32, r1 uint32, r2 uint32, r3 uint32) uint32

func OP_SRR(op uint32, s uint32, r2 uint32, r3 uint32) uint32

func Rconv(r int) string

func SP(x uint32, y uint32) uint32

