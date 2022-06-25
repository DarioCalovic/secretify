const resource = 'secret'
export default ($axios, apiURL) => ({
  create(
    cipher,
    expiresAt,
    revealOnce,
    hasPassphrase,
    fileIdentifier,
    email,
    webhookAddr
  ) {
    return $axios.post(
      `${apiURL}/${resource}`,
      {
        cipher,
        expires_at: expiresAt,
        reveal_once: revealOnce,
        has_passphrase: hasPassphrase,
        file_identifier: fileIdentifier,
        email,
        webhook_addr: webhookAddr,
      },
      {
        headers: {},
      }
    )
  },
  view(identifier) {
    return $axios.get(
      `${apiURL}/${resource}/${identifier}`,
      {},
      {
        headers: {},
      }
    )
  },
  cipher(identifier) {
    return $axios.get(
      `${apiURL}/${resource}/${identifier}/_cipher`,
      {},
      {
        headers: {},
      }
    )
  },
  delete(identifier) {
    return $axios.delete(`${apiURL}/${resource}/${identifier}`, {
      headers: {},
    })
  },
})
