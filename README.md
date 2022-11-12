# Static update file server

simple https server for update files.

## Install

Or you can install from source using Go:

    $ go install github.com/vanishs/updatefile@latest

## Usage

To serve the "./root" directory on https://127.0.0.1:8080:

    $ updatefile -ssl example.com

To use a different addr specify with the `-addr` flag:

    $ updatefile -ssl example.com -addr 127.0.0.1:5000

To serve a different directory use the `-root` flag:

    $ updatefile -ssl example.com -root public

To change pwd use the `-pwd` flag:

    $ updatefile -ssl example.com -pwd 654321

## Options

`-addr` Defines the addr to serve. (Defaults to 127.0.0.1:8080).

`-root` Defines the directory to serve. (Defaults to root directory).

`-pwd` validate password. (Default:123456).

`-ssl` ssl file name. (Default:example.com).

## update file demo

To upload file:
- curl --insecure -X POST --data-binary @js/main.js https://127.0.0.1:8080/main.js?pwd=123456

To get file:
- md5 is local md5
- return status>=300 if error.
- return status==200 if post body(md5) is empty.
- return status>=200 if post body(md5) are different(Get counts=status-200)
- curl --insecure -X GET https://127.0.0.1:8080/main.js?pwd=123456&md5=591d8a89d6bb4e07bb714495d8cfc0ef

To del file:
- curl --insecure -X DELETE https://127.0.0.1:8080/main.js?pwd=123456

To make PUT counts:
- curl --insecure -X PUT https://127.0.0.1:8080/main.js?pwd=123456

To get file info:
- If md5 is set, create a file marker when the marker does not exist
- curl --insecure -X TRACE https://127.0.0.1:8080/main.js?pwd=123456&md5=591d8a89d6bb4e07bb714495d8cfc0ef
- - {"Md5":"","Get":0,"Put":0}
- - - Md5: md5 of this file
- - - Get: GET counts(GET post body(md5) isn't empty and different will count)
- - - Put: PUT counts
