package main

import (
	"Supernova/Arguments"
	"Supernova/Converters"
	"Supernova/Encryptors"
	"Supernova/Utils"
	"flag"
	"fmt"
	"log"
	"os"
)

// Structure
type FlagOptions struct {
	outFile     string
	inputFile   string
	language    string
	encryption  string
	obfuscation string
	variable    string
	key         string
	debug       bool
}

// global variables
var (
	__version__ = "1.0.0"
	__license__ = "MIT"
	__author__  = "@nickvourd"
	__github__  = "https://github.com/nickvourd/Supernova"
)

var __ascii__ = `

███████╗██╗   ██╗██████╗ ███████╗██████╗ ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗ 
██╔════╝██║   ██║██╔══██╗██╔════╝██╔══██╗████╗  ██║██╔═══██╗██║   ██║██╔══██╗
███████╗██║   ██║██████╔╝█████╗  ██████╔╝██╔██╗ ██║██║   ██║██║   ██║███████║
╚════██║██║   ██║██╔═══╝ ██╔══╝  ██╔══██╗██║╚██╗██║██║   ██║╚██╗ ██╔╝██╔══██║
███████║╚██████╔╝██║     ███████╗██║  ██║██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║
╚══════╝ ╚═════╝ ╚═╝     ╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝

Supernova v%s - A real fucking shellcode encryptor.
Supernova is an open source tool licensed under %s.
Written with <3 by %s...
Please visit %s for more...

`

// Options function
func Options() *FlagOptions {
	inputFile := flag.String("i", "", "Path to the raw 64-bit shellcode.")
	encryption := flag.String("enc", "", "Shellcode encryption (i.e., XOR, RC4, AES)")
	obfuscation := flag.String("obs", "", "Shellcode obfuscation")
	language := flag.String("lang", "", "Programming language to translate the shellcode (i.e., Nim, Rust, C, CSharp)")
	outFile := flag.String("o", "", "Name of output file")
	variable := flag.String("v", "shellcode", "Name of shellcode variable")
	debug := flag.Bool("d", false, "Enable Debug mode")
	key := flag.String("k", "", "Key phrase for encryption")
	flag.Parse()

	return &FlagOptions{outFile: *outFile, inputFile: *inputFile, language: *language, encryption: *encryption, obfuscation: *obfuscation, variable: *variable, debug: *debug, key: *key}
}

// main function
func main() {
	// Print ascii
	fmt.Printf(__ascii__, __version__, __license__, __author__, __github__)

	// Check GO version of the current system
	Utils.Version()

	// Retrieve command-line options using the Options function
	options := Options()

	// Check Arguments Length
	Arguments.ArgumentLength()

	// Call function named ArgumentEmpty
	Arguments.ArgumentEmpty(options.inputFile, 1)

	// Call function name ConvertShellcode2String
	rawShellcode, err := Converters.ConvertShellcode2String(options.inputFile)
	if err != nil {
		fmt.Println("[!] Error:", err)
		return
	}

	// Call function named ArgumentEmpty
	Arguments.ArgumentEmpty(options.language, 2)

	// Check for valid values of language argument
	foundLanguage := Arguments.ValidateArgument("lang", options.language, []string{"Nim", "Rust", "C", "CSharp"})

	// Checks if either encryption or obfuscation options are provided.
	if options.encryption == "" && options.obfuscation == "" {
		logger := log.New(os.Stderr, "[!] ", 0)
		logger.Fatal("Please provide at least -enc or -obs option with a valid value...")
	}

	// Call function named ConvertShellcode2Hex
	convertedShellcode, payloadLength := Converters.ConvertShellcode2Hex(rawShellcode, foundLanguage)

	// Print payload size and choosen language
	fmt.Printf("[+] Payload size: %d bytes\n\n[+] Converted payload to %s language\n\n", payloadLength, foundLanguage)

	if options.debug != false {
		// Call function named ConvertShellcode2Template
		template := Converters.ConvertShellcode2Template(convertedShellcode, foundLanguage, payloadLength, options.variable)
		// Print original template
		fmt.Printf("[+] The original payload:\n\n%s\n\n", template)
	}

	if options.encryption != "" {
		Arguments.ValidateArgument("enc", options.encryption, []string{"XOR", "RC4", "AES"})
		// Call function named DetectEncryption
		Encryptors.DetectEncryption(options.encryption, rawShellcode)
	}

	// Check for valid values of obfuscation argument
	if options.obfuscation != "" {
		Arguments.ValidateArgument("obs", options.obfuscation, []string{"IPv4", "IPv6", "MAC", "UUID"})
	}

}
