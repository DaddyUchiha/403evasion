package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"net/url"
)

	
const (
	blue = "\x1b[34m"
	reset = "\x1b[0m"
)
func main() {
	logo := `  _    _                     _                         
 | |  | |                   | |                        
 | |__| |  ___ __  __ _ __  | |  ___   _ __  ___  _ __ 
 |  __  | / _ \\ \/ /| '_ \ | | / _ \ | '__|/ _ \| '__|
 | |  | ||  __/ >  < | |_) || || (_) || |  |  __/| |   
 |_|  |_| \___|/_/\_\| .__/ |_| \___/ |_|   \___||_|   
                     | |                               
                     |_|                               `
																												   
	fmt.Println("")
	fmt.Println(string(logo))		
	fmt.Println("")																								   	
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("| 1: Header Bypass Injection   			 		    |")
	fmt.Println("| 2: URL Bypass 403						    |")
	fmt.Println("---------------------------------------------------------------------")	
	fmt.Printf("Choose Attack Option: ")
	read := bufio.NewReader(os.Stdin)
	reader,err := read.ReadString('\n')
	if err != nil{
		fmt.Println(err)
	}
	reader = strings.TrimSpace(reader)
	if reader == "1"{
		bypass()
	}else if reader == "2"{
		encoder()
	}else{
		fmt.Println("No Valid option")
		fmt.Println("Please Choose From Above Options!")
		os.Exit(0)
	}

}


func bypass(){
	read := bufio.NewReader(os.Stdin)
	fmt.Printf("%s","Enter your URL here : ")
	reade,_ := read.ReadString('\n')
	reade = strings.TrimSpace(reade)

	if reade == ""{
		fmt.Println("Invalid or Empty Value")
	}

	// Headers Injection File
	// Open the text file containing the headers
	file, err := os.Open("HeadersAttack.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	blue := "\x1b[34m"
	reset := "\x1b[0m"
	fmt.Println("")
	fmt.Println("==========================================")
	fmt.Println(blue,"Starting Headers Attack on URL",reset)
	fmt.Println("==========================================")
	fmt.Println("")

	// Define the URL for testing
	url,err := url.Parse(reade)
	if err != nil{
		fmt.Println("INVALID URL",err)
	}

	// Read the headers from the file one by one
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into key and value based on the colon
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			// Create a new HTTP request for each header
			req, err := http.NewRequest("GET", url.String(), nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}
			
			// Set only one header at a time
			req.Header.Set(key, value)
			req.Header.Set("User-Agent", "Mozilla/5.0")

			// Create an HTTP client and send the request
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error making request:", err)
				return
			}
			defer resp.Body.Close()

			time.Sleep(1000 * time.Millisecond)
			red := "\x1b[31m"
			green := "\x1b[32m"
			reset := "\x1b[0m"

			// Output the response status and header information
			fmt.Printf("Request Header: [ %s: %s ]\n", key, value)
			fmt.Printf("Request User-Agent : %s\n", req.UserAgent())
			fmt.Println("Response Status:", resp.Status)
			

			//Check Response Status Code 
			if resp.Status == "200 OK"{
				fmt.Println(green,"Successfull Bypassed [403]",reset)
			}else{
				fmt.Println(red,"[ Not Bypassed ]",reset)
			}

			// Print a separator between requests for readability
			fmt.Println("--------------------------------------------------")
		}
		
	}
	

	// Handle any errors during file reading
	if err := scanner.Err(); 
	err != nil {
		fmt.Println("Error reading file:", err)
	}
}


func encoder(){
	//fmt.Println("Enter your URL like this : 'example.com'")
	fmt.Printf("Enter you URL : ")
	reader := bufio.NewReader(os.Stdin)

	read,_ := reader.ReadString('\n')

	read = strings.TrimSpace(read)

	if read == ""{
		fmt.Println("Empty or Invalid URL")
	}

	by,err := os.Open("403bypassencode.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer by.Close()
	blue := "\x1b[34m"
	red := "\x1b[31m"
	green := "\x1b[32m"
	reset := "\x1b[0m"
	fmt.Println("")
	fmt.Println("==========================================")
	fmt.Println(blue,"*STARTING URL ATTACK ON URL*",reset)
	fmt.Println("==========================================")
	fmt.Println("")


		// Loop to read the file line by line
		readerr := bufio.NewScanner(by)
		for readerr.Scan(){
			tarr := readerr.Text()

			u, err := url.Parse(read)
			if err != nil {
				fmt.Println("Error parsing URL:", err)
				break
			}
			u.Path += tarr // Append the payload to the URL path
	
			req, err := http.NewRequest("GET", u.String(), nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				continue
			}
	
		client := &http.Client{}
		resp ,err := client.Do(req)
		if err != nil{
			fmt.Println("Error while the sending request")
			fmt.Println("Check your URL!")
			return
		}
		defer resp.Body.Close()

		time.Sleep(1000 * time.Millisecond)

		
		fmt.Println("===============================================")
		if resp.StatusCode != 200 {
			fmt.Printf("%v %v [%v] %v\n",red,read+tarr, resp.StatusCode, reset)
		}else if resp.StatusCode == 200 {
			fmt.Printf("%v %v [%v] %v \n",green,read+tarr, resp.StatusCode, reset)
		}
		
	}

}

