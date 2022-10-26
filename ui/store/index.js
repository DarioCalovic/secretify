export const state = () => ({
  /* Settings */
  metaSetting: {
    hoster: {
      name: '',
      address: '',
    },
  },
  policySetting: {
    secret: {
      max_length: 500,
    },
    passphrase: {
      required: false,
    },
    storage: {
      enabled: false,
      filesystem: {
        max_filesize: 0,
        allowed_file_extensions: '',
      },
    },
  },
  offline: true,
})

export const mutations = {
  /* Settings */
  SET_METASETTINGS(state, settings) {
    state.metaSetting = settings
  },
  SET_POLICYSETTINGS(state, settings) {
    state.policySetting = settings
  },
  SET_OFFLINE(state, flag) {
    state.offline = flag
  },
}

export const actions = {
  async nuxtServerInit({ dispatch }) {
    await dispatch('get_metaSettings')
    await dispatch('get_policySettings')
  },
  async get_metaSettings({ commit }) {
    const res = await this.$repositories.settings.meta().catch((e) => {
      if (e.response) {
        if (Object.prototype.hasOwnProperty.call(e.response, 'status')) {
          console.log(e.response.status)
        } else {
          console.log(e.response)
        }
      } else {
        console.log(e)
      }
    })
    if (res) {
      const { status, data } = res
      if (status === 200 && data.data) {
        const all = data.data
        commit('SET_METASETTINGS', all)
        commit('SET_OFFLINE', false)
      }
    }
  },
  async get_policySettings({ commit }) {
    const res = await this.$repositories.settings.policy().catch((e) => {
      if (e.response) {
        if (Object.prototype.hasOwnProperty.call(e.response, 'status')) {
          console.log(e.response.status)
        } else {
          console.log(e.response)
        }
      } else {
        console.log(e)
      }
    })
    if (res) {
      const { status, data } = res
      if (status === 200 && data.data) {
        const all = data.data
        commit('SET_POLICYSETTINGS', all)
        commit('SET_OFFLINE', false)
      }
    }
  },
}
