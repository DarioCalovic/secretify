<template>
  <div class="container">
    <div class="columns is-vcentered is-centered dotted">
      <div class="shape"></div>
      <div class="column is-12">
        <h2 class="title is-1">The safe way to share secrets</h2>
        <h3 class="subtitle is-3">
          Keep sensitive information out of your email and chat logs. Read more
          <a href="/about">about</a> it.
        </h3>
        <div class="card mt-6">
          <b-tabs v-model="activeTab">
            <b-tab-item label="Text">
              <ValidationObserver ref="form">
                <form @submit.prevent="onCreateSecret">
                  <ValidationProvider rules="required" name="Secret" slim>
                    <b-field
                      slot-scope="{ errors }"
                      label="Secret"
                      :type="{ 'is-danger': errors[0] }"
                      :message="
                        errors.length > 0 ? errors : fieldMessages.secret
                      "
                    >
                      <b-input
                        id="secretTextarea"
                        v-model="form.secret.value"
                        :maxlength="policySetting.secret.max_length"
                        type="textarea"
                        class="autosize"
                        placeholder="Tell me a secret.."
                        @input="autoresize"
                      ></b-input>
                    </b-field>
                  </ValidationProvider>
                  <b-field
                    v-show="policySetting.passphrase.required"
                    label="Passphrase"
                  >
                    <b-input
                      v-model="form.passphrase.value"
                      type="password"
                      :required="policySetting.passphrase.required"
                    ></b-input>
                  </b-field>
                  <b-field>
                    <b-collapse
                      v-model="isOpen"
                      aria-id="collapseAddInfos"
                      animation="slide"
                    >
                      <div>
                        <b-field
                          v-show="!policySetting.passphrase.required"
                          label="Passphrase"
                          message="Only the one with the passphrase can reveal the secret. Not even an administrator can restore it."
                        >
                          <b-input
                            v-model="form.passphrase.value"
                            type="password"
                            :required="policySetting.passphrase.required"
                          ></b-input>
                        </b-field>
                        <b-field
                          label="Expires in"
                          message="Once the secret expires, it will be deleted automatically."
                        >
                          <b-select
                            v-model="form.expiresAt.value"
                            placeholder="Select a time"
                          >
                            <option value="5m">5 minutes</option>
                            <option value="1h">1 hour</option>
                            <option value="24h">1 day</option>
                            <option value="168h">7 days</option>
                          </b-select>
                        </b-field>
                        <b-field
                          message="Revealed once, the secret will be deleted automatically."
                        >
                          <b-switch v-model="form.revealOnce.value"
                            >Reveal secret only once</b-switch
                          >
                        </b-field>
                        <b-field
                          v-show="!form.revealOnce.value"
                          message="Disabled, the secret will be deleted automatically when it expires."
                        >
                          <b-switch v-model="form.destroyManual.value"
                            >Secret can be destroyed manually</b-switch
                          >
                        </b-field>
                      </div>
                    </b-collapse>
                  </b-field>

                  <div class="level">
                    <div class="level-left">
                      <div class="level-item">
                        <b-icon
                          class="action-icon"
                          icon="dots-horizontal"
                          @click.native="isOpen = !isOpen"
                        >
                        </b-icon>
                      </div>
                      <div class="level-item">
                        <b-icon
                          class="action-icon"
                          icon="shuffle-variant"
                          @click.native="shuffleSecret"
                        >
                        </b-icon>
                      </div>
                    </div>
                    <div class="level-right">
                      <div class="level-item">
                        <b-field class="is-grouped is-grouped-right">
                          <div class="control">
                            <b-button native-type="submit" type="is-primary"
                              >Create secret link</b-button
                            >
                          </div>
                        </b-field>
                      </div>
                    </div>
                  </div>
                </form>
              </ValidationObserver>
            </b-tab-item>
            <b-tab-item v-if="policySetting.storage.enabled" label="File">
              <ValidationObserver ref="formFile">
                <form @submit.prevent="onCreateSecretFile">
                  <ValidationProvider
                    :rules="
                      'required|size:' +
                      policySetting.storage.filesystem.max_filesize / 1000 +
                      (policySetting.storage.filesystem.allowed_file_extensions
                        ? '|ext:' +
                          policySetting.storage.filesystem
                            .allowed_file_extensions
                        : '')
                    "
                    name="File"
                    slim
                  >
                    <b-field
                      slot-scope="{ errors }"
                      :message="errors.length > 0 ? errors : ''"
                      type="is-danger"
                    >
                      <b-upload
                        v-model="file"
                        drag-drop
                        expanded
                        validation-message="Please select a file"
                        :type="errors.length > 0 ? 'is-danger' : ''"
                      >
                        <section class="section">
                          <div class="content has-text-centered">
                            <p>
                              <b-icon icon="upload" size="is-large"> </b-icon>
                            </p>
                            <span v-if="!file">
                              <p>
                                Drop your file here or click to upload (Max.
                                {{
                                  policySetting.storage.filesystem.max_filesize
                                    | formatBytes
                                }})
                              </p>
                              <p
                                v-if="
                                  policySetting.storage.filesystem
                                    .allowed_file_extensions
                                "
                                class="help"
                              >
                                Allowed file extensions are
                                {{
                                  policySetting.storage.filesystem
                                    .allowed_file_extensions
                                }}
                              </p>
                            </span>
                            <span v-else>
                              <p>
                                {{ file.name }} ({{ file.size | formatBytes }})
                              </p>
                            </span>
                          </div>
                        </section>
                      </b-upload></b-field
                    > </ValidationProvider
                  ><b-field>
                    <b-collapse
                      v-model="isOpenFile"
                      aria-id="collapseAddInfosFile"
                      animation="slide"
                    >
                      <div class="file--block">
                        <b-field
                          label="Message"
                          message="The message will be encrypted client-side as well as the file."
                        >
                          <b-input
                            v-model="form.secret.value"
                            :maxlength="policySetting.secret.max_length / 2"
                            type="textarea"
                            placeholder=""
                            :required="false"
                          ></b-input>
                        </b-field>
                        <b-field
                          v-show="!policySetting.passphrase.required"
                          label="Passphrase"
                          message="Only the one with the passphrase can reveal the secret. Not even an administrator can restore it."
                        >
                          <b-input
                            v-model="form.passphrase.value"
                            type="password"
                            :required="policySetting.passphrase.required"
                          ></b-input>
                        </b-field>
                        <b-field
                          label="Expires in"
                          message="Once the secret expires, it will be deleted automatically."
                        >
                          <b-select
                            v-model="form.expiresAt.value"
                            placeholder="Select a time"
                          >
                            <option value="5m">5 minutes</option>
                            <option value="1h">1 hour</option>
                            <option value="24h">1 day</option>
                            <option value="168h">7 days</option>
                          </b-select>
                        </b-field>
                        <b-field
                          message="Revealed once, the secret will be deleted automatically."
                        >
                          <b-switch v-model="form.revealOnce.value"
                            >Reveal secret only once</b-switch
                          >
                        </b-field>
                      </div>
                    </b-collapse>
                  </b-field>
                  <div class="level">
                    <div class="level-left">
                      <div class="level-item">
                        <b-icon
                          class="action-icon"
                          icon="dots-horizontal"
                          @click.native="isOpenFile = !isOpenFile"
                        >
                        </b-icon>
                      </div>
                    </div>
                    <div class="level-right">
                      <div class="level-item">
                        <b-field class="is-grouped is-grouped-right">
                          <div class="control">
                            <b-button native-type="submit" type="is-primary"
                              >Create secret link</b-button
                            >
                          </div>
                        </b-field>
                      </div>
                    </div>
                  </div>
                </form>
              </ValidationObserver></b-tab-item
            >
          </b-tabs>
        </div>
        <div class="pt-2">
          <nav class="level">
            <!-- Left side -->
            <div class="level-left">
              <div class="level-item">
                <div class="icon-text">
                  <span class="icon"><i class="mdi mdi-lock-outline"></i></span>
                  <span>End-to-end encrypted</span>
                </div>
              </div>
            </div>
          </nav>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { ValidationProvider, ValidationObserver } from 'vee-validate'
import { ToastProgrammatic as Toast } from 'buefy'
import { mapState } from 'vuex'
const generator = require('generate-password')

export default {
  name: 'Home',
  layout: 'hero',
  components: {
    ValidationProvider,
    ValidationObserver,
  },
  data() {
    return {
      activeTab: 0,
      radio: 'none',
      isOpen: false,
      isOpenFile: false,
      toastText: 'You just copied the secret link!',
      form: {
        secret: {
          value: '',
        },
        passphrase: {
          value: '',
        },
        expiresAt: {
          value: '24h',
        },
        revealOnce: {
          value: true,
        },
        destroyManual: {
          value: true,
        },
      },
      fieldMessages: {
        secret:
          'The secret will be encrypted client-side and only the cipher and meta information will be stored safe until it gets deleted. The key will not be stored anywhere but added only to the link.',
      },
      file: {
        name: null,
        size: 0,
      },
    }
  },
  computed: {
    ...mapState({
      policySetting: (state) => {
        return state.policySetting
      },
      offline: (state) => {
        return state.offline
      },
    }),
  },
  mounted() {
    if (this.offline) {
      Toast.open({
        message: 'API seems offline. Please contact administrator.',
        type: 'is-danger',
      })
    }
    // Track
    this.$track.pageview({})
  },
  methods: {
    autoresize() {
      const element = document.getElementById('secretTextarea')
      element.style.height = '5px'
      element.style.height = element.scrollHeight + 2 + 'px'
    },
    shuffleSecret() {
      const password = generator.generate({
        length: 30,
        numbers: true,
        symbols: true,
        lowercase: true,
        uppercase: true,
        strict: true,
      })
      this.form.secret.value = password
    },
    onCreateSecret() {
      this.$refs.form.validate().then((success) => {
        if (!success) {
          return
        }
        this.createSecret()
      })
    },
    onCreateSecretFile() {
      this.$refs.formFile.validate().then((success) => {
        if (!success) {
          return
        }
        this.createSecret()
      })
    },
    async createSecret() {
      let passphrase = this.form.passphrase.value
      let cipher = this.form.secret.value

      // Hash passphrase if present
      if (passphrase) {
        passphrase = this.$crypto.createHashWithoutPadding(passphrase)
      }

      // Client encryption
      const key = await this.$crypto.generateEncryptionKeyString()
      // Handle file input
      let fileIdentifier
      if (this.file && this.file.name && this.file.siz > 0) {
        let fileCipher = this.file
        if (passphrase) {
          fileCipher = await this.$crypto.encryptFile(fileCipher, passphrase)
        }
        const encryptedFile = await this.$crypto.encryptFile(fileCipher, key)

        const formData = new FormData()
        formData.append('Content-type', 'application/octet-stream')
        formData.append('file', encryptedFile)

        // Upload file
        const res = await this.$repositories.file
          .upload(this.file.name, this.file.size, this.file.type, formData)
          .catch((e) => {
            Toast.open({
              message:
                'Error while trying to upload file: ' + e.response.data.error,
              type: 'is-danger',
            })
          })
        if (res) {
          const { status, data } = res
          if (status === 200 && data) {
            fileIdentifier = data.data.identifier
          }
        } else {
          return
        }

        // Prepare meta information for cipher since it is a shared file
        cipher = JSON.stringify({
          text: cipher,
          file_id: fileIdentifier,
        })
      }

      if (cipher) {
        if (passphrase) {
          cipher = await this.$crypto.encryptString(cipher, passphrase)
        }
        cipher = await this.$crypto.encryptString(cipher, key)

        // Create secret
        const res = await this.$repositories.secret
          .create(
            cipher,
            this.form.expiresAt.value,
            this.form.revealOnce.value,
            this.form.destroyManual.value,
            !!passphrase,
            fileIdentifier
          )
          .catch((err) => {
            Toast.open({
              message: err,
              type: 'is-danger',
            })
          })
        if (res) {
          const { status, data } = res
          if (status === 200 && data) {
            // Change route
            this.$router.push({
              name: 'created',
              params: {
                link: `${this.$config.uiURL}/s/${data.data.identifier}#${key}`,
                expiresAt: this.form.expiresAt.value,
                passphrase,
                revealOnce: this.form.revealOnce.value,
              },
            })
          }
        }
      } else {
        Toast.open({
          message: 'Something went wrong while preparing the cipher',
          type: 'is-danger',
        })
      }
    },
  },
  head() {
    return {
      title: 'Home - Secretify',
    }
  },
}
</script>
