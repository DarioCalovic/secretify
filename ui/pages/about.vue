<template>
  <div class="blank">
    <section class="section">
      <div class="container">
        <div class="columns is-centered">
          <div class="column">
            <section class="section" id="overview">
              <h1 class="title">A brief introduction.</h1>
              <h3 class="subtitle is-3">
                Think twice about where you store sensitive information. It
                might be there forever.
              </h3>
              <p class="is-medium">
                Has your personal information been ever exposed in a data
                breach? I hope not. Because the most serious consequences to you
                can occur when sensitive information like pass­words or social
                security numbers are leaked. Criminals usually obtain this data
                either by physical theft, accessing public information, or
                through data breaches. In order to minimize the risk, you should
                think twice wether to put sensitive information like a password
                in an e-mail or a chat conversation.
              </p>
              <p class="is-medium">
                <strong>
                  Your privacy should be concerning you, especially in these
                  days.
                </strong>
                Technologies are changing the world and the daily lives of each
                and every one of us. We are surrounded by technology at every
                moment of our lives. Some mistake will be made along the way.
                People give away credit card information, passwords, social
                security numbers or any other sensitive information with ease.
                It might seem safe at first glance, but once you have fallen
                victim, you would have wished being more careful with your
                sensitive information.
              </p>
            </section>
            <section class="section" id="how-it-works">
              <h3 class="subtitle is-3">How does it work?</h3>
              <div class="cl">
                {{ $config.uiURL }}/s/<span class="tag is-warning is-medium"
                  >n9w7fECxMKK8JwdHLB</span
                >#<span class="tag is-black is-medium"
                  >QVA6dhJgmxyycdUjdRxDKzhB7x</span
                >
              </div>
              <p class="is-medium">
                Secretify will generate a secret link you can share with others.
                The secret link is composed of the
                <span class="tag is-warning">identifier</span> and the
                <span class="tag is-black">encryption key</span>. The key is
                only known to you and the people you share it with. As soon as
                the secret expires, it will be deleted automatically. Since
                everything is encrypted end-to-end, not even an administrator
                with privileged access can make sense of any of the secrets
                being stored.
              </p>
              <p class="is-medium">
                Do not take our word for it. All of the code is open-source and
                available on
                <a href="https://github.com/DarioCalovic/secretify">Github</a>.
              </p>
            </section>
            <section class="section" id="security-aspects">
              <h3 class="subtitle is-3">Security concerns you? Good!</h3>
              <p class="is-medium">
                All secrets are <strong>end-to-end encrypted</strong> using
                state of the art algorithmic encryption
                <strong>AES-256</strong> (Advanced Encryption Standard). The
                generated key is never being stored on the servers, but added to
                the secret as a hash link. The contents of the hashes are never
                being sent to the server, but are only available for the client
                (the browser).
              </p>
            </section>
            <section class="section" id="premium-features">
              <h3 class="subtitle is-3">Why Secretify Premium?</h3>
              <p class="is-medium">
                Add enterprise level features like upgrade assistance and a
                technical manager as well as enterprise readiness features like
                high availability. Among other things also integrate Secretify
                in your business landscape with our Microsoft Outlook and
                Plugins plugins.
              </p>
            </section>
            <section class="section" id="get-it">
              <h3 class="subtitle is-3">
                You want Secretify for your own business?
              </h3>
              <p class="is-medium">
                Take a look at the
                <a href="https://www.secretify.io/pricing">available options</a
                >.
              </p>
            </section>
          </div>
          <div class="column is-one-fifth">
            <aside class="menu">
              <ul class="menu-list">
                <li>
                  <nuxt-link :to="{ hash: '#overview' }">Overview</nuxt-link>
                </li>
                <li>
                  <nuxt-link :to="{ hash: '#how-it-works' }"
                    >How it works</nuxt-link
                  >
                </li>
                <li>
                  <nuxt-link :to="{ hash: '#security-aspects' }"
                    >Security aspects</nuxt-link
                  >
                </li>
                <li>
                  <nuxt-link :to="{ hash: '#premium-features' }"
                    >Premium features</nuxt-link
                  >
                </li>
                <li>
                  <nuxt-link :to="{ hash: '#get-it' }">Get it</nuxt-link>
                </li>
              </ul>
            </aside>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
export default {
  name: 'About',
  mounted() {
    // Track
    this.$track.pageview({})
  },
  head() {
    return {
      title: 'About - Secretify',
    }
  },
}
</script>
