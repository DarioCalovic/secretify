export default function ({ $axios, $config }) {
  $axios.onRequest((config) => {
    if ($config.apiKey) {
      const key = $config.apiKey
      config.headers.common.Authorization = `Bearer ${key}`
    }
  })
}
