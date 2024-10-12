const url = location.hostname
var currentRoom
var conn

async function addUser(login, password) {
          console.log("data: " + JSON.stringify({
                  login: login,
                  password: password
          }));
          let res = await fetch(`${'http://' + url + ':10017'}/addUser`, {
                  method: 'POST',
                  headers: {
                          'Content-Type': 'application/json',
                  },
                  body: JSON.stringify({
                          login: login,
                          password: password
                  })
          })
          return (await res.json())
}

async function authUser(login, password) {
  console.log("data: " + JSON.stringify({
          login: login,
          password: password
  }));
  let res = await fetch(`${'http://' + url + ':10017'}/authUser`, {
          method: 'POST',
          headers: {
                  'Content-Type': 'application/json',
          },
          body: JSON.stringify({
                  login: login,
                  password: password
          })
  })
  return (await res.json())
}

async function createRoom(login, password, username) {
  console.log("data: " + JSON.stringify({
          login: login,
          password: password,
          username: username
  }));
  let res = await fetch(`${'http://' + url + ':10017'}/createRoom`, {
          method: 'POST',
          headers: {
                  'Content-Type': 'application/json',
          },
          body: JSON.stringify({
                  login: login,
                  password: password,
                  username: username
          })
  })
  return (await res.json())
}

async function connectRoom(login, password, username, roomid) {
  console.log("data: " + JSON.stringify({
          login: login,
          password: password,
          username: username,
          secret: roomid,
  }));
  let res = await fetch(`${'http://' + url + ':10017'}/connectRoom`, {
          method: 'POST',
          headers: {
                  'Content-Type': 'application/json',
          },
          body: JSON.stringify({
                  login: login,
                  password: password,
                  username: username,
                  secret: roomid,
          })
  })
  return (await res.json())
}

async function getRooms(login, password) {
  console.log("data: " + JSON.stringify({
          login: login,
          password: password
  }));
  let res = await fetch(`${'http://' + url + ':10017'}/rooms`, {
          method: 'POST',
          headers: {
                  'Content-Type': 'application/json',
          },
          body: JSON.stringify({
                  login: login,
                  password: password
          })
  })
  return (await res.json())
}

async function sendMessage(secret, messageText, roomId) {
  console.log("data: " + JSON.stringify({
          message: messageText,
          secret: secret,
          roomid: roomId,
  }));
  let res = await fetch(`${'http://' + url + ':10017'}/sendMessage`, {
          method: 'POST',
          headers: {
                  'Content-Type': 'application/json',
          },
          body: JSON.stringify({
                  message: messageText,
                  secret: secret,
                  roomid: roomId,
          })
  })
  return (await res.json())
}

async function getMessages(secret, roomid) {
  console.log("data: " + JSON.stringify({
    secret: secret,
    roomid: roomid
  }));
  let res = await fetch(`${'http://' + url + ':10017'}/getMessages`, {
          method: 'POST',
          headers: {
                  'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            secret: secret,
            roomid: roomid
          })
  })
  return (await res.json())
}

function getCookie(name) {
  let matches = document.cookie.match(new RegExp(
    "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
  ));
  return matches ? decodeURIComponent(matches[1]) : undefined;
}

function deleteCookie(name) {
  setCookie(name, "", {
    'max-age': -1
  })
}

function setCookie(name, value, options = {}) {

  options = {
    path: '/',
    ...options
  };

  if (options.expires instanceof Date) {
    options.expires = options.expires.toUTCString();
  }

  let updatedCookie = encodeURIComponent(name) + "=" + encodeURIComponent(value);

  for (let optionKey in options) {
    updatedCookie += "; " + optionKey;
    let optionValue = options[optionKey];
    if (optionValue !== true) {
      updatedCookie += "=" + optionValue;
    }
  }

  document.cookie = updatedCookie;
}

function insertAfter(newNode, existingNode) {
  existingNode.parentNode.insertBefore(newNode, existingNode.nextSibling);
}

function SignOut() {
  deleteCookie('login');
  deleteCookie('password');
  deleteCookie('token');
  window.location.replace(location.origin + "/");
}

function ListAdd(id, element) {
  let ul = document.getElementById(id);
  let li = document.createElement("li");
  li.setAttribute('onclick', `OpenChat(\'${element}\');`)
  li.appendChild(document.createTextNode(element));
  ul.appendChild(li);
}

function ListAddMessage(idList, username, message, originalUsername) {
  let mdiv = document.getElementById(idList);
  let div = document.createElement("div");
  let p1 = document.createElement("p");
  let p2 = document.createElement("p");
  if (originalUsername == username) {
    div.setAttribute('id', 'message-block-user')
  } else {
    div.setAttribute('id', 'message-block')
  }
  p1.appendChild(document.createTextNode(message));
  p2.appendChild(document.createTextNode(username));

  mdiv.appendChild(div);
  div.appendChild(p1);
  div.appendChild(p2);
}
//document.getElementById("yourH1_element_Id").innerHTML = "yourTextHere";
function ListDelete(id) {
  document.getElementById(id).innerHTML = "";
}

async function OpenChat(roomId) {
  
  ListDelete("message-list")
  let title = document.getElementById("chat-title");
  title.innerHTML = "Текущий чат: " + roomId;
  currentRoom = roomId;

  let button = document.getElementById("message-button");
  button.className = "";
  let liTags = document.getElementsByTagName("li");
  for (let i = 0; i < liTags.length; i++) {
    if (liTags[i].textContent == roomId) {
      liTags[i].className = "disabled";
      continue;
    }
    liTags[i].className = "";
  }

 // found.className = "disabled"

  var res = await getMessages(getCookie("token"), roomId)
    if (res.error != "") {
        console.log("Error: " + res.message)
    } else {
      if (res.messagelist != null) {
        res.messagelist.forEach(element => {
          ListAddMessage("message-list", element.username, element.message, res.username)
        });
    }

    if (window["WebSocket"]) {
      conn = new WebSocket("ws://" + document.location.hostname + ":10015/ws");
      conn.addEventListener("open", (ev) => {
        conn.send(JSON.stringify({
          secret: getCookie("token"),
          roomid: roomId
        }));
        });
      
      conn.onclose = function (evt) {
          var item = document.createElement("div");
          item.innerHTML = "<b>Connection closed.</b>";
          //appendLog(item);
      };
      conn.onmessage = function (evt) {
          
        let message = evt.data;  
        let res = JSON.parse(message);
        //console.log(res);
        res.messagelist.forEach(element => {
          ListAddMessage("message-list", element.username, element.message, "")
        });
      };
  } else {
      var item = document.createElement("div");
      item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
      //appendLog(item);
  }
        //console.log("Success: " + res.message)
    }
}