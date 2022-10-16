<template>
  <section id="app" class="hero is-fullheight">
    <!-- Hero head: will stick at the top -->
    <div class="hero-head head">
      <nav-bar />
    </div>
    <div class="hero-body">
      <nuxt />
    </div>
    <div class="hero-foot">
      <footer-bar />
    </div>
  </section>
</template>
<script>
// @ is an alias to /src
import NavBar from '@/components/NavBar'
import FooterBar from '@/components/FooterBar'

export default {
  name: 'Marketing',
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
    })
  },
}
</script>
