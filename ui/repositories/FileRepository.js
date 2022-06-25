const resource = 'file'
export default ($axios, apiURL) => ({
  upload(filename, size, type, data) {
    return $axios.post(
      `${apiURL}/${resource}?filename=${filename}&size=${size}&type=${type}`,
      data,
      {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      }
    )
  },
  download(identifier) {
    return $axios.get(`${apiURL}/${resource}/${identifier}`, {
      responseType: 'blob',
      headers: {},
    })
  },
})
