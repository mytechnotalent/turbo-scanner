<a rel="me" href="https://defcon.social/@kevinthomas">Mastodon</a>
<a rel="me" href="https://ioc.exchange/@kevinthomas"></a>
<a rel="me" href="https://infosec.exchange/@kevinthomas"></a>
<a rel="me" href="https://windows11sucks.com/@kevinthomas"></a>

![image](https://github.com/mytechnotalent/turbo-scanner/blob/main/turbo-scanner.png?raw=true)

## FREE Reverse Engineering Self-Study Course [HERE](https://github.com/mytechnotalent/Reverse-Engineering-Tutorial)

<br>

# turbo-scanner
A port scanner and service detection tool that uses 1000 goroutines at once to scan any hosts's ip or fqdn with the sole purpose of testing your own network to ensure there are no malicious services running.

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

## Terms Of Use
* Do NOT use this on any computer you do not own or are not allowed to run this on.<br>
* You may NEVER attempt to sell this, it is free and open source.<br>
* The authors and publishers assume no responsibility.<br>
* For educational purposes only.

## Run Tests
```bash
go test -v -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)
