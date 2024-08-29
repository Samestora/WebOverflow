# Trivia Web  
Change `.env.example` to `.env`

### Requirements  
- Go 1.23  
- go migrate on the host system  
```bash
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
- Docker  

### Usage  
To run WebOverflow :  
```bash
    sudo make all
```
