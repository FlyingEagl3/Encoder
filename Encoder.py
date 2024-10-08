import argparse
import urllib.parse
import html
import base64
import codecs
import idna  # Using idna for Punycode encoding
import re

class Encoder:
    def __init__(self, payload):
        self.payload = payload

    def url_encode(self):
        return urllib.parse.quote(self.payload)

    def html_entity_encode(self):
        return html.escape(self.payload)

    def unicode_encode(self):
        return self.payload.encode('unicode_escape').decode()

    def ascii_encode(self):
        return self.payload.encode('ascii', 'ignore').decode()

    def utf7_encode(self):
        return self.payload.encode('utf-7').decode('utf-7', errors='ignore')

    def utf8_encode(self):
        return self.payload.encode('utf-8').decode('utf-8', errors='ignore')

    def utf16_encode(self):
        return self.payload.encode('utf-16').decode('utf-16', errors='ignore')

    def utf32_encode(self):
        return self.payload.encode('utf-32').decode('utf-32', errors='ignore')

    def iso_encode(self, encoding):
        return self.payload.encode(encoding).decode(encoding, errors='ignore')

    def windows_encode(self, encoding):
        return self.payload.encode(encoding).decode(encoding, errors='ignore')

    def binary_encode(self):
        return ' '.join(format(ord(char), '08b') for char in self.payload)

    def hex_encode(self):
        return self.payload.encode().hex()

    def octal_encode(self):
        return ' '.join(format(ord(char), '03o') for char in self.payload)

    def decimal_encode(self):
        return ' '.join(str(ord(char)) for char in self.payload)

    def js_encode(self):
        return ''.join(f'\\u{ord(char):04x}' for char in self.payload)

    def js_unicode_encode(self):
        return self.js_encode()

    def js_octal_encode(self):
        return ''.join(f'\\{ord(char):03o}' for char in self.payload)

    def double_encode(self, encode_func):
        # Call the encoding function twice with the payload
        first_encode = encode_func()  # Call the function without self
        return encode_func()  # Call the function again

    def base64_encode(self):
        return base64.b64encode(self.payload.encode()).decode()

    def base32_encode(self):
        return base64.b32encode(self.payload.encode()).decode()

    def base16_encode(self):
        return base64.b16encode(self.payload.encode()).decode()

    def punycode_encode(self):
        # Extract domain from URL if present
        domain = self.extract_domain(self.payload)
        if domain:
            return idna.encode(domain).decode()  # Using idna for Punycode encoding
        return "Invalid domain for Punycode encoding"

    def extract_domain(self, url):
        # Use regex to extract the domain from a URL
        match = re.search(r'^(?:https?://)?([^/]+)', url)
        return match.group(1) if match else None

    def rot13_encode(self):
        return codecs.encode(self.payload, 'rot_13')

    def caesar_cipher_encode(self, shift):
        result = []
        for char in self.payload:
            if char.isalpha():
                shift_base = ord('a') if char.islower() else ord('A')
                result.append(chr((ord(char) - shift_base + shift) % 26 + shift_base))
            else:
                result.append(char)
        return ''.join(result)

    def xor_encode(self, key):
        return ''.join(chr(ord(c) ^ key) for c in self.payload)

    def get_encoded_payloads(self):
        encodings = {
            "URL Encoding": self.url_encode(),
            "HTML Entity Encoding": self.html_entity_encode(),
            "Unicode Encoding": self.unicode_encode(),
            "ASCII Encoding": self.ascii_encode(),
            "UTF-7 Encoding": self.utf7_encode(),
            "UTF-8 Encoding": self.utf8_encode(),
            "UTF-16 Encoding": self.utf16_encode(),
            "UTF-32 Encoding": self.utf32_encode(),
            "ISO-8859-1 Encoding": self.iso_encode('iso-8859-1'),
                        "ISO-8859-2 Encoding": self.iso_encode('iso-8859-2'),
            "ISO-8859-5 Encoding": self.iso_encode('iso-8859-5'),
            "ISO-8859-6 Encoding": self.iso_encode('iso-8859-6'),
            "ISO-8859-7 Encoding": self.iso_encode('iso-8859-7'),
            "ISO-8859-8 Encoding": self.iso_encode('iso-8859-8'),
            "ISO-8859-9 Encoding": self.iso_encode('iso-8859-9'),
            "ISO-8859-10 Encoding": self.iso_encode('iso-8859-10'),
            "ISO-8859-11 Encoding": self.iso_encode('iso-8859-11'),
            "ISO-8859-13 Encoding": self.iso_encode('iso-8859-13'),
            "ISO-8859-14 Encoding": self.iso_encode('iso-8859-14'),
            "ISO-8859-15 Encoding": self.iso_encode('iso-8859-15'),
            "ISO-8859-16 Encoding": self.iso_encode('iso-8859-16'),
            "Windows-1250 Encoding": self.windows_encode('windows-1250'),
            "Windows-1251 Encoding": self.windows_encode('windows-1251'),
            "Windows-1252 Encoding": self.windows_encode('windows-1252'),
            "Windows-1253 Encoding": self.windows_encode('windows-1253'),
            "Windows-1254 Encoding": self.windows_encode('windows-1254'),
            "Windows-1255 Encoding": self.windows_encode('windows-1255'),
            "Windows-1256 Encoding": self.windows_encode('windows-1256'),
            "Windows-1257 Encoding": self.windows_encode('windows-1257'),
            "Windows-1258 Encoding": self.windows_encode('windows-1258'),
            "Binary Encoding": self.binary_encode(),
            "Hex Encoding": self.hex_encode(),
            "Octal Encoding": self.octal_encode(),
            "Decimal Encoding": self.decimal_encode(),
            "JavaScript Encoding": self.js_encode(),
            "JavaScript Unicode Encoding": self.js_unicode_encode(),
            "JavaScript Octal Encoding": self.js_octal_encode(),
            "Double URL Encoding": self.double_encode(self.url_encode),
            "Double HTML Entity Encoding": self.double_encode(self.html_entity_encode),
            "Double Unicode Encoding": self.double_encode(self.unicode_encode),
            "Double ASCII Encoding": self.double_encode(self.ascii_encode),
            "Base64 Encoding": self.base64_encode(),
            "Base32 Encoding": self.base32_encode(),
            "Base16 Encoding": self.base16_encode(),
            "Punycode Encoding": self.punycode_encode(),
            "ROT13 Encoding": self.rot13_encode(),
            "Caesar Cipher Encoding": self.caesar_cipher_encode(3),
            "XOR Encoding": self.xor_encode(42),
            "Null Character Encoding": self.null_character_encode(),
            "Tab Character Encoding": self.tab_character_encode(),
            "Newline Character Encoding": self.newline_character_encode(),
            "Carriage Return Character Encoding": self.carriage_return_character_encode(),
            "Space Character Encoding": self.space_character_encode(),
        }
        return encodings

    def null_character_encode(self):
        return '\0'.join(self.payload)

    def tab_character_encode(self):
        return '\t'.join(self.payload)

    def newline_character_encode(self):
        return '\n'.join(self.payload)

    def carriage_return_character_encode(self):
        return '\r'.join(self.payload)

    def space_character_encode(self):
        return ' '.join(self.payload)

def main():
    # Print tool name and author
    print("Encoder Tool by Flying_Eagle")

    # Set up argument parser
    parser = argparse.ArgumentParser(description='Encoder Tool for various encoding schemes.')
    parser.add_argument('-p', '--payload', type=str, required=True, help='The payload to encode')
    parser.add_argument('-v', '--verbose', action='store_true', help='Increase verbosity')

    args = parser.parse_args()

    # Create an Encoder instance
    encoder = Encoder(args.payload)
    encoded_payloads = encoder.get_encoded_payloads()

    # Print encoded payloads
    if args.verbose:
        for scheme, encoded in encoded_payloads.items():
            print(f"{scheme}: {encoded}")
    else:
        print(encoded_payloads)

if __name__ == "__main__":
    main()
