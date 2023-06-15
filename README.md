## Diplomski-steam-clone-BE
School project for my BSc thesis

## Prerequisites
Go 1.19 or later installed
Gin framework (github.com/gin-gonic/gin) installed

## Installation
Clone the repository: git clone https://github.com/stonalampa/diplomski-steam-clone-BE
Change into the project directory: $ cd diplomski-steam-clonse-BE

## Usage
There are 2 flags:
1. local || deployment <- this will choose if we start the local server or we connect to deployed env
2. true || false <- true will run seeding, false will run the server
3.
## Start server with
go run main.go local (or deployment) false

## Start seeds with
go run main.go local (or deployment) true
