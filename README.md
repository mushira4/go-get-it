# go-get-it
Go + Redis app that allows to get and cache information via an API

## Dependencies
- go dep - Used to manage 3rd party libs dependencies
- github.com/garyburd/redigo - Redis lib, used to comunicate with redis
- github.com/ghodss/yaml - Yaml file reader lib, used to reaad the configuration file

## Lifecycle
This project uses the Makefile to manage this project lifecycle  
- Use `make dependecies` to download the libraries the project uses
- Use `make test` to run the app tests the
- Use `make install` to generate the go executable file

## How to setup
 1. Configure your `$GOPATH` to your go workspace is;
 2. Execute the `make dependencies` command to dowload the project dependencies to your go workspace;
 3. Execute the `make install` command to generate the app executable
 4. Start a Redis server
 5. Configure the local.yaml setting the redis client properties.
 6. Run the executable generated at the step 3

## Project API

| Resource       | Method | Description                                                                                           |
|----------------|--------|--------------------------------------------------------------------------------------------------------|
|/save           |(POST)  | Save all parameters that is sent at the redis storage.                                                |
|/retrieve       |(GET)   |Get all the registers that the server has (Sorry it uses KEYS redis command, dont use it in production)|
|/retrieve?query |(GET)   |Get registers that matches the pattern (Sorry it uses KEYS redis command, dont use it in production)|

### Samples
/save - request sample
`{
    "aa" : "a",
    "ab" : "b",
    "CC" : "c",
    "cD" : "d"
}`

## Logs
The project uses the default go lang logger. If you want log rotation, you can use the default linux logrotate to do this.

In case of any doubt, mail me: mushira4@gmail.com    =)
