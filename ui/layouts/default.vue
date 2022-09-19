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
  mounted() {
    this.$nextTick(() => {
      document.documentElement.style.setProperty(
        '--main-bg-color',
        this.$config.branding.primary_color
      )
      document.documentElement.style.setProperty('display', 'block')
    })
  },
}
</script>
