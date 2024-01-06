# Prefix Packet Header Example

Based On [Stephen Clear article](https://blog.stephencleary.com/2009/04/sample-code-length-prefix-message.html): 

Sometimes I don't like to use IIS or APACHE for hosting RESTFul APIs, or Real-Time operations, So I go to write the server from scratch.
you can use it in anything else, Bluetooth communication [ for unlimited data transfer without limitation ; ) ], Arduino, even if you decide to write your own protocol, WebSockets built on top of this method, you can use it as you want.

Some useful knowledge: 
[HTTP Built Top Of TCP](https://www.khanacademy.org/computing/computers-and-internet/xcae6f4a7ff015e7d:the-internet/xcae6f4a7ff015e7d:web-protocols/a/hypertext-transfer-protocol-http), 

If they lie to you and tell you It's completely different, they do not know anything, Most of them use HTTP for its simplicity, 
They do not know what is under the hood, for me, It's one of the greatest protocols, but not suitable for all stuff,

Whatever.. most network Protocols Built on top of TCP.. for example, FTP, SMTP, gRPC ( Using HTTP/2 and it is also built on top of TCP) WebSocket, even Torrent Protocol.. etc
