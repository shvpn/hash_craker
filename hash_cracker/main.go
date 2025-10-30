package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/shvpn/md5cracker/utils/crack"
)

func main() {
	wordlistPath := flag.String("wordlist", "pass.txt", "path to password wordlist file")
	targetHash := flag.String("hash", "6a85dfd77d9cb35770c9dc6728d73d3f", "target MD5 hash to crack")
	verboseOut := flag.String("out", "", "optional path to save verbose output (e.g., verbose.txt); if empty, prints to stdout only")
	hashType := flag.String("type", "md5", "hash type: md5, sha1, sha512")
	quitOnFound := flag.Bool("quit", true, "quit when a match is found (true) or continue scanning (false)")
	flag.Parse()

	// Open wordlist file
	f, err := os.Open(*wordlistPath)

	if err != nil {
		log.Fatalf("failed to open wordlist %q: %v", *wordlistPath, err)
	}
	defer f.Close()

	// Optional verbose output file
	var outFile *os.File
	fmt.Println(*verboseOut)
	if *verboseOut != "" {
		outFile, err = os.Create(*verboseOut)
		if err != nil {
			log.Fatalf("failed to create output file %q: %v", *verboseOut, err)
		}
		defer outFile.Close()
	}

	printVerbose := func(s string) {
		fmt.Println(s)
		if outFile != nil {
			_, _ = outFile.WriteString(s + "\n")
		}
	}

	printVerbose(fmt.Sprintf("Starting MD5 crack: wordlist=%s target=%s", *wordlistPath, *targetHash))

	scanner := bufio.NewScanner(f)

	lineNo := 0
	found := false

	for scanner.Scan() {
		lineNo++
		// Trim whitespace and any leftover CR/LF
		password := strings.TrimSpace(scanner.Text())
		if password == "" {
			continue
		}
		var hash string
		switch *hashType {
		case "md5":
			hash = crack.Hash(password)
		case "sha1":
			hash = crack.HashSHA1(password)
		case "sha512":
			hash = crack.HashSHA512(password)
		}

		// Verbose: show attempt
		printVerbose(fmt.Sprintf("line %d: try '%s' -> %s", lineNo, password, hash))

		if hash == strings.ToLower(*targetHash) {
			printVerbose("======================================")
			printVerbose(fmt.Sprintf("FOUND! password = %q", password))
			printVerbose(fmt.Sprintf("md5(password) = %s", hash))
			printVerbose(fmt.Sprintf("line: %d", lineNo))
			printVerbose("======================================")
			found = true
			if *quitOnFound {
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading wordlist: %v", err)
	}

	if !found {
		printVerbose("Finished: no match found.")
	}
}
