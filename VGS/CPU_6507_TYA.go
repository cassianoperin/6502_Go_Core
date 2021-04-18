package VGS

import "fmt"

// TYA Transfer Index Y to Accumulator
//
//     Y -> A                            N Z C I D V
//                                       + + - - - -
//
//     addressing	   assembler    opc  bytes  cyles
//     --------------------------------------------
//     implied	     TYA	         98    1     2
func opc_TYA(bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		A = Y

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tTYA  Transfer Index Y to Accumulator.\tA = Y (%d)\n", opcode, Y)
			fmt.Println(dbg_show_message)
		}

		flags_Z(A)
		flags_N(A)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
