function EchoTest()
{
  if ("WebSocket" in window)
  {
	console.log("WebSocket is supported by your Browser.");
	
	// open a web socket
	var ws = new WebSocket("ws://localhost:8080/echo-ws"); // Connect to a websocket server.
	
	$(window).on('beforeunload', function(){ // Defer socket.close() until page is closed or we browse away.
	    ws.close();
	});
	
	ws.onopen = function()
	{
		// Web Socket is connected, send data using send()
		ws.send("Message to send");
		console.log("Message sent...");
	};
	
	ws.onmessage = function (evt) 
	{ 
		var received_msg = evt.data;
		console.log("Message received: " + received_msg);
	};
	
	ws.onclose = function()
	{ 
		// websocket is closed.
		console.log("Connection is closed..."); 
	};
  }
  else
  {
	// The browser doesn't support WebSocket
	console.log("WebSocket not supported by your Browser. Please use a modern web browser.");
  }
}

function EmitTest()
{
  if ("WebSocket" in window)
  {
	console.log("WebSocket is supported by your Browser.");
	
	// open a web socket
	var ws = new WebSocket("ws://localhost:8080/emit-ws"); // Connect to a websocket server.
	
	$(window).on('beforeunload', function(){ // Defer socket.close() until page is closed or we browse away.
	    ws.close();
	});
	
	ws.onopen = function()
	{

	};
	
	ws.onmessage = function (evt) 
	{ 
		var received_msg = evt.data;
		console.log("Message received: " + received_msg);
	};
	
	ws.onclose = function()
	{ 
		// websocket is closed.
		console.log("Connection is closed..."); 
	};
  }
  else
  {
	// The browser doesn't support WebSocket
	console.log("WebSocket not supported by your Browser. Please use a modern web browser.");
  }
}

$(function() {	// Runs once the DOM is loaded.
	//EchoTest();
	EmitTest();
});