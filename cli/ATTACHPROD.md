## Helpful Notes
- The project you attach **must** be configured on prod. You should be able to login to https://console.taubyte.com and access it.
- All repositories don't **have** to be registered on production auth, only the config repository.
- Once repositories are registered through `pushSpecific` you can use `pushAll` to push them all.  There's also `pushAll --branch <someBranch>` which will push all the repos you have pushed to the current universe instance.



### Steps to success: 
<br/>

#### Start your multiverse:
```shell
$ dream new multiverse --bind node@8888/http
```

#### Attach your project:


> Open a new terminal after dreamland is completely started


```shell
$ dream inject attachProdProject  \
    --project-id QmYfMsCDvC9geoRRMCwRxvW1XSn3VQQoevBC48D9scmLJX  \
    --git-token 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  \
    --branch dreamland
```

#### Push the code repository:
```shell
$ dream inject pushSpecific \
    --repository-id 517160745 \
    --repository-fullname taubyte-test/tb_code_prodproject \
    --project-id QmYfMsCDvC9geoRRMCwRxvW1XSn3VQQoevBC48D9scmLJX \
    --branch dreamland
```

#### Push another repository (website):
```shell
$ dream inject pushSpecific \
    --repository-id 517418095 \
    --repository-fullname taubyte-test/tb_website_prodSocketWebsite \
    --project-id QmYfMsCDvC9geoRRMCwRxvW1XSn3VQQoevBC48D9scmLJX \
    --branch dreamland
```

## Now use it!

Currently this repository uses the domain "hal.computers.com"  so you'll need to add `127.0.0.1 hal.computers.com` or any relative domain you've added to your relative added project.

### Websocket site:
http://hal.computers.com:8888

### Ping Pong:
http://hal.computers.com:8888/ping

### Socket URL:
http://hal.computers.com:8888/getsocketurl