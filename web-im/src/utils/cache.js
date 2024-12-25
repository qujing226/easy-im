import store from '@/store'
import Lockr from 'lockr'
import axios from 'axios'

const cache = {
    loadingCache: function() {
        if (Lockr.get('authToken') && !axios.defaults.headers['authToken']) {
            const userInfo = Lockr.get('UserInfo')
            if (userInfo) {
                store.commit('SET_USERINFO', userInfo)
            }
        }
        store.commit('SET_APPNAME', Lockr.get('systemName'))
        store.commit('SET_APPLOGO', Lockr.get('systemLogo'))
    },

	updateAxiosCache: function() {
	    axios.defaults.headers['authToken'] = Lockr.get('authToken')
	    axios.defaults.headers['sessionId'] = Lockr.get('sessionId')
	},
	updateAxiosHeaders: function() {
	    axios.defaults.headers['authToken'] = Lockr.get('authToken')
	    axios.defaults.headers['sessionId'] = Lockr.get('sessionId')
	},

    rmAxiosCache: function() {
        Lockr.rm('authToken');
        Lockr.rm('sessionId')
    }
}

export default cache