<template>
  <div class="container">
    <div class="columns is-vcentered is-centered dotted">
      <div class="column">
        <div :v-show="secret.valid">
          <div v-show="!reveal">
            <h3 class="title is-1">
              You shall not pass <span style="font-weight: 500">🧙</span>
            </h3>
            <h3 class="subtitle is-3">
              This secret is proteced with a passphrase.
            </h3>
            <div class="mt-6">
              <ValidationObserver ref="form">
                <form @submit.prevent="onRetrieveChipher">
                  <div v-if="secret.hasPassphrase">
                    <ValidationProvider rules="required" name="Passphrase" slim>
                      <b-field
                        slot-scope="{ errors }"
                        :type="{ 'is-danger': errors[0] }"
                        :message="errors.length > 0 ? errors : ''"
                      >
                        <b-input
                          v-model="form.passphrase"
                          placeholder="Passphrase*"
                          type="password"
                          password-reveal
                        >
                        </b-input>
                      </b-field>
                    </ValidationProvider>
                    <b-field>
                      <b-button
                        native-type="submit"
                        type="is-primary"
                        label="Reveal secret"
                        :disabled="form.passphrase == ''"
                      />
                    </b-field>
                  </div>
                </form>
              </ValidationObserver>
            </div>
          </div>
          <transition name="fade">
            <div v-show="reveal">
              <h3 class="title is-1">
                Poof <span style="font-weight: 500">🧙💨</span>
              </h3>
              <h3 class="subtitle is-3">The secret is revealed.</h3>
              <div class="card mt-6">
                <progress-bar
                  v-if="secret.revealOnce"
                  :duration="policySetting.secret.reveal_duration"
                />
                <div class="card-content">
                  <div class="content">
                    <div v-if="secret.file.identifier">
                      <b-notification
                        type="is-danger"
                        has-icon
                        aria-close-label="Close notification"
                        role="alert"
                        :closable="false"
                      >
                        Only download files from trusted sources. All files are
                        encrypted on our servers. Therefore we have absolutely
                        no knowledge about the contents of the file.
                      </b-notification>
                      <table class="table">
                        <tbody>
                          <tr>
                            <td><strong>Filename</strong></td>
                            <td>
                              {{ secret.file.filename }}
                            </td>
                          </tr>
                          <tr>
                            <td><strong>Type</strong></td>
                            <td>
                              {{ secret.file.type }}
                            </td>
                          </tr>
                          <tr>
                            <td><strong>Size</strong></td>
                            <td>
                              {{ secret.file.size | formatBytes }}
                            </td>
                          </tr>
                          <tr>
                            <td><strong>Message</strong></td>
                            <td>
                              {{ decryptedValue }}
                            </td>
                          </tr>
                        </tbody>
                      </table>
                    </div>
                    <cb-textarea
                      v-else
                      :text="decryptedValue"
                      :first="decryptedValueFirst"
                      :rest="decryptedValueRest"
                      toastText="You just copied the secret!"
                    />
                    <nav class="level">
                      <!-- Left side -->
                      <div class="level-left"></div>

                      <!-- Right side -->
                      <div class="level-right buttons">
                        <b-button
                          v-if="!secret.revealOnce && secret.destroyManual"
                          type="is-text"
                          @click="destroySecret"
                        >
                          Destroy
                        </b-button>
                        <b-button
                          v-if="!secret.file.identifier"
                          class="withIcon"
                          icon-left="content-copy"
                          type="is-primary"
                          v-clipboard:copy="decryptedValue"
                          v-clipboard:success="onCopy"
                        >
                          <template v-if="!decryptedValueFirst">Copy</template
                          ><template v-else>Copy all</template>
                        </b-button>
                        <b-button
                          v-if="!secret.file.identifier && decryptedValueFirst"
                          class="withIcon"
                          icon-left="content-copy"
                          type="is-primary"
                          v-clipboard:copy="decryptedValueFirst"
                          v-clipboard:success="onCopy"
                        >
                          Copy first line
                        </b-button>
                        <b-button
                          v-if="secret.file.identifier"
                          class="withIcon"
                          icon-left="file-download"
                          type="is-primary"
                          @click="downloadFile"
                        >
                          Download
                        </b-button>
                      </div>
                    </nav>
                  </div>
                </div>
              </div>

              <div class="white pt-2">
                <nav class="level">
                  <!-- Left side -->
                  <div class="level-left">
                    <div class="level-item">
                      <b-button
                        tag="router-link"
                        type="is-text"
                        icon-left="reply"
                        :to="{ name: 'index', force: true }"
                      >
                        Reply with a secret
                      </b-button>
                    </div>
                  </div>
                  <!-- Right side -->
                  <div class="level-right">
                    <p v-if="secret.deleted">
                      Any traces were already deleted on the system.
                      <span v-if="secret.file.identifier && secret.revealOnce">
                        The file will be deleted soon as well.</span
                      >
                    </p>
                    <p v-if="!secret.revealOnce">
                      Secret will be destroyed
                      {{ secret.expires_at | diffValidity }}.
                    </p>
                  </div>
                </nav>
              </div>
            </div>
          </transition>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { ValidationProvider, ValidationObserver } from 'vee-validate'
import { mapState } from 'vuex'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import CbTextarea from '@/components/CbTextarea'
import ProgressBar from '@/components/ProgressBar'
import { ToastProgrammatic as Toast } from 'buefy'
dayjs.extend(relativeTime)
export default {
  name: 'Harpocrates',
  layout: 'hero',
  components: {
    CbTextarea,
    ProgressBar,
    ValidationProvider,
    ValidationObserver,
  },
  computed: {
    ...mapState({
      policySetting: (state) => {
        return state.policySetting
      },
    }),
  },
  async asyncData({ error, params, $repositories }) {
    const res = await $repositories.secret.view(params.id).catch((err) => {
      return error({
        statusCode: err.response.status,
        message: err.message,
      })
    })
    if (res) {
      const { status, data } = res
      if (status === 200 && data) {
        return {
          reveal: !data.data.has_passphrase,
          secret: {
            valid: true,
            expires_at: data.data.expires_at,
            revealOnce: data.data.reveal_once,
            destroyManual: data.data.destroy_manual,
            hasPassphrase: data.data.has_passphrase,
            deleted: data.data.deleted,
            file: data.data.file,
          },
        }
      }
    }
  },
  filters: {
    diffValidity(val) {
      return dayjs(val).fromNow()
    },
  },
  data() {
    return {
      form: {
        passphrase: '',
      },
      decryptedValue: '',
      decryptedValueFirst: '',
      decryptedValueRest: '',
      encryptionKey: '',
      toastText: 'You just copied the secret!',
    }
  },
  methods: {
    async downloadFile() {
      const res = await this.$repositories.file
        .download(this.secret.file.identifier)
        .catch((err) => {
          Toast.open({
            message: 'File could not be downloaded',
            type: 'is-danger',
          })
          return err
        })
      if (res) {
        const { status, data } = res
        if (status === 200) {
          const decryptedFile = await this.$crypto.decryptFile(
            data,
            this.encryptionKey,
            this.secret.file.filename
          )

          const url = window.URL.createObjectURL(decryptedFile)
          const link = document.createElement('a')
          link.href = url
          link.setAttribute('download', this.secret.file.filename)
          document.body.appendChild(link)
          link.click()
        }
      }
    },
    onCopy(e) {
      Toast.open(this.toastText)
    },
    async getCipher(id) {
      const res = await this.$repositories.secret.cipher(id).catch((err) => {
        return err
      })
      if (res) {
        const { status, data } = res
        if (status === 200 && data) {
          this.secret.deleted = data.data.deleted
          return data.data.cipher
        }
      }
    },
    async destroySecret() {
      const res = await this.$repositories.secret
        .delete(this.$route.params.id)
        .catch((err) => {
          Toast.open({
            message: 'Secret could not be destroyed',
            type: 'is-danger',
          })
          return err
        })
      if (res) {
        const { status } = res
        if (status === 200) {
          Toast.open('Secret destroyed!')
          setTimeout(function () {
            window.location.reload(true)
          }, 1500)
        }
      }
    },
    onRetrieveChipher() {
      this.$refs.form.validate().then((success) => {
        if (!success) {
          return
        }
        this.retrieveCipher()
      })
    },
    retrieveCipher() {
      this.getCipher(this.$route.params.id)
        .then((cipher) => {
          this.revealSecret(cipher)
        })
        .catch((err) => {
          console.log(err)
        })
    },
    async revealSecret(cipher) {
      if (cipher) {
        // Decrypt first with key from url hash
        let decrypted = await this.$crypto.decryptString(
          cipher,
          this.encryptionKey
        )

        // Decrypt with passphrase
        if (this.secret.hasPassphrase) {
          const passphraseHash = this.$crypto.createHashWithoutPadding(
            this.form.passphrase
          )
          try {
            decrypted = await this.$crypto.decryptString(
              decrypted,
              passphraseHash
            )
          } catch (e) {
            Toast.open({
              message: 'Wrong passphrase!',
              type: 'is-danger',
            })
            return
          }
        }
        if (this.secret.file.identifier) {
          try {
            const d = JSON.parse(decrypted)
            this.decryptedValue = d.text
          } catch (e) {
            Toast.open({
              message: 'Secret is corrupt',
              type: 'is-danger',
            })
            return
          }
        } else {
          this.decryptedValue = decrypted
          if (/\r|\n/.exec(this.decryptedValue)) {
            const parts = this.decryptedValue.split('\n')
            this.decryptedValueFirst = parts[0]
            if (parts.length > 1) {
              this.decryptedValueRest = '\n' + parts.slice(1).join('\n')
            }
          }
        }
      }
      this.reveal = true
    },
  },
  mounted() {
    this.encryptionKey = this.$route.hash.substring(1)

    if (!this.secret.hasPassphrase) {
      this.retrieveCipher()
    }

    // Track
    this.$track.pageview({ url: '/s/{id}' }) // generalize secret link for tracking
  },
  head() {
    return {
      title: 'Reveal Secret - Secretify',
    }
  },
}
</script>
