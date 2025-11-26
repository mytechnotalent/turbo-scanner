![image](https://github.com/mytechnotalent/turbo-scanner/blob/main/turbo-scanner.png?raw=true)

## FREE Reverse Engineering Self-Study Course [HERE](https://github.com/mytechnotalent/Reverse-Engineering-Tutorial)

<br>

# turbo-scanner
A port scanner and service detection tool that uses 1000 goroutines at once to scan any hosts IP or FQDN with the sole purpose of testing your own network to ensure there are no malicious services running.

<br>

# Windows
```bash
usage: turbo-scanner_010w.exe localhost
```

# MAC
```bash
usage: ./turbo-scanner_010m localhost
```

# Linux
```bash
usage: sudo ./turbo-scanner_010l localhost
```

<br>

## Terms Of Use
* Do NOT use this on any computer you do not own or are not allowed to run this on.<br>
* You may NEVER attempt to sell this, it is free and open source.<br>
* The authors and publishers assume no responsibility.<br>
* For educational purposes only.

<br>

## Run Tests
```bash
go test -v -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

<br>

## License
[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)
