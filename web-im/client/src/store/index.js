import Vue from 'vue'
import Vuex from 'vuex'
import Lockr from 'lockr'
import {
    addAuth,
    removeAuth
} from '@/utils/auth'
import {
    resetRouter
} from '@/router'
import imApi from '@/api/im'

Vue.use(Vuex)


const state = {
    userInfo: null, // 用户信息
    // 权限信息
    allAuth: null, // 总权限信息 默认空 调整动态路由
    socketAction: '',
    contactSync: '',
    toContactId: 0,
    unread: 0,
    allContacts: [],
    globalConfig:[], // 全局配置
    setting: {
        sendKey: "1",
        theme: "default",
        isVoice: true,
        avatarCricle: false,
        hideMessageName: false,
        hideMessageTime: false,
    }
}

const mutations = {
    SET_AUTH: (state, data) => {
        const token = data.token
		    const id = data.user.id
        Lockr.set('authToken', token)
		    Lockr.set('id', id)
        addAuth(token, "sessionId")
    },
    catchSocketAction(state, data) {
        state.socketAction = data;
    },

    initContacts: (state, data) => {
        state.allContacts = data;
    },
    openChat: (state, data) => {
        state.toContactId = data;
        state.contactSync = Math.random().toString(36).substr(-8);
    },
    updateSetting(state, data) {
        state.userInfo.setting = data;
        state.setting = data;
    },
    setGlobalConfig(state, data) {
        state.globalConfig = data;
    }
}

const actions = {
    // 登录
    Login({commit,dispatch}, userInfo) {
        return new Promise((resolve, reject) => {
            imApi.loginAPI(userInfo).then(res => {
				console.log(res);
                commit('SET_AUTH', res.data)
				resolve()
            }).catch(error => {
				console.log(error);
                dispatch('LogOut')
                reject(error)
            })
        })
    },
    // 登出
    LogOut({
        commit
    }) {
        return new Promise((resolve, reject) => {
            commonApi.logoutAPI().then(() => {
                /** flush 清空localStorage .rm('authToken') 按照key清除 */
                Lockr.rm('authToken');
                Lockr.rm('sessionId');
                Lockr.rm('UserInfo');
                removeAuth()
                resetRouter()
                resolve()
            }).catch(error => {
                reject(error)
            })
        })
    },
    getSystemInfo({
        commit
    }) {
        return new Promise((resolve, reject) => {
            imApi.getSystemInfo().then(res=>{
                if(res.code==0){
                  Lockr.set('globalConfig',res.data);
                  commit('setGlobalConfig', res.data);
                  resolve(res)
                }
              }).catch(error => {
                reject(error)
            })
        })
    }
}

export default new Vuex.Store({
    state,
    mutations,
    actions
})
