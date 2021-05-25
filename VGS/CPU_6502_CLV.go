package VGS

import "fmt"

// CLV  Clear Overflow Flag
//
//      0 -> V                           N Z C I D V
//                                       - - - - - 0
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       CLV           B8     1    2

func opc_CLV(bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		P[6] = 0

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tCLV  Clear Overflow Flag.\tP[6]=%d\n", opcode, P[6])
			fmt.Println(dbg_show_message)
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
