<template>
  <div class="container">
    <b-navbar>
      <template #brand>
        <!--<b-navbar-item
          class="title"
          tag="router-link"
          :to="{ path: '/', params: { yourCustomParam: Date.now() } }"
        >
          Secretify
        </b-navbar-item>-->
        <a href="/" class="navbar-item title nuxt-link-active"
          ><img src="~/assets/img/logo.svg" /> secretify</a
        >
      </template>
      <template #start>
        <b-navbar-item tag="router-link" :to="{ path: '/about' }">
          About
        </b-navbar-item>
      </template>
      <template #end>
        <b-navbar-item v-show="$config.marketing.show === 'true'" tag="div">
          <b-button tag="router-link" to="/pricing" type="is-primary">
            Pricing
          </b-button>
        </b-navbar-item>
        <b-navbar-item v-if="$config.branding.logo" tag="div"
          ><img :src="$config.branding.logo"
        /></b-navbar-item>
      </template>
    </b-navbar>
  </div>
</template>
<script>
export default {
  name: 'NavBar',
}
</script>
