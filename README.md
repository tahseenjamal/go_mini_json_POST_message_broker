This a simple post request message queue written in Go Lang

Early I had written something similar in Python

Application asks for 3 parameters on command line


threads, port, target URL to which you want the program to submit the POST JSON data

Compiling is simple

go build main.go


Running with 5 threads on port 8080 with target service URL as http://URL:PORT


./main 5 8080 http://www.abc.com:9090

So this means this mini message broker would listen on port 8080 and any json submitted would be POSTed on target URL on port 9090

As said this is simple mini message broker, which can be used to cushion a service, which cannot take spike loads
