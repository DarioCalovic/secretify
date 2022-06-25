<template>
  <div class="container">
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
                <h3 class="title">Whoops!</h3>
                <div class="box__description-text">
                  It seems like we couldn't find the page you were looking for
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
              class="button is-primary"
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

<style lang="scss" scoped>
//variables
$l-purple: #0862ff;
$t-purple: #8c8aa7;
$pink: #ff5e65;
$white: #fff;

.box {
  background: transparent;
  border-radius: 0;
  box-shadow: none;

  .box__ghost {
    padding: 15px 25px 25px;
    width: 200px;
    position: relative;
    margin: auto;

    .symbol {
      &:nth-child(1) {
        opacity: 0.2;
        animation: shine 4s ease-in-out 3s infinite;

        &:before,
        &:after {
          content: '';
          width: 12px;
          height: 4px;
          background: $white;
          position: absolute;
          border-radius: 5px;
          bottom: 65px;
          left: 0;
        }
        &:before {
          transform: rotate(45deg);
        }
        &:after {
          transform: rotate(-45deg);
        }
      }
      &:nth-child(2) {
        position: absolute;
        left: -5px;
        top: 30px;
        height: 18px;
        width: 18px;
        border: 4px solid;
        border-radius: 50%;
        border-color: $white;
        opacity: 0.2;
        animation: shine 4s ease-in-out 1.3s infinite;
      }
      &:nth-child(3) {
        opacity: 0.2;
        animation: shine 3s ease-in-out 0.5s infinite;

        &:before,
        &:after {
          content: '';
          width: 12px;
          height: 4px;
          background: $white;
          position: absolute;
          border-radius: 5px;
          top: 5px;
          left: 40px;
        }
        &:before {
          transform: rotate(90deg);
        }
        &:after {
          transform: rotate(180deg);
        }
      }
      &:nth-child(4) {
        opacity: 0.2;
        animation: shine 6s ease-in-out 1.6s infinite;

        &:before,
        &:after {
          content: '';
          width: 15px;
          height: 4px;
          background: $white;
          position: absolute;
          border-radius: 5px;
          top: 10px;
          right: 30px;
        }
        &:before {
          transform: rotate(45deg);
        }
        &:after {
          transform: rotate(-45deg);
        }
      }
      &:nth-child(5) {
        position: absolute;
        right: 5px;
        top: 40px;
        height: 12px;
        width: 12px;
        border: 3px solid;
        border-radius: 50%;
        border-color: $white;
        opacity: 0.2;
        animation: shine 1.7s ease-in-out 7s infinite;
      }
      &:nth-child(6) {
        opacity: 0.2;
        animation: shine 2s ease-in-out 6s infinite;

        &:before,
        &:after {
          content: '';
          width: 15px;
          height: 4px;
          background: $white;
          position: absolute;
          border-radius: 5px;
          bottom: 65px;
          right: -5px;
        }
        &:before {
          transform: rotate(90deg);
        }
        &:after {
          transform: rotate(180deg);
        }
      }
    }

    .box__ghost-container {
      background: $white;
      width: 100px;
      height: 100px;
      border-radius: 100px 100px 0 0;
      position: relative;
      margin: 0 auto;
      animation: upndown 3s ease-in-out infinite;

      .box__ghost-eyes {
        position: absolute;
        left: 50%;
        top: 45%;
        //transform: translate(-50%, -45%);
        height: 12px;
        width: 70px;

        .box__eye-left {
          width: 12px;
          height: 12px;
          background: $l-purple;
          border-radius: 50%;
          margin: 0 10px;
          position: absolute;
          left: 0;
        }
        .box__eye-right {
          width: 12px;
          height: 12px;
          background: $l-purple;
          border-radius: 50%;
          margin: 0 10px;
          position: absolute;
          right: 0;
        }
      }

      .box__ghost-bottom {
        display: flex;
        position: absolute;
        top: 100%;
        left: 0;
        right: 0;

        div {
          flex-grow: 1;
          position: relative;
          top: -10px;
          height: 20px;
          border-radius: 100%;
          background-color: $white;

          &:nth-child(2n) {
            top: -12px;
            margin: 0 -0px;
            border-top: 15px solid $l-purple;
            background: transparent;
          }
        }
      }
    }

    .box__ghost-shadow {
      height: 20px;
      box-shadow: 0 50px 15px 5px #3b3769;
      border-radius: 50%;
      margin: 0 auto;
      animation: smallnbig 3s ease-in-out infinite;
    }
  }

  .box__description {
    margin-top: 5rem;
    margin-bottom: 5rem;
    .box__description-container {
      color: $white;
      text-align: center;
      font-size: 16px;
      margin: 0 auto;
      margin-bottom: 2rem;

      .box__description-title {
        font-size: 24px;
        letter-spacing: 0.5px;
      }

      .box__description-text {
        color: $white;
        line-height: 20px;
        margin-top: 20px;
      }
    }

    .box__button {
      display: block;
      position: relative;
      background: $pink;
      border: 1px solid transparent;
      border-radius: 50px;
      height: 50px;
      text-align: center;
      text-decoration: none;
      color: $white;
      line-height: 50px;
      font-size: 18px;
      padding: 0 70px;
      white-space: nowrap;
      margin-top: 25px;
      transition: background 0.5s ease;
      overflow: hidden;
      -webkit-mask-image: -webkit-radial-gradient(white, black);

      &:before {
        content: '';
        position: absolute;
        width: 20px;
        height: 100px;
        background: $white;
        bottom: -25px;
        left: 0;
        border: 2px solid $white;
        transform: translateX(-50px) rotate(45deg);
        transition: transform 0.5s ease;
      }

      &:hover {
        background: transparent;
        border-color: $white;

        &:before {
          transform: translateX(250px) rotate(45deg);
        }
      }
    }
  }
}

h3.title {
  color: $white;
}

//keyframes
@keyframes upndown {
  0% {
    transform: translateY(5px);
  }
  50% {
    transform: translateY(15px);
  }
  100% {
    transform: translateY(5px);
  }
}
@keyframes smallnbig {
  0% {
    width: 90px;
  }
  50% {
    width: 100px;
  }
  100% {
    width: 90px;
  }
}
@keyframes shine {
  0% {
    opacity: 0.2;
  }
  25% {
    opacity: 0.1;
  }
  50% {
    opacity: 0.2;
  }
  100% {
    opacity: 0.2;
  }
}
</style>

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
