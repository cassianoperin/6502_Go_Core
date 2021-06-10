package main

import (
	"6502/CORE"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/faiface/pixel/pixelgl"
)

// Function used by readROM to avoid 'bytesread' return
func ReadContent(file *os.File, bytes_number int) []byte {

	bytes := make([]byte, bytes_number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

// Read ROM and write it to the RAM
func readROM(filename string) {

	var (
		fileInfo os.FileInfo
		err      error
	)

	// Get ROM info
	fileInfo, err = os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loading ROM:", filename)
	romsize := fileInfo.Size()
	fmt.Printf("Size in bytes: %d\n", romsize)

	// Open ROM file, insert all bytes into memory
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Call ReadContent passing the total size of bytes
	data := ReadContent(file, int(romsize))
	// Print raw data
	//fmt.Printf("%d\n", data)
	//fmt.Printf("%X\n", data)

	// // 4KB roms
	// if romsize == 4096 {
	// 	// Load ROM to memory
	// 	for i := 0; i < len(data); i++ {
	// 		// F000 - FFFF // Cartridge ROM
	// 		VGS.Memory[0xF000+i] = data[i]
	// 	}
	// }

	// // 2KB roms (needs to duplicate it in memory)
	// if romsize == 2048 {
	// 	// Load ROM to memory
	// 	for i := 0; i < len(data); i++ {
	// 		// F000 - F7FF (2KB Cartridge ROM)
	// 		VGS.Memory[0xF000+i] = data[i]
	// 		// F800 - FFFF (2KB Mirror Cartridge ROM)
	// 		VGS.Memory[0xF800+i] = data[i]
	// 	}
	// }

	if romsize == 65536 {
		// Load ROM to memory
		for i := 0; i < len(data); i++ {
			// F000 - F7FF (2KB Cartridge ROM)
			CORE.Memory[i] = data[i]
			// F800 - FFFF (2KB Mirror Cartridge ROM)
			CORE.Memory[i] = data[i]
		}
	}

	// // Print Memory -  Fist 2kb
	// for i := 0xF7F0; i <= 0xF7FF; i++ {
	// 	fmt.Printf("%X ", VGS.Memory[i])
	// }
	// fmt.Println()
	// //
	// for i := 0xFFF0; i <= 0xFFFF; i++ {
	// 	fmt.Printf("%X ", VGS.Memory[i])
	// }
	// fmt.Println()

	// // Print Memory
	// for i := 0; i < len(VGS.Memory); i++ {
	// 	fmt.Printf("%X ", VGS.Memory[i])
	// }
	// os.Exit(2)
}

// func checkArgs() {
// 	if len(os.Args) != 2 {
// 		fmt.Printf("Usage: %s ROM_FILE\n\n", os.Args[0])
// 		os.Exit(0)
// 	}
// }

func checkArgs() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [options] ROM_FILE\n\n%s -help for a list of available options\n\n", os.Args[0], os.Args[0])
		os.Exit(0)
	}

	cliHelp := flag.Bool("help", false, "Show this menu")
	cliConsole := flag.Bool("console", false, "Open program in interactive console")
	cliDebug := flag.Bool("debug", false, "Enable Debug Mode")
	cliPause := flag.Bool("pause", false, "Start emulation Paused")

	// wordPtr := flag.String("word", "foo", "a string")
	// numbPtr := flag.Int("numb", 42, "an int")
	// var svar string
	// flag.StringVar(&svar, "ROM_FILE", "bar", "ROM_FILE")
	// fmt.Println("word:", *wordPtr)
	// fmt.Println("numb:", *numbPtr)
	// fmt.Println("svar:", svar)
	// fmt.Println("tail:", flag.Arg(0))
	flag.Parse()

	if *cliHelp {
		fmt.Printf("Usage: %s [options] ROM_FILE\n  -console\n    	Open program in interactive console\n  -debug\n    	Enable Debug Mode\n  -pause\n    	Start emulation Paused\n  -help\n    	Show this menu\n\n", os.Args[0])
		os.Exit(0)
	}

	if *cliConsole {
		fmt.Printf("CHAMAR CONSOLEEE!")
		os.Exit(0)
	}

	if *cliDebug {
		CORE.Debug = true
	}

	if *cliPause {
		CORE.Pause = true
	}

}

func testFile(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("File '%s' not found.\n\n", os.Args[1])
		os.Exit(0)
	}
}

func main() {
	// Validate the Arguments
	checkArgs()

	// testFile(os.Args[1])
	if len(flag.Args()) != 0 { // Ensure that there is an last argument (rom name)
		// Check if file exist
		testFile(flag.Arg(0))
	} else {
		fmt.Printf("Usage: %s [options] ROM_FILE\n  -console\n    	Open program in interactive console\n  -debug\n    	Enable Debug Mode\n  -pause\n    	Start emulation Paused\n  -help\n    	Show this menu\n\n", os.Args[0])
		os.Exit(0)
	}

	fmt.Printf("\nMOS 6502 CPU Emulator\n\n")

	// Set initial variables values
	CORE.Initialize()
	// Initialize Timers
	CORE.InitializeTimers()

	// Read ROM to the memory
	// readROM(os.Args[1])
	readROM("/Users/cassiano/go/src/6502/TestPrograms/6502_functional_test.bin")
	// readROM("/Users/cassiano/go/src/6502/TestPrograms/6502_decimal_test.bin")

	// Reset system
	CORE.Reset()

	// Start Window System and draw Graphics
	pixelgl.Run(CORE.Run)

}
