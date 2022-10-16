<template>
  <div class="blank">
    <section class="section">
      <div class="container">
        <div class="is-centered">
          <h1 class="title is-1">Get Secretify for your business</h1>
          <h3 class="subtitle is-3">
            Fair pricing that scales with your business
          </h3>
        </div>
        <div class="columns is-centered mt-6">
          <div class="column">
            <div class="card">
              <div class="card-content">
                <div class="content back">
                  <span class="tag is-white">Self-managed</span>
                  <h3 class="title is-3">Free</h3>
                  <h5 class="subtitle is-5">
                    Ideal for proof of concepts and developing ideas.
                  </h5>
                  <div class="price-details">
                    <span class="currency">CHF</span>
                    <span class="price">0</span>
                  </div>
                  <a
                    href="https://hub.docker.com/r/dariocalovic/secretify"
                    class="button"
                    >Getting started</a
                  >
                </div>
              </div>
              <div class="card-content pt-1">
                <div class="content is-medium">
                  <p>Free features:</p>
                  <ul class="fake">
                    <li>
                      <b-icon icon="check"> </b-icon> Secretify CE (Community
                      Edition)
                    </li>
                    <li><b-icon icon="check"> </b-icon> Community Support</li>
                    <li><b-icon icon="check"> </b-icon> Open Source</li>
                    <li><b-icon icon="check"> </b-icon> For Development Use</li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
          <div class="column">
            <div class="card">
              <div class="card-content">
                <div class="content back colorful">
                  <span class="tag is-white">Self-managed</span>
                  <span class="tag is-white">SaaS</span>
                  <h3 class="title is-3">Premium</h3>
                  <h5 class="subtitle is-5">
                    For production-ready with the option to scale.
                  </h5>
                  <div class="price-details">
                    <span class="currency">CHF</span>
                    <span class="price">1.50</span>
                    <span class="details"
                      ><p>per employee/month</p>
                      <p>
                        or billed annually at CHF 15.- (2 months free)
                      </p></span
                    >
                  </div>
                  <div class="button-group">
                    <a
                      class="button is-primary"
                      href="mailto:sales@secretify.io?subject=I am interessted in Secretify Premium!&body=I am interessted to get more information about Secretify Premium. Our company has about 100 employee's. Let's talk!"
                      >Contact sales</a
                    >
                  </div>
                  <p class="note">
                    From 300 employees there is a discount, contact sales for
                    more information.
                  </p>
                </div>
              </div>
              <div class="card-content pt-1">
                <div class="content is-medium">
                  <p>Everything from Free, plus:</p>
                  <ul class="fake">
                    <li>
                      <b-icon icon="check"> </b-icon> Enterprise readiness
                    </li>
                    <li>
                      <b-icon icon="check"> </b-icon> Customizable branding
                    </li>
                    <li><b-icon icon="check"> </b-icon> Upgrade assistance</li>
                    <li v-show="false">
                      <b-icon icon="check"> </b-icon>
                      Integration in Microsoft Teams
                    </li>
                    <li>
                      <b-icon icon="check"> </b-icon> Integration in Microsoft
                      Outlook
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
export default {
  name: 'Pricing',
  mounted() {
    // Track
    this.$track.pageview({})
  },
  data() {
    return {
      noOfEmployees: 20,
    }
  },
  head() {
    return {
      title: 'Pricing - Secretify',
    }
  },
}
</script>
