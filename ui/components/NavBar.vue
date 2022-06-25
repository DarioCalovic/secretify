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
          ><img src="~/assets/img/logo-white.png" /> secretify</a
        >
      </template>
      <template #end>
        <!-- <b-navbar-item tag="router-link" :to="{ path: '/login' }">
          Login
        </b-navbar-item>
        -->
        <b-navbar-item tag="router-link" :to="{ path: '/about' }">
          About
        </b-navbar-item>
      </template>
    </b-navbar>
  </div>
</template>
<script>
export default {
  name: 'NavBar',
}
</script>
