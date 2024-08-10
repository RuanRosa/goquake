```
       _           _
    _--/           \--_
   /  /             \  \
  /  |               |  \
 /   |               |   \
 |   |    _______    |   |
 |   |    \     /    |   | _______       _______     _____          _____           _______    _____
 \    \_   |   |   _/    / \     /       \     /    <     \         \    |        | \     /    ~-_  |
  \     -__|   |__-     /   |   |         |   |     /      \         |   |       ||  |   |        ~\|
   \_                 _/    |   |         |   |    /  /^\   \        |   |      / |  |   |   _/|    `
     --__         __--      |   |         |   |   /  /   \   \       |   |     /  /  |    > <_ |
         --|   |--          |   |         |   |  /   \   /    \      |    \  _-  <   |   |    \|    ,
           |   |             |   |       |   |  /  /~~   ~~\   \     |   |~  ~-_  \   |   |        /|
            | |               \_  \_   _/  _/ _-   \       /    -_   |   |      \ |    \_  \_   _-~ |
             |                  --__> <__--  <______>     <_______> <_____>      \|      --__> <__-\|
                                                                                  `                 `

```
<h1>GoQuake</h1>

<p>goquake was created to make it possible to export data from quake servers through positional files</p>

<h2>⚡ Run</h2>
replace quake.log file for your log and execute
<code>go run cmd/main.go</code>

<h2>project tree</h2>
<pre>
.    
├── cmd
│   └── main.go
├── domain
│   ├── ports
│   │   └── game.go
│   └── quake.go
├── example.json
├── gateways
│   └── file.go
├── go.mod
├── quake.log
├── quake.logls
├── teste.json
└── utils
    └── string.go
</pre>
