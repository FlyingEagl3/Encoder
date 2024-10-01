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
	authorName = "Assistant"
	toolName   = "Encoding Demonstration Tool"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-h" {
		printHelp()
		return
	}

	fmt.Printf("Author: %s\nTool: %s\n\n", authorName, toolName)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to encode: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	fmt.Println("\nEncoding Schemes:")
	fmt.Println("----------------")

	encodeAndPrint("URL Encoding", url.QueryEscape, text)
	encodeAndPrint("HTML Entity Encoding", htmlEntityEncode, text)
	encodeAndPrint("Unicode Encoding", unicodeEncode, text)
	encodeAndPrint("ASCII Encoding", asciiEncode, text)
	encodeAndPrint("UTF-7 Encoding", utf7Encode, text)
	encodeAndPrint("UTF-8 Encoding", utf8Encode, text)
	encodeAndPrint("UTF-16 Encoding", utf16Encode, text)
	encodeAndPrint("UTF-32 Encoding", utf32Encode, text)
	encodeAndPrint("ISO-8859-1 Encoding", iso88591Encode, text)
	encodeAndPrint("ISO-8859-2 Encoding", iso88592Encode, text)
	encodeAndPrint("ISO-8859-5 Encoding", iso88595Encode, text)
	encodeAndPrint("ISO-8859-6 Encoding", iso88596Encode, text)
	encodeAndPrint("ISO-8859-7 Encoding", iso88597Encode, text)
	encodeAndPrint("ISO-8859-8 Encoding", iso88598Encode, text)
	encodeAndPrint("ISO-8859-9 Encoding", iso88599Encode, text)
	encodeAndPrint("ISO-8859-10 Encoding", iso885910Encode, text)
	encodeAndPrint("ISO-8859-11 Encoding", iso885911Encode, text)
	encodeAndPrint("ISO-8859-13 Encoding", iso885913Encode, text)
	encodeAndPrint("ISO-8859-14 Encoding", iso885914Encode, text)
	encodeAndPrint("ISO-8859-15 Encoding", iso885915Encode, text)
	encodeAndPrint("ISO-8859-16 Encoding", iso885916Encode, text)
	encodeAndPrint("Windows-1250 Encoding", windows1250Encode, text)
	encodeAndPrint("Windows-1251 Encoding", windows1251Encode, text)
	encodeAndPrint("Windows-1252 Encoding", windows1252Encode, text)
	encodeAndPrint("Windows-1253 Encoding", windows1253Encode, text)
	encodeAndPrint("Windows-1254 Encoding", windows1254Encode, text)
	encodeAndPrint("Windows-1255 Encoding", windows1255Encode, text)
	encodeAndPrint("Windows-1256 Encoding", windows1256Encode, text)
	encodeAndPrint("Windows-1257 Encoding", windows1257Encode, text)
	encodeAndPrint("Windows-1258 Encoding", windows1258Encode, text)
	encodeAndPrint("Binary Encoding", binaryEncode, text)
	encodeAndPrint("Hex Encoding",	hexEncode, text)
	encodeAndPrint("Octal Encoding", octalEncode, text)
	encodeAndPrint("Decimal Encoding", decimalEncode, text)
	encodeAndPrint("JavaScript Encoding", javascriptEncode, text)
	encodeAndPrint("JavaScript Unicode Encoding", javascriptUnicodeEncode, text)
	encodeAndPrint("JavaScript Octal Encoding", javascriptOctalEncode, text)
	encodeAndPrint("Double URL Encoding", doubleURLEncode, text)
	encodeAndPrint("Double HTML Entity Encoding", doubleHTMLEntityEncode, text)
	encodeAndPrint("Double Unicode Encoding", doubleUnicodeEncode, text)
	encodeAndPrint("Double ASCII Encoding", doubleASCIIEncode, text)
	encodeAndPrint("Base64 Encoding", base64Encode, text)
	encodeAndPrint("Base32 Encoding", base32Encode, text)
	encodeAndPrint("Base16 Encoding", base16Encode, text)
	encodeAndPrint("Punycode Encoding", punycodeDecode, text)
	encodeAndPrint("ROT13 Encoding", rot13Encode, text)
	encodeAndPrint("Caesar Cipher Encoding", caesarCipherEncode, text)
	encodeAndPrint("Vigenère Cipher Encoding", vigenereCipherEncode, text)
	encodeAndPrint("XOR Encoding", xorEncode, text)
	encodeAndPrint("Null Character Encoding", nullCharEncode, text)
	encodeAndPrint("Tab Character Encoding", tabCharEncode, text)
	encodeAndPrint("Newline Character Encoding", newlineCharEncode, text)
	encodeAndPrint("Carriage Return Character Encoding", crCharEncode, text)
	encodeAndPrint("Space Character Encoding", spaceCharEncode, text)
}

func printHelp() {
	fmt.Printf("Author: %s\nTool: %s\n\n", authorName, toolName)
	fmt.Println("Usage:")
	fmt.Println("  go run encoding_demo.go")
	fmt.Println("  go run encoding_demo.go -h")
	fmt.Println("\nThis tool demonstrates various encoding schemes for a given input text.")
}

func encodeAndPrint(name string, encoder func(string) string, text string) {
	encoded := encoder(text)
	fmt.Printf("%s: %s\n", name, encoded)
}

func htmlEntityEncode(s string) string {
	return strings.Replace(s, "&", "&amp;", -1)
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
	return strings.Replace(s, "+", "\\+", -1)
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
		sb.WriteString(fmt.Sprintf("\\U%08X
", v))
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

func punycodeDecode(s string) string {
	// Punycode decoding is not implemented in the standard library,
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
			sb.WriteRune(((r-'a'+shift)%26) + 'a')
		} else if r >= 'A' && r <= 'Z' {
			sb.WriteRune(((r-'A'+shift)%26) + 'A')
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func vigenereCipherEncode(s, key string) string {
	var sb strings.Builder
	keyLen := len(key)
	for i, r := range s {
		if r >= 'a' && r <= 'z' {
			shift := int(key[i%keyLen]) - 'a'
			sb.WriteRune(((r-'a'+shift)%26) + 'a')
		} else if r >= 'A' && r <= 'Z' {
			shift := int(key[i%keyLen]) - 'A'
			sb.WriteRune(((r-'A'+shift)%26) + 'A')
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
	return strings.Replace(s, "\x00", "\\x00", -1)
}

func tabCharEncode(s string) string {
	return strings.Replace(s, "\t", "\\t", -1)
}

func newlineCharEncode(s string) string {
	return strings.Replace(s, "\n", "\\n", -1)
}

func crCharEncode(s string) string {
	return strings.Replace(s, "\r", "\\r", -1)


	}

func spaceCharEncode(s string) string {
	return strings.Replace(s, " ", "\\s", -1)
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
	// so we can't provide a complete implementation here.
	return s
}


