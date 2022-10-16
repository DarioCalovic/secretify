<template>
  <section>
    <footer class="footer">
      <div class="container">
        <div class="level">
          <div class="level-left">
            <div class="level-item pr-3">
              <div class="footer-copyright">
                <div>
                  &copy;
                  <a href="https://www.secretify.io">secretify.io</a>
                  {{ year }} under GNU license . Build with a lot of coffee ☕.
                  <span v-if="metaSetting && metaSetting.hoster.name">
                    Hosted for
                    <a :href="metaSetting.hoster.address">{{
                      metaSetting.hoster.name
                    }}</a
                    >.
                  </span>
                </div>
              </div>
            </div>
          </div>
          <div class="level-right">
            <p
              class="level-item"
              v-if="
                $config.track.enabled === 'true' &&
                $config.track.apiURL &&
                $config.track.domain
              "
            >
              <a :href="$config.track.apiURL + '/' + $config.track.domain"
                >Stats</a
              >
            </p>
            <p class="level-item">
              <a href="/imprint">Imprint</a>
            </p>
            <p class="level-item">
              <a href="/privacy">Privacy</a>
            </p>
            <p class="level-item">
              <a href="https://github.com/DarioCalovic/secretify"
                ><b-icon icon="github-circle"> </b-icon
              ></a>
            </p>
          </div>
        </div>
      </div>
    </footer>
  </section>
</template>
<style>
p.level-item {
  margin-bottom: 0;
}
</style>
<script>
import { mapState } from 'vuex'
import dayjs from 'dayjs'

export default {
  name: 'FooterBar',
  computed: {
    ...mapState({
      metaSetting: (state) => {
        return state.metaSetting
      },
      policySetting: (state) => {
        return state.policySetting
      },
    }),
    year() {
      return dayjs().year()
    },
  },
}
</script>
