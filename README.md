# autocomplete-service
Creator: Jonathan Kenney (jonathankenney12@gmail.com)

Solution for code challege: https://github.com/zesty-io/gists/blob/master/code-challenge-autocomplete

# Setup
This module was created using `go1.10.4` on Ubuntu18.04 LTS and has not been verified on any other go version or operating system. The implementation is pure go and requires no 3rd party packages.

Clone the repository locally and ensure you have a program like [curl](https://curl.haxx.se/) or [Postman](https://www.postman.com/) for testing.

# Usage
In your terminal, switch into the `autocomplete-service` directory and build the binary with `go build`.

Next, start the service using `./autocomplete-service`.

e.g.
```
cd ~/code/autocomplete-service
go build
./autocomplete-service
```

### Request with curl
In another terminal window, use curl to send a request to the server as follows:

```
curl localhost:9000/autocomplete?term=th
```

Set the term parameter to the desired term (only alphabetical terms allowed). Note that this endpoint will only accept GET requests.

### Request with Postman

In a new Postman request, set the method to GET and enter the URL as follows: 

```
localhost:9000/autocomplete?term=th
```

Set the term parameter to the desired term (only alphabetical terms allowed) and hit `Send`. Note that this endpoint will only accept GET requests.

# Solution Verification and Performance
Experimental results required by the [CHALLENGE.md](./CHALLENGE.md) can be found in [this Google Doc](https://docs.google.com/document/d/1qyLToFWBO2_EcWGMGZ9VSDEs9jTIdemhlohRlGtmEhQ/edit?usp=sharing).