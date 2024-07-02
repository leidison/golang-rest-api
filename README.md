# Install air package

```sh
go install github.com/air-verse/air@latest
```

Adicione no bashrc o comando air

```sh
nano ~/.bashrc

alias air='$(go env GOPATH)/bin/air'
```

Inicialize o servidor com o comando:
```sh
air
```