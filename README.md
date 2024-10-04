# 403Evasion

403Evasion is a security tool written in Go designed to bypass *HTTP 403 Forbidden* responses. It allows penetration testers and security researchers to test access controls and bypass restrictions using custom headers and various URL encodings. The tool implements two primary attack strategies:

* Header Injection Bypass: Sends HTTP requests with custom headers (like User-Agent, Referer, etc.) to try bypassing the 403 restriction. You can specify a wordlist of headers to test different possibilities.
* URL Encoding Bypass: Appends various encoded payloads to URLs and sends requests to discover if alternative URL encodings can bypass restrictions.
This tool can be useful in web application testing where legitimate resources are blocked due to server-side access control mechanisms.

## Features
Header-based bypass attack using a Pre-designed wordlist of headers.
URL encoding attack to trick the server using different encoded payloads.
Output highlighting for successful or failed attempts.
Easy-to-use and customizable.
## Installation

#### Clone the Repository:
```bash
git clone https://github.com/DaddyUchiha/403evasion
```

#### Navigate to the project directory:
```bash
cd 403evasion
```
#### Run the Script:
```bash
go run 403evasion.go
```
#### Follow the Prompts:

The tool will perform requests based on your input and display the response status code for each attempt, highlighting successful 403 bypasses.
## Usage
When you run the tool, you’ll be presented with two options:

* Header Bypass Injection: This option allows you to perform header-based bypass attempts. You can use a wordlist containing various HTTP headers to test for potential bypasses.
* URL Bypass 403: This option performs URL-based bypassing by appending encoded payloads to the URL and checking for a 403 bypass.

### Choose Attack Option:

The tool will ask for an attack option:

 * Header Bypass Injection : [1]
 * URL Bypass 403 :	       [2]					                            

### Enter the target URL : http://example.com

## Contributions
Feel free to contribute to the project by submitting pull requests or opening issues. Whether it's new features, bug fixes, or general improvements, any help is appreciated!

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgements
Thanks to the security research community for the numerous methods and ideas for bypassing 403 restrictions.
