<template>
  <section id="app">
    <div class="head">
      <nav-bar />
    </div>
    <nuxt />
    <footer-bar />
  </section>
</template>
<script>
// @ is an alias to /src
import NavBar from '@/components/NavBar'
import FooterBar from '@/components/FooterBar'

export default {
  name: 'Home',
  components: {
    FooterBar,
    NavBar,
  },
}
</script>
