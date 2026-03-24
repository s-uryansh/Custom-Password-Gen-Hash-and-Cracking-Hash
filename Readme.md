# Assignment 2
# CSD356 Foundation of Information Security

## Submitted by: Suryansh Rohil

# Algorithm Approach

generate.go uses DFS and backtracking to generate all possible permutation of base word with leetMap (Detail explanation in generate.go)

# Usage

`git clone https://github.com/s-uryansh/Custom-Password-Cracking-Pipeline-Leet-Based-Generation-SHA-1-Hashing-Hashcat-Analysis.git`

`cd Assignment2`

`go mod tidy`

### Generate Passwords and Hashes

`go run generate.go`

### HashCat Usage

`hashcat -m 100 -a 0 hashes.txt dictionary.txt -o recovered.txt`

- `-m 100`: Sets the hash type to SHA-1.
- `-a 0`: Sets the attack mode to "Straight" => Dictonary Attack.
- `hashes.txt`: Hash list to target.
- `dictonary.txt`: Word list being used to attack.
- `-o recovered.txt`: Saves the cracked hashes and plain text.
