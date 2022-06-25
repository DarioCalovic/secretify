import Peer from 'peerjs'
// TODO: Find a way to make peer load before inject occures (e.g. async / await?)
// TODO: as currently there is a race condition between vue mounting and peer
// TODO: being ready
export default (ctx, inject) => {
  inject('peer', Peer)
}
