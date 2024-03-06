# ccWebServer

This is my implementation of the Web Server project from [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-webserver)

## Building

To build the Web Server use the following command

'''bash
go build -o ccWebServer ./cmd/main.go
'''

## Configuration

The Web Server takes the running address and then www directory path from a config.yaml file.
Amend the example_config.yaml with your details and rename it to config.yaml for it to be read by the Web Server.

The config.yaml file must be kept in the location that the binary is being run from.

## Usage

After building the binary, setting up the www directory in your location of choice and configuring the config.yaml file run the command below to start the server.

'''bash
./ccWebServer
'''
