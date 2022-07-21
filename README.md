# Static update file server

simple http server for update files.

## Install

Or you can install from source using Go:

    $ go get github.com/vanishs/updatefile

## Usage

To serve the "./root" directory on port 8080:

    $ updatefile

To use a different port specify with the `-port` flag:

    $ updatefile -port 5000

To serve a different directory use the `-root` flag:

    $ updatefile -root public

## Options

`-port` Defines the TCP port to listen on. (Defaults to 8080).

`-root` Defines the directory to serve. (Defaults to root directory).

## update file demo

To upload file:
- curl -X POST -d 'name=linuxize' http://127.0.0.1:8080/55.txt

To get file: return status 200 and data if md5 are different
- curl -X GET -d '591d8a89d6bb4e07bb714495d8cfc0ef' http://127.0.0.1:8080/55.txt

To del file:
- curl -X DELETE http://127.0.0.1:8080/55.txt

