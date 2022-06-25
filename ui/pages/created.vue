<template>
  <div class="container">
    <div class="columns is-vcentered is-centered dotted">
      <div class="created column is-12">
        <transition name="fade">
          <div>
            <h2 class="title is-1">
              Well done <span style="font-weight: 500">🥳</span>
            </h2>
            <h3 class="subtitle is-3">
              Your secret link has been created and is now waiting to be shared.
            </h3>
            <div class="white mt-6 mb-2">
              <nav class="level">
                <!-- Left side -->
                <div class="level-left">
                  <div class="level-item">
                    <b-button
                      tag="router-link"
                      type="is-text"
                      icon-left="reload"
                      :to="{ name: 'index', force: true }"
                    >
                      Create secret
                    </b-button>
                  </div>
                </div>
              </nav>
            </div>
            <div class="card">
              <div class="card-content">
                <div class="content">
                  <cb-textarea :text="link" />

                  <nav class="level">
                    <!-- Left side -->
                    <div class="level-left"></div>

                    <!-- Right side -->
                    <div class="level-right">
                      <!--<b-button
                        tag="router-link"
                        type="is-text"
                        :to="{ name: 'index', force: true }"
                      >
                        Share
                      </b-button>-->

                      <b-button
                        class="withIcon"
                        icon-left="content-copy"
                        type="is-primary"
                        v-clipboard:copy="link"
                        v-clipboard:success="onCopy"
                      >
                        Copy
                      </b-button>
                    </div>
                  </nav>
                </div>
              </div>
            </div>
            <div class="white pt-2">
              <nav class="level">
                <!-- Left side -->
                <div class="level-left"></div>

                <!-- Right side -->
                <div class="level-right">
                  <p>
                    Expires in
                    <strong>{{ expiresAt }}</strong> | Passphrase needed?
                    <strong>{{ passphrase | yesOrNo }}</strong> | Reveal only
                    once?
                    <strong>{{ revealOnce | yesOrNo }}</strong>
                  </p>
                </div>
              </nav>
            </div>
          </div>
        </transition>
      </div>
    </div>
  </div>
</template>
<style>
.created .textarea:not([rows]) {
  min-height: 3em;
}
</style>
<script>
import { ToastProgrammatic as Toast } from 'buefy'
import CbTextarea from '@/components/CbTextarea'

export default {
  name: 'LinkCreated',
  layout: 'hero',
  components: { CbTextarea },
  filters: {
    yesOrNo(val) {
      return val ? 'yes' : 'no'
    },
  },
  data() {
    return {
      toastText: 'You just copied the secret link!',
      link: '',
      expiresAt: '',
      passphrase: '',
      revealOnce: true,
    }
  },
  methods: {
    onCopy(e) {
      Toast.open(this.toastText)
    },
  },
  mounted() {
    if (!this.$route.params.link) {
      this.$router.push({ name: 'index' })
    }

    this.link = this.$route.params.link
    this.expiresAt = this.$route.params.expiresAt
    this.passphrase = this.$route.params.passphrase
    this.revealOnce = this.$route.params.revealOnce

    // Track
    this.$track.pageview({})
  },
  head() {
    return {
      title: 'Link created - Secretify',
    }
  },
}
</script>
