package main

import "io/ioutil"

const (
	FBOD_FILE_MODE = 0650
)

// 4BOD Instructions
const (
	FBOD_ASM_NOP = 0x0 // Do nothing
	FBOD_ASM_MVA = 0x1 // Move arg1 mem value to acc
	FBOD_ASM_MVM = 0x2 // Move acc value to arg1 mem
	FBOD_ASM_STA = 0x3 // Set acc value to arg1
	FBOD_ASM_INA = 0x4 // Set acc to arrow key state (1 L; 2 R; 4 U; 8 D)
	FBOD_ASM_INC = 0x5 // Increment acc
	FBOD_ASM_CLS = 0x6 // Clear screen
	FBOD_ASM_SHL = 0x7 // Bitwise shift acc left
	FBOD_ASM_SHR = 0x8 // Bitwise shift acc right
	FBOD_ASM_RDP = 0x9 // Read pixel at x=mem arg1, y=mem arg2 to acc
	FBOD_ASM_FLP = 0xA // Toggle pixel at x=mem arg1, y=mem arg2
	FBOD_ASM_FLG = 0xB // Create flag named arg1
	FBOD_ASM_JMP = 0xC // Jump to flag mem arg1 (last in file)
	FBOD_ASM_CEQ = 0xD // Only perform next instruction if mem arg1 == acc
	FBOD_ASM_CGT = 0xE // Only perform next instruction if mem arg1 > acc
	FBOD_ASM_CLT = 0xF // Only perform next instruction if mem arg1 < acc
)

type Instruction struct {
	ins  byte
	arg1 byte
	arg2 byte
}

type FBOD struct {
	acc     byte
	mem     []byte
	flags   []byte   // List of program addresses
	screen  []uint16 // 16 16-bit columns
	program []Instruction
}

func NewFBOD() *FBOD {
	f := FBOD{
		acc:     0,
		mem:     make([]byte, 16),
		flags:   make([]byte, 16),
		screen:  make([]uint16, 16),
		program: make([]Instruction, 256),
	}
	return &f
}

func (f *FBOD) ClearMem() {
	f.acc = 0
	f.mem = make([]byte, 16)
	f.screen = make([]uint16, 16)
}

func (f *FBOD) ClearScreen() {
	f.screen = make([]uint16, 16)
}

func (f *FBOD) FlipPixel(x, y byte) {
	f.screen[y] ^= 1 << (15 - x)
}

func (f *FBOD) GetPixel(x, y byte) byte {
	return byte((f.screen[y] << x) % 2)
}

func (f *FBOD) SaveProgram(filename string) error {
	data := make([]byte, len(f.program)*2)
	for i := 0; i < len(f.program)*2; i += 2 {
		ins := f.program[i/2]
		data[i] = ins.ins
		data[i+1] = (ins.arg1 << 4) | ins.arg2
	}

	return ioutil.WriteFile(filename, data, FBOD_FILE_MODE)
}

func (f *FBOD) LoadProgram(filename string) error {
	f.ClearMem()
	f.program = make([]Instruction, 256)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	for i := 0; i < len(data); i += 2 {
		f.program[i/2] = Instruction{
			ins:  data[i],
			arg1: (data[i+1] >> 4) % 16,
			arg2: data[i+1] % 16,
		}
	}
	return nil
}

// Reads through the program and indexes all flags
func (f *FBOD) ReadFlags() {
	for i, ins := range f.program {
		if ins.ins == FBOD_ASM_FLG {
			f.flags[ins.arg1] = byte(i)
		}
	}
}

// Returns the index of the next instruction to perform
func (f *FBOD) PerformInstruction(progIndex byte) byte {
	nextIndex := progIndex + 1
	ins := f.program[progIndex]
	x := f.mem[ins.arg1] // Handy reference to resolved addresses
	y := f.mem[ins.arg2]

	switch ins.ins {

	case FBOD_ASM_NOP:
		// Does Nothing

	case FBOD_ASM_MVA:
		f.acc = x

	case FBOD_ASM_MVM:
		f.mem[ins.arg1] = f.acc

	case FBOD_ASM_STA:
		f.acc = ins.arg1

	case FBOD_ASM_INA:
		f.acc = GetArrowsNybl()

	case FBOD_ASM_INC:
		f.acc++
		if f.acc > 15 {
			f.acc = 0
		}
	case FBOD_ASM_CLS:
		f.ClearScreen()

	case FBOD_ASM_SHL:
		f.acc <<= 1
		f.acc %= 16 // Chop off shifted bits outside of nybl

	case FBOD_ASM_SHR:
		f.acc >>= 1

	case FBOD_ASM_RDP:
		f.acc = f.GetPixel(x, y)

	case FBOD_ASM_FLP:
		f.FlipPixel(x, y)

	case FBOD_ASM_FLG:
		// Does nothing; flags are read before execution

	case FBOD_ASM_JMP:
		nextIndex = f.flags[ins.arg1]

	case FBOD_ASM_CEQ:
		if !(x == f.acc) {
			nextIndex++
		}

	case FBOD_ASM_CGT:
		if !(x > f.acc) {
			nextIndex++
		}

	case FBOD_ASM_CLT:
		if !(x < f.acc) {
			nextIndex++
		}

	}

	return nextIndex
}
