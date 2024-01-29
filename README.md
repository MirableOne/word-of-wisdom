# Word Of Wisdom

## Description
Design and implement “Word of Wisdom” tcp server.
- TCP server should be protected from DDOS attacks with the [Proof of Work](https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
- The choice of the POW algorithm should be explained.
- After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
- Docker file should be provided both for the server and for the client that solves the POW challenge


## Algorithm

I've chosen [Hashcash](https://en.wikipedia.org/wiki/Hashcash) as PoW algorithm. 

Reasons:

- It has clear description;
- Simplicity;
- Proved effectiveness (by Bitcoin).


## How project works

Server listens for tcp connection and requires hash for every message from client.
On the valid message it responses with random quote.

Client starts endless loop then makes PoW for every message and sends it to the server.
It prints received quotes into log. 

On macbook pro m1 client makes about 3 messages per second. 

### Tradeoffs
- The quotes are hard-coded to avoid writing boilerplate code for the database;
- For storage, I used naive in-memory storage (hash map) for the same reasons;
- Messages use json for serialisation to make everything simpler.

## How to run

To run everything together and see the mess of logs:
```shell
docker-compose up 
```

Start server:
```shell
docker-compose up server 
```

Start client:
```shell
docker-compose up client 
```