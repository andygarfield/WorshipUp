import Vue from 'vue'
import Vuex from 'vuex'
import App from './App.vue'

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        writing: false
    },
    mutations: {
        toggleMode (state) {
            state.writing = !state.writing;
        }
    }
})

new Vue({
    el: '#app',
    store,
    render: h => h(App)
})
