<template>
  <div class="container">
    <div class="columns is-centered dotted">
      <div class="column is-5" v-show="seen">
        <div class="card">
          <div class="card-content">
            <div class="content">
              <b-field label="Secret">
                <b-input
                  class="dark"
                  v-model="formSecret.secret.value"
                  type="textarea"
                  readonly
                  expanded
                ></b-input>
              </b-field>

              <p>
                Expires in
                <strong>{{ formSecret.expiresAt.value }}</strong> | Passphrase
                needed?
                <strong>{{ formSecret.passphrase.value | yesOrNo }}</strong> |
                Reveal only once?
                <strong>{{ formSecret.revealOnce.value | yesOrNo }}</strong>
              </p>
            </div>
          </div>
        </div>
      </div>
      <div class="column is-5" v-show="seen">
        <div class="card">
          <div class="content">
            <b-tabs v-model="activeTab">
              <b-tab-item label="Link">
                <form ref="form">
                  <b-field
                    label="Email"
                    message="Your email address is used only for security reasons. This will also allow us to send you a copy of the link and notify you when somebody reveals the secret."
                  >
                    <b-input
                      v-model="form.email.value"
                      type="email"
                      placeholder="Your email"
                    ></b-input>
                  </b-field>
                  <b-field class="is-grouped is-grouped-right">
                    <div class="control">
                      <b-button type="is-primary" @click="generateLink"
                        >Get a shared link</b-button
                      >
                    </div>
                  </b-field>
                </form>
              </b-tab-item>

              <b-tab-item label="Email" disabled>Not supported yet</b-tab-item>
            </b-tabs>
          </div>
        </div>
      </div>
      <div class="column is-5" v-show="!seen">
        <transition name="fade">
          <div class="card">
            <div class="card-content">
              <div class="content">
                <h3 class="title">Hurray!</h3>
                <p>
                  The secret has been successfully encrypted and is now waiting
                  to be shared:
                </p>
                <cb-textarea :text="response.secretURL" />
                <p>
                  Expires in
                  <strong>{{ formSecret.expiresAt.value }}</strong> | Passphrase
                  needed?
                  <strong>{{ formSecret.passphrase.value | yesOrNo }}</strong> |
                  Reveal only once?
                  <strong>{{ formSecret.revealOnce.value | yesOrNo }}</strong>
                </p>
              </div>
            </div>
          </div>
        </transition>
      </div>
    </div>
  </div>
</template>

<script>
import CbTextarea from '@/components/CbTextarea'

export default {
  name: 'Prepare',
  layout: 'hero',
  components: { CbTextarea },
  filters: {
    yesOrNo(val) {
      return val ? 'yes' : 'no'
    },
  },
  created() {
    if (!('form' in this.$route.params)) {
      this.$router.push({ path: '/' })
      return
    }
    this.formSecret = this.$route.params.form
  },
  mounted() {
    const email = window.localStorage.getItem('email')
    if (email) {
      this.form.email.value = email
    }
  },
  data() {
    return {
      activeTab: 0,
      seen: true,
      response: {
        secretURL: '',
        createdAt: null,
      },
      formSecret: {
        secret: {
          value: '',
        },
        passphrase: {
          value: '',
        },
        expiresAt: {
          value: '5m',
        },
        revealOnce: {
          value: true,
        },
      },
      form: {
        email: {
          value: '',
        },
      },
    }
  },
  methods: {
    handleSubmit(e) {
      console.log('submit')
    },
    async generateLink(event) {
      if (!this.$refs.form.checkValidity()) {
        this.$refs.form.reportValidity()
        return
      }

      window.localStorage.setItem('email', this.form.email.value)

      const res = await this.$repositories.secret
        .encrypt(
          this.form.email.value,
          this.formSecret.secret.value,
          this.formSecret.passphrase.value,
          this.formSecret.expiresAt.value,
          this.formSecret.revealOnce.value
        )
        .catch((e) => {
          console.log(e)
        })
      if (res) {
        const { status, data } = res
        if (status === 200 && data) {
          this.response.createdAt = data.data.created_at
          this.response.secretURL = `${this.$config.uiURL}/s/${data.data.identifier}`
          this.seen = !this.seen
        }
      }
    },
  },
  head() {
    return {
      title: 'Share - Secretify',
    }
  },
}
</script>
