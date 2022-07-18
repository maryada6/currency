# Currency

## Problem description

A library to model money which can be expressed in Rupees, paise or a combination thereof, that we can add values of money.
>As a fan of wealth,
I want to model money which can be expressed in Rupees, paise or a combination thereof,
So that I can add values of money.

## Development environment setup

After checking out the repo, import `github.com/stretchr/testify` to use `assert`.

### Prerequisite

- Go (1.18.2)

## Build instructions

To build this project use the following command:

    $ go build


## Test instructions

To run the test use the following command:

    $ go test ./...

## Library usage

__Money entity__, To use money entity import it in the required file. Create money object as:

    object name := NewMoney( rupee, paise )

To use __Balance__ function to get the current balance(moneyOne and moneyTwo to be objects of type Money):

    money.Balance()
_Returns the balance amount as __rupee.paise__._

To use __Equals__ function to compare two moneys:

    moneyOne.Equals(moneyTwo)
_Returns True or False depending on equality._

To use __Add__ function to add two moneys:

    moneyOne.Add(moneyTwo)
_The sum is stored in moneyOne._

To use __Subtract__ function to subtract one money from another:

    moneyOne.Subtract(moneyTwo)
_The difference is stored in moneyOne._
