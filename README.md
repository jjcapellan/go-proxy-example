This example shows a simple way to separate backend and frontend within the same domain using 3 servers: 
* Reverse proxy ([jjc-reverse-proxy](https://github.com/jjcapellan/jjc-reverse-proxy))
* Backend 
* Frontend

All requests are received by the proxy, and it will forward them to one server or another depending on the route prefix. In this case, requests whose path begins with "/api" are sent to the backend server, and the rest to the frontend server.  

This architecture is intended to be deployed in a secure network, such as Heroku, which only exposes a single port and encrypts all connections to the outside.