<template>
  <div class="blank">
    <section class="section">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-12">
            <h1 class="title">Privacy</h1>
            <nuxt-content :document="doc" />
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
export default {
  name: 'Privacy',
  async asyncData({ $content }) {
    const doc = await $content('privacy').fetch()
    return {
      doc,
    }
  },
  mounted() {
    // Track
    this.$track.pageview({})
  },
  head() {
    return {
      title: 'Privacy - Secretify',
    }
  },
}
</script>
