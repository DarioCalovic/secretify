<template>
  <div>
    <section class="section is-medium">
      <div class="container darkness">
        <div class="columns is-centered">
          <div class="column">
            <h1 class="title">A brief introduction.</h1>
            <h3 class="subtitle is-3">
              Think twice where you put your sensitive information. It might be
              there forever.
            </h3>
            <p>
              Has your personal information been ever exposed in a data breach?
              I hope not. Because the most serious consequences to you can occur
              when sensitive information like pass­words or social security
              numbers are leaked. Criminals usually obtain this data either by
              physical theft, accessing public information, or through data
              breaches. In order to minimize the risk, you should think twice
              wether to put sensitive information like a password in an e-mail
              or a chat conversation.
            </p>
            <p>
              <strong>
                Your privacy should be concerning you, especially in these days.
              </strong>
              Technologies are changing the world and the daily lives of each
              and every one of us. We are surrounded by technology at every
              moment of our lives. Some mistake will be made along the way.
              People give away credit card information, passwords, social
              security numbers or any other sensitive information with ease. It
              might seem safe at first glance, but once you have fallen victim,
              you would have wished being more careful with your sensitive
              information.
            </p>
            <h3 class="subtitle is-3">How does it work?</h3>
            <div class="cl">
              {{ $config.uiURL }}/s/<span class="tag is-white is-medium"
                >n9w7fECxMKK8JwdHLB</span
              >#<span class="tag is-dark is-medium"
                >QVA6dhJgmxyycdUjdRxDKzhB7x</span
              >
            </div>
            <p>
              Secretify will generate a secret link you can share with others.
              As soon as the secret expires, it will be deleted automatically.
              To be more safe, you can have the secret only revealed once or
              even demand a password. The secret link is composed of the
              <span class="tag is-white">identifier</span> and the
              <span class="tag is-dark">encryption key</span>. The key is only
              known to you and the people you share it with. Since everything is
              encrypted end-to-end, not even an administrator with privileged
              access can make sense of any of the secrets being stored.
            </p>
            <p>
              Do not take our word for it. All of the code is open-source and
              available on
              <a href="https://github.com/DarioCalovic/secretify">Github</a>.
            </p>
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
