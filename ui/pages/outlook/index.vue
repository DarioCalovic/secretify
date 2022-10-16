<template>
  <div class="container">
    <div class="columns is-vcentered is-centered dotted">
      <div class="shape"></div>
      <div class="column is-12">
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
                  <div class="error-message" v-show="errorMessage">
                    {{ errorMessage }}
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
                  <div class="error-message" v-show="errorMessage">
                    {{ errorMessage }}
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
  name: 'Outlook',
  layout: 'office',
  components: {
    ValidationProvider,
    ValidationObserver,
  },
  data() {
    return {
      errorMessage: '',
      isOfficeInitialized: false,
      selectedText: '',
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
      file: null,
    }
  },
  computed: {
    ...mapState({
      policySetting: (state) => {
        return state.policySetting
      },
    }),
  },
  mounted() {
    const This = this
    if (window.Office) {
      try {
        /*eslint-disable */
        Office.initialize = (_reason) => {
          this.isOfficeInitialized = true

          Office.context.mailbox.item.getSelectedDataAsync(
            Office.CoercionType.Text,
            {},
            function callback(result) {
              This.form.secret.value = result.value.data
            }
          )
        }
      } catch (e) {
        // do something
        console.warn('[office.js] ', e)
      }
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
      if (this.file) {
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
            const link = `${this.$config.uiURL}/s/${data.data.identifier}#${key}`
            const linkText = `<a href="${link}">${link}</a>`
            const uiYerOrNo = !!passphrase ? 'yes' : 'no'
            const uiRevealOnce = this.form.revealOnce.value ? 'yes' : 'no'
            const uiExpiresAt = this.form.expiresAt.value
            const prefixText = fileIdentifier
              ? `File <strong>${this.file.name}</strong>`
              : `Secret`
            let replaceText = `${linkText}<br/>
    ${prefixText} expires in <strong>${uiExpiresAt}</strong> | Passphrase needed? <strong>${uiYerOrNo}</strong> | Reveal only once? <strong>${uiRevealOnce}</strong>`
            // Put in email
            const item = Office.context.mailbox.item

            item.body.getTypeAsync(function (result) {
              if (result.status == Office.AsyncResultStatus.Failed) {
                This.errorMessage = result.error.message
              } else {
                // Successfully got the type of item body.
                // Set data of the appropriate type in body.
                if (result.value == Office.MailboxEnums.BodyType.Html) {
                  // Body is of HTML type.
                  // Specify HTML in the coercionType parameter
                  // of setSelectedDataAsync.
                  item.body.setSelectedDataAsync(
                    replaceText,
                    {
                      coercionType: Office.CoercionType.Html,
                    },
                    function (asyncResult) {
                      if (
                        asyncResult.status == Office.AsyncResultStatus.Failed
                      ) {
                        This.errorMessage = asyncResult.error.message
                      } else {
                        // Successfully set data in item body.
                        Office.context.ui.closeContainer()
                      }
                    }
                  )
                } else {
                  // Body is of text type.
                  item.body.setSelectedDataAsync(
                    link,
                    {
                      coercionType: Office.CoercionType.Text,
                    },
                    function (asyncResult) {
                      if (
                        asyncResult.status == Office.AsyncResultStatus.Failed
                      ) {
                        This.errorMessage = asyncResult.error.message
                      } else {
                        // Successfully set data in item body.
                        Office.context.ui.closeContainer()
                      }
                    }
                  )
                }
              }
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
      title: 'Outlook - Secretify',
      script: [
        {
          innerHTML: `
        window._historyCache = {
            replaceState: window.history.replaceState,
            pushState: window.history.pushState
        };
     `,
        },
        {
          src: 'https://appsforoffice.microsoft.com/lib/1/hosted/office.js',
        },
        {
          innerHTML: `
        // And restore them
        window.history.replaceState = window._historyCache.replaceState;
        window.history.pushState = window._historyCache.pushState;
     `,
        },
      ],
    }
  },
}
</script>
