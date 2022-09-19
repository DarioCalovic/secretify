<template>
  <section>
    <b-modal
      v-model="isComponentModalActive"
      has-modal-card
      trap-focus
      :destroy-on-hide="false"
      aria-role="dialog"
      aria-label="Example Modal"
      close-button-aria-label="Close"
      aria-modal
    >
      <template>
        <form action="">
          <div class="modal-card" style="width: auto">
            <header class="modal-card-head">
              <p class="modal-card-title">Login</p>
              <button type="button" class="delete" @click="$emit('close')" />
            </header>
            <section class="modal-card-body">
              <b-field label="Email">
                <b-input
                  type="email"
                  :value="email"
                  placeholder="Your email"
                  required
                >
                </b-input>
              </b-field>

              <b-field label="Password">
                <b-input
                  type="password"
                  :value="password"
                  password-reveal
                  placeholder="Your password"
                  required
                >
                </b-input>
              </b-field>

              <b-checkbox>Remember me</b-checkbox>
            </section>
            <footer class="modal-card-foot">
              <b-button label="Close" @click="$emit('close')" />
              <b-button label="Login" type="is-primary" />
            </footer>
          </div>
        </form>
      </template>
    </b-modal>
  </section>
</template>

<script>
export default {
  data() {
    return {
      isComponentModalActive: false,
      formProps: {
        email: 'evan@you.com',
        password: 'testing',
      },
    }
  },
}
</script>
