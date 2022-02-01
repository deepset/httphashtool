# httphashtool
httphashtool is a simple tool which makes http request and prints the address of the request along with MD5 hash of the response.

Tool performs the request in parallel, with a default of 10 parallel requests
http request has a timeout of 5 seconds


# Install
git clone github.com/deepset/httphashtool

cd github.com/deepset/httphashtool

go build


# Example
$ ./httphashtool adjust.com

$ ./httphashtool -parallel 3 google.com adjust.com twitter.com www.facebook.com http://reddit.com

# Testing
cd github.com/deepset/httphashtool

go test -v ./...
      
---
      
cd github.com/deepset/httphashtool/script

go test -v




