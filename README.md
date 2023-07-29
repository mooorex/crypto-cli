# crypto-cli
A command-line interface to securely transfer important data using [ECIES](https://www.secg.org/).

## Examples
Help command:
```
$./cypto-cli --help
NAME:
   crypto-cli - Secruely transfer important data using ecies.

USAGE:
   crypto-cli [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   keygen, g, kg   Generate an ecies key pair
   encrypt, e, en  Encrypt a plaintext to a ciphertext using ecies
   decrypt, d, de  Decrypt a ciphertext to a plaintext using ecies
   help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
