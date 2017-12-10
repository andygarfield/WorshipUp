# WorshipUp

## General idea
Mostly copy OpenSong in organization abilities, but have a central server so that multiple people can work on the worship service throughout the week.

## Configurations
People who are setting up the system can choose from a few different configurations. These are ordered by simplest to most complex.

1. A basic single-system setup where the data is stored and presented locally. The client and server are implimented on the same machine.
2. A client and a server implementation on separate machines. The server has the master copy of the data and the client syncs with it. The desktop client presents with a current copy of the data. Also have a web interface that has all the capabilities of the desktop application.
3. Same as #2 but with devices set up to connect to a client server and control the presentation and potentially make changes mid-presentation. Would like to do this over wi-fi direct or bluetooth.

The desire is to have the ability to smoothly upgrade to a more complicated configuration if a church's needs grow.

## General Architechture
![Architecture Chart](./architecture.svg)

[source](https://www.lucidchart.com/invitations/accept/c311a15e-7224-4a86-ba2f-b82a73967ef2)