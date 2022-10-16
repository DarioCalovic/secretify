<template>
  <div class="container">
    <div class="columns is-vcentered is-centered dotted">
      <div class="shape"></div>
      <div class="created column is-12">
        <transition name="fade">
          <div>
            <h2 class="title is-1">
              Well done <span style="font-weight: 500">🥳</span>
            </h2>
            <h3 class="subtitle is-3">
              Your secret link has been created and is now waiting to be shared.
            </h3>
            <div class="mt-6 mb-2">
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
                      Create new secret
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
                    <div class="level-left">
                      <b-field v-show="false" v-if="!backupLink">
                        <p class="control">
                          <b-button @click="splitKey">Split key</b-button>
                        </p>
                        <p class="control">
                          <b-dropdown
                            aria-role="list"
                            position="is-bottom-left"
                          >
                            <b-button slot="trigger" slot-scope="{ active }">
                              <b-icon
                                :icon="active ? 'menu-up' : 'menu-down'"
                              ></b-icon>
                            </b-button>

                            <b-dropdown-item custom aria-role="listitem"
                              ><i
                                >Additional split key settings coming soon
                              </i></b-dropdown-item
                            >
                          </b-dropdown>
                        </p>
                      </b-field>
                      <b-field v-else>
                        <b-button @click="undoSplitKey"
                          >Undo key splitting</b-button
                        ></b-field
                      >
                    </div>

                    <!-- Right side -->
                    <div class="level-right">
                      <b-button
                        v-clipboard:copy="link"
                        v-clipboard:success="onCopy"
                        class="withIcon"
                        icon-left="content-copy"
                        type="is-primary"
                      >
                        Copy
                      </b-button>
                    </div>
                  </nav>
                </div>
              </div>
            </div>
            <div class="pt-2">
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
      backupLink: '',
      expiresAt: '',
      passphrase: '',
      revealOnce: true,
    }
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
  methods: {
    onCopy(e) {
      Toast.open({
        message: this.toastText,
        type: 'is-black',
      })
    },
    undoSplitKey() {
      this.link = this.backupLink
      this.backupLink = ''
    },
    splitKey() {
      this.backupLink = this.link

      const needed = 2
      const total = 3
      const keys = [
        'a98fsd7fdas987adsf',
        'dhgf879afds7869d44',
        'kljöfbgx89654kjhdf',
      ]

      const parts = this.link.split('#')
      if (parts.length > 1) {
        let exampleLink = parts[0] + '#'
        keys.forEach((k, i) => {
          if (i < needed) {
            exampleLink += k
            if (i + 1 < needed) exampleLink += '+'
          }
        })
        let newText = parts[0] + '\n\nKeys:\n'
        newText += keys.join('\n')
        newText += `\n\nShare each key with another person. In order to retrieve the secret, you need to append a # and at least ${needed} out of the ${total} keys separated by a + to the link.`
        newText += `\n\ne.g. ${exampleLink}`
        this.link = newText
      } else {
        Toast.open({
          message: 'Could not extract link with key',
          type: 'is-danger',
        })
      }
    },
  },
  head() {
    return {
      title: 'Link created - Secretify',
    }
  },
}
</script>
