## Prerequisites
- Go installed on the local machine
- MySQL intalled on the local machine

### Installation
1. Clone this repository:
    `git clone https://github.com/danielwetan/kdigital-backend`
2. Database configuration:
    * Open http://localhost/phpmyadmin in the browser
    * Create a new table with the name `kdigital-backend`
    * Import database to current table, select `kdigital-backend.sql` file from project folder
3. Install depedencies:
    `cd kdigital-backend & go get -u`
4. Start the server:
    `go run main.go`
5. Test the API -> `Please open api.http file`