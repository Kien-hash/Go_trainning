let hostname = window.location.hostname;
let port = 8000;
let socket = new WebSocket("ws://" + hostname + ":" + port + "/ws");
let info = JSON.parse(sessionStorage.getItem('userInfo'));
let textarea = document.getElementById('text-area');
let inputMessage = document.getElementById('inputMessage');
let infoName = info.firstname + " " + info.lastname;
let msg;


document.cookie = "session_token" + '=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';

socket.onopen = () => {
   console.log("Successfully Connected");
   let message = { Event: "initial", ID: info.ID, Name: infoName };
   socket.send(JSON.stringify(message));
};

socket.onmessage = (event) => {
   msg = JSON.parse(event.data);
   console.log(msg);
   // Hien thi ra cac user-online
   if (msg.Event == "user-online") {
      $('#ulUser').empty();
      for (let i = 0; i < msg.Users.length; i++) {
         let user = msg.Users[i]
         if (user.ID != info.ID) {
            $('#ulUser').append(`<input type="checkbox" id="${user.ID}" name="${user.Name}">${user.Name}<br />`)
         } else {
            $('#ulUser').append(`<input type="checkbox" id="${user.ID}" name="${user.Name}">${user.Name + " **YOU"}<br />`)
         }
      }
   } else if (msg.Event == "has-message") {
      // sự kiện nhận phòng
      let roomname = msg.Name
      let i = confirm(roomname + ' has message for you!!!');
      // lưu phòng 
      if (i == true) {
         document.getElementById("roomName").innerHTML = roomname;
         document.getElementById("openChat").hidden = false;
         // xóa các tin nhắn trong text area
         textarea.value = "";
         sessionStorage.setItem('room', msg.Context)
      } else {
         console.log("this is else function")
      }
   } else if (msg.Event == "message") {
      displayMessage(msg)
   }
}

socket.onclose = event => {
   console.log("Socket Closed Connection: ", event);
   alert("Socket Connection Closed!!!")
};

socket.onerror = error => {
   console.log("Socket Error: ", error);
};


//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
function Chat() {
   var selected = [];
   var roomname = "Room: ";
   $('#ulUser input:checked').each(function () {
      let ID = parseInt($(this).attr('id'))
      if (ID == info.ID) {
      } else {
         selected.push(ID);
         roomname += $(this).attr('name').split(" ").pop() + " ";
      }
   });

   $("#ulUser input:checked").prop("checked", false);
   if (selected.length == 0) {
      alert('cant chat alone!!!');
      return;
   }
   else {
      selected.push(info.ID);
      roomname += info.lastname.split(" ").pop() + " ";
   }

   document.getElementById("roomName").innerHTML = roomname;
   document.getElementById("openChat").hidden = false;

   // xóa các tin nhắn trong text area
   textarea.value = "";

   // Send(sự kiện khởi tạo phong)
   let message = { Event: "initial-single-room", ID: info.ID, Name: roomname, Receiver: selected };
   // let roomsId = selected.split(',').map(Number);
   selected.sort((a, b) => a - b);
   console.log(selected)
   sessionStorage.setItem('room', selected.toString());
   socket.send(JSON.stringify(message));
}

function Send() {
   let inputMessage = document.getElementById("inputMessage");
   let receiver = sessionStorage.getItem('room').split(',').map(Number);
   receiver.sort();
   let message = { Event: "message", ID: info.ID, Name: infoName, Receiver: receiver, Context: inputMessage.value };
   socket.send(JSON.stringify(message));
   inputMessage.value = "";
}

function displayMessage(msg) {
   let receiver = sessionStorage.getItem('room').split(',').map(Number);
   if (JSON.stringify(msg.Receiver) == JSON.stringify(receiver)) {
      if (msg.ID == info.ID) {
         textarea.value += 'YOU: ' + msg.Context + '\n';
      } else {
         textarea.value += msg.Name + ': ' + msg.Context + '\n';
      }
      setInterval(function () {
         textarea.scrollTop = textarea.scrollHeight;
      }, 1000);
   }

}

function outRoom() {
   location.reload()
}

var input = inputMessage;

input.addEventListener("keyup", function (event) {
   // Number 13 is the "Enter" key on the keyboard
   if (event.keyCode === 13) {
      // Cancel the default action, if needed
      event.preventDefault();
      // Trigger the button element with a click
      document.getElementById("send-button").click();
   }
});

