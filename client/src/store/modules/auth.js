import axios from 'axios'

export default {
  state: {
    isLoggedIn: false,
    token: '',
    role: '',
  },
  mutations: {
    auth(state, loginStatus) {
      state.isLoggedIn = loginStatus
    },
    setToken(state, token) {
      state.token = token
      // axios.defaults.headers.common['Authorization'] = token
    },
    setRole(state, role) {
      state.role = role
      // axios.defaults.headers.common['Authorization'] = token
    },
  },
  actions: {
    async login({commit, state}, payload) {
      let options = {
        method: 'POST',
        url: '/auth/login',
        data: payload || '',
      }

      return axios(options).then((data) => {
        if (data) {
          console.log(data)
          commit('setRootId', data.user.id)
          commit('setToken', data.token)
          commit('setRole', data.role)
          commit('auth', true)
          commit('profile/initProfile', data.user, {root: true})
          return true
        }
      })
    },
    signup(_, data) {
      data.birthdate = data.birthdate.toISOString()
      data.gender = data.gender[0]
      console.log(data)
      let options = {
        method: 'POST',
        url: '/auth/register',
        data: data,
      }
      return axios(options)
    },
    logout({commit}) {
      axios({method: 'DELETE', url: '/auth/logout'}).then(() => {
        commit('auth', false)
        commit('setToken', '')
      })
    },
  },
  getters: {
    isAuthenticated(state) {
      return state.isLoggedIn
    },
  },
}