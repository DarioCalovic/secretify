export const state = () => ({
  /* Settings */
  metaSetting: {},
  policySetting: {},
})

export const mutations = {
  /* Settings */
  SET_METASETTINGS(state, settings) {
    state.metaSetting = settings
  },
  SET_POLICYSETTINGS(state, settings) {
    state.policySetting = settings
  },
}

export const actions = {
  async nuxtServerInit({ dispatch }) {
    await dispatch('get_metaSettings')
    await dispatch('get_policySettings')
  },
  async get_metaSettings({ commit }) {
    const res = await this.$repositories.settings.meta().catch((e) => {
      console.log(e)
    })
    if (res) {
      const { status, data } = res
      if (status === 200 && data.data) {
        const all = data.data
        commit('SET_METASETTINGS', all)
      }
    }
  },
  async get_policySettings({ commit }) {
    const res = await this.$repositories.settings.policy().catch((e) => {
      console.log(e)
    })
    if (res) {
      const { status, data } = res
      if (status === 200 && data.data) {
        const all = data.data
        commit('SET_POLICYSETTINGS', all)
      }
    }
  },
}
