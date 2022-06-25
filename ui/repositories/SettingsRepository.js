const resource = 'setting'
export default ($axios, apiURL) => ({
  meta() {
    return $axios.get(
      `${apiURL}/${resource}/_meta`,
      {},
      {
        headers: {},
      }
    )
  },
  policy() {
    return $axios.get(
      `${apiURL}/${resource}/_policy`,
      {},
      {
        headers: {},
      }
    )
  },
})
