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
- curl -X POST --data-binary @js/main.js http://127.0.0.1:8080/main.js

To get file:
- md5 is local md5
- return status>=300 if error.
- return status==200 if post body(md5) is empty.
- return status>=200 if post body(md5) are different(Get counts=status-200)
- curl -X GET http://127.0.0.1:8080/main.js?md5=591d8a89d6bb4e07bb714495d8cfc0ef

To del file:
- curl -X DELETE http://127.0.0.1:8080/main.js

To make PUT counts:
- curl -X PUT http://127.0.0.1:8080/main.js

To get file info:
- If md5 is set, create a file marker when the marker does not exist
- curl -X TRACE http://127.0.0.1:8080/main.js?md5=591d8a89d6bb4e07bb714495d8cfc0ef
- - {"Md5":"","Get":0,"Put":0}
- - - Md5: md5 of this file
- - - Get: GET counts(GET post body(md5) isn't empty and different will count)
- - - Put: PUT counts
