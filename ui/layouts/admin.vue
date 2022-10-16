<template>
  <section id="app" class="hero is-fullheight">
    <div class="hero-body">
      <nuxt />
    </div>
  </section>
</template>
<script>
export default {
  name: 'Home',
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
