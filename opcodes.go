package Cpu6502

const (
// Load/Store Operations
	// LDA - Load Accumulator
	LDA_Imm_A9 = 0xA9
	LDA_Zp_A5 = 0xA5
	LDA_ZpX_B5 = 0xB5
	LDA_Abs_AD = 0xAD
	LDA_AbsX_BD = 0xBD
	LDA_AbsY_B9 = 0xB9
	LDA_IndexIndirX_A1 = 0xA1
	LDA_IndirIndexY_B1 = 0xB1
	
	// LDX - Load X Register
	LDX_Imm_A2 = 0xA2
	LDX_Zp_A6 = 0xA6
	LDX_ZpY_B6 = 0xB6
	LDX_Abs_AE = 0xAE
	LDX_AbsY_BE = 0xBE
		
	// LDY - Load Y Register
	LDY_Imm_A0 = 0xA0
	LDY_Zp_A4 = 0xA4
	LDY_ZpX_B4 = 0xB4
	LDY_Abs_AC = 0xAC
	LDY_AbsX_BC = 0xBC

	// STA - Store Accumulator
	STA_Zp_85 = 0x85
	STA_ZpX_95 = 0x95
	STA_Abs_8D = 0x8D
	STA_AbsX_9D = 0x9D
	STA_AbsY_99 = 0x99
	STA_IndexIndirX_81 = 0x81
	STA_IndirIndexY_91 = 0x91

	// STX - Store X Register
	STX_Zp_86 = 0x86
	STX_ZpY_96 = 0x96
	STX_Abs_8E = 0x8E
		
	// STY - Store Y Register
	STY_Zp_84 = 0x84
	STY_ZpX_94 = 0x94
	STY_Abs_8C = 0x8C

// Register Transfers
	TAX_AA = 0xAA // TAX - Transfer Accumulator to X
	TAY_A8 = 0xA8 // TAY - Transfer Accumulator to Y
	TXA_8A = 0x8A // TXA - Transfer X to Accumulator
	TYA_98 = 0x98 // TYA - Transfer Y to Accumulator

// Stack Operations
	TSX_BA = 0xBA // TSX - Transfer Stack Pointer to X
	TXS_9A = 0x9A // TXS - Transfer X to Stack Pointer
	PHA_48 = 0x48 // PHA - Push Accumulator
	PHP_08 = 0x08 // PHP - Push Processor Status
	PLA_68 = 0x68 // PLA - Pull Accumulator
	PLP_28 = 0x28 // PLP - Pull Processor Status
	
// Logical
	// AND - Logical AND
	AND_Imm_29 = 0x29
	AND_Zp_25 = 0x25
	AND_ZpX_35 = 0x35
	AND_Abs_2D = 0x2D
	AND_AbsX_3D = 0x3D
	AND_AbsY_39 = 0x39
	AND_IndexIndirX_21 = 0x21
	AND_IndirIndexY_31 = 0x31
		
	// EOR - Exclusive OR
	EOR_Imm_49 = 0x49
	EOR_Zp_45 = 0x45
	EOR_ZpX_55 = 0x55
	EOR_Abs_4D = 0x4D
	EOR_AbsX_5D = 0x5D
	EOR_AbsY_59 = 0x59
	EOR_IndexIndirX_41 = 0x41
	EOR_IndirIndexY_51 = 0x51
	
	// ORA - Logical Inclusive OR
	ORA_Imm_09 = 0x09
	ORA_Zp_05 = 0x05
	ORA_ZpX_15 = 0x15
	ORA_Abs_0D = 0x0D
	ORA_AbsX_1D = 0x1D
	ORA_AbsY_19 = 0x19
	ORA_IndexIndirX_01 = 0x01
	ORA_IndirIndexY_11 = 0x11
		
	// BIT - Bit Test
	BIT_Zp_24 = 0x24
	BIT_Abs_2C = 0x2C

// Arithmetic
	// ADC - Add with Carry
	ADC_Imm_69 = 0x69
	ADC_Zp_65 = 0x65
	ADC_ZpX_75 = 0x75
	ADC_Abs_6D = 0x6D
	ADC_AbsX_7D = 0x7D
	ADC_AbsY_79 = 0x79
	ADC_IndexIndirX_61 = 0x61
	ADC_IndirIndexY_71 = 0x71
		
	// SBC - Subtract with Carry
	SBC_Imm_E9 = 0xE9
	SBC_Zp_E5 = 0xE5
	SBC_ZpX_F5 = 0xF5
	SBC_Abs_ED = 0xED
	SBC_AbsX_FD = 0xFD
	SBC_AbsY_F9 = 0xF9
	SBC_IndexIndirX_E1 = 0xE1
	SBC_IndirIndexY_F1 = 0xF1

	// CMP - Compare accumulator
	CMP_Imm_C9 = 0xC9
	CMP_Zp_C5 = 0xC5
	CMP_ZpX_D5 = 0xD5
	CMP_Abs_CD = 0xCD
	CMP_AbsX_DD = 0xDD
	CMP_AbsY_D9 = 0xD9
	CMP_IndexIndirX_C1 = 0xC1
	CMP_IndirIndexY_D1 = 0xD1
		
	// CPX - Compare X register
	CPX_Imm_E0 = 0xE0
	CPX_Zp_E4 = 0xE4
	CPX_Abs_EC = 0xEC
	
	// CPY - Compare Y register
	CPY_Imm_C0 = 0xC0
	CPY_Zp_C4 = 0xC4
	CPY_Abs_CC = 0xCC

	// INC - Increment a memory location
	INC_Zp_E6 = 0xE6
	INC_ZpX_F6 = 0xF6
	INC_Abs_EE = 0xEE
	INC_AbsX_FE = 0xFE

	INX_E8 = 0xE8 // INX - Increment the X register
	INY_C8 = 0xC8 // INY - Increment the Y register
		
	// DEC - Decrement a memory location
	DEC_Zp_C6 = 0xC6
	DEC_ZpX_D6 = 0xD6
	DEC_Abs_CE = 0xCE
	DEC_AbsX_DE = 0xDE

	DEX_CA = 0xCA // DEX - Decrement the X register
	DEY_88 = 0x88 // DEY - Decrement the Y register

// Shifts
	// ASL - Arithmetic Shift Left
	ASL_Accumulator_0A = 0x0A
	ASL_Zp_06 = 0x06
	ASL_ZpX_16 = 0x16
	ASL_Abs_0E = 0x0E
	ASL_AbsX_1E = 0x1E

	// LSR - Logical Shift Right
	LSR_Accumulator_4A = 0x4A
	LSR_Zp_46 = 0x46
	LSR_ZpX_56 = 0x56
	LSR_Abs_4E = 0x4E
	LSR_AbsX_5E = 0x5E
	
	// ROL - Rotate Left
	ROL_Accumulator_2A = 0x2A
	ROL_Zp_26 = 0x26
	ROL_ZpX_36 = 0x36
	ROL_Abs_2E = 0x2E
	ROL_AbsX_3E = 0x3E

	// ROR - Rotate Right
	ROR_Accumulator_6A = 0x6A
	ROR_Zp_66 = 0x66
	ROR_ZpX_76 = 0x76
	ROR_Abs_6E = 0x6E
	ROR_AbsX_7E = 0x7E
		
// Jumps & Calls
	// JMP - Jump
	JMP_Abs_4C = 0x4C
	JMP_Indirect_6C = 0x6C
		
	// JSR - Jump to Subroutine
	JSR_20 = 0x20
	
	// RTS - Return from Subroutine
	RTS_60 = 0x60
		
// Branches
	BCC_Relative_90 = 0x90 // BCC - Branch if Carry Clear
	BCS_Relative_B0 = 0xB0 // BCS - Branch if Carry Set
	BEQ_Relative_F0 = 0xF0 // BEQ - Branch if Equal
	BMI_Relative_30 = 0x30 // BMI - Branch if Minus
	BNE_Relative_D0 = 0xD0 // BNE - Branch if Not Equal
	BPL_Relative_10 = 0x10 // BPL - Branch if Positive
	BVC_Relative_50 = 0x50 // BVC - Branch if Overflow Clear
	BVS_Relative_70 = 0x70 // BVS - Branch if Overflow Set

// Status Flag Changes
	CLC_18 = 0x18 // CLC - Clear Carry Flag
	CLD_D8 = 0xD8 // CLD - Clear Decimal Mode
	CLI_58 = 0x58 // CLI - Clear Interrupt Disable
	CLV_B8 = 0xB8 // CLV - Clear Overflow Flag
	SEC_38 = 0x38 // SEC - Set Carry Flag
	SED_F8 = 0xF8 // SED - Set Decimal Flag
	SEI_78 = 0x78 // SEI - Set Interrupt Disable
		
// System Functions
	BRK_00 = 0x00 // BRK - Force Interrupt
	NOP_EA = 0xEA // NOP - No Operation
	RTI_40 = 0x40 // RTI - Return from Interrupt
)