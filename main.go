package main

import (
	"bufio"
	crand "crypto/rand"
	"flag"
	"log"
	"math/big"
	"os"
	"strings"

	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func dieThrow(sides *big.Int) (*big.Int, error) {
	randValue, err := crand.Int(crand.Reader, sides)
	if err != nil {
		return nil, fmt.Errorf("throwing die: %w", err)
	}
	return randValue.Add(randValue, big.NewInt(1)), nil
}

func diceWordKey() (string, error) {
	var b strings.Builder
	var err error
	b.Grow(5)
	for i := 0; i < 5 && err == nil; i++ {
		throw, e := dieThrow(big.NewInt(6))
		if e != nil {
			return "", fmt.Errorf("building dw key: %w", err)
		}
		fmt.Fprintf(&b, "%v", throw)
	}
	return b.String(), nil
}

func parseWordList(fileName string) (dwMap map[string]string, err error) {
	dwMap = make(map[string]string)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		dwMap[fields[0]] = fields[1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dwMap, err
}

func main() {
	// Simulating dice has got to be the most inefficient way to select
	// random words from a list

	viper.SetEnvPrefix("diceware")
	viper.BindEnv("length")
	viper.BindEnv("wordList")
	flag.Int("length", 6, "length of passphrase")
	flag.String("wordlist", "wordlist.txt", "die-throw indexed wordlist")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	dwMap, err := parseWordList(viper.GetString("wordlist"))
	if err != nil {
		log.Fatal(err)
	}

	passPhrase := make([]string, 0)

	for i := 0; i < viper.GetInt("length"); i++ {
		key, err := diceWordKey()
		if err != nil {
			log.Fatal(err)
		}
		passPhrase = append(passPhrase, dwMap[key])
	}

	for _, v := range passPhrase {
		fmt.Printf("%v ", v)
	}
	fmt.Print("\n")
}
