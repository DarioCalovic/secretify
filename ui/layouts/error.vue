<template>
  <div class="container error">
    <div class="columns is-vcentered is-centered has-text-centered dotted">
      <div class="column is-5">
        <div class="box">
          <div class="box__ghost">
            <div class="symbol"></div>
            <div class="symbol"></div>
            <div class="symbol"></div>
            <div class="symbol"></div>
            <div class="symbol"></div>
            <div class="symbol"></div>

            <div class="box__ghost-container">
              <div class="box__ghost-eyes" id="ghost-eyes">
                <div class="box__eye-left"></div>
                <div class="box__eye-right"></div>
              </div>
              <div class="box__ghost-bottom">
                <div></div>
                <div></div>
                <div></div>
                <div></div>
                <div></div>
              </div>
            </div>
            <div class="box__ghost-shadow"></div>
          </div>

          <div class="box__description">
            <div class="box__description-container">
              <template v-if="error.statusCode === 404" class="title">
                <h1 class="title">Whoops!</h1>
                <div class="box__description-text">
                  <p class="is-medium">
                    It seems like we couldn't find the page you were looking for
                  </p>
                </div>
              </template>
              <template v-else>
                <h3 class="title">Whoops!</h3>
                <div class="box__description-text">
                  Something went wrong
                </div></template
              >
            </div>

            <nuxt-link
              class="button is-white"
              :to="{
                path: '/',
              }"
              >Share a secret</nuxt-link
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
let pageX = 0
let pageY = 0
let mouseY = 0
let mouseX = 0
if (process.browser) {
  pageX = window.innerWidth
  pageY = window.innerHeight
}

export default {
  props: {
    error: {
      type: Object,
      default: null,
    },
  },
  mounted() {
    window.addEventListener('mousemove', (e) => {
      // verticalAxis
      mouseY = e.pageY
      const yAxis = ((pageY / 2 - mouseY) / pageY) * 300
      // horizontalAxis
      mouseX = e.pageX / -pageX
      const xAxis = -mouseX * 100 - 100

      document.getElementById('ghost-eyes').style.transform =
        'translate(' + xAxis + '%,-' + yAxis + '%)'
    })

    window.addEventListener('resize', (e) => {
      pageX = window.innerWidth
      pageY = window.innerHeight
    })
  },
  layout: 'hero', // you can set a custom layout for the error page

  head() {
    return {
      title: this.error.statusCode + ' Error - Secretify',
    }
  },
}
</script>
