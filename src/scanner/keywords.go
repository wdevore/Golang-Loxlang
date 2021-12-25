package scanner

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

var Keywords = map[string]api.TokenType{
	"const":    api.CONST,
	"import":   api.IMPORT,
	"print":    api.PRINT,
	"var":      api.VAR,
	"nil":      api.NIL,
	"true":     api.TRUE,
	"false":    api.FALSE,
	"and":      api.AND,
	"or":       api.OR,
	"if":       api.IF,
	"else":     api.ELSE,
	"while":    api.WHILE,
	"for":      api.FOR,
	"break":    api.BREAK,
	"continue": api.CONTINUE,
	"fun":      api.FUN,
	"code":     api.CODE,
	"alignTo":  api.ALIGN_TO,
	"global":   api.GLOBAL,
	"at":       api.AT,
	"as":       api.AS,
	"use":      api.USE,
	"readOnly": api.READ_ONLY,
	"byte":     api.BYTE,
	"half":     api.HALF,
	"word":     api.WORD,
	"data":     api.DATA,
	"int":      api.INT,
	"hi":       api.HI,
	"lo":       api.LO,

	// Instructions
	"add": api.ADD,
	"sub": api.SUB,
	"xor": api.XOR,
	// "or":     OR,
	// "and":    AND,
	"sll":    api.SLL,
	"srl":    api.SRL,
	"sra":    api.SRA,
	"slt":    api.SLT,
	"sltu":   api.SLTU,
	"addi":   api.ADDI,
	"xori":   api.XORI,
	"ori":    api.ORI,
	"andi":   api.ANDI,
	"slli":   api.SLLI,
	"srli":   api.SRLI,
	"srai":   api.SRAI,
	"slti":   api.SLTI,
	"sltiu":  api.SLTIU,
	"lb":     api.LB,
	"lh":     api.LH,
	"lw":     api.LW,
	"lbu":    api.LBU,
	"lhu":    api.LHU,
	"sb":     api.SB,
	"sh":     api.SH,
	"sw":     api.SW,
	"beq":    api.BEQ,
	"bne":    api.BNE,
	"blt":    api.BLT,
	"bge":    api.BGE,
	"bltu":   api.BLTU,
	"bgeu":   api.BGEU,
	"jal":    api.JAL,
	"jalr":   api.JALR,
	"lui":    api.LUI,
	"auipc":  api.AUIPC,
	"ecall":  api.ECALL,
	"ebreak": api.EBREAK,
	// Pseudo instructions
	"la":   api.LA,
	"nop":  api.NOP,
	"li":   api.LI,
	"mv":   api.MV,
	"not":  api.NOT,
	"neg":  api.NEG,
	"negw": api.NEGW,
	"sext": api.SEXT,
	"seqz": api.SEQZ,
	"snez": api.SNEZ,
	"sltz": api.SLTZ,
	"sgtz": api.SGTZ,
	"beqz": api.BEQZ,
	"bnez": api.BNEZ,
	"blez": api.BLEZ,
	"bgez": api.BGEZ,
	"bltz": api.BLTZ,
	"bgtz": api.BGTZ,
	"bgt":  api.BGT,
	"ble":  api.BLE,
	"bgtu": api.BGTU,
	"bleu": api.BLEU,
	"j":    api.J,
	"ret":  api.RET,
	"call": api.CALL,
	"tail": api.TAIL,
}
