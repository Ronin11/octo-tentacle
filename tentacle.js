// var WebSocket = require('ws')
//   , ws = new WebSocket('ws://localhost:8080');
// ws.on('open', function() {
//     ws.send('something');
// });
// ws.on('message', function(message) {
//     console.log('received: %s', message);
// });

// var express = require('express');

// var config = require('./config');
// var clog = require('./util/clog');

// // Express Setup
// var router = express.Router();
// var app = express();

// // app.use('/api', apiRoutes);


// app.listen(config.port, function(){
// 	clog.i("Tentacle listening on port ", config.port);
// });

const io = require('socket.io-client');
const socket = io('http://localhost:8080');

socket.emit('chat message', "HIIIII!!!");
