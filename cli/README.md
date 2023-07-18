## Commands:

### Install:

Install as dream
```bash
go build -o ~/go/bin/dream
```

If you want bash completion, add this to your .bashrc:

`PROG=dream source ~/go/pkg/mod/github.com/urfave/cli/v2@v2.11.1/autocomplete/bash_autocomplete`




### Start
```
`dream new multiverse`
    --empty          // By default starts a universe with all services named "blackhole"
                     // empty creates a multiverse with an empty universe named "blackhole"
    --daemon         // Starts it in the background
    --bind           // [service@0000/http,...] [service@0000/p2p,...] [service@0000,...]
    --disable        // [service,...]
    --enable         // [service,...]
    --fixtures       // [fixtureName,...]
    --simples        // [name,...]  // Starts simples of `name` with all clients available
    --listen-on-all  // Changes dreamland default url to 0.0.0.0 rather than 127.0.0.1
                     // for hosting on your local ip
    --id             // Sets or opens a universe on a specific id
    --keep           // If set will store the universe in $HOME/.cache/dreamland rather than /tmp"

`dream new universe`  // attaches to the started multiverse, errors if no multiverse is started
    --name      // defaults to "blackhole"
    --empty     // defaults with having all services, and a simple named `client` with all clients
                // empty removes all of this, and gives an empty universe.
    --bind      // [service@0000/http,...] [service@0000/p2p,...] [service@0000,...]
    --disable   // [service,...]
    --enable    // [service,...]
    --fixtures  // [fixtureName,...]
    --simples   // [name,...]  // Starts simples of `name` with all clients available
```
    

### Inject
```
[universe] // defaults to `blackhole`

`dream inject fixture [name] [universe] [params],...`
    --name [name]      // name of the fixture
    --universe [name]  // defaults to blackhole if not set


// Service
`dream inject [name] [universe]`
    [name]      // name of the service
    --universe [name]  // defaults to blackhole if not set
    --http [int]       //

`dream inject services [name,...] [universe]`
    --name [name,...]      // name of the services to kill
    --universe [name]  // defaults to blackhole if not set

`dream inject simple [name] [universe]`
    --name [name]      // name of the simple
    --universe [name]  // defaults to blackhole if not set
    --disable          // [service,...]
    --empty            // defaults to having all clients available
        --enable       // [service,...]
```
    
### Kill

```
`dream kill multiverse`  // kills the multiverse...

`dream kill [name] [universe]`
    [name]      // name of the service
    --universe [name]  // defaults to blackhole if not set

`dream kill simple [name] [universe]`
    --name [name]      // name of the simple
    --universe [name]  // defaults to blackhole if not set

`dream kill services [name,...] [universe]`
    --name [name,...]      // name of the services to kill
    --universe [name]  // defaults to blackhole if not set
```

### Status 
```
`dream status` // Shows available status commands 

`dream status u/universe [universe]`
    --universe [name]   // defaults to blackhole if not set ex:`dream status u`

`dream status [node]` 
--node [name] // name of the node to render 

`dream status id` // Shows the id of the given universe
```

### Functions Demo
- Run 
```shell
$ dream new multiverse --bind console@2022/http,node@2023/http
```

- Connect to "http://127.0.0.1:2022"

- <span id="Domain">Create a domain</span>
```yaml
id: Qmf7yRv4Jg5bXAxHDLgzGdRYWKJjUjAD8DxMVGLTAvs6Ai
description: ''
tags: []
fqdn: hal.computers.com
certificate:
    type: auto
```

#### Functions
- <span id="etsy">Add a domain to `/etc/hosts`</span>
```shell
# sed -i 'd/hal.computers.com/' /etc/hosts
# echo "127.0.0.1 hal.computers.com" >> /etc/hosts
```

- Create a function with some inline code
```yaml
id: QmXKt1xBnCUjwDzLFztxk4w34qqwe9wNcx2kxqLm5iSxDs
description: ''
tags: []
trigger:
    type: https
    method: GET
    paths:
        - /ping
domains:
    - someDomain
includes:
    - ''
execution:
    timeout: 10s
    memory: 10MB
    call: pingerMethod

```
```go
package funcsomething


import (
    "github.com/taubyte/go-sdk/event"
)

//export pingerMethod
func pingerMethod(e event.Event) uint32 {
    h, _ := e.HTTP()
    h.Write([]byte("Pong"))

    return 0
}
```

- Push code on web-console
- More instructions below website

- Inject the pushAll fixture to fake the github payloads or open [odval](http://127.0.0.1:1421/network/blackhole) and push the fixture `pushAll
```shell
$ dream inject fixture pushAll
```

- Wait for jobs to finish,  if the job fails on jobStatus push again.  Sometimes code will push first and cause it to fail due to not seeing the config.

- http://hal.computers.com:2023/ping


#### Websites Demo
- Create a repo for your website on github using test account, you can do this through console, or through github 

- Clone the repo, build the website, and ensure the production build builds properly
- Add a taubyte folder in the root of the repo 
- In the taubyte folder add a config.yaml file

- Example: 
```yaml
version: 1.00
enviroment: 
  image: node
  variables:
workflow:
    - build
```

- For image: choose whatever docker container image name you would like to use. It is reccomended to use the node docker image

- For Workflow: This is the list of bash scripts you would like to execute to build your website.

- Add bash scripts for all workflow elements in the taubyte folder. If you workflow is: -test,-build you would add a test.sh, and build.sh file to the taubyte folder.

- The goal is that after executing your scripts all of your production build files will be placed an out directory in the root of your repo

- Example: 
```bash
# !/bin/bash

node build/build.js
```

- In the VueUtilsWithTest repo, node executes the build/build.js file, which will build to an out file

- It is recommended that you test your bash script and see if production build is built to out file locally

- Push your repos changes to github

- On github, go to your repo settings>webhooks

- Click on the most recent push payload > recent deliveries, click the most recent delivery

- If the payload failed to deliver, push another commit, or click the redilver button

- Find this new payload

- Copy this Payload and then go to dreamland-test repo

- Update the website-payload.json file to this payload

- In common/variables.go update website repo hook id and repo id, you can find this in the payload json

- Replace dreamland-test, in patrick

- Replace patrick and dreamland-test in cli

- Run dreamland using go run . new multiverse --bind console@2022/http,node@2023/http

- Connect to "http://127.0.0.1:2022" and create a project, or connect to "http://127.0.0.1:1421" and push project with jobs fixture and select testproject as your project 

- Create a domain like in the function [example](#Domain), and add domain to [/etc/hosts](#etsy) or use prexisting domain if using testProject 

- Add website and do not click generate repository button, enter full name of your repo, and the id. The full name is taubyte-test/"your-repo-name"
- You can get your repo Id from hitting f12 on the repo's github page then in elements, search for "octolytics-dimension-repository_id"

``` yaml 
id: QmaMk4tWYHwPBiJyxtamTd3VPFzbLZd2ssVutcqnYkwff9
description: ''
tags: []
domains:
    - somedomain
libraries: []
source:
    paths:
        - /
    includes: []
    branch: main
    github:
        id: '512852992'
        fullname: taubyte-test/tb_website_reactdemo
```


- Open [odval](http://127.0.0.1:1421/network/blackhole) and push the fixture `pushAll` and wait for jobs to finish
- Check job status on console, there should be 3 jobs, if it didn't go through then try pushWebsite fixture again. 

- Check http://hal.computers.com:2023/
for your website once job has been succesfully completed
