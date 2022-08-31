# demo-golang
For demo Golang researching

## Setup Golang
- Golang can be setup here: https://go.dev/learn
- Choose your OS, download the installer and install the application.

## Run project
- Clone the project
- In the terminal at the root of project, run following command to start the application
`go run .`
- Golang will get and download all dependencies that the application required. Then run the application on debug mode.

## Compile and install project
- You can read a documentation that present how to build and install a golang application here: https://go.dev/doc/tutorial/compile-install

### Demo

## APIs:
- The APIs will run at http://localhost:8080
1. Get albums: [GET] http://localhost:8080/albums
2. Get album detail: [GET] http://localhost:8080/albums/{id} (id = 1 to 4)
3. Add album: [POST] http://localhost:8080/album
