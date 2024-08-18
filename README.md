# Den Golang Aristkur Template 

A golang template architecture using fiber framework version 2, the goal of this architecture is to make it easy to create rest api based services using the golang programming language as easy as using php, nodejs or python

## Structur Directory

.
├── config
│   └── config.go
├── entity
│   └── *.go
├── migrate
│   ├── db.go
│   ├── seed.go
├── test
│   └── *_test.go
├── main.go
└── README.md


- Config/config.go  : Place for all global varible configuration 
- Controller/*.go   : Place for all logic process
- Entity/*.entiy    : Place for all model like strukur data
- Migrate/db.go     : Place for auto migration create database and table by reference model  in entity directory
- Migrate/seed.go   : Place for auto populate date di database
- Test              : Place for unit test each controller 

## Installation

Steps to install this framework:

1. Clone the repository:
   git clone https://github.com/dendie-sanjaya/golang-arsitekur-template

##  Contact

If you have question, you can contact this email   
Email: dendie.sanjaya@gmail.com


License

This project is licensed under the MIT License.
