# Autoswitcher for GPU mining scripts 

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


Example:

```
@echo off
title gobyte
ccminer.exe -o stratum+tcp://....... -a ... -u ... 
```
