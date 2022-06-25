<template>
  <div class="hero is-fullheight">
    <div class="hero-head head">
      <nav-bar />
    </div>
    <div class="hero-body">
      <div class="container dotted">
        <div class="columns">
          <div class="column is-12">
            <h3 class="subtitle is-3">Room {{ room_id }}</h3>
            <div>
              <span class="tag is-dark" data-tooltip="Tooltip Text"
                >{{ peers.length }}/{{ max_peers }} online</span
              >
              <span class="tag is-danger">4:49 timer</span>
            </div>
            <div class="mt-1">
              <span
                class="tag is-link is-normal mr-1"
                v-for="peer in peers"
                :key="peer.id"
                >@{{ peer.id | prettyUsername(user.id) }}</span
              >
            </div>
          </div>
        </div>
        <div class="columns">
          <div class="column is-12">
            <div class="chat--messages">
              <div
                :class="
                  'chat--message mb-2' +
                  (message.userId === user.id ? ' me' : '')
                "
                v-for="(message, index) in messages"
                :key="index"
              >
                <div
                  :class="
                    'chat--message-content' +
                    (message.userId === 0 ? ' bot' : '')
                  "
                >
                  <span class="chat--message-user">{{
                    message.userId | prettyUsername(user.id)
                  }}</span>
                  <span class="chat--message-text">{{ message.text }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="hero-foot">
      <div class="container">
        <div class="columns">
          <div class="column is-12">
            <b-field>
              <b-input
                v-model="inputText"
                id="chat_message"
                placeholder="Type message here..."
                expanded
              ></b-input>
              <p class="control">
                <b-button id="send" type="is-primary">Send</b-button>
              </p>
            </b-field>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<style>
.columns {
  margin: 0 !important;
}

.title,
.subtitle {
  color: white;
  font-weight: 500;
}

.chat--message {
  width: 100%;
}

.chat--message-content {
  display: inline-block;
  background: rgba(255, 255, 255, 0.9);
  padding: 0.75rem;
  border-radius: 0.25rem;
}

.bot {
  border-left: 5px solid #f14668;
}

.me {
  display: flex;
  justify-content: flex-end;
}

.me > .chat--message-content {
  background: rgba(0, 0, 0, 0.2);
  color: #ffffff;
}
.chat--message-user {
  font-weight: bold;
  display: block;
}
</style>
<script>
import NavBar from '@/components/NavBar'
const io = require('socket.io-client')
let socket = null
export default {
  filters: {
    prettyUsername(val, userId) {
      if (val === 0) {
        return 'Secretify Bot'
      } else if (val === userId) {
        return 'Me'
      }
      return val
    },
  },
  name: 'Socket',
  components: {
    NavBar,
  },
  layout: 'chat',
  data() {
    return {
      peer: null,
      user: {
        id: '',
        name: 'Tony',
      },
      peerConns: [],
      peers: [],
      room_id: '',
      max_peers: 2,
      inputText: '',
      messages: [
        {
          text: `This conversation is end-to-end that only you and your contact can access. As soon as you leave or reload this page, the messages are gone forever and the room will be destroyed.`,
          userId: 0,
        },
      ],
    }
  },
  mounted() {
    window.addEventListener('beforeunload', this.beforeWindowUnload)

    this.peer = new this.$peer({
      path: this.$config.peer.path,
      host: this.$config.peer.host,
      port: this.$config.peer.port,
    })

    const This = this

    this.room_id = this.$route.params.id

    socket = io.connect(this.$config.peer.socketAddress, {
      path: this.$config.peer.socketPath,
    })

    socket.on('connect', () => {
      console.log('Got socket connection!')
    })

    socket.on('disconnect', () => {
      console.log('Got socket disconnect!')
    })

    socket.on('add-peer', function (userId) {
      console.log('Add peer with id', userId)
      if (!This.addPeer(userId) && This.user.id !== userId) {
        // Add peer
        const conn = This.peer.connect(userId)
        conn.on('open', () => {
          console.log('New friend connected with id', userId)
          This.peerConns.push(conn)
        })
        conn.on('error', (e) => {
          console.log('Could not connect to friend')
        })
      }
    })

    socket.on('remove-peer', function (userId) {
      console.log('Remove peer')
      This.removePeer(userId)
    })

    this.peer.on('open', (id) => {
      This.user.id = id
      console.log('Peer opened', This.user.id)
      socket.emit('join-room', this.room_id, This.user.id, This.user.name)
    })

    this.peer.on('error', (id) => {
      console.log('Could not open peer', id)
    })

    this.peer.on('connection', (conn) => {
      conn.on('data', (data) => {
        const userName = conn.peer

        This.addMessage(userName, data)
      })
    })

    this.peer.on('error', (e) => {
      console.log('Peer error')
    })

    const send = document.getElementById('send')

    send.addEventListener('click', (e) => {
      if (This.inputText.length !== 0) {
        // Send to other peers
        This.addMessage(This.user.id, This.inputText)
        This.peerConns.forEach((peerConn) => {
          if (This.user.id !== peerConn.peer) {
            peerConn.send(This.inputText)
          }
        })
        This.inputText = ''
      }
    })

    document.querySelector('#chat_message').addEventListener('keydown', (e) => {
      if (e.key === 'Enter' && This.inputText.length !== 0) {
        // Send to other peers
        This.addMessage(This.user.id, This.inputText)
        This.peerConns.forEach((peerConn) => {
          if (This.user.id !== peerConn.peer) {
            peerConn.send(This.inputText)
          }
        })
        This.inputText = ''
      }
    })
  },
  methods: {
    beforeWindowUnload(e) {
      socket.emit('leave-room', this.room_id, this.user.id)
      socket.disconnect()
    },
    addMessage(userId, msg) {
      this.messages.push({
        userId,
        text: msg,
      })
    },
    addPeer(id) {
      let found = false
      this.peers.forEach((p) => {
        if (p.id === id) {
          found = true
        }
      })
      if (!found) {
        this.peers.push({
          id,
        })
      }
      return found
    },
    removePeer(id) {
      this.peers.forEach((p, i) => {
        if (p.id === id) {
          this.peers.splice(i)
        }
      })
    },
  },
  beforeRouteLeave(to, from, next) {
    socket.emit('leave-room', this.room_id, this.user.id)
    socket.disconnect()
  },
  head() {
    return {
      title: 'Transfer — secretify.io',
    }
  },
}
</script>
