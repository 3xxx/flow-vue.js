<template>
<div>
<!-- <head>
  <meta charset="UTF-8">
  <title>Simple Chat</title>
</head>
<body> -->
<!-- <header>
  <nav>
    <div class="nav-wrapper">
      <a href="/" class="brand-logo right">Simple Chat</a>
    </div>
  </nav>
</header> -->
<!-- <main id="app"> -->
  <div class="row">
    <div class="col s12">
      <div class="card horizontal">
        <div id="chat-messages" class="card-content" v-html="chatContent">
        </div>
      </div>
    </div>
  </div>
    <div class="row" v-if="joined">
      <div class="input-field col s8">
        <input type="text" v-model="newMsg" @keyup.enter="send">
      </div>
      <div class="input-field col s4">
        <button class="waves-effect waves-light btn" @click="send">
          <i class="material-icons right">chat</i>
            Send
        </button>
      </div>
    </div>
    <div class="row" v-if="!joined">
      <div class="input-field col s8">
        <input type="email" v-model.trim="email" placeholder="Email">
      </div>
      <div class="input-field col s8">
        <input type="text" v-model.trim="username" placeholder="Username">
      </div>
      <div class="input-field col s4">
        <button class="waves-effect waves-light btn" @click="join()">
          <i class="material-icons right">done</i>
            Join
        </button>
      </div>
    </div>
<!-- </main> -->
<!-- <footer class="page-footer">
</footer> -->
<!-- <remote-script src="https://cdn.jsdelivr.net/emojione/2.2.6/lib/js/emojione.min.js"></remote-script>
<remote-script src="https://code.jquery.com/jquery-2.1.1.min.js"></remote-script>
<remote-script src="https://cdnjs.cloudflare.com/ajax/libs/crypto-js/3.1.2/rollups/md5.js"></remote-script>
<remote-script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.8/js/materialize.min.js"></remote-script> -->

</div>
</template>

<script>
// import '../assets/js/swiper.min.js';
// import "https://unpkg.com/vue@2.1.3/dist/vue.min.js";
// import '../assets/js/emojione.min.js';
import '../assets/js/jquery-2.1.1.min.js';
import '../assets/js/md5.js';
// import '../assets/js/materialize.min.js';

const axios = require('axios');
export default {
  name: 'chatme',
  data() {
    return {
      ws: null, // Our websocket
      newMsg: '', // Holds new messages to be sent to the server
      chatContent: '', // A running list of chat messages displayed on the screen
      email: null, // Email address used for grabbing an avatar
      username: null, // Our username
      joined: false // True if email and username have been filled in
    };
  },
  mounted() {},
  computed: {},
    created: function() {
      var self = this;
      this.ws = new WebSocket('ws://' + window.location.host + '/ws');
      this.ws.addEventListener('message', function(e) {
        var msg = JSON.parse(e.data);
        self.chatContent += '<div class="chip">'
                    + '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
                    + msg.username
                + '</div>'
                + emojione.toImage(msg.message) + '<br/>'; // Parse emojis

          var element = document.getElementById('chat-messages');
          element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
      });
    },

    methods: {
      send: function () {
          if (this.newMsg != '') {
              this.ws.send(
                  JSON.stringify({
                      email: this.email,
                      username: this.username,
                      message: $('<p>').html(this.newMsg).text() // Strip out html
                  }
              ));
              this.newMsg = ''; // Reset newMsg
          }
      },

      join: function () {
          if (!this.email) {
              Materialize.toast('You must enter an email', 2000);
              return
          }
          if (!this.username) {
              Materialize.toast('You must choose a username', 2000);
              return
          }
          this.email = $('<p>').html(this.email).text();
          this.username = $('<p>').html(this.username).text();
          this.joined = true;
      },

      gravatarURL: function(email) {
          return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
      }
    },
  watch: {
    adminInfo: function(newValue) {}
  }
};
</script>

<style scoped>
/*@import'../assets/css/swiper.min.css';*/
/*@import '../assets/css/materialize.min.css';*/
/*@import"https://fonts.googleapis.com/icon?family=Material+Icons";*/
@import '../assets/css/emojione.min.css';
@font-face {
  font-family: 'Material Icons';
  font-style: normal;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/s/materialicons/v43/flUhRq6tzZclQEJ-Vdg-IuiaDsNc.woff2) format('woff2');
}

.material-icons {
  font-family: 'Material Icons';
  font-weight: normal;
  font-style: normal;
  font-size: 24px;
  line-height: 1;
  letter-spacing: normal;
  text-transform: none;
  display: inline-block;
  white-space: nowrap;
  word-wrap: normal;
  direction: ltr;
  -webkit-font-feature-settings: 'liga';
  -webkit-font-smoothing: antialiased;
}

/* fallback */
/*@font-face {
  font-family: 'Material Icons';
  font-style: normal;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/s/materialicons/v47/flUhRq6tzZclQEJ-Vdg-IuiaDsNc.woff2) format('woff2');
}*/

/*.material-icons {
  font-family: 'Material Icons';
  font-weight: normal;
  font-style: normal;
  font-size: 24px;
  line-height: 1;
  letter-spacing: normal;
  text-transform: none;
  display: inline-block;
  white-space: nowrap;
  word-wrap: normal;
  direction: ltr;
  -webkit-font-feature-settings: 'liga';
  -webkit-font-smoothing: antialiased;
}*/
/*<style scoped>*/
body {
    display: flex;
    min-height: 100vh;
    flex-direction: column;
}

main {
    flex: 1 0 auto;
}

#chat-messages {
    min-height: 10vh;
    height: 60vh;
    width: 100%;
    overflow-y: scroll;
}
</style>