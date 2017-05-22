# go-get-it
Go + Redis app that allows to get and cache information via an API

This project uses the Makefile to manage this project lifecycle  
- Use `make dependecies` to download the libraries the project uses
- Use `make install` to generate the go executable file

## How to setup
 1. Configure your $GOPATH to your go workspace is;
 2. Execute the `make dependencies` command to dowload the project dependencies to your go workspace;
 3. Execute the `make install` command to generate the app executable
 4. Start a Redis server
 5. Configure the local.yaml setting the redis client properties.
 6. Run the executable generated at the step 3

## Project API

| Resource | Method | Description                                                                                           |
|----------|--------|--------------------------------------------------------------------------------------------------------|
|/save     |(POST)  | Save all parameters that is sent at the redis storage.                                                |
|/retrieve |(GET)   |Get all the registers that the server has (Sorry it uses KEYS redis command, dont use it in production)|


In case of any doubt, mail me: mushira4@gmail.com    =)
