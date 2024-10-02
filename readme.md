# assignment-go-rest-api

## API-Documentation
access the documentation [here](https://documenter.getpostman.com/view/16000432/2s9YJhwKdt)

## entity-relationship-diagram
<a href="https://drive.google.com/file/d/1mr5Td9mu9GFAAel0iq4SbQ9U4x4i9QOC/view?usp=sharing" >ERD Link</a>
<img src="https://drive.google.com/file/d/1mr5Td9mu9GFAAel0iq4SbQ9U4x4i9QOC/view?usp=sharing">


## Running the apps
- copy and populate the `env.example`
```bash
cp .env.example .env
```
- run the server
- when the server running in ENV=dev, it will automatically create table and seed the database
```go
go run main.go
```
- to run the test
```go
go test ./...
```
## Assumptions
1. there are 9 level of prize for gacha games, user have 5% chance to get level 1, 55% chance to get level 2, 20% chance to get level 3, 10% chance to get level 4, 2% chance to get level 5, 1.5% chance to get level 6, 1% chance to get level 7, 0.4% chance to get level 8, 0.1% chance to get level 9
2. different level of prize in gacha games has different prize, level 1 got 0 rupiah, level 2 got 10000 rupiah, level 3 got 20000 rupiah, level 4 got 50000 rupiah, level 5 got 100000 rupiah, level 6 got 150000 rupiah, level 7 got 200000 rupiah, level 8 got 250000 rupiah, level 9 got 300000 rupiah.
3. the truth is, it doesn't matter what number the user did choose, the system will choose the prize randomly based on the percentage, just like a "<b>real life cases</b>"
4. the minimum length of password is 6 chars
5. the maximum length of password is 20 chars
6. the minimum amount of top-up is 50000
7. the maximum amount of top-up is 10000000
8. the minimum amount of transfer is 1000
9. the maximum amount of transfer is 50000000
10. the maximum length of transfer description is 35 char
11. forget password token will expired in 15 minutes
12. the wallet number format is 4200000000000