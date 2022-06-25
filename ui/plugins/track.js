import Plausible from 'plausible-tracker'
let cfg
const pageview = (options) => {
  if (!cfg.track.enabled) {
    return
  }
  const { trackPageview } = Plausible({
    domain: cfg.track.domain,
    trackLocalhost: cfg.track.localhost,
    apiHost: cfg.track.apiURL,
  })

  trackPageview(options)
}

export default (ctx, inject) => {
  cfg = ctx.$config
  inject('track', {
    pageview,
  })
}
