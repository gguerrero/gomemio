# GoMemIO
A GoLang experiment to create a TCP server to store key/value pairs in memory (much REDIS like).

## Commands
Basic commands to read/write values on Memory:

### Write values
```bash
SET max 100
> OK

SET min 10
> OK

SET name "Uncle Bob"
> OK
```

### Read values
```bash
GET max min
> 100, 10

GET name
> Uncle Bob
```

## Persistence
...

## Configuration
...
