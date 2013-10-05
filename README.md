pgen
====

Password generator

[![Build Status](https://travis-ci.org/JonPulfer/pgen.png)](https://travis-ci.org/JonPulfer/pgen)

Overview
--------

Generates a random string suitable for use as a password. The default
behaviour is to produce a string of 12 chars. This can be controlled
by passing a number to the command.

`pgen --pwdlen 15` will produce a string of 15 chars. 

Each char in the password will be unique. This limits the length to 92 chars.


