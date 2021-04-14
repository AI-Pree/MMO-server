# MMO-server
MMO server for  2D or 3D indie games
 
## Requirements
- for handling packet loss, add another layer on top of udp connection that would verify the missing packets (keeps sending the packets to the recipient until ACK is sent back for the respective packet).
- Integrate with other programming language.
- Adding daemons for traffic handling, checking nodes condition.
- front end process for handling client connection to the server.
- Game logic distributor that manages non-real time aspect of the game.
