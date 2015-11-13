package main

import (
	"fmt"
	"math/big"
	"UJSPackage/crypto"
	"UJSPackage/file"
	"os"
	"bufio"
	"io/ioutil"
)


func main() {
	var input string = ""
	// Ask user for what they want to do
	for input != "4" {
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("1) Create new keypair")
		fmt.Println("2) Encrypt Message")
		fmt.Println("3) Decrypt Message")
		fmt.Println("4) Quit")
		fmt.Print("> ")
		fmt.Scanln(&input)
		// Call appropriate function
		switch {
			// Create new keypair
			case input == "1" :
				generateKeyPair()
				break
			// Encrypt message with public key
			case input == "2" :
				encryptMessage()
				break
			// Decrypt message with private key
			case input == "3" :
				decryptMessage()
				break
			case input == "4" :
				break
			default :
				fmt.Println("Invalid input")
				break
		}
	}
}

// Function to read in a file and encrypt it to a ciphertext
func encryptMessage() {
	// Strings for user input
	var messageFile, keyFile string;
	messageFile = "";
	keyFile 	= "";
	// Collect said input
	fmt.Println("\nEnter file containing key");
	fmt.Print("> ");
	fmt.Scanf("%s\n", &keyFile);
	fmt.Println("\nEnter file to encrypt");
	fmt.Print("> ");
	fmt.Scanf("%s\n", &messageFile);
	// Read the file provided by user and store
	//  it in a byte slice
	input, err	:= ioutil.ReadFile(messageFile)
	// If ReadFile encountered an error, let the user know
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	// Turn our input into a slice of big Ints
	//  First we need to typecast it to a string, which
	//  is essentially a free operation in Go
	plain 		:= stringToBigInts(string(input))
	// Read in the public key provided by user
	n, e, err 	:= readKey(keyFile);
	// If readKey encountered an error, let the user know
	if err != nil {
		return;
	}
	// new slice of big Ints to hold our encrypted text
	ciphers 	:= make([]*big.Int, 0)
	// encrypt each of our plaintext chunks
	for _, v := range plain {
		ciphers = append(ciphers, big.NewInt(0).Exp(v,e,n))
	}
	// write the encrypted data to a file
	writeCipher(ciphers, "secret.message");
}

// Function to decrypt messages from and encrypted file that will
//  be specified by the user
func decryptMessage() {
	// Strings to store user input
	var cipherFile, keyFile string;
	cipherFile = "";
	keyFile = "";
	// Collect said input
	fmt.Println("\nEnter file containing key");
	fmt.Print("> ");
	fmt.Scanf("%s\n", &keyFile);
	fmt.Println("\nEnter file containing ciphertext");
	fmt.Print("> ");
	fmt.Scanf("%s\n", &cipherFile);
	// Read in ciphertext and private key file
	ciphers 	:= readCipher(cipherFile);
	n, d, err 	:= readKey(keyFile);
	// If there was a file error, just return
	if err != nil {
		return;
	}
	// Make a slice of big Ints to hold our plaintext chunks
	plain 		:= make([]*big.Int, 0)
	// Decrypt each of our cipher blocks into plaintext big Ints
	for _, v 	:= range ciphers {
		plain	= append(plain, big.NewInt(0).Exp(v,d,n))
	}
	// Convert big Ints back into the original text
	output	:= bigIntsToString(plain)
	// Write the output to a file
	writePlainText(output, "plain.message")
}

// Function to generate an RSA private/public keypair
func generateKeyPair() {
	// We're working with big Ints here so we can have large enough keys
	var p,q,n,phi,e,d *big.Int;
	// Find a couple of big prime numbers about 1024 bits in length
	p = crypto.BigPrime(1024);
	q = crypto.BigPrime(1024);
	// Multiply them to create n, which should be about 2048 bits
	n = big.NewInt(0).Mul(p,q);
	// Generate phi(n)
	phi = big.NewInt(0).Mul(big.NewInt(0).Sub(p,big.NewInt(1)),
							big.NewInt(0).Sub(q,big.NewInt(1)));
	// We're just going to pick this arbitrary prime number for our e value
	e = big.NewInt(65537);
	// get the modular inverse of e
	d = crypto.ModInverse(e, phi)
	// Write the two keys to files	
	writeKey(n, e, "public.key");
	fmt.Println("\nPublic key written to public.key");
	
	writeKey(n, d, "private.key");
	fmt.Println("Private key written to private.key\n");

}

// Function to write encryption keys into a file
func writeKey(modulus, exponent *big.Int, fileName string) {
	// Create writer and check for error
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("\nError, could not open file!")
		return
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	// Write out the modulus and exponent with a little formatting to make
	//  it more readable
	w := bufio.NewWriter(file)
	w.WriteString("----------Modulus----------\n")
	w.WriteString(fmt.Sprintln(modulus))
	w.WriteString("--------End Modulus--------\n")
	w.WriteString("----------Exponent---------\n")
	w.WriteString(fmt.Sprintln(exponent))
	w.WriteString("-------End Exponent--------\n")
	w.Flush()
}

// Function to read encryption keys from a file
func readKey(fileName string) (*big.Int, *big.Int, error) {
	// Get the contents of file as a slice of strings
	keyStrings, err := file.FileLines(fileName)
	// If we ran into an i/o error, tell user and return
	if err != nil {
		fmt.Println("\nError, key could not be read from file")
		return nil, nil, err
	}
	// indices 1 and 4 of the keyfile will hold the modulus and exponent
	modulus, _ := big.NewInt(0).SetString(keyStrings[1], 0)
	exponent, _ := big.NewInt(0).SetString(keyStrings[4], 0)
	// Return to the user
	return modulus, exponent, nil
}

// Function to write encrypted texts to a file
func writeCipher(message []*big.Int, fileName string) {
		// Create writer and check for error
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("\nError, could not open file!\n");
		return;
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	// Create a writer on the opened file and write the encrypted
	// data to it with some formatting
	w := bufio.NewWriter(file);
	w.WriteString("----------Message----------\n");
	// iterate over the slice of encrypted messages and write them
	//  on separate lines with dividers
	for _, v := range message {
		w.WriteString(fmt.Sprintln(v));
		w.WriteString("-----------------------------------\n");
	}
	w.WriteString("--------End Message--------\n");
	w.Flush();
	fmt.Printf("\nMessage written to file %s\n", fileName);
}

// Function to write our decrypted plaintexts to a file
func writePlainText(message string, fileName string) {
		// Create writer and check for error
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("\nError, could not open file!\n");
		return;
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	// Output the string to a file, with a little formatting
	w := bufio.NewWriter(file);
	w.WriteString("----------Message----------\n");
	w.WriteString(fmt.Sprintln(message));
	w.WriteString("--------End Message--------\n");
	w.Flush();
	fmt.Printf("\nMessage written to file %s\n", fileName);
}

// Function to read in an encrypted ciphertext from a file
func readCipher(fileName string) []*big.Int {
	// Read the file in as a slice of strings
	cipherStrings, err	:= file.FileLines(fileName)
	// If there was an error, let the user know and return
	if err != nil {
		fmt.Println("\nError, ciphertext could not be read from file!");
		return nil
	}
	// Create a slice of big Ints to hold our ciphertext
	cipherText 	:= make([]*big.Int, 0)
	// Every second line, except the last, should hold an encrypted
	//  chunk of data - convert them to big Ints and store them
	for i := 1; i < len(cipherStrings) - 2; i += 2 {
		next,_		:=	big.NewInt(0).SetString(cipherStrings[i], 0)
		cipherText	= append(cipherText, next)
	}
	return cipherText
}

// Function to convert a string of arbitrary length into a slice of
//  big Ints with bitlength 2000
func stringToBigInts(input string) []*big.Int {
	// Some variables to hold our data
	var index int
	bigInts 	:= make([]*big.Int, 0)
	// convert input to slice of bytes
	bytes 		:= []byte(input)
	// take chunks of 250 bytes and convert to bigInts
	for index = 0; index + 250 < len(bytes); index += 250 {
		bigInts = append(bigInts, big.NewInt(0).SetBytes(bytes[index:index+250]))
	}
	bigInts = append(bigInts,big.NewInt(0).SetBytes(bytes[index:len(bytes)]))
	// return slice
	return bigInts
}

// Function to take a slice of big Ints, convert them, and
//  concatenate them into a single string
func bigIntsToString(bigInts []*big.Int) string {
	bytes 	:= make([]byte, 0)
	// convert bigInts to byte slices
	// concatenate the slices together
	for _, v 	:= range bigInts {
		bytes = append(bytes, v.Bytes()...)
	}
	// convert slice of bytes to output string
	output	:= string(bytes)
	// return string
	return output
}