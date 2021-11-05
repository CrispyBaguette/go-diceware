# go-diceware
Dirty, poorly written, and ineffcient diceware passphrase generator

Built mainly on code shamelessly copypasted from StackOverflow.

Expects a wordlist with a two column, whitespace-separated format (see [this file](https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt), also available on IPFS [here](https://ipfs.io/ipfs/bafybeibxmzvzztddr4boxbcfmbuv3zi2uvl7rowoa2v6kg5qfnzsyx6nj4/eff_large_wordlist.txt)). The first column is the result of 5 die rolls, and the second is the associated word. 

# Example
```bash
git clone "https://github.com/CrispyBaguette/go-diceware.git"
cd go-diceware
wget "https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt"
go run . --length 6 --wordlist "eff_large_wordlist.txt"
```
