# Encoder Tool

================

A command-line utility that takes a payload as input and encodes it using various encoding schemes.

### Overview

The Encoder Tool is a Python script designed to encode various types of input strings using multiple encoding schemes. It can handle links, symbols, emails, numbers, and characters, providing a wide range of encoding options such as URL encoding, HTML entity encoding, Base64, and more.

### Installation

Clone the repository:

     git clone https://github.com/FlyingEagl3/Encoder.git
     cd Encoder
     
Ensure you have Python 3.x installed on your system.

Install any required libraries if not already available:

    pip install idna

### Usage

To use the Encoder tool, simply run the command with the payload as an argument:


    python Encoder.py -p "YourPayloadHere" -v

For example:

    python Encoder.py -p "https://example.com" -v


This will encode the payload using all supported encoding schemes and print the encoded payload for each scheme.

The Encoder tool supports the following options:

    
    -p: [payload] The string you want to encode.
    -v: (Optional) Increases verbosity, printing the name of each encoding scheme along with the encoded payload.


Requirements
                 
    Python 3.x


Required libraries:

    argparse
    urllib
    html
    base64
    codecs
    idna
    re

Vigenère Cipher Encoding: Not implemented in the provided script.

### Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue for any bugs or feature requests.

### Author

The Encoder tool was written by [Flying_Eagle](https://x.com/flying_eagl3).
