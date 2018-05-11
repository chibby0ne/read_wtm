# Autoswitcher for GPU mining scripts 
[![Build Status](https://travis-ci.org/chibby0ne/read_wtm.svg?branch=master)](https://travis-ci.org/chibby0ne/read_wtm)
[![Coverage Status](https://coveralls.io/repos/github/chibby0ne/read_wtm/badge.svg?branch=master)](https://coveralls.io/github/chibby0ne/read_wtm?branch=master)

Let's say you have a GPU mining rig and you want to mine the most profitable
coin.
Since this might change day to day (or even hour to hour), this implies you
would have to constantly check for the most profitable one, remote access your
miner, close the current running mining program, and open the other more
profitable coin.

This program automates just that process, gathering using the wonderful
[whattomine](https://whattomine.com) page.

You can customize your rig setup in the `conf.json`.

## NOTE: For Windows!!

Each batfile should change the title of the command prompt with the name of
the coin that it mines. Otherwise there's no easy way to find out which coin a
miner program that can mine several coin is actually mining at the moment.

You need to rename (and copy if necessary) the miner executables, so that its
name contains the coin that you would like to mine in order to find it, when
trying to start/kill the coin currently being mined.


Example:

```
@echo off
title gobyte
ccminer.exe -o stratum+tcp://....... -a ... -u ... 
```
