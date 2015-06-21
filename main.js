var express = require('express');
var app = express();
var server = require('http').Server(app);
var io = require('socket.io')(server);

app.use(express.static(__dirname + '/asset'));

io.on('connection', function (socket) {
	console.log("connected");
	socket.on('authorization', function (data) {
		console.log("auth: ",data);
	});
});

io.of("/profile").on('connection', function (socket) {
	console.log("connected");
	socket.on('authorization', function (data) {
		console.log("auth2: ",data);
	});
});


server.listen(3000);