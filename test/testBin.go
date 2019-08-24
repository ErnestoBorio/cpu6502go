package main

import (
  . "github.com/petruza/cpu6502/cpu6502"
  _ "fmt"
  "os"
  "io/ioutil"
)

var mem []byte

func main() {
  var err error
	mem, err = ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	cpu := new(Cpu).Init()
	cpu.HookMemoryReader(0, 0xFFFF, memRead)
	cpu.HookMemoryWriter(0, 0xFFFF, memWrite)
}


func memRead(address uint16) byte {
	return mem[address]
}

func memWrite(address uint16, value byte) {
	mem[address] = value
}


// Instruction mnemonic for each opcode
var mnemonics = [0x100] string {
//  0      1      2      3      4      5      6      7      8      9      A      B      C      D      E      F
  "BRK", "ORA", "...", "...", "...", "ORA", "ASL", "...", "PHP", "ORA", "ASL", "...", "...", "ORA", "ASL", "...", // 00
  "BPL", "ORA", "...", "...", "...", "ORA", "ASL", "...", "CLC", "ORA", "...", "...", "...", "ORA", "ASL", "...", // 10
  "JSR", "AND", "...", "...", "BIT", "AND", "ROL", "...", "PLP", "AND", "ROL", "...", "BIT", "AND", "ROL", "...", // 20
  "BMI", "AND", "...", "...", "...", "AND", "ROL", "...", "SEC", "AND", "...", "...", "...", "AND", "ROL", "...", // 30
  "RTI", "EOR", "...", "...", "...", "EOR", "LSR", "...", "PHA", "EOR", "LSR", "...", "JMP", "EOR", "LSR", "...", // 40
  "BVC", "EOR", "...", "...", "...", "EOR", "LSR", "...", "CLI", "EOR", "...", "...", "...", "EOR", "LSR", "...", // 50
  "RTS", "ADC", "...", "...", "...", "ADC", "ROR", "...", "PLA", "ADC", "ROR", "...", "JMP", "ADC", "ROR", "...", // 60
  "BVS", "ADC", "...", "...", "...", "ADC", "ROR", "...", "SEI", "ADC", "...", "...", "...", "ADC", "ROR", "...", // 70
  "...", "STA", "...", "...", "STY", "STA", "STX", "...", "DEY", "...", "TXA", "...", "STY", "STA", "STX", "...", // 80
  "BCC", "STA", "...", "...", "STY", "STA", "STX", "...", "TYA", "STA", "TXS", "...", "...", "STA", "...", "...", // 90
  "LDY", "LDA", "LDX", "...", "LDY", "LDA", "LDX", "...", "TAY", "LDA", "TAX", "...", "LDY", "LDA", "LDX", "...", // A0
  "BCS", "LDA", "...", "...", "LDY", "LDA", "LDX", "...", "CLV", "LDA", "TSX", "...", "LDY", "LDA", "LDX", "...", // B0
  "CPY", "CMP", "...", "...", "CPY", "CMP", "DEC", "...", "INY", "CMP", "DEX", "...", "CPY", "CMP", "DEC", "...", // C0
  "BNE", "CMP", "...", "...", "...", "CMP", "DEC", "...", "CLD", "CMP", "...", "...", "...", "CMP", "DEC", "...", // D0
  "CPX", "SBC", "...", "...", "CPX", "SBC", "INC", "...", "INX", "SBC", "NOP", "...", "CPX", "SBC", "INC", "...", // E0
  "BEQ", "SBC", "...", "...", "...", "SBC", "INC", "...", "SED", "SBC", "...", "...", "...", "SBC", "INC", "..."} // F0

// Addressing modes enum
const (
	// 1 byte instructions
	_imp = 1
	_acu = 2
	_imm = 3
	// 2 bytes instructions
	_2bytes = 10
	_zrp = 11
	_zpx = 12
	_zpy = 13
	_iix = 14
	_iiy = 15
	_rel = 16
	// 3 bytes instructions
	_3bytes = 20
	_abs = 21
	_abx = 22
	_aby = 23
	_ind = 24
)

// Addressing modes for each opcode
var addressing = [0x100] byte {
//  0     1    2     3      4     5     6    7      8     9     A    B      C     D     E    F
  _imp, _iix,  0  ,  0  , _zrp, _zrp, _zrp,  0  , _imp, _imm, _acu,  0  , _abs, _abs, _abs,  0, // 00
  _rel, _iiy,  0  ,  0  , _zpx, _zpx, _zpx,  0  , _imp, _aby, _imp,  0  , _abx, _abx, _abx,  0, // 10
  _abs, _iix,  0  ,  0  , _zrp, _zrp, _zrp,  0  , _imp, _imm, _acu,  0  , _abs, _abs, _abs,  0, // 20
  _rel, _iiy,  0  ,  0  , _zpx, _zpx, _zpx,  0  , _imp, _aby, _imp,  0  , _abx, _abx, _abx,  0, // 30
  _imp, _iix,  0  ,  0  , _zrp, _zrp, _zrp,  0  , _imp, _imm, _acu,  0  , _abs, _abs, _abs,  0, // 40
  _rel, _iiy,  0  ,  0  , _zpx, _zpx, _zpx,  0  , _imp, _aby, _imp,  0  , _abx, _abx, _abx,  0, // 50
  _imp, _iix,  0  ,  0  , _zrp, _zrp, _zrp,  0  , _imp, _imm, _acu,  0  , _ind, _abs, _abs,  0, // 60
  _rel, _iiy,  0  ,  0  , _zpx, _zpx, _zpx,  0  , _imp, _aby, _imp,  0  , _abx, _abx, _abx,  0, // 70
  _imm, _iix, _imm,  0  , _zrp, _zrp, _zrp,  0  , _imp, _imm, _imp,  0  , _abs, _abs, _abs,  0, // 80
  _rel, _iiy,  0  ,  0  , _zpx, _zpx, _zpy,  0  , _imp, _aby, _imp,  0  ,  0  , _abx,  0  ,  0, // 90
  _imm, _iix, _imm,  0  , _zrp, _zrp, _zrp,  0  , _imp, _imm, _imp,  0  , _abs, _abs, _abs,  0, // A0
  _rel, _iiy,  0  ,  0  , _zpx, _zpx, _zpy,  0  , _imp, _aby, _imp,  0  , _abx, _abx, _aby,  0, // B0
  _imm, _iix, _imm,  0  , _zrp, _zrp, _zrp,  0  , _imp, _imm, _imp,  0  , _abs, _abs, _abs,  0, // C0
  _rel, _iiy,  0  ,  0  , _zpx, _zpx, _zpx,  0  , _imp, _aby, _imp,  0  , _abx, _abx, _abx,  0, // D0
  _imm, _iix, _imm,  0  , _zrp, _zrp, _zrp,  0  , _imp, _imm, _imp,  0  , _abs, _abs, _abs,  0, // E0
  _rel, _iiy,  0  ,  0  , _zpx, _zpx, _zpx,  0  , _imp, _aby, _imp,  0  , _abx, _abx, _abx,  0} // F0