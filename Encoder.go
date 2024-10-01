package main

import (
	"bufio"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"strings"
	"unicode/utf16"
	"unicode/utf32"
)

const (
	authorName = "Flying_Eagle" // Updated author name
	toolName   = "Encoder Tool"
)

func main() {
	// Check if at least one argument is provided
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	// Get the payload from command-line arguments
	payload := os.Args[1]

	// Check for verbosity option
	verbose := false
	if len(os.Args) > 2 && os.Args[2] == "-v" {
		verbose = true
	}

	fmt.Printf("Author: %s\nTool: %s\n\n", authorName, toolName)

	// Print the payload if verbosity is enabled
	if verbose {
		fmt.Printf("Payload: %s\n", payload)
	}

	fmt.Println("\nEncoding Schemes:")
	fmt.Println("----------------")

	// Call encoding functions and print results
	encodeAndPrint("URL Encoding", url.QueryEscape, payload, verbose)
	encodeAndPrint("HTML Entity Encoding", htmlEntityEncode, payload, verbose)
	encodeAndPrint("Unicode Encoding", unicodeEncode, payload, verbose)
	encodeAndPrint("ASCII Encoding", asciiEncode, payload, verbose)
	encodeAndPrint("UTF-7 Encoding", utf7Encode, payload, verbose)
	encodeAndPrint("UTF-8 Encoding", utf8Encode, payload, verbose)
	encodeAndPrint("UTF-16 Encoding", utf16Encode, payload, verbose)
	encodeAndPrint("UTF-32 Encoding", utf32Encode, payload, verbose)
	encodeAndPrint("ISO-8859-1 Encoding", iso88591Encode, payload, verbose)
	encodeAndPrint("ISO-8859-2 Encoding", iso88592Encode, payload, verbose)
	encodeAndPrint("ISO-8859-5 Encoding", iso88595Encode, payload, verbose)
	encodeAndPrint("ISO-8859-6 Encoding", iso88596Encode, payload, verbose)
	encodeAndPrint("ISO-8859-7 Encoding", iso88597Encode, payload, verbose)
	encodeAndPrint("ISO-8859-8 Encoding", iso88598Encode, payload, verbose)
	encodeAndPrint("ISO-8859-9 Encoding", iso88599Encode, payload, verbose)
	encodeAndPrint("ISO-8859-10 Encoding", iso885910Encode, payload, verbose)
	encodeAndPrint("ISO-8859-11 Encoding", iso885911Encode, payload, verbose)
	encodeAndPrint("ISO-8859-13 Encoding", iso885913Encode, payload, verbose)
	encodeAndPrint("ISO-8859-14 Encoding", iso885914Encode, payload, verbose)
	encodeAndPrint("ISO-8859-15 Encoding", iso885915Encode, payload, verbose)
	encodeAndPrint("ISO-8859-16 Encoding", iso885916Encode, payload, verbose)
	encodeAndPrint("Windows-1250 Encoding", windows1250Encode, payload, verbose)
	encodeAndPrint("Windows-1251 Encoding", windows1251Encode, payload, verbose)
	encodeAndPrint("Windows-1252 Encoding", windows1252Encode, payload, verbose)
	encodeAndPrint("Windows-1253 Encoding", windows1253Encode, payload, verbose)
	encodeAndPrint("Windows-1254 Encoding", windows1254Encode, payload, verbose)
	encodeAndPrint("Windows-1255 Encoding", windows1255Encode, payload, verbose)
	encodeAndPrint("Windows-1256 Encoding", windows1256Encode, payload, verbose)
	encodeAndPrint("Windows-1257 Encoding", windows1257Encode, payload, verbose)
	encodeAndPrint("Windows-1258 Encoding", windows1258Encode, payload, verbose)
	encodeAndPrint("Binary Encoding", binaryEncode, payload, verbose)
	encodeAndPrint("Hex Encoding", hexEncode, payload, verbose)
	encodeAndPrint("Octal Encoding", octalEncode, payload, verbose)
	encodeAndPrint("Decimal Encoding", decimalEncode, payload, verbose)
	encodeAndPrint("JavaScript Encoding", javascriptEncode, payload, verbose)
	encodeAndPrint("JavaScript Unicode Encoding", javascriptUnicodeEncode, payload, verbose)
	encodeAndPrint("JavaScript Octal Encoding", javascriptOctalEncode, payload, verbose)
	encodeAndPrint("Double URL Encoding", doubleURLEncode, payload, verbose)
	encodeAndPrint("Double HTML Entity Encoding", doubleHTMLEntityEncode, payload, verbose)
	encodeAndPrint("Double Unicode Encoding", doubleUnicodeEncode, payload, verbose)
	encodeAndPrint("Double ASCII Encoding", doubleASCIIEncode, payload, verbose)
	encodeAndPrint("Base64 Encoding", base64Encode, payload, verbose)
	encodeAndPrint("Base32 Encoding", base32Encode, payload, verbose)
	encodeAndPrint("Base16 Encoding", base16Encode, payload, verbose)
	encodeAndPrint("Punycode Encoding", punycodeEncode, payload, verbose)
	encodeAndPrint("ROT13 Encoding", rot13Encode, payload, verbose)
	encodeAndPrint("Caesar Cipher Encoding", func(s string) string { return caesarCipherEncode(s, 3) }, payload, verbose) // Using a shift of 3 for Caesar Cipher
	encodeAndPrint("Vigenère Cipher Encoding", func(s string) string { return vigenereCipherEncode(s, "key") }, payload, verbose) // Using "key" as the Vigenère key
	encodeAndPrint("XOR Encoding", func(s string) string { return xorEncode(s, "key") }, payload, verbose) // Using "key" for XOR encoding
	encodeAndPrint("Null Character Encoding", nullCharEncode, payload, verbose)
	encodeAndPrint("Tab Character Encoding", tabCharEncode, payload, verbose)
	encodeAndPrint("Newline Character Encoding", newlineCharEncode, payload, verbose)
	encodeAndPrint("Carriage Return Character Encoding", crCharEncode, payload, verbose)
	encodeAndPrint("Space Character Encoding", spaceCharEncode, payload, verbose)
}

func printHelp() {
	fmt.Printf("Author: %s\nTool: %s\n\n", authorName, toolName)
	fmt.Println("Usage:")
	fmt.Println("  Encoder [payload]")
	fmt.Println("  Encoder [payload] -v")
	fmt.Println("\nThis tool demonstrates various encoding schemes for a given input text.")
}

func encodeAndPrint(name string, encoder func(string) string, text string, verbose bool) {
	encoded := encoder(text)
	if verbose {
		fmt.Printf("%s: %s\n", name, encoded)
	} else {
		fmt.Println(encoded)
	}
}

// Encoding functions (placeholders for the actual implementations)
func htmlEntityEncode(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}

func unicodeEncode(s string) string {
	var sb strings.Builder
	for _, r := range s {
		sb.WriteString(fmt.Sprintf("\\u%04X", r))
	}
	return sb.String()
}

func asciiEncode(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if r <= 127 {
			sb.WriteRune(r)
		} else {
			sb.WriteString(fmt.Sprintf("\\x%02X", byte(r)))
		}
	}
	return sb.String()
}

func utf7Encode(s string) string {
	return strings.ReplaceAll(s, "+", "\\+")
}

func utf8Encode(s string) string {
	return s
}

func utf16Encode(s string) string {
	encoded := utf16.Encode([]rune(s))
	var sb strings.Builder
	for _, v := range encoded {
		sb.WriteString(fmt.Sprintf("\\u%04X", v))
	}
	return sb.String()
}

func utf32Encode(s string) string {
	encoded := utf32.Encode([]rune(s))
	var sb strings.Builder
	for _, v := range encoded {
		sb.WriteString(fmt.Sprintf("\\U%08X", v))
	}
	return sb.String()
}

func iso88591Encode(s string) string {
	return encodeISO8859(s, 1)
}

func iso88592Encode(s string) string {
	return encodeISO8859(s, 2)
}

func iso88595Encode(s string) string {
	return encodeISO8859(s, 5)
}

func iso88596Encode(s string) string {
	return encodeISO8859(s, 6)
}

func iso88597Encode(s string) string {
	return encodeISO8859(s, 7)
}

func iso88598Encode(s string) string {
	return encodeISO8859(s, 8)
}

func iso88599Encode(s string) string {
	return encodeISO8859(s, 9)
}

func iso885910Encode(s string) string {
	return encodeISO8859(s, 10)
}

func iso885911Encode(s string) string {
	return encodeISO8859(s, 11)
}

func iso885913Encode(s string) string {
	return encodeISO8859(s, 13)
}

func iso885914Encode(s string) string {
	return encodeISO8859(s, 14)
}

func iso885915Encode(s string) string {
	return encodeISO8859(s, 15)
}

func iso885916Encode(s string) string {
	return encodeISO8859(s, 16)
}

func windows1250Encode(s string) string {
	return encodeWindows(s, 1250)
}

func windows1251Encode(s string) string {
	return encodeWindows(s, 1251)
}

func windows1252Encode(s string) string {
	return encodeWindows(s, 1252)
}

func windows1253Encode(s string) string {
	return encodeWindows(s, 1253)
}

func windows1254Encode(s string) string {
	return encodeWindows(s, 1254)
}

func windows1255Encode(s string) string {
	return encodeWindows(s, 1255)
}

func windows1256Encode(s string) string {
	return encodeWindows(s, 1256)
}

func windows1257Encode(s string) string {
	return encodeWindows(s, 1257)
}

func windows1258Encode(s string) string {
	return encodeWindows(s, 1258)
}

func binaryEncode(s string) string {
	var sb strings.Builder
	for _, r := range s {
		sb.WriteString(fmt.Sprintf("%08b ", r))
	}
	return sb.String()
}

func hexEncode(s string) string {
	return hex.EncodeToString([]byte(s))
}

func octalEncode(s string) string {
	var sb strings.Builder
	for _, r := range s {
		sb.WriteString(fmt.Sprintf("%03o ", r))
	}
	return sb.String()
}

func decimalEncode(s string) string {
	var sb strings.Builder
	for _, r := range s {
		sb.WriteString(fmt.Sprintf("%d ", r))
	}
	return sb.String()
}

func javascriptEncode(s string) string {
	var sb strings.Builder
	for _, r := range s {
		sb.WriteString(fmt.Sprintf("\\x%02X", r))
	}
	return sb.String()
}

func javascriptUnicodeEncode(s string) string {
	var sb strings.Builder
	for _, r := range s {
		sb.WriteString(fmt.Sprintf("\\u%04X", r))
	}
	return sb.String()
}

func javascriptOctalEncode(s string) string {
	var sb strings.Builder
	for _, r := range s {
		sb.WriteString(fmt.Sprintf("\\%03o", r))
	}
	return sb.String()
}

func doubleURLEncode(s string) string {
	return url.QueryEscape(url.QueryEscape(s))
}

func doubleHTMLEntityEncode(s string) string {
	return htmlEntityEncode(htmlEntityEncode(s))
}

func doubleUnicodeEncode(s string) string {
	return unicodeEncode(unicodeEncode(s))
}

func doubleASCIIEncode(s string) string {
	return asciiEncode(asciiEncode(s))
}

func base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func base32Encode(s string) string {
	return base32.StdEncoding.EncodeToString([]byte(s))
}

func base16Encode(s string) string {
	return hex.EncodeToString([]byte(s))
}

func punycodeEncode(s string) string {
	// Punycode encoding is not implemented in the standard library,
	// so we can't provide a complete implementation here.
	return s
}

func rot13Encode(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			sb.WriteRune(((r-'a'+13)%26) + 'a')
		} else if r >= 'A' && r <= 'Z' {
			sb.WriteRune(((r-'A'+13)%26) + 'A')
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func caesarCipherEncode(s string, shift int) string {
	var sb strings.Builder
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			sb.WriteRune(((r-'a'+rune(shift))%26) + 'a')
		} else if r >= 'A' && r <= 'Z' {
			sb.WriteRune(((r-'A'+rune(shift))%26) + 'A')
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func vigenereCipherEncode(s, key string) string {
	var sb strings.Builder
	keyLen := len(key)
	keyIndex := 0
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			shift := int(key[keyIndex%keyLen]) - 'a'
			sb.WriteRune(((r-'a'+rune(shift))%26) + 'a')
			keyIndex++
		} else if r >= 'A' && r <= 'Z' {
			shift := int(key[keyIndex%keyLen]) - 'A'
			sb.WriteRune(((r-'A'+rune(shift))%26) + 'A')
			keyIndex++
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func xorEncode(s, key string) string {
	var sb strings.Builder
	keyLen := len(key)
	for i, r := range s {
		sb.WriteRune(rune(int(r) ^ int(key[i%keyLen])))
	}
	return sb.String()
}

func nullCharEncode(s string) string {
	return strings.ReplaceAll(s, "\x00", "\\x00")
}

func tabCharEncode(s string) string {
	return strings.ReplaceAll(s, "\t", "\\t")
}

func newlineCharEncode(s string) string {
	return strings.ReplaceAll(s, "\n", "\\n")
}

func crCharEncode(s string) string {
	return strings.ReplaceAll(s, "\r", "\\r")
}

func spaceCharEncode(s string) string {
	return strings.ReplaceAll(s, " ", "\\s")
}

func encodeISO8859(s string, index int) string {
	var sb strings.Builder
	for _, r := range s {
		if r <= 255 {
			sb.WriteRune(r)
		} else {
			sb.WriteString(fmt.Sprintf("\\x%02X", byte(r)))
		}
	}
	return sb.String()
}

func encodeWindows(s string, codepage int) string {
	// Windows codepage encoding is not implemented in the standard library,
	// so we not implemented here.
	return s
}

func main() {
	// The main function is already defined above.
	// This is just to ensure the program is complete.
}
